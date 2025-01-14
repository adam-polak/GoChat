package controller

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Controller struct {
		// Temporary until db connection set up
		DB *pgxpool.Pool
	}
)
