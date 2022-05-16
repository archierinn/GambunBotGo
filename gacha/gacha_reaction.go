package gacha

import (
	"math/rand"
)

func HappyReaction() (pkg, sticker string) {
	var stickerArr []string
	var pickSticker int

	pkgHappy := []string{"446", "789", "6370", "8522", "11537"}
	caseHappy := rand.Intn(len(pkgHappy) - 1)

	switch caseHappy {
	case 0:
		stickerArr = []string{"1989", "1993"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 1:
		stickerArr = []string{"10859"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 2:
		stickerArr = []string{"11088016", "11088036"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 3:
		stickerArr = []string{"16581266", "16581269", "16581271", "16581289"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 4:
		stickerArr = []string{"52002734", "52002735"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	}

	return pkgHappy[caseHappy], stickerArr[pickSticker]
}

func SadReaction() (pkg, sticker string) {
	var stickerArr []string
	var pickSticker int

	pkgSad := []string{"446", "789", "11537", "11538"}
	caseSad := rand.Intn(len(pkgSad) - 1)

	switch caseSad {
	case 0:
		stickerArr = []string{"2008", "2022"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 1:
		stickerArr = []string{"10860", "10879"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 2:
		stickerArr = []string{"52002751", "52002763"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	case 3:
		stickerArr = []string{"51626504", "51626526"}
		pickSticker = rand.Intn(len(stickerArr) - 1)
	}

	return pkgSad[caseSad], stickerArr[pickSticker]
}
