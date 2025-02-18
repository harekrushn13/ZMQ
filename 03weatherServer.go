package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"os"
	"time"
)

type WeatherData struct {
	Zipcode     string  `json:"zipcode"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func readWeatherData(filename string) ([]WeatherData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var weatherData []WeatherData
	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		return nil, err
	}

	return weatherData, nil
}

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUB)
	defer socket.Close()
	defer context.Term()
	socket.Bind("tcp://*:5555")
	//socket.Bind("ipc://weather.ipc")

	rand.Seed(time.Now().UnixNano())

	weatherData, err := readWeatherData("03weather_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		index := rand.Intn(len(weatherData))
		entry := weatherData[index]

		msg := fmt.Sprintf("%s %.2f %v", entry.Zipcode, entry.Temperature, entry.Humidity)

		socket.Send(msg, 0)
		fmt.Println("Sent : ", msg)

		time.Sleep(5 * time.Second)
	}
}
