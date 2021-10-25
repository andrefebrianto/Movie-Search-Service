package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/database/mysql"
	"github.com/andrefebrianto/Search-Movie-Service/externalservice/omdbapi"
	moviegrpchandler "github.com/andrefebrianto/Search-Movie-Service/movie/delivery/grpc"
	moviehttphandler "github.com/andrefebrianto/Search-Movie-Service/movie/delivery/http"
	movieusecase "github.com/andrefebrianto/Search-Movie-Service/movie/usecase"
	searchlogcommand "github.com/andrefebrianto/Search-Movie-Service/searchlog/repository/command"
	searchlogusecase "github.com/andrefebrianto/Search-Movie-Service/searchlog/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var GlobalConfig = viper.New()

func setCorsHeader(nextHandler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return nextHandler(context)
	}
}

func responsePing(requestContext echo.Context) error {
	return requestContext.String(http.StatusOK, "Service is running properly")
}

func init() {
	GlobalConfig.SetConfigFile(`config.json`)
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize database
	mysql.SetupConnection()

	// gRPC instance
	grpcServer := grpc.NewServer()

	// Echo instance
	httpServer := echo.New()

	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(setCorsHeader)

	// Routes
	httpServer.GET("/", responsePing)

	timeoutInSecond := time.Duration(GlobalConfig.GetInt("context.timeout") * int(time.Second))
	searchLogMySqlCommandRepository := searchlogcommand.CreateMySqlCommandRepository(mysql.GetConnection())
	searchLogUseCase := searchlogusecase.CreateSearchLogUseCase(searchLogMySqlCommandRepository, timeoutInSecond)

	var movieDataProvider omdbapi.MovieDataProvider
	err := GlobalConfig.UnmarshalKey("omdbApiConfig", &movieDataProvider)
	if err != nil {
		panic(err)
	}
	movieDataProvider.SearchLog = searchLogUseCase

	movieUseCase := movieusecase.CreateMovieUseCase(movieDataProvider, timeoutInSecond)

	// Init handler
	moviehttphandler.HandleHttpRequest(httpServer, movieUseCase)
	moviegrpchandler.HandleGrpcRequest(grpcServer, movieUseCase)

	g := new(errgroup.Group)
	g.Go(func() error { return httpServer.Start(GlobalConfig.GetString("httpServer.port")) })
	g.Go(func() error {
		listen, err := net.Listen("tcp", GlobalConfig.GetString("grpcServer.port"))
		if err != nil {
			log.Fatalln(err)
		}
		return grpcServer.Serve(listen)
	})

	g.Wait()
}
