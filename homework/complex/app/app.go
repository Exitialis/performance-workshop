package app

import (
	"context"
	"fmt"
	"github.com/exitialis/workshop/homework/complex/handlers"
	lru2 "github.com/exitialis/workshop/homework/complex/internal/cache"
	"github.com/exitialis/workshop/homework/complex/internal/clients/calltracking"
	phones2 "github.com/exitialis/workshop/homework/complex/internal/clients/phones"
	phones_gateway "github.com/exitialis/workshop/homework/complex/internal/gateway"
	calltracking2 "github.com/exitialis/workshop/homework/complex/internal/service/calltracking"
	"github.com/exitialis/workshop/homework/complex/internal/service/phones"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func StartApp(ctx context.Context, withPprof bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	router := mux.NewRouter()

	phonesClient := phones2.NewPhonesClient()
	phonesService := phones.NewPhonesService(phonesClient)
	calltrackingClient := calltracking.NewCalltrackingClient()
	calltrackingService := calltracking2.NewCalltrackingService(calltrackingClient)
	lru := lru2.NewLruCache()
	phonesGateway := phones_gateway.NewPhonesGateway(phonesService, calltrackingService, lru, ctx)
	if withPprof {
		go ServePProf(ctx)
	}

	router.HandleFunc("/getPhone", handlers.New(phonesGateway).Handle)

	log.Println("start server on :8890")

	srv := http.Server{
		Addr: ":8890",
		Handler: router,
	}

	go func () { _ = srv.ListenAndServe() }()

	<-ctx.Done()
	srv.SetKeepAlivesEnabled(false)
	_ = srv.Shutdown(context.Background())
}

// ServePProf serves pprof endpoints.
func ServePProf(ctx context.Context) {
	srv := http.Server{
		Addr:         ":3366",
		Handler:      pprofHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Minute,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("start pprof on :3366")

	go func() { _ = srv.ListenAndServe() }()

	<-ctx.Done()
	srv.SetKeepAlivesEnabled(false)
	_ = srv.Shutdown(context.Background())
}

func pprofHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	return mux
}


