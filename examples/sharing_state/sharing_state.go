package sharingstate

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var validAgent = regexp.MustCompile(`(?i)(chrome|firefox)`)

func uaMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		log.Println("User Agent: ", userAgent)
		if !validAgent.MatchString(userAgent) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "agent", userAgent)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

func uaStatusHandler(w http.ResponseWriter, r *http.Request) {
	ua := r.Context().Value("agent").(string)
	fmt.Fprintf(w, "%s", fmt.Sprintf("congratulations, you are using: %s", ua))
}

func SharingState() {
	http.HandleFunc("GET /withcontext", uaMiddleware(uaStatusHandler))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic("could not start server")
	}
}
