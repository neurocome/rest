package services

import "github.com/neurocome/rest/entity"

type EmployeService interface {
	Save(entity.Employe) entity.Employe
	FindAll() []entity.Employe
}

type employeService struct {
	employe []entity.Employe
}

func New() EmployeService {
	return &employeService{}
}

func (service *employeService) Save(employe entity.Employe) entity.Employe {
	service.employe = append(service.employe, employe)
	return employe
}

func (service *employeService) FindAll() []entity.Employe {
	return service.employe
}
