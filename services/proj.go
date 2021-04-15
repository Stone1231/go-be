package services

import (
	. "go-be/database"
	Model "go-be/models"
)

type ProjService struct {
}

var Proj = new(ProjService)

func (self *ProjService) Init() error {
	tx := DB.Begin()
	if err := self.Clear(); err != nil {
		tx.Rollback()
		return err
	}
	names := []string{"proj1", "proj2", "proj3"}
	count := len(names)
	for i := 0; i < count; i++ {
		model := &Model.Proj{ID: uint(i + 1), Name: names[i]}
		if err := tx.Create(model).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (self *ProjService) ReadAll() []Model.Proj {
	var list []Model.Proj
	DB.Find(&list)
	return list
}

func (self *ProjService) Delete(id uint) error {
	var model Model.Proj
	if err := DB.First(&model, id).Error; err != nil {
		return err
	}
	tx := DB.Begin()
	if err := tx.Model(&model).Association("Users").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (self *ProjService) Clear() error {
	var list []Model.Proj
	DB.Find(&list)
	tx := DB.Begin()
	for i := 0; i < len(list); i++ {
		var model Model.Proj = list[i]
		if err := tx.Model(&model).Association("Users").Clear(); err != nil {
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
