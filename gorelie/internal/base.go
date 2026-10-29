package internal

type KeyValueStorage struct{
	
}

// A simple key-value database
type Database struct {
	log    Log
	memory map[string][]byte
}

func (k *Database) Open() error {
	k.memory = map[string][]byte{}
	return nil
}

func (k *Database) Close() error {
	return nil
}

func (kv *Database) Set(key []byte, val []byte) (updated bool, err error) {
	return true, nil
}

func (kv *Database) Del(key []byte) (deleted bool, err error) {
	return true, nil
}

func (db *Database) Select(schema *Schema, row Row) (ok bool, err error) {
	return true, nil
}

func (db *Database) Insert(schema *Schema, row Row) (updated bool, err error) {
	key := row.EncodeKey(schema)
	value := row.EncodeValue(schema)
	return db.SetEx(key, value, ModeInsert)
}

func (db *Database) Upsert(schema *Schema, row Row) (updated bool, err error) {
	return true, nil
}

func (db *Database) Update(schema *Schema, row Row) (updated bool, err error) {
	return true, nil
}

func (db *Database) Delete(schema *Schema, row Row) (deleted bool, err error) {
	return true, nil
}

func (kv *Database) SetEx(key []byte, val []byte, mode UpdateMode) (bool, error) {
	return true, nil
}
