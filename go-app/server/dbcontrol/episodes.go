package dbcontrol

import (
	"context"
	"fmt"
)

type Episode struct {
	ID   int
	Date int
}

// this get all episodes which have programID and date between fromDate and toDate
func GetEpisodes(ctx context.Context, programID, fromDate, toDate int) ([]Episode, error) {
	query := "select id, cast(date as unsigned) from episodes where program_id = ? and date between ? and ? order by date desc"

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed PrepareContext: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, programID, fromDate, toDate)
	if err != nil {
		return nil, fmt.Errorf("failed QueryContext: %w", err)
	}
	defer rows.Close()

	var episodes []Episode
	for rows.Next() {
		var e Episode
		if err := rows.Scan(&e.ID, &e.Date); err != nil {
			return nil, fmt.Errorf("failed rows.Scan: %w", err)
		}
		episodes = append(episodes, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Error: %w", err)
	}
	return episodes, nil
}
