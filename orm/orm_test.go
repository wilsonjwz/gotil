package orm

import (
	"testing"
)

var (
	dsn = "root:root@tcp(mysql:3306)/school?parseTime=true&loc=Local&charset=utf8mb4"
)

type Student struct {
	Id   int    `gorm:"column:id;primary_key" xorm:"'id' pk"`
	Name string `gorm:"column:name" xorm:"'name'"`
}

func (m *Student) TableName() string {
	return "student"
}

func TestInitGorm(t *testing.T) {
	// 初始化 Gorm
	//gc1 := &MySQLConfig{DSN: dsn, Active: 10, Idle: 10, IdleTimeout: 10, ShowSQL: true}
	//g := InitGorm(gc1)
	//
	//student := &Student{Name: "Jerry"}
	//g.AutoMigrate(student)
	//
	//err := g.Create(student).Error
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//var sts []*Student
	//err = g.Model(student).Select([]string{"id", "name"}).Where("id = ?", 1).Find(&sts).Error
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//for _, v := range sts {
	//	xlog.Debug("gorm:", v)
	//}
}

func TestInitXorm(t *testing.T) {
	// 初始化 Xorm
	//gc1 := &MySQLConfig{DSN: dsn, Active: 10, Idle: 10, IdleTimeout: 10, ShowSQL: true}
	//x := InitXorm(gc1)
	//
	//student := new(Student)
	//x.Sync2(student)
	//
	//_, err := x.Insert(&Student{Name: "Jerry"})
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//_, err = x.Table(student.TableName()).Select("id,name").Where("id = ?", 1).Get(student)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debug("xorm:", student)
}
