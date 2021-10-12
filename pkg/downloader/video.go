package downloader

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/shana0440/niconico-downloader-go/pkg/domain"
	"github.com/shana0440/niconico-downloader-go/pkg/domain/api_session"
)

func DownloadVideo(url, outDir string, session domain.Session) {
	client := session.WithClient(&http.Client{})
	log.Println("Fetching API data...")
	apiData := fetchAPIData(url, client)
	log.Println("Fetched API data")
	log.Println("Fetching session data...")
	sessionData := fetchSessionData(apiData, client)
	log.Println(sessionData)
	log.Println("Fetched session data")
	// Start heart beat
	// Download video
	// Cancel heart beat
}

func fetchAPIData(url string, client *http.Client) domain.APIData {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Status code error: %s %s", url, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	initialData := doc.Find("#js-initial-watch-data")
	if initialData.Size() == 0 {
		log.Fatalln("Can't found #js-inital-watch-data on page")
	}
	rawAPIData, exists := initialData.Attr("data-api-data")
	if !exists {
		log.Fatalln("Can't found data-api-data on js-inital-watch-data")
	}
	apiData := domain.MakeAPIData(rawAPIData)
	return apiData
}

func fetchSessionData(apiData domain.APIData, client *http.Client) api_session.APISessionBody {
	url := apiData.Media.Delivery.Movie.Session.URLs[0].URL
	sessionPayload := api_session.MakePayload(apiData)
	payload, err := json.Marshal(sessionPayload)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Post(url+"?_format=json", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode > 299 {
		log.Fatalf("Status code error: %s %s", url, resp.Status)
	}
	defer resp.Body.Close()
	body := api_session.MakeBody(resp)
	return body
}

