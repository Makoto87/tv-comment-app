package dbcontrol_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Makoto87/tv-comment-app/go-app/server/dbcontrol"
)

func TestGetPrograms(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		mockClosure func(mock sqlmock.Sqlmock)
		want        []dbcontrol.Program
	}{
		{
			name:  "test1",
			input: "test",
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare("select id, program_name from programs where program_name like ?")
				column := []string{"id", "program_name"}
				ep.ExpectQuery().WithArgs("test").WillReturnRows(sqlmock.NewRows(column).AddRow(1, "test1").AddRow(2, "test2"))
			},
			want: []dbcontrol.Program{
				{ID: 1, Name: "test1"},
				{ID: 2, Name: "test2"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			dbcontrol.DB = db

			c.mockClosure(mock)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
			defer cancel()

			programs, err := dbcontrol.GetPrograms(ctx, c.input)
			if err != nil {
				t.Errorf("error was not expected while select programs: %s", err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(programs, c.want) {
				t.Errorf("Want: %#v \nGot: %#v", c.want, programs)
			}
		})
	}
}
