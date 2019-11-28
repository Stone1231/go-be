package services

type InitService struct {
}

var Init = new(InitService)

func (self *InitService) All() error {
	if err := Dept.Init(); err != nil {
		return err
	}
	if err := Proj.Init(); err != nil {
		return err
	}
	if err := User.Init(); err != nil {
		return err
	}
	if err := File.Clear("img"); err != nil {
		return err
	}
	return nil
}

func (self *InitService) Clear() error {
	if err := Dept.Clear(); err != nil {
		return err
	}
	if err := Proj.Clear(); err != nil {
		return err
	}
	if err := User.Clear(); err != nil {
		return err
	}
	if err := File.Clear("img"); err != nil {
		return err
	}
	return nil
}
