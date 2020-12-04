package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//此处可以使用db而没有引入包，是因为同一个包下所以db *gorm.DB可以直接使用
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	//默认返回tags
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	//默认返回count
	return
}

//判断该文章是否存在该标签
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//添加文章标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

//在创建前执行
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

//在更新前执行
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

//通过id 判断是否存在该标签
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	fmt.Println(tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//编辑标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}

//删除标签
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
