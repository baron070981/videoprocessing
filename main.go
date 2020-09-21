package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

var videoname string = "testfile.mp4"

func download(url string) error {
	fl, err := os.Create("testfile.mp4")
	if err != nil {
		fmt.Println("Error create file...")
		return err
	}
	defer fl.Close()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return err
	}
	defer resp.Body.Close()
	io.Copy(fl, resp.Body)
	return nil
}

func getvideo(videoname string, timeframe string) error {

	video, err := os.Open(videoname)
	if err != nil {
		fmt.Println("Error read video", err)
		return err
	}
	defer video.Close()
	videobuff := bufio.NewReader(video)
	fmt.Println(videobuff)

	return nil
}

func main() {

	fmt.Println("Start")
	download("https://whatsaper.ru/wp-content/plugins/download-attachments/includes/download.php?id=1630")
	err := getvideo(videoname, "12:11")
	if err != nil {
		fmt.Println("Error from getvideo\n", err)
	}
}
