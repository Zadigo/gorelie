package internal

const (
	ModeUpsert UpdateMode = 0 // insert or update
	ModeInsert UpdateMode = 1 // insert new
	ModeUpdate UpdateMode = 2 // update existing
)

const (
	TypeI64 CellType = 1
	TypeStr CellType = 2
)
