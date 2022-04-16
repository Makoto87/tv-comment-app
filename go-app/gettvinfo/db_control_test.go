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
				mock.ExpectExec(`INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES ("hoge", NOW(), NOW()),("fuga", NOW(), NOW())`).WillReturnResult(sqlmock.NewResult(2, 2))
			},
		},
		{
			name:  "test2",
			input: []string{"hoge"},
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES ("hoge", NOW(), NOW())`).WillReturnResult(sqlmock.NewResult(1, 1))
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
			gettvinfo.Db = db

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
				mock.ExpectQuery(`select id from programs where program_name in ("hoge","fuga")`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))
				mock.ExpectExec(`INSERT IGNORE INTO episodes(program_id, date, episode_number, episode_title, created_at, updated_at) VALUES (1, NOW(), null, null, NOW(), NOW()),(2, NOW(), null, null, NOW(), NOW())`).WillReturnResult(sqlmock.NewResult(2, 2))
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
			gettvinfo.Db = db

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
