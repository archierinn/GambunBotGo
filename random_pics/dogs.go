package random_pics

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetDogs() (image, preview, errs string) {
	var client http.Client

	req, err := http.NewRequest("GET", "https://api.thedogapi.com/v1/images/search?has_breeds=true&mime_type=jpg,png&size=med&sub_id=sub_id&limit=1", nil)
	req.Header.Add("x-api-key", "2fd56482-7da8-46ba-adae-8f9f4bce39c0")
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
		temp1 := strings.Split(bodyString, "https")
		temp2 := strings.Split(temp1[1], `"`)
		image = "https" + temp2[0]
		return image, image, ""
	} else {
		out := fmt.Sprint(resp.StatusCode)
		return "", "", "dog server error " + out
	}
}
