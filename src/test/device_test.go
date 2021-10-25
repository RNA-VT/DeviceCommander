package test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rna-vt/devicecommander/graph/model"
	p "github.com/rna-vt/devicecommander/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DeviceSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	// repository Repository
	// person     *model.Person
	service p.DeviceService
}

func (s *DeviceSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.service = p.DeviceService{
		DBConnection: s.DB,
		Initialized:  true,
	}
}

func (s DeviceSuite) TestGetAll() {

	const sqlSelectAll = `SELECT * FROM "devices"`

	s.mock.ExpectQuery(sqlSelectAll).
		WillReturnRows(sqlmock.NewRows(nil))

	devices, err := s.service.GetAll()

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), devices, []*model.Device{}, "they should be equal")
}

// func (s *DeviceSuite) AfterTest(_, _ string) {
// 	require.NoError(s.T(), s.mock.ExpectationsWereMet())
// }

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceSuite))
}
