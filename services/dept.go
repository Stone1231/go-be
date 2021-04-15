package services

import (
	. "go-be/database"
	Model "go-be/models"
)

type DeptService struct {
}

var Dept = new(DeptService)

func (self *DeptService) Init() error {
	tx := DB.Begin()
	if err := self.Clear(); err != nil {
		tx.Rollback()
		return err
	}
	names := []string{"dept1", "dept2", "dept3"}
	count := len(names)
	for i := 0; i < count; i++ {
		model := &Model.Dept{ID: uint(i + 1), Name: names[i]}
		if err := tx.Create(model).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (self *DeptService) ReadAll() []Model.Dept {
	var list []Model.Dept
	DB.Find(&list)
	return list
}

func (self *DeptService) ReadOne(id uint) Model.Dept {
	var model Model.Dept
	DB.First(&model, id)
	DB.Model(&model).Association("Users").Find(&model.Users)
	return model
}

func (self *DeptService) Delete(id uint) error {
	var model Model.Dept
	// maybe record not found will delete all dept
	if err := DB.First(&model, id).Error; err != nil {
		return err
	}
	DB.Model(&model).Association("Users").Find(&model.Users)
	tx := DB.Begin()
	// delete users's user_proj
	for i := 0; i < len(model.Users); i++ {
		if err := tx.Model(&(model.Users[i])).Association("Projs").Clear(); err != nil {
			tx.Rollback()
			return err
		}
	}
	//delete dept's users
	if err := tx.Delete(Model.User{}, "dept_id = ?", &model.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (self *DeptService) Clear() error {
	var list []Model.Dept
	DB.Find(&list)
	tx := DB.Begin()
	for i := 0; i < len(list); i++ {
		var model Model.Dept = list[i]
		DB.Model(&model).Association("Users").Find(&model.Users)
		// delete users's user_proj
		for j := 0; j < len(model.Users); j++ {
			if err := tx.Model(&(model.Users[j])).Association("Projs").Clear(); err != nil {
				tx.Rollback()
				return err
			}
		}
		//delete dept's users
		if err := tx.Delete(Model.User{}, "dept_id = ?", &model.ID).Error; err != nil {
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
