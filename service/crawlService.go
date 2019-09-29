package service

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type PrettyReport struct{
	Date string `json:"date"`
	Description string `json:"description"`
	Temperature string `json:"temperature"`
}
func PrettyForecast()  {
	urls := "https://www.bbc.com/weather/4887398"
	c := colly.NewCollector(
		colly.AllowedDomains("bbc.com", "www.bbc.com"),
		colly.CacheDir("_weather_cache"))

	c.OnHTML("a[id^=daylink-]", func(el *colly.HTMLElement) {

		date := el.ChildAttr(".wr-day__title", "aria-label")
		wrDes := el.ChildText(".wr-day__body .wr-day__weather-type-description-container")
		highTemp := el.ChildText(".wr-day__body .wr-day__details .wr-day-temperature__high-value .wr-value--temperature--c")
		lowTemp := el.ChildText(".wr-day__body .wr-day__details .wr-day-temperature__low-value .wr-value--temperature--c")
		f := PrettyReport{
			Date: date,
			Description: wrDes,
			Temperature: fmt.Sprintf("%sC - %sC",lowTemp,highTemp),
		}
		res, _ := json.MarshalIndent(f, "", "\t")
		log.Println(string((res)))

	})

	c.Visit(urls)
}