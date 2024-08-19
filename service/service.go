package service

import (
	"context"
	"database/sql"

	wrapper "github.com/perisynctechnologies/formSpree/utils"
)

type IService interface {
	ProjectSetUp(ctx context.Context, p Project) error
	DeleteProject(ctx context.Context, f Filter) error
	CreatedForm(ctx context.Context, f Form) error
}

type Service struct {
	jwtKey string
	db     *sql.DB
	mail   wrapper.ISetting
}

func New(jwtKey string, db *sql.DB, mail wrapper.ISetting) IService {
	return &Service{
		jwtKey: jwtKey,
		db:     db,
		mail:   mail,
	}
}
