package dbcontrol

import "fmt"

type Program struct {
	ID   int
	Name string
}

// This function get all programs which part of program name matches argument.
func GetPrograms(search string) ([]Program, error) {
	query := "select id, program_name from programs where program_name like ?"
	if search == "" {
		search = "%"
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select id & program_name: %w", err)
	}

	rows, err := stmt.Query(search)
	if err != nil {
		return nil, fmt.Errorf("failed to select program id, program_name by Query: %w", err)
	}

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
