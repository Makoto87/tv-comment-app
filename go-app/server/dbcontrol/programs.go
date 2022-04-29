package dbcontrol

import (
	"context"
	"fmt"
)

type Program struct {
	ID   int
	Name string
}

// This function get all programs which part of program name matches argument.
func GetPrograms(ctx context.Context, search string) ([]Program, error) {
	query := "select id, program_name from programs where program_name like ?"
	if search == "" {
		search = "%"
	}

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select id & program_name: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, search)
	if err != nil {
		return nil, fmt.Errorf("failed to select program id, program_name by Query: %w", err)
	}
	defer rows.Close()

	var programs []Program
	for rows.Next() {
		var p Program
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, fmt.Errorf("failed rows.Scan id & program_name from program: %w", err)
		}
		programs = append(programs, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Error: %w", err)
	}

	return programs, nil
}
