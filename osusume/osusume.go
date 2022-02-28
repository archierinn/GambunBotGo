package osusume

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const INITIAL_VALUE int = 0
const ONE int = 1
const MAX_NUMBER_MAL_ANIME int = 50000
const MAX_NUMBER_ANIDB_ANIME int = 20000
const MAX_NUMBER_ANN_ANIME int = 30000
const MIN_NUMBER_MAL_MANGA int = 115000
const MAX_NUMBER_MAL_MANGA int = 135000
const MAX_NUMBER_ANN_MANGA int = 30000
const MAX_NUMBER_VNDB_VN int = 35000
const BASE_URL_MAL string = "https://myanimelist.net/"
const BASE_URL_ANIDB string = "https://anidb.net/"
const BASE_URL_ANN string = "https://www.animenewsnetwork.com/encyclopedia/"
const BASE_URL_VNDB string = "https://vndb.org/"

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func ArrayRand(elements []string) int {
	randomIndex := INITIAL_VALUE
	if len(elements) > ONE {
		randomIndex = rand.Intn(len(elements))
	}
	return randomIndex
}

func GetOsusumeAnime() string {
	sourceAnime := []string{"mal", "anidb", "ann"}
	randomIndex := ArrayRand(sourceAnime)
	rng := INITIAL_VALUE
	code := ""
	if sourceAnime[randomIndex] == "mal" {
		rng = randInt(ONE, MAX_NUMBER_MAL_ANIME)
		rngToStr := strconv.Itoa(rng)
		code = BASE_URL_MAL + "anime/" + rngToStr
	} else if sourceAnime[randomIndex] == "anidb" {
		rng = randInt(ONE, MAX_NUMBER_ANIDB_ANIME)
		rngToStr := strconv.Itoa(rng)
		code = BASE_URL_ANIDB + "anime/" + rngToStr
	} else {
		rng = randInt(ONE, MAX_NUMBER_ANN_ANIME)
		rngToStr := strconv.Itoa(rng)
		code = BASE_URL_ANN + "anime.php?id=" + rngToStr
	}
	return code
}

func GetOsusumeManga() string {
	sourceManga := []string{"mal", "ann"}
	randomIndex := ArrayRand(sourceManga)
	rng := INITIAL_VALUE
	code := ""
	if sourceManga[randomIndex] == "mal" {
		rng = randInt(MIN_NUMBER_MAL_MANGA, MAX_NUMBER_MAL_MANGA)
		rngToStr := strconv.Itoa(rng)
		code = BASE_URL_MAL + "manga/" + rngToStr
	} else {
		rng = randInt(ONE, MAX_NUMBER_ANN_MANGA)
		rngToStr := strconv.Itoa(rng)
		code = BASE_URL_ANN + "manga.php?id=" + rngToStr
	}
	return code
}

func GetOsusumeVN() string {
	rng := randInt(ONE, MAX_NUMBER_VNDB_VN)
	rngToStr := strconv.Itoa(rng)
	code := BASE_URL_VNDB + "v" + rngToStr
	return code
}

func GetRandomOsusume(message string) (randomCode string) {
	rand.Seed(int64(time.Now().Nanosecond()))
	code := ""
	osusumeType := ""
	if strings.Contains(message, "#osusumeanime") {
		osusumeType = "anime"
		code = GetOsusumeAnime()
	} else if strings.Contains(message, "#osusumemanga") {
		osusumeType = "manga"
		code = GetOsusumeManga()
	} else if strings.Contains(message, "#osusumevn") {
		osusumeType = "vn"
		code = GetOsusumeVN()
	}

	result := "Rekomendasi " + osusumeType + " hari ini: \n\n" + code
	return result
}
