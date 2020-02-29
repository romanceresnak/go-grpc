package repos_test

import (
	"database/sql"
	"github.com/DATA-dog/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/romanceresnak/go-grpc/repos"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err   error
	db    *xorm.Engine
	dbSql *sql.DB
	mock  sqlmock.Sqlmock

	gr            repos.GlobalRepository
	truncateUsers = func() {
		mock.ExpectQuery("TRUNCATE user").
			WillReturnRows(sqlmock.NewRows([]string{}))

		_, err = db.Query("TRUNCATE user")
		Ω(err).To(BeNil()) //will it be nil ?
	}

	clearDatabase = func() {
		if db == nil {
			Fail("unable to run test because database is missing")
		}

		truncateUsers()

		return
	}
)

var _ = BeforeSuite(func() {
	db, err = xorm.NewEngine("mysql", "") //real database
	Ω(err).To(BeNil())
	dbSql, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //fake database
	Ω(err).To(BeNil())
	db.DB().DB = dbSql

	gr = repos.GlobalRepo(db)
})

var _ = AfterSuite(func() {
	err = mock.ExpectationsWereMet()
	Ω(err).To(BeNil())
})

func TestGoGrpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGrpc Suite")
}
