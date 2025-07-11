package server

import (
	"encoding/json"
	"log"
	"net/http"

	db "diffinlist/internal/db/generated"
	"diffinlist/internal/middleware"

	"diffinlist/web"
	"diffinlist/web/views/pages"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes(queries *db.Queries) http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/aaa", s.HelloWorldHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/login", middleware.WithOptionalUser(queries,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pageData := getPageAuthInfo(r)

			component := pages.Auth(pageData)
			component.Render(r.Context(), w)

		}),
	))
	mux.Handle("/sign-up", middleware.WithOptionalUser(queries,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pageData := getPageAuthInfo(r)
			component := pages.Signup(pageData)
			component.Render(r.Context(), w)

		}),
	))
	setupAuthRoutes(mux, queries)

	// mux.Handle("/preferences", templ.Handler(pages.()))

	mux.Handle("/home", middleware.WithOptionalUser(queries,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pageData := getPageAuthInfo(r)

			component := pages.Home(pageData)
			component.Render(r.Context(), w)
		}),
	))
	mux.Handle("/test", templ.Handler(pages.TestPage()))
	mux.Handle("/nav", templ.Handler(pages.Nav()))

	mux.Handle("/{$}", middleware.WithOptionalUser(queries,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pageData := getPageAuthInfo(r)
			component := pages.Index(pageData)
			component.Render(r.Context(), w)
		}),
	))

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
