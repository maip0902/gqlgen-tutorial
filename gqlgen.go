package gqlgen_todos

import (
"database/sql"
"log"
"net/http"
	"os"

	"./database"

_ "github.com/go-sql-driver/mysql"
"github.com/gorilla/context"
"github.com/gorilla/csrf"
)

// Server is whole server implementation for this wiki app.
// This holds database connection and router settings.
type Server struct {
	db      *sql.DB
	handler http.Handler
}

// Close makes the database connection to close.
func (s *Server) Close() error {
	return s.db.Close()
}

// Init initialize server state. Connecting to database, compiling templates,
// and settings router.
func (s *Server) Init(dbconf, env string, debug bool) {
	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	db, err := cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}

	s.db = db
	s.Route()
}

// New returns server object.
func New() *Server {
	return &Server{}
}

// csrfProtectKey should have 32 byte length.
var csrfProtectKey = []byte("32-byte-long-auth-key")

// Run starts running http server.
func (s *Server) Run(addr string) {
	log.Printf("start listening on %s", addr)

	// NOTE: when you serve on TLS, make csrf.Secure(true)
	CSRF := csrf.Protect(
		csrfProtectKey, csrf.Secure(false))
	http.ListenAndServe(addr, context.ClearHandler(CSRF(s.handler)))
}

const defaultPort = "8080"

func Route() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen_todos.NewExecutableSchema(gqlgen_todos.Config{Resolvers: &gqlgen_todos.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
