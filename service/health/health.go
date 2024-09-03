package health

import (
	"context"

	"bitbucket.org/moladinTech/go-lib-common/logger"
	"bitbucket.org/moladinTech/moladin-go-skeleton-service/model"

	"github.com/jmoiron/sqlx"
)

const (
	OK  = "OK"
	BAD = "BAD"
)

type IHealth interface {
	Check(ctx context.Context) model.HTTPResponse
}

type Health struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

func NewHealth(master, slave *sqlx.DB) *Health {
	return &Health{
		master: master,
		slave:  slave,
	}
}

func (s *Health) Check(ctx context.Context) model.HTTPResponse {
	var response = model.HTTPResponse{
		Master: OK,
		Slave:  OK,
	}

	// Check master connection
	err := s.master.PingContext(ctx)
	if err != nil {
		logger.Error(ctx, "failed ping database master", err)
		response.Master = BAD
	}

	// Check slave connection
	err = s.slave.PingContext(ctx)
	if err != nil {
		logger.Error(ctx, "failed ping database slave", err)
		response.Slave = BAD
	}

	return response
}
