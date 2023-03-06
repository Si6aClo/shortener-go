package v1

import (
	"shortener/db"
)

type TaskServerV1 struct {
	PgContext *db.PgContext
}

// NewTaskServerV1 creates a new TaskServerV1 with created database connection
func NewTaskServerV1() *TaskServerV1 {
	pgContext, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	return &TaskServerV1{
		PgContext: pgContext,
	}
}
