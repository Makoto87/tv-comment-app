package dbcontrol_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Makoto87/tv-comment-app/go-app/server/dbcontrol"
)

func TestGetEpisodes(t *testing.T) {
	type args struct {
		programID int
		fromDate  int
		toDate    int
	}

	cases := []struct {
		name        string
		args        args
		mockClosure func(mock sqlmock.Sqlmock)
		want        []dbcontrol.Episode
	}{
		{
			name: "test1",
			args: args{programID: 10, fromDate: 20220420, toDate: 20220430},
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare("select id, cast(date as unsigned) from episodes where program_id = ? and date between ? and ? order by date desc")
				column := []string{"id", "date"}
				ep.ExpectQuery().WithArgs(10, 20220420, 20220430).WillReturnRows(sqlmock.NewRows(column).AddRow(1, 20220428).AddRow(2, 20220430))
			},
			want: []dbcontrol.Episode{
				{ID: 1, Date: 20220428},
				{ID: 2, Date: 20220430},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			dbcontrol.DB = db

			c.mockClosure(mock)

			a := c.args
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()
			episodes, err := dbcontrol.GetEpisodes(ctx, a.programID, a.fromDate, a.toDate)
			if err != nil {
				t.Errorf("error was not expected while select episodes: %s", err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(episodes, c.want) {
				t.Errorf("Want: %#v \nGot: %#v", c.want, episodes)
			}
		})
	}
}
