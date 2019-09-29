package model

import "time"

type DayReport struct {
	ID 					int64		`json:"-"`
	LocalDate			string		`json:"localDate" pg:",unique"`
	WeatherDescription 	string 		`json:"enhancedWeatherDescription"`
	MaxTempC			int			`json:"maxTempC"`
	MaxTempF			int			`json:"maxTempF"`
	MinTempC			int			`json:"minTempC"`
	MinTempF			int			`json:"minTempF"`
	Sunrise				string		`json:"sunrise" `
	Sunset				string		`json:"sunset"`
	UvIndex				int			`json:"uvIndex"`
	UvIndexBand			string		`json:"uvIndexBand"`
	WindDirection		string		`json:"windDirection"`
	WindDescription		string		`json:"windDescription"`
	WindSpeedKph		int		`json:"windSpeedKph"`
	WindSpeedMph		int		`json:"windSpeedMph"`
	WeatherType			int		`json:"weatherType"`
	LocationID 			string
	LastUpdated			time.Time	`pg:"default:now()"`
}
type Summary struct {
	Report 		DayReport 		`json:"report"`
}

type Forecasts struct {
	Summary 		Summary			`json:"summary"`

}

type Response struct {
	Forecasts		[]Forecasts  	`json:"forecasts"`
	Location		Location 		`json:"location"`
}

type Location struct {
	Container 	string 		`json:"container"`
	ID			string		`json:"id"`
	Latitude	float32		`json:"latitude"`
	Longtitude	float32		`json:"longtitude"`
	Name		string		`json:"name"`
}

type ForecastDom struct {
	Date		string
	WeatherDescription string
	MaxTempC	string
	MinTempC	string
	WindSpeed	string
	WindDescription string
}