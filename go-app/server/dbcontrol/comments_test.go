package dbcontrol_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Makoto87/tv-comment-app/go-app/server/dbcontrol"
)

func TestGetComments(t *testing.T) {

	cases := []struct {
		name        string
		args        int
		mockClosure func(mock sqlmock.Sqlmock)
		want        []dbcontrol.Comment
	}{
		{
			name: "test1",
			args: 1,
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare("select comments.id, comments.comment, comments.likes, cast(comments.post_date as unsigned), users.id, users.user_name from comments inner join users on comments.user_id = users.id where episode_id = ?")
				column := []string{"comments_id", "comment", "likes", "post_date", "users_id", "user_name"}
				ep.ExpectQuery().WithArgs(1).WillReturnRows(sqlmock.NewRows(column).AddRow(1, "test comment1", 3, 20220430, 1, "ゲストユーザー"))
			},
			want: []dbcontrol.Comment{
				{ID: 1, Comment: "test comment1", Likes: 3, User: &dbcontrol.User{ID: 1, Name: "ゲストユーザー"}, PostDate: "20220430"},
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

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()
			comments, err := dbcontrol.GetComments(ctx, c.args)
			if err != nil {
				t.Errorf("error was not expected while select comments: %s", err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(comments, c.want) {
				t.Errorf("Want: %#v \nGot: %#v", c.want, comments)
			}
		})
	}
}

func TestCreateComment(t *testing.T) {
	type args struct {
		episodeID int
		userID    int
		comment   string
	}

	cases := []struct {
		name        string
		args        args
		mockClosure func(mock sqlmock.Sqlmock)
		want        error
	}{
		{
			name: "test1",
			args: args{episodeID: 1, userID: 1, comment: "create test comment"},
			mockClosure: func(mock sqlmock.Sqlmock) {
				ep := mock.ExpectPrepare("insert into comments values(null, ?, ?, ?, now(), 0, now(), now())")
				ep.ExpectExec().WithArgs("create test comment", 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: nil,
		},
	}

	for _, c := range cases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		dbcontrol.DB = db
		c.mockClosure(mock)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		a := c.args
		err = dbcontrol.CreateComment(ctx, a.episodeID, a.userID, a.comment)
		if err != nil {
			t.Errorf("error was not expected while insert into comments: %s", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}
}
