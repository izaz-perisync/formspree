package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	global "github.com/perisynctechnologies/formSpree"
	mailmon "github.com/perisynctechnologies/formSpree/mail"
	"github.com/perisynctechnologies/formSpree/server/handler"
	"github.com/perisynctechnologies/formSpree/server/router"
	"github.com/perisynctechnologies/formSpree/service"
	wrapper "github.com/perisynctechnologies/formSpree/utils"
	"github.com/urfave/negroni"
)

func main() {
	log.Println("service started")
	conf := global.GlobalConfig()

	db, err := sql.Open("postgres", conf.DB)
	if err != nil {
		log.Println("connect2 err", err)
		return
	}

	sm := mailmon.New(conf.Mail.User, conf.Mail.Secret, conf.Mail.Host, conf.Mail.Port)
	mail := wrapper.New(conf.Mail.TemplatePath, sm)

	s := service.New(conf.JWT.Key, db, mail)
	h := handler.New(s)
	r := router.BuildRoute(h)
	n := negroni.Classic()
	n.UseHandler(r)
	server := http.Server{
		Addr: fmt.Sprintf(":%d", conf.Server.Port),
		Handler: handlers.CORS(
			handlers.ExposedHeaders(conf.Server.ExposedHeaders),
			handlers.AllowedHeaders(conf.Server.AllowedHeaders),
			handlers.AllowedMethods(conf.Server.AllowedMethods),
			handlers.AllowedOrigins(conf.Server.AllowedOrigins),
		)(n),
	}

	go func() {
		// Start listening.
		log.Println("listening at port", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Handle graceful shutdown.
	lock := make(chan os.Signal, 1)
	signal.Notify(lock, os.Interrupt, syscall.SIGTERM)
	<-lock

	server.Shutdown(context.TODO())

}
