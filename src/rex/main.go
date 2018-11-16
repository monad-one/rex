package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/appengine" // Required external App Engine library
	"google.golang.org/appengine/urlfetch"
)

const alsGCF = "https://us-central1-devfest18-221623.cloudfunctions.net/recommend"

// const fakeResp = `[
// 	{
// 	  "movie_id": 121029,
// 	  "prediction": 6.3348446,
// 	  "title": "No Distance Left to Run (2010)",
// 	  "youtube_id": "asdfqwerty"
// 	},
// 	{
// 	  "movie_id": 129536,
// 	  "prediction": 6.07897,
// 	  "title": "Code Name Coq Rouge (1989)",
// 	  "youtube_id": "asdfqwerty"
// 	},
// 	{
// 	  "movie_id": 77736,
// 	  "prediction": 5.9684463,
// 	  "title": "Crazy Stone (Fengkuang de shitou) (2006)",
// 	  "youtube_id": "asdfqwerty"
// 	},
// 	{
// 	  "movie_id": 117907,
// 	  "prediction": 5.914757,
// 	  "title": "My Brother Tom (2001)",
// 	  "youtube_id": "asdfqwerty"
// 	},
// 	{
// 	  "movie_id": 112577,
// 	  "prediction": 5.868834,
// 	  "title": "Willie & Phil (1980)",
// 	  "youtube_id": "asdfqwerty"
// 	}
// ]`

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's alive!")
}

// Session contains information about the current active user and/or item being browsed.
type Session struct {
	UserID  int `json:"user_id,omitempty"`
	MovieID int `json:"movie_id,omitempty"`
}

// Movie contains the movie's details including IDs and prediction score.
type Movie struct {
	Title      string  `json:"title"`
	MovieID    int     `json:"movie_id"`
	Prediction float32 `json:"prediction"`
	YoutubeID  string  `json:"youtube_id"`
}

func getSingleParameter(r *http.Request, param string) string {
	params := r.URL.Query()[param]
	if len(params) < 1 {
		log.Printf("request missing parameter: %v\n", param)
		return ""
	}

	if len(params) > 1 {
		log.Printf("too many parameters: %v\nusing first: %v\n", params, params[0])
	}

	return params[0]
}

func list(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(getSingleParameter(r, "user_id"))
	movieID, _ := strconv.Atoi(getSingleParameter(r, "movie_id"))

	sess := Session{UserID: userID, MovieID: movieID}
	js, err := json.Marshal(sess)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	rd := bytes.NewReader(js)
	resp, err := client.Post(alsGCF, "application/json", rd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// body := []byte(fakeResp)

	log.Printf("received body: %v", string(body))

	var movies []Movie
	err = json.Unmarshal(body, &movies)
	if err != nil {
		log.Println(err)
		return
	}

	t, err := template.ParseFiles("list.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, movies)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/list", list)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	appengine.Main()
}
