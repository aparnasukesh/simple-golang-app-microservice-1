package di

import (
	"log"
	"microservice-one/config"
	"microservice-one/internals/app/user"
	grpcclient "microservice-one/pkg/grpcClient"
	"microservice-one/pkg/redis"
	"microservice-one/pkg/sql"
)

func InitResources(cfg config.Config) (*user.Handler, error) {

	// Db initialization
	db, err := sql.NewSql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Redis initialization
	redisClient, err := redis.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	serviceTwoClient, err := grpcclient.NewMicroServiceTwoServiceClient(cfg.MicroserviceTwoPort)
	if err != nil {
		return nil, err
	}
	// User Module initialization
	repo := user.NewRepository(db)
	service := user.NewService(repo, redisClient, serviceTwoClient)
	userHandler := user.NewHttpHandler(service)

	return userHandler, nil
}
