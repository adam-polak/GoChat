package controller

import "database/sql"

type (
	Controller struct {
		// Temporary until db connection set up
		DB *sql.DB
	}
)
