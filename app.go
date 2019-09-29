package main

import (
	"cic-project/dao"
	"cic-project/model"
	"cic-project/service"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct{
	DbHost			string  `toml:"db_host"`
	DbUser			string  `toml:"db_user"`
	DbPassword		string  `toml:"db_password"`
	LocationID 		string	`toml:"location_id"`
	DbName 			string	`toml:"db_name"`
}

var config = Config{}

func init()  {
	log.Println("Reading Configuration")
	if _, err := toml.DecodeFile("configuration.toml", &config); err != nil {
		log.Fatal(err)
	}
}

func main()  {
	start := time.Now()
	var queryDate = ""
	if len(os.Args) > 1 {
		queryDate = os.Args[1]
	} else {
		log.Println("Summary forecast...")
		service.PrettyForecast()
	}

	urls := "https://weather-broker-cdn.api.bbci.co.uk/en/forecast/aggregated/4887398"
	resp, err := service.SendGetRequest(urls)

	if err != nil {
		log.Fatal(err)
	}
	var report model.Response
	pgDao := dao.PgDao{
		Host: config.DbHost,
		User: config.DbUser,
		Password: config.DbPassword,
		DbName: config.DbName,
	}
	pgDao.GetConnection()
	defer pgDao.Db.Close()
	err = json.NewDecoder(strings.NewReader(string(resp))).Decode(&report)
	locationID := report.Location.ID
	for _, item := range report.Forecasts {
		newReport := item.Summary.Report
		newReport.LocationID = locationID
		pgDao.InsertForecast(newReport)
	}
	log.Println("Detail...")
	r := pgDao.FindForecast(queryDate)
	for _, item := range r {
		res, _ := json.MarshalIndent(item, "", "\t")
		log.Print(string(res))
	}
	end := time.Now()
	log.Println("Finished! Time elapsed:", end.Sub(start))
}

