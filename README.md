# Weather forecast crawler from _bbc.com

This is a simple app that craw weather forecast from bbc.com and store to db 

## Informations

This app will call directly to bbc api to get the summary weather forecast and store to postgresql db (table day_reports by default).<br>

It uses some Golang packages: <br/>
* [go-pg](https://github.com/go-pg/pg): a Golang package implement a PostgreSQL client, it also support ORM.
* [gocolly](https://github.com/gocolly/colly): this package is used for scraping purpose

The application wil take about 500ms - 1s to complete. By default, it will result next 14 days weather forecast, you can also query
for a specific day by parsing parameter to run.
## How to run the app
1. Install dependencies by this command (execution permission may be needed to run the script):
```bash
./install_dependencies.sh
```
<br/>
2. You must install postgres on your machine and update information to configuration.toml file.The file may look like this:

```bash
db_host = "host:port"
db_user = "postgre username"
db_password = "postgre password"
db_name = "database name"
```
<br/>
3. Get weather forecast for next 14 days:

```bash
./start_app.sh
```
The result will be somethings like this, you can also find the detail in table _**day_reports**_:

```bash
{

"date": "Sunday 29th September",

"description": "Thundery showers and a moderate breeze",

"temperature": "19°C - 23°C"
}
{

"date": "Monday 30th September",

"description": "Sunny intervals and a moderate breeze",

"temperature": "22°C - 31°C"
}
...

```
<br>
4. Get detail weather forecast for specific day by parsing day pamrameter with yyyy-mm-dd format:

```bash
./start_app.sh 2019-10-01
```
<br/>
Result:

```bash
2019/09/30 04:26:17 Reading Configuration
2019/09/30 04:26:18 Detail...
2019/09/30 04:26:18 {
        "localDate": "2019-10-01",
        "enhancedWeatherDescription": "Light cloud and a moderate breeze",
        "maxTempC": 30,
        "maxTempF": 85,
        "minTempC": 22,
        "minTempF": 71,
        "sunrise": "06:47",
        "sunset": "18:33",
        "uvIndex": 3,
        "uvIndexBand": "MODERATE",
        "windDirection": "SW",
        "windDescription": "A moderate breeze from the south west",
        "windSpeedKph": 24,
        "windSpeedMph": 15,
        "weatherType": 7,
        "LocationID": "4887398",
        "LastUpdated": "2019-09-30T04:26:18.165414+07:00"
}
2019/09/30 04:26:18 Finished! Time elapsed: 301.325348ms

```
