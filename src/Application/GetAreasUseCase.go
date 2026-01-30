package application

import (
	entities "software/src/Domain/Entities"
	repositories "software/src/Domain/Repositories"
)

type GetAreasUseCase struct {
	usecase repositories.Iarea
}

func NewGetAreasUseCase(app repositories.Iarea)*GetAreasUseCase{
	return &GetAreasUseCase{
		usecase: app,
	}
}

func(g *GetAreasUseCase)ExecuteGetAll()([]*entities.Maestria, error){
	area, err := g.usecase.GetAllAreas()
	if err != nil{
		return nil,err
	}

	if area == nil {
		return []*entities.Maestria{}, nil
	}

	return area, nil
}
