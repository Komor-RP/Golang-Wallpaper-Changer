package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// RedditResponse is Reddit outer response data
type RedditResponse struct {
	MetaData struct {
		Modhash string     `json:"modhash"`
		Dist    int        `json:"dist"`
		Posts   []PostData `json:"children"`
	} `json:"data"`
}

// PostData is Reddit inner data for individual posts
type PostData struct {
	Post struct {
		Title  string `json:"title"`
		Link   string `json:"url"`
		Domain string `json:"domain"`
	} `json:"data"`
}

func main() {

	getRandomImage(getRedditPosts())
	//i.imgur.com, i.redd.it

}

func getRedditPosts() ([]PostData, error) {
	url := "https://www.reddit.com/r/wallpaper/hot/.json?"
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "for-reddit-wallpaper-changer-komorrr")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data := RedditResponse{}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.MetaData.Posts, nil
}

func getRandomImage(posts []PostData, err error) error {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s \n", posts[0].Post)

	return downloadImage(posts[1].Post.Link)

}

func downloadImage(imgURL string) error {
	out, err := os.Create("wallpaper.jpg")
	if err != nil {
		return err
	}
	defer out.Close()

	res, err := http.Get(imgURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
