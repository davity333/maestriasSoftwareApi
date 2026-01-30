package application

import (
	entities "software/src/Domain/Entities"
	repositories "software/src/Domain/Repositories"
)

type GetByIdUseCase struct {
	usecase repositories.Iarea
}

func NewGetByIdUseCase(app repositories.Iarea)*GetByIdUseCase{
	return &GetByIdUseCase{
		usecase: app,
	}
}

func (g *GetByIdUseCase) ExecuteGetById(id int) (*entities.Maestria, error) {
    return g.usecase.GetById(id)
}
