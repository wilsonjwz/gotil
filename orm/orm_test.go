package orm

import (
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

func TestInitGorm(t *testing.T) {
	// 初始化 Gorm
	gc1 := &MySQLConfig{DSN: "dsn", Active: 100, Idle: 50, ShowSQL: true}
	g := InitGorm(gc1)
	var student = &struct {
		Id   int    `gorm:"column:id;primary_key"`
		Name string `gorm:"column:name"`
	}{}
	err := g.Table("student").Select([]string{"id", "name"}).Where("id = ?", 1).First(student).Error
	if err != nil {
		xlog.Error(err)
		return
	}
}

func TestInitXorm(t *testing.T) {
	// 初始化 Xorm
	gc1 := &MySQLConfig{DSN: "dsn", Active: 100, Idle: 50, ShowSQL: true}
	x := InitXorm(gc1)
	var student = &struct {
		Id   int    `xorm:"'id'"`
		Name string `xorm:"'name'"`
	}{}
	_, err := x.Table("student").Select("id,name").Where("id = ?", 1).Get(student)
	if err != nil {
		xlog.Error(err)
		return
	}
}
