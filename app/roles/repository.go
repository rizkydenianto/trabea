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
	if err != nil {
		defer conn.Close()
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.Id, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
