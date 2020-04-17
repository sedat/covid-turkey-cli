package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

type JSONBody struct {
	Country             string `json:"country"`
	Cases               int    `json:"cases"`
	TodayCases          int    `json:"todayCases"`
	Deaths              int    `json:"deaths"`
	TodayDeaths         int    `json:"todayDeaths"`
	Recovered           int    `json:"recovered"`
	Active              int    `json:"active"`
	Critical            int    `json:"critical"`
	CasesPerOneMillion  int    `json:"casesPerOneMillion"`
	DeathsPerOneMillion int    `json:"deathsPerOneMillion"`
	TotalTests          int    `json:"totalTests"`
	TestsPerOneMillion  int    `json:"testsPerOneMillion"`
}

func main() {
	app := &cli.App{
		Name:  "covid-turkey",
		Usage: "Covid Türkiye güncel bilgiler",
		Action: func(c *cli.Context) error {
			url := "https://coronavirus-19-api.herokuapp.com/countries/turkey"
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			var data JSONBody
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Son veriler; \n Toplam vaka: %d \n Toplam vefat: %d \n Bugün vaka: %d \n Bugün vefat: %d \n İyileşen: %d \n Aktif hasta: %d \n Kritik: %d \n Milyonda vaka: %d \n Milyonda ölüm: %d \n Toplam test: %d \n", data.Cases, data.Deaths, data.TodayCases, data.TodayDeaths, data.Recovered, data.Active, data.Critical, data.CasesPerOneMillion, data.DeathsPerOneMillion, data.TotalTests)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
