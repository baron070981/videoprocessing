package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

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
	data, err := ioutil.ReadFile("testfile.mp4")
	if err != nil {
		fmt.Println("Error read file")
		return err
	}
	nd := make([]byte, len(data))
	for _, b := range data {
		if b == '\x00' {
			nd = append(nd, b)
		}
	}

	fmt.Println(string(nd))
	return nil
}

func main() {

	fmt.Println("Start")
	download("https://whatsaper.ru/wp-content/plugins/download-attachments/includes/download.php?id=1630")

}
