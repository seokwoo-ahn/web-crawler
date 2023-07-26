package config

import (
	"os"

	toml "github.com/pelletier/go-toml"
)

type Configs struct {
	DptDay     string
	DptStation string
	ArvStation string
	StationMap map[string]string
}

func NewConfig(file string) *Configs {
	c := new(Configs)
	c.StationMap = SetStationMap()

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer f.Close()
		if err := toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}

func SetStationMap() map[string]string {
	stationMap := make(map[string]string)

	stationMap["수서"] = "0551"
	stationMap["동탄"] = "0552"
	stationMap["평택지제"] = "0553"
	stationMap["천안아산"] = "0502"
	stationMap["오송"] = "0297"
	stationMap["대전"] = "0010"
	stationMap["김천(구미)"] = "0507"
	stationMap["서대구"] = "0506"
	stationMap["동대구"] = "0015"
	stationMap["경주"] = "0508"
	stationMap["울산"] = "0509"
	stationMap["부산"] = "0020"
	stationMap["공주"] = "0514"
	stationMap["익산"] = "0030"
	stationMap["정읍"] = "0033"
	stationMap["광주송정"] = "0036"
	stationMap["나주"] = "0037"
	stationMap["목포"] = "0041"

	return stationMap
}
