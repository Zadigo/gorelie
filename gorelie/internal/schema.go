package internal

type CellType uint8

type UpdateMode int

type TableColumn struct {
	Name string
	Type CellType
}

type Cell struct{}

type Row []Cell

func (row Row) EncodeKey(schema *Schema) (key []byte) {
	return nil
}

func (row Row) EncodeValue(schema *Schema) (value []byte) {
	return nil
}

func (row Row) DecodeKey(schema *Schema, key []byte) (err error) {
	return nil
}

func (row Row) DecodeValue(schema *Schema, value []byte) (err error) {
	return nil
}

type Schema struct {
	Table       string
	Columns     []TableColumn
	PrimaryKeys []int
}

func (s *Schema) NewRow() Row {
	return make(Row, len(s.Columns))
}
