package http

//
//import (
//	"context"
//	"fmt"
//	"github.com/klepon46/edot-user-service/config"
//	"net/http"
//	//"bitbucket.org/moladinTech/moladin-go-skeleton-service/config"
//	//delivery "bitbucket.org/moladinTech/moladin-go-skeleton-service/delivery/http"
//	//"bitbucket.org/moladinTech/go-lib-common/logger"
//	//common "bitbucket.org/moladinTech/go-lib-common/registry"
//)
//
//type IServer interface {
//	Serve(ctx context.Context)
//}
//
//type Server struct {
//	router Router
//}
//
//func NewServer() *Server {
//	return &Server{
//		router: NewRouter(
//			nil,
//			nil,
//		),
//	}
//}
//
//func (s *Server) Serve(ctx context.Context) {
//	srv := &http.Server{
//		Addr:    fmt.Sprintf(":%d", config.Cold.AppPort),
//		Handler: s.router.Register(),
//	}
//
//	go func() {
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			panic(err)
//		}
//	}()
//
//	//// Wait for interrupt signal to gracefully shut down the server with
//	//// a timeout of config.Hot.ShutDownDelayInSecond seconds.
//	//quit := make(chan os.Signal)
//	//// kill (no param) default send syscanll.SIGTERM
//	//// kill -2 is syscall.SIGINT
//	//// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
//	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	//<-quit
//	//logger.Info(ctx, fmt.Sprintf("Shutdown Server in %d seconds", config.Hot.ShutDownDelayInSecond))
//	//
//	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Hot.ShutDownDelayInSecond)*time.Second)
//	//defer cancel()
//	//if err := srv.Shutdown(ctx); err != nil {
//	//	logger.Fatal(ctx, "Server Shutdown:", logger.Tag{
//	//		Key:   "error",
//	//		Value: err.Error(),
//	//	})
//	//}
//	//// catching ctx.Done(). timeout of config.Hot.ShutDownDelayInSecond seconds.
//	//select {
//	//case <-ctx.Done():
//	//	log.Println(fmt.Sprintf("timeout of %d seconds.", config.Hot.ShutDownDelayInSecond))
//	//}
//	//log.Println("Server exiting")
//}
