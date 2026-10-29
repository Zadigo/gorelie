package tests

import (
	"testing"

	"github.com/Zadigo/gorelie/internal"
	"github.com/stretchr/testify/suite"
)

type TestDatabase struct {
	suite.Suite
	db internal.Database
}

func (d *TestDatabase) SetupTest() {
	d.db = internal.Database{}
}

func (d *TestDatabase) TestOpen() {
	returnValue := d.db.Open()
	d.NoError(returnValue)
}

func (d *TestDatabase) TestClose() {
	returnValue := d.db.Close()
	d.NoError(returnValue)
}

func TestRunDatabaseSuite(t *testing.T) {
	suite.Run(t, new(TestDatabase))
}

type TestDatabaseCrud struct {
	suite.Suite
	db internal.Database
}

func (d *TestDatabaseCrud) SetupTest() {
	d.db = internal.Database{}
}

func (d *TestDatabaseCrud) TestSelect() {

}

func (d *TestDatabaseCrud) TestInsert() {

}

func (d *TestDatabaseCrud) TestUpsert() {

}

func (d *TestDatabaseCrud) TestUpdate() {

}

func (d *TestDatabaseCrud) TestDelete() {

}

func TestRunDatabaseCrudSuite(t *testing.T) {
	suite.Run(t, new(TestDatabaseCrud))
}
