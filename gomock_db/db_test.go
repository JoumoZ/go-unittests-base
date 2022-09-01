package gomock_db

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // assert if Get function used

	m := NewMockDB(ctrl)
	//m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	//m.EXPECT().Get(gomock.Any()).Return(630, nil)
	//m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	//m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))
	m.EXPECT().Get(gomock.Eq("Tom2")).Return(630, errors.New("not exist"))
	//
	//if v := GetFromDB(m, "Tom"); v != -1 {
	//	t.Fatal("expected -1, but got", v)
	//}
	if v := GetFromDB(m, "Tom2"); v != 630 {
		t.Fatalf("expected 630, but %+v", v)
	}
	//if v := GetFromDB(m, "Tom"); v != -1 {
	//	t.Fatal("expected -1, but got", v)
	//}
	//if v := GetFromDB(m, "Tom"); v != -1 {
	//	t.Fatal("expected -1, but got", v)
	//}
}
