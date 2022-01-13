package gacha

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

const pool = 100

func Gacha(total_gacha, rate, bulk_draw int) int {
	rand.Seed(int64(time.Now().Nanosecond()))

	shuffle := rand.Perm(pool)
	ssr := make([]int, rate)
	for i := range ssr {
		ssr[i] = shuffle[i]
	}

	lucky_hit := 0
	for x := 0; x < total_gacha; x++ {
		pull_result := make([]int, bulk_draw)
		for i := range pull_result {
			pull_result[i] = rand.Intn(pool)
		}

		for y := range ssr {
			for z := range pull_result {
				if ssr[y] == pull_result[z] {
					lucky_hit += 1
				}
			}
		}
	}

	return int(lucky_hit)
}

func GachaPercentage() string {
	result := Gacha(100, 3, 10) //gacha 100x, rate 3% (rate/pool), 10pull/gacha
	if result >= 33 {
		percentage := (math.Floor((((rand.Float64() * (100 - 75)) + 75) * 100))) / 100
		percentageStr := strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
		if percentage >= 90 {
			message := "Laksek! Luck kamu:\n" + percentageStr
			return message
		} else {
			message := "Ya! Luck kamu:\n" + percentageStr
			return message
		}
	} else if result < 28 {
		percentage := (math.Floor(((rand.Float64() * (74 - 45)) + 45) * 100)) / 100
		percentageStr := strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
		message := "Biasa saja, luck kamu:\n" + percentageStr
		return message
	} else {
		percentage := (math.Floor((rand.Float64() * (44 - 0)) * 100)) / 100
		percentageStr := strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
		if percentage < 10 {
			message := "AMPAS! Luck kamu:\n" + percentageStr
			return message
		} else {
			message := "Sebaiknya tidak, luck kamu:\n" + strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
			return message
		}
	}
}

func GachaSim(total_gacha, rate, bulk_draw int) string {
	balancer := 0
	for x := 0; x <= 3; x++ {
		if balancer <= rate-1 {
			gacha_result := Gacha(total_gacha, rate, bulk_draw)
			balancer = gacha_result
		}
	}

	// message := "Jumlah Rarity Tertinggi yang kamu dapat:\n" + strconv.Itoa(balancer)
	message := "Draw: " + strconv.Itoa(total_gacha) + " Rate:" + strconv.Itoa(rate)
	return message
}
