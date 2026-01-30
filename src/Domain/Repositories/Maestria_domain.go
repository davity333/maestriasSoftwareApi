package repositories

import entities "software/src/Domain/Entities"

type Iarea interface {
	GetAllAreas() ([]*entities.Maestria, error)
	GetById(id int) (*entities.Maestria, error)
}