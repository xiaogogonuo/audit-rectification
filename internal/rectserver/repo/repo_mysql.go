package repo

import (
	"audit-rectification/internal/rectserver/model"
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

type mysqlRepository struct {
	*sql.DB
}

func NewMysqlRepository(db *sql.DB) Repository {
	return &mysqlRepository{db}
}

func (m *mysqlRepository) Fetch(ctx context.Context, str string) (model.Response, error) {
	// TODO: 查询逻辑
	fmt.Println(str)
	//_, _ = m.Query("SELECT * FROM student", "")

	// 返回假数据
	return model.Response{
		StatusCode: http.StatusOK,
		ErrorMessage: "",
		Result: "succeed",
	}, nil
}