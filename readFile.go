package tmp

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type csvFile struct {
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

func readCSV() {
	csvFile, _ := os.Open("/home/shiun/Documents/Masterarbeit/Data/FinalForm.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var csvs []csvFile
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		csvs = append(csvs, csvFile{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}
