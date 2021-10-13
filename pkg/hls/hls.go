package hls

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/grafov/m3u8"
)

func DownloadVOD(masterURL, videoID, destPath string) {
	tmpDir := "./tmp/" + videoID
	os.MkdirAll(tmpDir, os.ModePerm)
	os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	playlistURL := fetchPlaylist(masterURL)
	segmentURLs := fetchSegmentList(playlistURL)

	segmentPathsCount := len(segmentURLs)
	segmentPaths := make([]string, segmentPathsCount)
	for i, url := range segmentURLs {
		segmentPaths[i] = downloadSegment(url, tmpDir)
		log.Printf("%d/%d\n", i, segmentPathsCount)
	}
	mergedFile := mergeTsFiles(tmpDir, segmentPaths)
	os.Rename(mergedFile, destPath)
}

func fetchPlaylist(masterURL string) string {
	resp, err := http.Get(masterURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Status code error: %s %s", masterURL, resp.Status)
	}
	playlist, listType, err := m3u8.DecodeFrom(resp.Body, true)
	if err != nil {
		log.Fatalln(err)
	}
	if listType != m3u8.MASTER {
		log.Fatalf("The URL should be master URL: %s", masterURL)
	}
	masterPlaylist := playlist.(*m3u8.MasterPlaylist)
	playlistURL := getURL(masterURL, masterPlaylist.Variants[0].URI)
	return playlistURL
}

func fetchSegmentList(playlistURL string) []string {
	resp, err := http.Get(playlistURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Status code error: %s %s", playlistURL, resp.Status)
	}
	playlist, listType, err := m3u8.DecodeFrom(resp.Body, true)
	if listType != m3u8.MEDIA {
		log.Fatalf("The URL should be playlist URL: %s", playlistURL)
	}
	playlistPlaylist := playlist.(*m3u8.MediaPlaylist)
	var segmentURLs []string
	for _, segment := range playlistPlaylist.Segments {
		if segment == nil {
			break
		}
		segmentURLs = append(segmentURLs, getURL(playlistURL, segment.URI))
	}
	return segmentURLs
}

func downloadSegment(segmentURL, tmpDir string) string {
	filename := getFilename(segmentURL)
	segmentPath := tmpDir + "/" + filename
	segmentFile, err := os.Create(segmentPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer segmentFile.Close()
	resp, err := http.Get(segmentURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalln("Status code error: %s %s", segmentURL, resp.Status)
	}
	_, err = io.Copy(segmentFile, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return segmentPath
}

func mergeTsFiles(tmpDir string, segmentPaths []string) string {
	outputPath := tmpDir + "/out.ts"
	out, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()
	for _, path := range segmentPaths {
		func() {
			src, err := os.Open(path)
			if err != nil {
				log.Fatalln(err)
			}
			defer src.Close()
			io.Copy(out, src)
		}()
	}
	if err != nil {
		log.Fatalln(err)
	}
	return outputPath
}

func getURL(oldURL string, targetPath string) string {
	paths := strings.Split(oldURL, "/")
	newURL := strings.Join(paths[:len(paths)-1], "/")
	return newURL + "/" + targetPath
}

func getFilename(url string) string {
	paths := strings.Split(url, "/")
	filenameWithQuery := paths[len(paths)-1]
	filename := strings.Split(filenameWithQuery, "?")[0]
	return filename
}
