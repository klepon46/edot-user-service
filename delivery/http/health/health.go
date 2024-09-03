package health

import (
	"net/http"

	service "bitbucket.org/moladinTech/moladin-go-skeleton-service/service/health"

	common "bitbucket.org/moladinTech/go-lib-common/registry"
	"bitbucket.org/moladinTech/go-lib-common/response"

	"github.com/gin-gonic/gin"
)

type IHealth interface {
	Check(c *gin.Context)
}

type Health struct {
	common common.IRegistry
	health service.IHealth
}

func NewHealth(common common.IRegistry, health service.IHealth) *Health {
	return &Health{
		common: common,
		health: health,
	}
}

// Check Health
// @Summary Health check
// @Schemes
// @Description do health check for databases
// @Tags        check
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response
// @Router      /health [get]
func (h *Health) Check(c *gin.Context) {
	const logCtx = "delivery.http.health.Health.Check"

	var (
		ctx     = c.Request.Context()
		status  = http.StatusOK
		message = http.StatusText(status)
	)

	sentry := h.common.GetSentry()
	span := sentry.StartSpan(ctx, logCtx)
	ctx = sentry.SpanContext(*span)
	defer sentry.Finish(span)

	c.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Data:    h.health.Check(ctx),
		Message: message,
	})
	return
}
