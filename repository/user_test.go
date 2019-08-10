package repository

import (
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var userRepo User
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mo, err := getDBMock()
	if err != nil {
		os.Exit(-1)
	}

	defer db.Close()
	db.LogMode(true)

	userRepo = NewUserRepository(db)
	mock = mo

	code := m.Run()
	os.Exit(code)
}

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestFind(t *testing.T) {

	rows := mock.NewRows([]string{"id", "name", "uid", "account_id"}).
		AddRow(1, "hogehoge", "uid1", "hogehoge").
		AddRow(2, "fugafuga", "uid2", "fugafuga")
	mock.ExpectQuery(
		"^SELECT (.+) FROM `users` (.+)$").WillReturnRows(rows)

	users, err := userRepo.Find()
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 2 {
		t.Fatal("返り値が不正です")
	}
}

func TestFindById(t *testing.T) {

	rows := mock.NewRows([]string{"id", "name", "uid", "account_id"}).
		AddRow(1, "hogehoge", "uid1", "hogehoge")
	mock.ExpectQuery(
		"^SELECT (.+) FROM `users` (.+)$").WillReturnRows(rows)

	_, err := userRepo.FindById(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindByAccountId(t *testing.T) {

	rows := mock.NewRows([]string{"id", "name", "uid", "account_id"}).
		AddRow(1, "hogehoge", "uid1", "hogehoge")
	mock.ExpectQuery(
		"^SELECT (.+) FROM `users` (.+)account_id = ?(.+)$").WithArgs("hogehoge").WillReturnRows(rows)

	_, err := userRepo.FindByAccountId("hogehoge")
	if err != nil {
		t.Fatal(err)
	}
}
