package main

import (
	"awesomeProject1/member/domain/service"
	db "awesomeProject1/member/infrastructure"
	memberRepository "awesomeProject1/member/infrastructure/repository"
	memberRouter "awesomeProject1/member/router"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"net/http"
	"time"
)

func main() {
	RunConfigService()
}

func RunConfigService() {

	// ----------Router---------
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Write([]byte("Couchspace is coming"))
	}).Methods(http.MethodGet)
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("PONG" + "0.0.2"))
	}).Methods(http.MethodGet)
	memberRouter.NewMemberRouter(
		service.NewMemberService(
			memberRepository.NewMemberRepository(
				db.NewMongoDb()))).GetApi(router)

	fmt.Println("-----------Background_Service---------")

	//----------CORS---------
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*",
		},
	})

	//router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	srv := &http.Server{
		Handler:      corsOpts.Handler(router),
		Addr:         ":9001",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	srv.ListenAndServe()
}
