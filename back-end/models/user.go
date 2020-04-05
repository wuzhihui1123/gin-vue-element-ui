package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModel

	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (ret []User, err error) {

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&ret).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&ret).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}
