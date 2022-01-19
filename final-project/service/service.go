package service

import (
	"final-project/entity"
	"final-project/repository"
	"final-project/utils/error_utils"
	"log"
)

type ServiceInterface interface {
	Create(*entity.TodoList) (*entity.TodoList, error_utils.MessageErr)
	Update(*entity.TodoList) (*entity.TodoList, error_utils.MessageErr)
	Delete(int)
	GetAll() (*[]entity.TodoList, error_utils.MessageErr)
	GetById(int) (*entity.TodoList, error_utils.MessageErr)
}

type service struct{}

var Serv ServiceInterface = &service{}

func (s *service) Create(todoReq *entity.TodoList) (*entity.TodoList, error_utils.MessageErr) {
	res, err := repository.Repo.Create(todoReq)

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}
func (s *service) Update(todoReq *entity.TodoList) (*entity.TodoList, error_utils.MessageErr) {
	res, err := repository.Repo.Update(todoReq)

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}
func (s *service) Delete(id int) {
	repository.Repo.Delete(id)
}
func (s *service) GetAll() (*[]entity.TodoList, error_utils.MessageErr) {
	res, err := repository.Repo.GetAll()

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}
func (s *service) GetById(id int) (*entity.TodoList, error_utils.MessageErr) {
	res, err := repository.Repo.GetById(id)

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}
