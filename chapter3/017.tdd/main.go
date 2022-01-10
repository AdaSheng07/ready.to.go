package main

import (
	"math"
	"sort"
)

var (
	personInfo = map[string]float64{}
)

func inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if item < minFatRate {
			minFatRate = item
		}
	}
	personInfo[name] = minFatRate
}

func getRank(name string) (rank int, fatRate float64) {
	fatRate2name := map[float64][]string{}
	fatRateArr := make([]float64, 0, len(personInfo))
	for name, fr := range personInfo {
		fatRate2name[fr] = append(fatRate2name[fr], name)
		fatRateArr = append(fatRateArr, fr)
	}
	sort.Float64s(fatRateArr)
	for i, fr := range fatRateArr {
		_names := fatRate2name[fr]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fatRate = fr
				return
			}
		}
	}
	return 0, 0
}
