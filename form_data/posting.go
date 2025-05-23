package formdata

import (
	"fmt"
	"net/http"
	"time"
)

type comment struct {
	username   string
	text       string
	dateString string
}

var comments []comment

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	body := `
		<html>
			<head>
				<!-- style and metadata -->
			</head>
		<body>`
	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("<div class='comment'>%s (%s) - @%s</div>",
			comments[i].text, comments[i].dateString, comments[i].username)
	}
	body += fmt.Sprintf(`
		<h1>Comments</h1>
		%s 
		<form method="POST" action="/comments">
			<div><input type="text" placeholder="Username" name="username" /></div>
			<textarea placeholder="Comment text" name="comment" ></textarea>
			<div><input type="submit" value="Submit" /></div>
		</form>
		</body>
		</html>
	`, commentBody)
	w.Write([]byte(body))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	commentText := r.Form.Get("comment")
	comments = append(comments,
		comment{
			username:   username,
			text:       commentText,
			dateString: time.Now().Format(time.RFC3339)})
	http.Redirect(w, r, "/comments", http.StatusFound)
}

func PostingFormData() {
	http.HandleFunc("GET /comments", commentsHandler)
	http.HandleFunc("POST /comments", postHandler)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}
