package service

import (
	"audit-rectification/internal/rectserver/global"
	"audit-rectification/internal/rectserver/model"
	"audit-rectification/internal/rectserver/repo"
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type rectService struct {
	rep repo.Repository
}

func NewRectService(rs repo.Repository) Service {
	return &rectService{rs}
}

func (rs *rectService) Fetch(ctx context.Context, str string) (model.Response, error) {
	var qv = reflect.ValueOf(rs)
	for name, pattern := range global.Patterns {
		if pattern.MatchString(str) {
			ret := qv.MethodByName(name).Call([]reflect.Value{reflect.ValueOf(pattern), reflect.ValueOf(str)})
			return ret[0].Interface().(model.Response), nil
		}
	}
	return model.Response{}, fmt.Errorf("no match found")
}

// P1 札幌营业部未按时完成的整改报告有哪些？
func (rs *rectService) P1(reg *regexp.Regexp, str string) (model.Response, error) {
	// TODO: 查询nebula数据库

	// 测试假数据
	fmt.Println(reg.FindStringSubmatch(str))
	fmt.Println("P1 called", reg.FindAllString(str, -1))
	return model.Response{Result: strings.Join(reg.FindAllString(str, -1), "")}, nil
}