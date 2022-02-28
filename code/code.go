package code

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const FILE_JSON string = "json/code.json"
const PAD_STRING string = "0"
const INITIAL_VALUE int = 0
const ONE int = 1
const THREE int = 3
const FIFTY int = 50
const NINE_NINE_NINE int = 999
const THOUSAND int = 1000
const MIN_RANDOM_NUMBER int = 100000
const MAX_RANDOM_NUMBER int = 400000

type JSONCode struct {
	Code  string
	Total string
}

func FileGetContents(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		log.Print(err)
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	return data, err
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func StrPadLeft(input string, padLength int, padString string) string {
	output := padString

	for padLength > len(output) {
		output += output
	}

	if len(input) >= padLength {
		return input
	}

	return output[:padLength-len(input)] + input
}

func GetCodeThree() string {
	fileByte, _ := FileGetContents(FILE_JSON)
	var decodedJSON []JSONCode
	if _err := json.Unmarshal(fileByte, &decodedJSON); _err != nil {
		log.Print(_err)
	}
	randomIndex := rand.Intn(len(decodedJSON))
	randomTotalCode := decodedJSON[randomIndex].Total
	randomCode := decodedJSON[randomIndex].Code
	totalCode, _ := strconv.Atoi(randomTotalCode)
	rng := INITIAL_VALUE
	if totalCode > FIFTY {
		jumlah := totalCode + FIFTY
		if jumlah > THOUSAND {
			rng = randInt(ONE, NINE_NINE_NINE)
		} else {
			rng = randInt(ONE, jumlah)
		}
	} else {
		rng = randInt(ONE, totalCode)
	}

	rngToStr := strconv.Itoa(rng)
	code := randomCode + "-" + StrPadLeft(rngToStr, THREE, PAD_STRING)

	return code
}

func GetCodeSix() string {
	rng := randInt(MIN_RANDOM_NUMBER, MAX_RANDOM_NUMBER)
	code := strconv.Itoa(rng)

	return code
}

func GetRandomCode(message string) (randomCode string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	code := ""
	if strings.Contains(message, "#kodenuklir3") {
		code = GetCodeThree()
	} else {
		code = GetCodeSix()
	}

	result := "Kode Anda hari ini: \n\n" + code

	return result
}
