package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getChecksum(filePath string) (string, error) {
	var md5String string

	file, err := os.Open(filePath)
	if err != nil {
		return md5String, err
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return md5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	md5String = hex.EncodeToString(hashInBytes)
	return md5String, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("body.html", responseData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	htmlHash, err := getChecksum("body.html")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("body.md5", []byte(htmlHash), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
