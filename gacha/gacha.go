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

func GachaPercentage() (luck string, pity int) {
	result := Gacha(100, 3, 10) // gacha 100x, rate 3% (rate/pool), 10pull/gacha
	if result >= 34 {
		percentage := (math.Floor((((rand.Float64() * (100 - 75)) + 75) * 100))) / 100
		percentageStr := strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
		if percentage >= 90 {
			message := "Laksek! Luck kamu:\n" + percentageStr
			return message, int(percentage)
		} else {
			message := "Ya! Luck kamu:\n" + percentageStr
			return message, int(percentage)
		}
	} else if result < 28 {
		percentage := (math.Floor((rand.Float64() * (44 - 0)) * 100)) / 100
		percentageStr := strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
		if percentage < 10 {
			message := "AMPAS! Luck kamu:\n" + percentageStr
			return message, int(percentage)
		} else {
			message := "Sebaiknya tidak, luck kamu:\n" + strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
			return message, int(percentage)
		}
	} else {
		percentage := (math.Floor(((rand.Float64() * (74 - 45)) + 45) * 100)) / 100
		percentageStr := strconv.FormatFloat(percentage, 'f', -1, 32) + "%"
		message := "Luck kamu biasa saja\n" + percentageStr
		return message, int(percentage)
	}
}

func GachaSim(total_gacha, rate, bulk_draw, luck int) string {
	balancer := 0
	repeat := 0
	if luck < 7 {
		repeat = 2
	} else {
		repeat = luck
	}
	for x := 0; x <= repeat; x++ {
		if balancer <= rate-1 {
			gacha_result := Gacha(total_gacha, rate, bulk_draw)

			if balancer < gacha_result {
				balancer = gacha_result
			}
		}
	}

	message := "SSR yang kamu dapat:\n" + strconv.Itoa(balancer)
	return message
}
