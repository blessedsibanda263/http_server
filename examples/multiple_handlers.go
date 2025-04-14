package examples

import "net/http"

func MultipleHandlers() {
	usersRouter := http.NewServeMux()
	commentsRouter := http.NewServeMux()

	usersRouter.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("new user form"))
	})

	commentsRouter.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("new comment form"))
	})

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/users/", http.StripPrefix("/users", usersRouter))
	mainRouter.Handle("/comments/", http.StripPrefix("/comments", commentsRouter))

	if err := http.ListenAndServe(":8000", mainRouter); err != nil {
		panic(err)
	}
}
