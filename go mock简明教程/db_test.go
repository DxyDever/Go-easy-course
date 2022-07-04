package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB2(t *testing.T) {
	ctrl := gomock.NewController(t)

	//断言DB.get方法是否被调用
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	//这里预定义返回了错误error
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}

	m.EXPECT().Get(gomock.Eq("Jam")).Return(1, nil)

	if v := GetFromDB(m, "Jam"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
