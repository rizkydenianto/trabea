package roles

import (
	"trabea/app"
)

func GetRolesData() ([]Role, error) {
	conn, err := app.InitDB()
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(`
		SELECT
			id,
			name
		FROM roles
	`)
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	var roles []Role
	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.Id, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if len(roles) <= 0 {
		return nil, nil
	}

	return roles, nil
}
