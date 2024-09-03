package http

//
//import (
//	"net/http"
//	"time"
//
//	"bitbucket.org/moladinTech/moladin-go-skeleton-service/docs"
//
//	commonMiddleware "bitbucket.org/moladinTech/go-lib-common/middleware/gin"
//	common "bitbucket.org/moladinTech/go-lib-common/registry"
//	commonResponse "bitbucket.org/moladinTech/go-lib-common/response"
//	delivery "bitbucket.org/moladinTech/moladin-go-skeleton-service/delivery/http"
//
//	sentryGin "github.com/getsentry/sentry-go/gin"
//	"github.com/gin-gonic/gin"
//	swaggerFiles "github.com/swaggo/files"
//	ginSwagger "github.com/swaggo/gin-swagger"
//)
//
//type Router interface {
//	Register() *gin.Engine
//	swagger()
//}
//
//type router struct {
//	engine   *gin.Engine
//	common   common.IRegistry
//	delivery delivery.IRegistry
//}
//
//func NewRouter(
//	common common.IRegistry,
//	delivery delivery.IRegistry,
//) Router {
//	return &router{
//		engine:   gin.Default(),
//		common:   common,
//		delivery: delivery,
//	}
//}
//
//// @title          moladin-go-skeleton-service Swagger API
//// @version        1.0
//// @description    moladin-go-skeleton-service Swagger API
//// @termsOfService http://swagger.io/terms/
//
//// @contact.name  API Support
//// @contact.url   http://www.swagger.io/support
//// @contact.email support@swagger.io
//
//// @license.name Apache 2.0
//// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
//
//// @BasePath /
//
//// @securityDefinitions.apikey BearerAuth
//// @in header
//// @name Authorization
//
//func (r *router) Register() *gin.Engine {
//	defer r.common.GetSentry().Flush(2 * time.Second)
//	r.engine.Use(sentryGin.New(sentryGin.Options{
//		Repanic: true,
//	}))
//
//	// Middleware
//	r.engine.Use(
//		commonMiddleware.CORS(),
//		commonMiddleware.RequestID(),
//		r.common.GetTraceMiddleware().Tracer(),
//		r.common.GetPanicRecoveryMiddleware().PanicRecoveryMiddleware(),
//	)
//
//	// handle no-route error (404 not found)
//	commonResponse.RouteNotFound(r.engine)
//
//	// Landing
//	r.engine.GET("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
//	})
//
//	// Health Check
//	r.engine.GET("/health", r.delivery.GetHealth().Check)
//
//	// v1
//	// Configuration
//	r.swagger()
//
//	return r.engine
//}
//
//func (r *router) swagger() {
//	docs.SwaggerInfo.Schemes = []string{"http", "https"}
//	// Route: /docs/index.html
//	r.engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//}
