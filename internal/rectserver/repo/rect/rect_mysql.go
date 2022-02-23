package rect

import (
	"audit-rectification/internal/rectserver/model"
	"audit-rectification/internal/rectserver/repo"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type mysqlRepository struct {
	*sql.DB
}

func NewMysqlRepository(db *sql.DB) repo.Repository {
	return &mysqlRepository{db}
}

func (mr *mysqlRepository) fetch(c *gin.Context, query string, args ...interface{}) ([]model.Graph, error) {
	rows, err := mr.QueryContext(c, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]model.Graph, 0)
	for rows.Next() {
		var data model.Graph
		err := rows.Scan(
			&data.VID,
			&data.Tag,
			&data.Edge,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (mr *mysqlRepository) Fetch(c *gin.Context, str string) (model.Response, error) {
	query := "Select id, title, content From posts limit ?"
	_, _ = mr.fetch(c, query, 10)
	return model.Response{}, nil
}

func (mr *mysqlRepository) GetByID(c *gin.Context, vid string) (model.Response, error) {
	query := "Select id, title, content From posts where id=?"
	_, _ = mr.fetch(c, query, vid)
	return model.Response{}, nil
}

func (mr *mysqlRepository) Create(c *gin.Context, p *model.Graph) (int64, error) {
	query := "Insert posts SET title=?, content=?"

	stmt, err := mr.PrepareContext(c, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(c, p.VID, p.Tag)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (mr *mysqlRepository) Update(c *gin.Context, p *model.Graph) (*model.Graph, error) {
	query := "Update posts set title=?, content=? where id=?"

	stmt, err := mr.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		c,
		p.VID,
		p.Tag,
		p.Edge,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (mr *mysqlRepository) Delete(c *gin.Context, id int64) (bool, error) {
	query := "Delete From posts Where id=?"

	stmt, err := mr.PrepareContext(c, query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(c, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
