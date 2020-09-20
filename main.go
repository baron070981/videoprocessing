package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)


// сервер с серверным приложением принимает от клиента видеофайл
// и время. Возвращает изображение кадра соответствуещего времени.

// клиент переходит по ссылке и получает видео и время. Обрабатывает.
// Отправляет обратно кадр соответствующий времени.
// 


func download(url string) (error){
	fl,err := os.Create("testfile.mp4")
	if err != nil{
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




func main(){
	
	fmt.Println("Start")
	download("https://whatsaper.ru/wp-content/plugins/download-attachments/includes/download.php?id=1630")
	
	
}





