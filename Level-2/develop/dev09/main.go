package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	bufSize = 1024 * 8
)

// Записываем результаты запроса GET в файл. Если для fileName задана пустая строка,
// то  последний фрагмент входного URL-адреса используется как имя файла.
func Wget(url, fileName string) {
	resp, err := getResponse(url)
	if err != nil {
		log.Fatal(err)
	}
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}
	if err := writeToFile(fileName, resp); err != nil {
		log.Fatal(err)
	}
}

//GET запрос на URL-адрес, возвращает ответ!
func getResponse(url string) (*http.Response, error) {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func writeToFile(fileName string, resp *http.Response) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriterSize(file, bufSize)

	_, err = io.Copy(bufferedWriter, resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {
	if len(os.Args) == 3 {
		Wget(os.Args[1], os.Args[2])
	} else if len(os.Args) == 2 {
		Wget(os.Args[1], "")
	}
}
