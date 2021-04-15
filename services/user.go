package services

import (
	. "go-be/database"
	Model "go-be/models"
	"time"

	. "github.com/ahmetb/go-linq"
)

type UserService struct {
}

var User = new(UserService)

func (self *UserService) Create(model *Model.User) error {
	DB.First(&(model.Dept), model.DeptID)
	DB.Where("ID in (?)", model.ProjIDs).Find(&(model.Projs))
	return DB.Create(model).Error
}

func (self *UserService) Update(model *Model.User) error {
	var ori Model.User
	if err := DB.First(&ori, model.ID).Error; err != nil {
		return err
	}

	// FIXME
	model.Dept.ID = model.DeptID
	DB.First(&model.Dept)
	//DB.First(&(model.Dept), model.DeptID)

	DB.Where("ID in (?)", model.ProjIDs).Find(&model.Projs)

	tx := DB.Begin()
	tx.Model(&ori).Association("Projs").Clear() //ori clear
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (self *UserService) ReadOne(id uint) Model.User {
	var model Model.User
	//DB.First(&model, id).Association("Dept").Find(&model.Dept)
	DB.First(&model, id)
	DB.Model(&model).Association("Dept").Find(&model.Dept)

	DB.Model(model).Association("Projs").Find(&model.Projs)
	// db.Preload("Projs").First(&model)
	model.BirthdayStr = model.Birthday.Format("2006-01-02")

	From(model.Projs).Select(func(c interface{}) interface{} {
		return c.(Model.Proj).ID
	}).ToSlice(&(model.ProjIDs))

	return model
}

func (self *UserService) ReadAll() []Model.User {
	var list []Model.User
	DB.Find(&list)
	return list
}

func (self *UserService) Delete(id uint) error {
	var model Model.User
	if err := DB.First(&model, id).Error; err != nil {
		return err
	}
	tx := DB.Begin()
	if err := tx.Model(&model).Association("Projs").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (self *UserService) Clear() error {
	var list []Model.User
	DB.Find(&list)
	tx := DB.Begin()
	for i := 0; i < len(list); i++ {
		model := list[i]
		if err := tx.Model(&model).Association("Projs").Clear(); err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Delete(&model).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (self *UserService) Init() error {
	tx := DB.Begin()

	if err := self.Clear(); err != nil {
		tx.Rollback()
		return err
	}

	model := &Model.User{}
	model.Name = "user1"
	model.Hight = 170
	model.Photo = ""
	model.Birthday = time.Date(1977, time.December, 31, 0, 0, 0, 0, time.UTC)

	var dept Model.Dept
	tx.First(&dept, 1)
	model.Dept = dept

	var projs []Model.Proj
	tx.Find(&projs)

	model.Projs = projs

	if err := tx.Create(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (self *UserService) Query(name string) []Model.User {
	var list []Model.User
	DB.Where("name LIKE ?", "%"+name+"%").Find(&list)
	return list
}
