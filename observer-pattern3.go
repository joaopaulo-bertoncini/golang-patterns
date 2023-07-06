package main

// Example 3: Weather Station

import (
	"fmt"
	"sync"
)

// WeatherData is our Subject
type WeatherData struct {
	temperature float64
	humidity    float64
	displays    []chan string
}

// NewWeatherData creates a new WeatherData
func NewWeatherData() *WeatherData {
	return &WeatherData{
		displays: make([]chan string, 0),
	}
}

// RegisterDisplay method for adding a new display
func (w *WeatherData) RegisterDisplay(display chan string) {
	w.displays = append(w.displays, display)
}

// UpdateWeather updates weather data and notifies all displays
func (w *WeatherData) UpdateWeather(temperature float64, humidity float64, wg *sync.WaitGroup) {
	w.temperature = temperature
	w.humidity = humidity
	for _, display := range w.displays {
		wg.Add(1) // add to the counter before sending a message
		display <- fmt.Sprintf("Temperature: %.2f, Humidity: %.2f", w.temperature, w.humidity)
	}
}

func main() {
	weatherData := NewWeatherData()

	// creating channels for displays
	display1 := make(chan string)
	display2 := make(chan string)

	weatherData.RegisterDisplay(display1)
	weatherData.RegisterDisplay(display2)

	var wg sync.WaitGroup // using WaitGroup

	go func() {
		for msg := range display1 {
			fmt.Println("Display 1 updated:", msg)
			wg.Done() // decrement counter when observer 1 is done
		}
	}()

	go func() {
		for msg := range display2 {
			fmt.Println("Display 2 updated:", msg)
			wg.Done() // decrement counter when observer 2 is done
		}
	}()

	weatherData.UpdateWeather(25.0, 60.0, &wg)
	weatherData.UpdateWeather(26.5, 55.0, &wg)

	wg.Wait() // wait until all observers are done
}
