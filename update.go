package update

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
}

func Check(url string) (bool, func() error, error) {
	latest, err := wget(url)
	if err != nil {
		return false, nil, err
	}

	latestSha := sha256.Sum256(latest)

	this, err := os.Executable()
	if err != nil {
		return false, nil, err
	}

	f, err := os.Open(this)
	if err != nil {
		return false, nil, err
	}
	defer f.Close()

	thisB, err := ioutil.ReadAll(f)
	thisSha := sha256.Sum256(thisB)

	return thisSha != latestSha, func() error {
		err := ioutil.WriteFile(this, latest, 0755)
		if err != nil {
			return err
		}
		return nil
	}, nil
}

func sha(in []byte) string {
	sum := sha256.Sum256(in)
	return fmt.Sprintf("%x", sum)
}

func wget(url string) ([]byte, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
