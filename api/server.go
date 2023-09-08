package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/tapansaikia/slot-book-server/storage"
	"github.com/urfave/negroni"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

type HandlerContext struct {
	store storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

// NewHandlerContext constructs a new HandlerContext,
// ensuring that the dependencies are valid values
func NewHandlerContext(store storage.Storage) *HandlerContext {
	if store == nil {
		panic("nil MongoDB session!")
	}
	return &HandlerContext{store}
}

func (s *Server) Start() error {
	// Default values
	err := envconfig.Process("SOCKETCAM", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Debug {
		log.Printf("==> SCHEME: %v", config.Scheme)
		log.Printf("==> ADDRESS: %v", config.ListenAddress)
		log.Printf("==> PRIVATEKEY: %v", config.PrivateKey)
		log.Printf("==> CERTIFICATE: %v", config.Certificate)
	}

	//construct the handler context
	hctx := NewHandlerContext(s.store)

	router := newRouter(hctx)
	n := negroni.Classic()

	n.UseHandler(router)
	if config.Scheme == "https" {
		log.Fatal(http.ListenAndServeTLS(config.ListenAddress, config.Certificate, config.PrivateKey, n))

	} else {
		log.Fatal(http.ListenAndServe(config.ListenAddress, n))

	}
	return nil
}

// NewRouter is the constructor for all my routes
func newRouter(hctx *HandlerContext) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/ws").
		Name("Communication Channel").
		HandlerFunc(hctx.serveWs)

	router.
		Methods("POST").
		Path("/book").
		Name("Book Slot").
		HandlerFunc(hctx.bookSlot)
	return router
}
