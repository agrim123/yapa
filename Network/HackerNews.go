package Network

import (
	"encoding/json"
	"fmt"
	_ "github.com/bclicn/color"
	"io/ioutil"
	_ "log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

type story struct {
	Author string `json:"by"`
	ID     int    `json:"id"`
	Score  int    `json:"score"`
	Time   int64  `json:"time"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	URL    string `json:"url"`
}

type stories []*story

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func HackerNews() {
	fmt.Println("Fetching Hacker News articles...")

	// Get ids of top stories
	rs, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}

	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var keys []int
	json.Unmarshal(bodyBytes, &keys)

	getTop10StoriesIds := keys[0:10]

	var stories []*story

	// Todo: Use concurrency
	for _, key := range getTop10StoriesIds {
		keyURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", key)
		res, err := http.Get(keyURL)
		// Process response
		if err != nil {
			panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}

		s := &story{}
		err = json.Unmarshal(body, s)
		if err != nil {
			return
		}

		stories = append(stories, s)
	}

	for k, v := range stories {
		fmt.Println(k, v.Title, v.URL, v.Score, time.Unix(v.Time, 0), v.Author)
	}

	var i int

	fmt.Printf("Which one to read? ")
	fmt.Scanf("%d", &i)

	open(stories[i].URL)
}
