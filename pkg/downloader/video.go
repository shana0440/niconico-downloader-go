package downloader

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/shana0440/niconico-downloader-go/pkg/domain"
	"github.com/shana0440/niconico-downloader-go/pkg/domain/api_session"
	"github.com/shana0440/niconico-downloader-go/pkg/domain/heartbeat"
	"github.com/shana0440/niconico-downloader-go/pkg/hls"
)

func DownloadVideo(url, outDir string, session domain.Session) {
	client := session.WithClient(&http.Client{})
	apiData := fetchAPIData(url, client)
	sessionURL := apiData.Media.Delivery.Movie.Session.URLs[0].URL
	sessionData := fetchSessionData(sessionURL, apiData, client)
	canceller := startHeartBeat(sessionURL, sessionData, client)
	// Replace / to - to avoid conflict with path separator
	title := apiData.Video.Title
	filename := strings.ReplaceAll(title, "/", "-") + ".ts"
	destPath := filepath.Join(outDir, filename)
	log.Println("Start downloading " + title + " ...")
	downloadVideo(sessionData, destPath, client)
	canceller <- struct{}{}
	close(canceller)
}

func fetchAPIData(url string, client *http.Client) domain.APIData {
	log.Println("Fetching API data...")
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
	log.Println("Fetched API data")
	return apiData
}

func fetchSessionData(sessionURL string, apiData domain.APIData, client *http.Client) api_session.APISessionBody {
	log.Println("Fetching session data...")
	apiURL := sessionURL + "?_format=json"
	sessionPayload := api_session.MakePayload(apiData)
	payload, err := json.Marshal(sessionPayload)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Post(apiURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode > 299 {
		log.Fatalf("Status code error: %s %s", apiURL, resp.Status)
	}
	defer resp.Body.Close()
	body := api_session.MakeBody(resp)
	log.Println("Fetched session data")
	return body
}

func startHeartBeat(sessionURL string, apiSessionBody api_session.APISessionBody, client *http.Client) chan<- struct{} {
	log.Println("Start heart beat")
	apiURL := sessionURL + "/" + apiSessionBody.Data.Session.ID + "/?_format=json&_method=PUT"
	canceller := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-canceller:
				log.Println("Cancelled heart beat")
				return
			default:
				log.Println("Send heart beat")
				heartBeatPayload := heartbeat.MakePayload(apiSessionBody)
				payload, err := json.Marshal(heartBeatPayload)
				_, err = client.Post(apiURL, "application/json", bytes.NewBuffer(payload))
				if err != nil {
					log.Fatalln(err)
				}
				// Send heart beat faster in case heart beat expired.
				lifetime := apiSessionBody.Data.Session.KeepMethod.Heartbeat.Lifetime / 2
				time.Sleep(time.Duration(lifetime) * time.Millisecond)
			}
		}
	}()
	return canceller
}

func downloadVideo(apiSessionBody api_session.APISessionBody, destPath string, client *http.Client) {
	masterURL := apiSessionBody.Data.Session.ContentURI
	hls.DownloadVOD(masterURL, apiSessionBody.Data.Session.RecipeID, destPath)
}
