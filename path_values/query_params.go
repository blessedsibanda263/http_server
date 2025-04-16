package path_values

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type comment struct {
	text      string
	username  string
	timestamp time.Time
}

func makeTime(val string) time.Time {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		panic(err)
	}
	return t
}

var comments = []comment{
	{text: "first!", username: "Bill", timestamp: makeTime("2023-09-01T00:00:00Z")},
	{text: "darn, I _just_ missed it", username: "Jill", timestamp: makeTime("2023-09-01T00:00:20Z")},
	{text: "😉 maybe next time", username: "Bill", timestamp: makeTime("2023-09-01T00:01:00Z")},
	{text: "ah, I see I'm late to the show, hello everyone", username: "Phil", timestamp: makeTime("2023-09-01T00:01:05Z")},
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	var result []comment
	params := r.URL.Query()
	if username := params.Get("username"); username != "" {
		filteredComments := []comment{}
		for k := range comments {
			if comments[k].username == username {
				filteredComments = append(filteredComments, comments[k])
			}
		}
		result = filteredComments
	}
	if search := params.Get("search"); search != "" {
		filteredComments := []comment{}
		re := regexp.MustCompile(search)
		for k := range comments {
			if re.MatchString(comments[k].text) {
				filteredComments = append(filteredComments, comments[k])
			}
		}
		result = filteredComments
	}
	commentString := ""
	for k := range result {
		commentString += fmt.Sprintf("%s (%s) @ %s\n", result[k].text, result[k].username, result[k].timestamp.Format("2006-01-02 15:04:05"))
	}
	fmt.Fprint(w, commentString)
}

func QueryParams() {
	http.HandleFunc("/comments", commentsHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
