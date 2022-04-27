package dbcontrol

import "fmt"

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

// this get comments which have episodeID
func GetComments(episodeID int) ([]Comment, error) {
	query := "select comments.id, comments.comment, comments.likes, cast(comments.post_date as unsigned), users.id, users.user_name from comments inner join users on comments.user_id = users.id where episode_id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select from comments: %w", err)
	}

	rows, err := stmt.Query(episodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to select comment by Query: %w", err)
	}

	var comments []Comment
	for rows.Next() {
		var c Comment
		var u User
		if err := rows.Scan(&c.ID, &c.Comment, &c.Likes, &c.PostDate, &u.ID, &u.Name); err != nil {
			return nil, fmt.Errorf("failed rows.Scan from comments: %w", err)
		}
		c.User = &u
		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Error: %w", err)
	}

	return comments, nil
}