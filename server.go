package main

import (
	"comments_service/config"
	"comments_service/graph"
	commentusecase "comments_service/internal/commentUseCase"
	secure_access "comments_service/internal/secureAccess"
	"comments_service/internal/storage/postgres"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	viperInst, err := config.LoadConfig()
	if err != nil{
		log.Fatal("Error in LoadConfig. Error:", err)
	}
	conf, err := config.ParseConfig(viperInst)
	if err != nil{
		log.Fatal("Error in ParseConfig. Error:", err)
	}
	pg, err := postgres.InitDb(conf.Postgres)
	if err != nil{
		log.Fatal(err)
	}
	defer pg.Db.Close()
	
	authorization := secure_access.NewAuthorization(conf.JWTKey)
	cuc := commentusecase.New(&pg, authorization, conf.MaxSubs)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Uc : &cuc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.Server.Port)
	log.Fatal(http.ListenAndServe(conf.Server.Host + ":" + conf.Server.Port, nil))
}
