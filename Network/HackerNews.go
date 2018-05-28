package Network

import (
	"encoding/json"
	"fmt"
	"github.com/bclicn/color"
	"io/ioutil"
	"log"
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

func GetArticleDetails(id int, ch chan<- *story) {
	keyURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	res, err := http.Get(keyURL)

	if err != nil {
		log.Fatal(err)
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

	ch <- s
}

func HackerNews() {
	fmt.Println("Fetching Hacker News articles...")

	// Get ids of top stories
	rs, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatal(err)
	}

	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var keys []int
	json.Unmarshal(bodyBytes, &keys)

	getTop10StoriesIds := keys[0:10]

	ch := make(chan *story)

	for _, key := range getTop10StoriesIds {
		go GetArticleDetails(key, ch)
	}

	var urls []string
	j := 0
	for range getTop10StoriesIds {
		i := <-ch
		fmt.Println(j, ")", color.Blue(i.Title), i.URL, i.Score, time.Unix(i.Time, 0))
		urls = append(urls, i.URL)
		j += 1
	}

	var k int
	fmt.Printf("Which one to read? ")
	fmt.Scanf("%d", &k)

	open(urls[k])
}
