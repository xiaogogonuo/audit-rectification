package rect

import (
	"audit-rectification/internal/rectserver/model"
	"audit-rectification/internal/rectserver/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	nebula "github.com/vesoft-inc/nebula-go/v3"
	"net/http"
)

type nebulaRepository struct {
	*nebula.Session
}

func NewNebulaRepository(s *nebula.Session) repo.Repository {
	return &nebulaRepository{s}
}

func (nr *nebulaRepository) checkResultSet(prefix string, res *nebula.ResultSet) error {
	if !res.IsSucceed() {
		return fmt.Errorf("%s, ErrorCode: %v, ErrorMsg: %s", prefix, res.GetErrorCode(), res.GetErrorMsg())
	}
	return nil
}

func (nr *nebulaRepository) Fetch(c *gin.Context, str string) (model.Response, error) {
	prefix := fmt.Sprintf("CREATE SPACE IF NOT EXISTS %s(vid_type=FIXED_STRING(20)); ", str)
	resultSet, err := nr.Execute(prefix)
	if err != nil {
		return model.Response{
			StatusCode:   http.StatusOK,
			ErrorMessage: err.Error(),
			Result:       "succeed",
		}, nil
	}
	err = nr.checkResultSet(prefix, resultSet)
	if err != nil {
		return model.Response{
			StatusCode:   http.StatusOK,
			ErrorMessage: err.Error(),
			Result:       "succeed",
		}, nil
	}
	return model.Response{}, nil
}
