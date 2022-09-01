package basedb

func (db *BaseDb) ExecStringRecord(rec *Record) error {
	key := string(rec.meta.key)
	member := string(rec.meta.member)
	switch rec.GetOper() {
	case StringSet:
		db.StringStorage.Set(key, member)
	}

	return nil
}
