package dbcontrol

import (
	"fmt"
)

type Episode struct {
	ID   int
	Date int
}

// this get all episodes which have programID and date between fromDate and toDate
func GetEpisodes(programID, fromDate, toDate int) ([]Episode, error) {
	query := "select id, cast(date as unsigned) from episodes where program_id = ? and date between ? and ? order by date desc"

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select id, date from episodes: %w", err)
	}

	rows, err := stmt.Query(programID, fromDate, toDate)
	if err != nil {
		return nil, fmt.Errorf("failed to select episode id, date by Query: %w", err)
	}

	var episodes []Episode
	for rows.Next() {
		var e Episode
		if err := rows.Scan(&e.ID, &e.Date); err != nil {
			return nil, fmt.Errorf("failed rows.Scan id, date from episodes: %w", err)
		}
		episodes = append(episodes, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Error: %w", err)
	}
	return episodes, nil
}
