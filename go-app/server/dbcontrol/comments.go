package dbcontrol

import (
	"context"
	"database/sql"
	"fmt"
)

type Comment struct {
	ID       int
	Comment  string
	Likes    int
	User     *User
	PostDate string
}

type User struct {
	ID   int
	Name string
}

// This get comments which have episodeID
func GetComments(ctx context.Context, episodeID int) ([]Comment, error) {
	query := "select comments.id, comments.comment, comments.likes, cast(comments.post_date as unsigned), users.id, users.user_name from comments inner join users on comments.user_id = users.id where episode_id = ?"

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed PrepareContext: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, episodeID)
	if err != nil {
		return nil, fmt.Errorf("failed QueryContext: %w", err)
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		var u User
		if err := rows.Scan(&c.ID, &c.Comment, &c.Likes, &c.PostDate, &u.ID, &u.Name); err != nil {
			return nil, fmt.Errorf("failed rows.Scan: %w", err)
		}
		c.User = &u
		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Error: %w", err)
	}

	return comments, nil
}

// This insert comment into comments table if programName and episodeDate is exist. episodeID is primary key, userID is foreign key of users table.
func CreateComment(ctx context.Context, comment, programName string, episodeDate, userID int) error {

	query := "select id from episodes where date = ? and program_id = (select id from programs where program_name = ?)"

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed PrepareContext for QueryRowContext: %w", err)
	}

	var episodeID int
	err = stmt.QueryRowContext(ctx, episodeDate, programName).Scan(&episodeID)
	if err == sql.ErrNoRows {
		return fmt.Errorf("query result is empty: %w", err)
	}
	if err != nil {
		return fmt.Errorf("failed QueryRowContext: %w", err)
	}

	stmt.Close()

	query = "insert into comments values(null, ?, ?, ?, now(), 0, now(), now())"

	stmt, err = DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed PrepareContext for ExecContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, comment, episodeID, userID)
	if err != nil {
		return fmt.Errorf("failed ExecContext: %w", err)
	}

	return nil
}

// This add 1 to likes of comment which has commentID. After update likes, this function return number of likes. If error is happened, this function return -1 and error
func UpdateCommentLikes(ctx context.Context, commentID int) (int, error) {
	query := "update comments set likes = likes + 1 where id = ?"

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return -1, fmt.Errorf("failed PrepareContext for update: %w", err)
	}

	_, err = stmt.ExecContext(ctx, commentID)
	if err != nil {
		return -1, fmt.Errorf("failed ExecContext: %w", err)
	}
	stmt.Close()

	query = "select likes from comments where id = ?"

	stmt, err = DB.PrepareContext(ctx, query)
	if err != nil {
		return -1, fmt.Errorf("failed PrepareContext for select: %w", err)
	}
	defer stmt.Close()

	var likes int
	err = stmt.QueryRowContext(ctx, commentID).Scan(&likes)
	if err != nil {
		return -1, fmt.Errorf("failed QueryRowContext: %w", err)
	}

	return likes, nil
}
