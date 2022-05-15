package random_pics

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetCats() (image, preview, errs string) {
	var client http.Client

	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search?mime_type=jpg,png&size=med&sub_id=sub_id&limit=1", nil)
	req.Header.Add("x-api-key", "3edb4670-ca5c-4dcd-bb6a-f660ff1722d7")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		temp := strings.Split(bodyString, `"`)
		return temp[9], temp[9], ""
	} else {
		out := fmt.Sprint(resp.StatusCode)
		return "", "", "cat server error " + out
	}
}
