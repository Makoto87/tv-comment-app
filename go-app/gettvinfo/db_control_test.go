package gettvinfo_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func TestProgramInsert(t *testing.T) {
	cases := []struct {
		name        string
		input       []string
		mockClosure func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "test1",
			input: []string{"hoge", "fuga"},
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare(`INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES (?, NOW(), NOW()),(?, NOW(), NOW())`)
				ep.ExpectExec().WithArgs("hoge", "fuga") .WillReturnResult(sqlmock.NewResult(2, 2))
			},
		},
		{
			name:  "test2",
			input: []string{"hoge"},
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare(`INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES (?, NOW(), NOW())`)
				ep.ExpectExec().WithArgs("hoge").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// mock database を用意
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			gettvinfo.DB = db

			c.mockClosure(mock)

			// 関数実行
			if err = gettvinfo.ProgramInsert(c.input); err != nil {
				t.Errorf("error was not expected while insert programs: %s", err)
			}

			// mockと実際の関数を比較
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestEpisodeInsert(t *testing.T) {
	cases := []struct {
		name        string
		input       []string
		mockClosure func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "test1",
			input: []string{"hoge", "fuga"},
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare(`select id from programs where program_name in (?,?)`)
				ep.ExpectQuery().WithArgs("hoge", "fuga").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

				ep = mock.ExpectPrepare(`INSERT IGNORE INTO episodes(program_id, date, episode_number, episode_title, created_at, updated_at) VALUES (?, NOW(), null, null, NOW(), NOW()),(?, NOW(), null, null, NOW(), NOW())`)
				ep.ExpectExec().WithArgs(1,2).WillReturnResult(sqlmock.NewResult(2,2))
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// mock database を用意
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			gettvinfo.DB = db

			c.mockClosure(mock)

			// 関数実行
			if err = gettvinfo.EpisodeInsert(c.input); err != nil {
				t.Errorf("error was not expected while insert programs: %s", err)
			}

			// mockと実際の関数を比較
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
