package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type row struct {
	CountryOrTerritory string `json:"country"`
	ReportingYear string `json:"report year"`
	Namespace string `json:"namespace"`
	ZoneCode string `json:"zone code"`
	ZoneId string `json:"zone id"`
	GeographicalName string `json:"geographical name"`
	BeginTime string `json:"begin time"`
	Website string `json:"wesite"`
	AQDZoneType string `json:"AQDzone type"`
	Pollutant string `json:"pollutat"`
	envelope string `json:"envelope"`
	ResidentPopulation string `json:"resident population"`
	ResidentPopulationYear string `json:"resident population year"`
	Area string `json:"area"`
	TimeExtensionExemption string `json:"time extension exemption"`
	CompetentAuthority string `json:"competent authority"`
	Telephone string `json:"telephone"`
	Address string `json:"address"`
	ProtectionTarget string `json:"protection target"`
	EndTime string `json:"end time"`
}

type CsvLine struct {
	Column1 string
	Column2 string
}

func readCSV() []byte {
	csvFile, _ := os.Open("/home/shiun/Documents/Masterarbeit/Data/FinalForm.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var rows []row
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Println(line)
		/*
		rows = append(rows, row{
			CountryOrTerritory: line[0],
			ReportingYear:  line[1],
			Namespace:  line[2],
			ZoneCode:  line[3],
			ZoneId:  line[4],
			GeographicalName:  line[5],
			BeginTime:  line[6],
			Website:  line[7],
			AQDZoneType:  line[8],
			Pollutant:  line[9],
			envelope:  line[10],
			ResidentPopulation:  line[11],
			ResidentPopulationYear:  line[12],
			Area:  line[13],
			TimeExtensionExemption:  line[14],
			CompetentAuthority:  line[15],
			Telephone:  line[16],
			Address:  line[17],
			ProtectionTarget:  line[18],
			EndTime:  line[19],
		})
		*/

	}
	rowsJson, _ := json.Marshal(rows)
	//fmt.Println(string(rowsJson))

	return rowsJson
}

func readCSV2() {
	//filename := "{{ ENTER FILE }}"

	// Open CSV file
	f, err := os.Open("/home/shiun/Documents/Masterarbeit/Data/FinalForm.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
		}
		fmt.Println(data.Column1 + " " + data.Column2)
	}
}