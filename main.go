package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mhristof/germ/log"
)

func main() {
}

func Check(url string) (bool, error) {
	latest, err := wget(url)
	if err != nil {
		return false, err
	}

	latestSha := sha256.Sum256(latest)

	this, err := os.Executable()
	if err != nil {
		return false, err
	}

	f, err := os.Open(this)
	if err != nil {
		return false, err
	}
	defer f.Close()

	thisB, err := ioutil.ReadAll(f)
	thisSha := sha256.Sum256(thisB)

	return thisSha != latestSha, nil
	// fmt.Println("Updating to new version")
	// err := ioutil.WriteFile(this, latest, 0755)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"err":  err,
	// 		"this": this,
	// 	}).Panic("Cannot write file")
	// }

	//}
}

func sha(in []byte) string {
	sum := sha256.Sum256(in)
	return fmt.Sprintf("%x", sum)
}

func wget(url string) ([]byte, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"url": url,
		}).Panic("Cannot download url")

	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
