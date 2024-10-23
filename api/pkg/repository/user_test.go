package repository

import (
	Telegram_Market "Telegram-Market"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var writeTest *kafka.Writer

func init() {
	writeTest = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})
}

func TestUserTelegramSql_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserTelegramSql(db, writeTest)

	type args struct {
		userId int64
	}

	type mockBehavior func(args args, id int64)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int64
		wantErr      bool
	}{
		{
			name: "ok",
			args: args{
				userId: int64(5213124123),
			},
			id: int64(5213124123),
			mockBehavior: func(args args, id int64) {
				mock.ExpectExec("INSERT INTO Users").WithArgs(args.userId).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			err = r.CreateUser(testCase.id)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserTelegramSql_GetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserTelegramSql(db, writeTest)

	type args struct {
		userId int64
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         Telegram_Market.Users
		wantErr      bool
	}{
		{
			name: "Ok",
			args: args{
				userId: int64(1),
			},
			want: Telegram_Market.Users{
				Id:      1,
				UserId:  1,
				Balance: 100.0,
			},
			mockBehavior: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "balance"}).AddRow(1, 1, 100.0)
				mock.ExpectQuery("SELECT id, user_id, balance FROM Users WHERE user_id = ?").WithArgs(args.userId).WillReturnRows(rows)
			},
		},
		{
			name: "User not found",
			args: args{
				userId: int64(2),
			},
			want: Telegram_Market.Users{},
			mockBehavior: func(args args) {
				mock.ExpectQuery("SELECT id, user_id, balance FROM Users WHERE user_id = ?").
					WithArgs(args.userId).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			user, err := r.GetUserById(testCase.args.userId)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, user)
			}
		})
	}
}

func TestUserTelegramSql_DeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserTelegramSql(db, writeTest)

	type args struct {
		userId int64
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         error
		wantErr      bool
	}{
		{
			name: "Ok",
			args: args{userId: int64(1)},
			mockBehavior: func(args args) {
				mock.ExpectExec("DELETE FROM Users WHERE user_id = ?").WithArgs(args.userId).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err = r.DeleteUser(testCase.args.userId)

			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, err)
			}
		})
	}
}

func TestUserTelegramSql_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserTelegramSql(db, writeTest)

	type args struct {
		userId int64
		user   Telegram_Market.Users
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         error
		wantErr      bool
	}{
		{
			name: "ok",
			args: args{
				userId: int64(1),
				user:   Telegram_Market.Users{Balance: 100.0},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("UPDATE Users SET balance = balance+? WHERE user_id = ?").
					WithArgs(args.user.Balance, args.userId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err = r.UpdateUser(testCase.args.userId, testCase.args.user)

			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, err)
			}
		})
	}
}
