package usecase

import (
	"time"

	"github.com/rodolfosouzamenezes/shipping-calculator-server/internal/freight/entity"
)

type ChangeRouteStatusInput struct {
	ID         string
	StartedAt  entity.CustomTime
	FinishedAt entity.CustomTime
	Event      string
}

type ChangeRouteStatusOutput struct {
	ID         string
	Status     string
	StartedAt  entity.CustomTime
	FinishedAt entity.CustomTime
}

type ChangeRouteStatusUseCase struct {
	Repository entity.RouteRepository
}

func NewChangeRouteStatusUseCase(repository entity.RouteRepository) *ChangeRouteStatusUseCase {
	return &ChangeRouteStatusUseCase{
		Repository: repository,
	}
}

func (u *ChangeRouteStatusUseCase) Execute(input ChangeRouteStatusInput) (*ChangeRouteStatusOutput, error) {
	route, err := u.Repository.FindByID(input.ID)

	if err != nil {
		return nil, err
	}

	if input.Event == "RouteStarted" {
		route.Start(time.Time(input.StartedAt))
	}

	if input.Event == "RouteFinished" {
		route.Finish(time.Time(input.FinishedAt))
	}

	err = u.Repository.Update(route)

	if err != nil {
		return nil, err
	}

	return &ChangeRouteStatusOutput{
		ID:         route.ID,
		Status:     route.Status,
		StartedAt:  entity.CustomTime(route.StartedAt),
		FinishedAt: entity.CustomTime(route.FinishedAt),
	}, nil
}
