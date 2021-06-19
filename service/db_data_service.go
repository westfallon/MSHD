package service

import (
	"github.com/westfallon/naseum/db"
	"log"
	"math/rand"
)

func GetDeathStatistics() (deathStatistics []db.DeathStatistics) {
	err := db.GetDB().Find(&deathStatistics).Error
	if err != nil {
		log.Fatal(err)
	}
	return deathStatistics
}

func GetCivilStructure() (civilStructure []db.CivilStructure) {
	err := db.GetDB().Find(&civilStructure).Error
	if err != nil {
		log.Fatal(err)
	}
	return civilStructure
}

func GetCommDisaster() (commDisasters []db.CommDisaster) {
	err := db.GetDB().Find(&commDisasters).Error
	if err != nil {
		log.Fatal(err)
	}
	return commDisasters
}

func GetCollapseRecord() (collapseRecord []db.CollapseRecord) {
	err := db.GetDB().Find(&collapseRecord).Error
	if err != nil {
		log.Fatal(err)
	}
	return collapseRecord
}

func GetDisasterInfo() (disasterInfo []db.DisasterInfo) {
	err := db.GetDB().Find(&disasterInfo).Error
	if err != nil {
		log.Fatal(err)
	}
	return disasterInfo
}

func GetDisasterRequest() (disasterRequest []db.DisasterRequest) {
	err := db.GetDB().Find(&disasterRequest).Error
	if err != nil {
		log.Fatal(err)
	}
	return disasterRequest
}

func RandSeq(n int) string {
	var letters = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
