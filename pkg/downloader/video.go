package downloader

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/shana0440/niconico-downloader-go/pkg/domain"
)

func DownloadVideo(url, outDir string, session domain.Session) {
	client := session.WithClient(&http.Client{})
	apiData := fetchAPIData(url, client)
	log.Println(apiData)
	// Start heart beat
	// Download video
}

func fetchAPIData(url string, client *http.Client) domain.APIData {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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
