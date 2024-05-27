package main

import (
	"comments_service/config"
	"comments_service/graph"
	commentusecase "comments_service/internal/commentUseCase"
	secure_access "comments_service/internal/secureAccess"
	"comments_service/internal/storage"
	inmemory "comments_service/internal/storage/inMemory"
	"comments_service/internal/storage/postgres"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	var stor storage.Storage
	viperInst, err := config.LoadConfig()
	if err != nil{
		log.Fatal("Error in LoadConfig. Error:", err)
	}
	conf, err := config.ParseConfig(viperInst)
	if err != nil{
		log.Fatal("Error in ParseConfig. Error:", err)
	}
	if conf.TypeMemory == 0{
		stor, err = postgres.InitDb(conf.Postgres)
		if err != nil{
			log.Fatal(err)
		}
	} else if conf.TypeMemory == 1{
		stor = inmemory.InitMemory()
	}
	
	authorization := secure_access.NewAuthorization(conf.JWTKey)
	cuc := commentusecase.New(stor, authorization, conf.MaxSubs, conf.MaxCommentSize)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Uc : &cuc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.Server.Port)
	log.Fatal(http.ListenAndServe(conf.Server.Host + ":" + conf.Server.Port, nil))
}
