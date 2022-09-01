package basedb

import (
	"encoding/binary"
	"errors"
	"time"
)

var (
	ErrInvalidEntry = errors.New("invalid entry")
)

const (
	entryHeaderSize = 22
)

type (
	Record struct {
		meta      *meta
		state     uint16
		timestamp uint64
	}

	meta struct {
		key        []byte
		member     []byte
		value      []byte
		keySize    uint32
		memberSize uint32
		valueSize  uint32
	}
)

func newRecordInternal(key, member, value []byte, state uint16, timestamp uint64) *Record {
	return &Record{
		state:     state,
		timestamp: timestamp,
		meta: &meta{
			key:        key,
			member:     member,
			value:      value,
			keySize:    uint32(len(key)),
			memberSize: uint32(len(member)),
			valueSize:  uint32(len(value)),
		},
	}
}

func NewRecord(key, member []byte, dataType, operation uint16) *Record {
	var state uint16 = 0
	state = state | (dataType << 8)
	state = state | operation
	return newRecordInternal(key, member, nil, state, uint64(time.Now().UnixNano()))
}

func (rec *Record) size() uint32 {
	return entryHeaderSize + rec.meta.keySize + rec.meta.memberSize + rec.meta.valueSize
}

func (rec *Record) GetType() uint16 {
	return rec.state >> 8
}

func (rec *Record) GetOper() uint16 {
	return rec.state & (2<<7 - 1)
}

func EncodeRecord(rec *Record) ([]byte, error) {
	if rec == nil || rec.meta.keySize == 0 {
		return nil, ErrInvalidEntry
	}

	buf := make([]byte, rec.size())
	binary.BigEndian.PutUint32(buf[0:4], rec.meta.keySize)
	binary.BigEndian.PutUint32(buf[4:8], rec.meta.memberSize)
	binary.BigEndian.PutUint32(buf[8:12], rec.meta.valueSize)
	binary.BigEndian.PutUint16(buf[12:14], rec.state)
	binary.BigEndian.PutUint64(buf[14:22], rec.timestamp)

	bufStart, bufEnd := entryHeaderSize, entryHeaderSize+rec.meta.keySize
	copy(buf[bufStart:bufEnd], rec.meta.key)

	bufStart = int(bufEnd)
	bufEnd += rec.meta.memberSize
	copy(buf[bufStart:bufEnd], rec.meta.member)

	if rec.meta.valueSize > 0 {
		bufStart = int(bufEnd)
		bufEnd += rec.meta.valueSize
		copy(buf[bufStart:bufEnd], rec.meta.value)
	}
	return buf, nil
}

func DecodeRecord(buf []byte) (*Record, error) {
	keySize := binary.BigEndian.Uint32(buf[0:4])
	memberSize := binary.BigEndian.Uint32(buf[4:8])
	valeSize := binary.BigEndian.Uint32(buf[8:12])
	state := binary.BigEndian.Uint16(buf[12:14])
	timestamp := binary.BigEndian.Uint64(buf[14:32])

	bufStart, bufEnd := entryHeaderSize, entryHeaderSize+keySize
	key := buf[bufStart:bufEnd]

	bufStart = int(bufEnd)
	bufEnd += memberSize
	member := buf[bufStart:bufEnd]

	bufStart = int(bufEnd)
	bufEnd += valeSize
	value := buf[bufStart:bufEnd]

	return &Record{
		meta: &meta{
			keySize:    keySize,
			memberSize: memberSize,
			valueSize:  valeSize,
			key:        key,
			member:     member,
			value:      value,
		},
		state:     state,
		timestamp: timestamp,
	}, nil
}
