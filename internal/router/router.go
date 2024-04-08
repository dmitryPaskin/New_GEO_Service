package router

import (
	"GeoServiseAppDate/internal/controller/authorizationHandlers"
	"GeoServiseAppDate/internal/controller/responder"
	"GeoServiseAppDate/internal/controller/searchGEOHandlers"
	"GeoServiseAppDate/internal/middleware/authMiddleware"
	"GeoServiseAppDate/internal/repository"
	"GeoServiseAppDate/internal/repository/authRepository"
	"GeoServiseAppDate/internal/service"
	"GeoServiseAppDate/internal/service/authService"
	"GeoServiseAppDate/pkg/database"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	chi          *chi.Mux
	authHandlers authorizationHandlers.HandlerAuth
	GEOHandler   searchGEOHandlers.Handler
}

func New() Router {
	var router Router
	router.chi = chi.NewRouter()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
		return Router{}
	}

	log := zap.NewExample()

	router.authHandlers = initAuthHandlers(db, log)
	router.GEOHandler = initGEOHandlers(db, log)

	return router
}

func (r *Router) Start() {
	tokenAuth := jwtauth.New("HS256", []byte("mySecret"), nil)

	r.chi.Use(middleware.Recoverer)
	r.chi.Use(middleware.Logger)

	r.chi.Post("/api/register", r.authHandlers.SingUpHandler)
	r.chi.Post("/api/login", r.authHandlers.SingInHandler)

	r.chi.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(tokenAuth))
		router.Use(authMiddleware.Authenticator)

		router.Post("/api/address/search", r.GEOHandler.SearchAddressHandler)
		router.Post("/api/address/geocode", r.GEOHandler.GeocodeHandler)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r.chi,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigChan
	stopCTX, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(stopCTX); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}

func initAuthHandlers(db database.Database, log *zap.Logger) authorizationHandlers.HandlerAuth {
	repository := authRepository.New(db.DB)
	proxyService := authService.NewAuthServiceProxy(repository)

	responder := responder.NewRespond(log)
	return authorizationHandlers.New(proxyService, responder)
}

func initGEOHandlers(db database.Database, log *zap.Logger) searchGEOHandlers.Handler {
	repository := repository.New(db.DB)
	proxyService := service.NewServiceProxy(service.NewService(&http.Client{}), repository)

	responder := responder.NewRespond(log)
	return searchGEOHandlers.New(proxyService, responder)
}
