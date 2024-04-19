package internal

import (
	"encoding/json"
	"log"
	"time"

	"github.com/tarm/serial"
)

const RX_BYTE = "1"

type TimeSeriesData struct {
	Status                 string  `json:"status,omitempty"`
	Runtime                int     `json:"runtime,omitempty"`
	Id                     int     `json:"id,omitempty"`
	HumidityValue          float32 `json:"humidity_value,omitempty"`
	HumidityUnit           string  `json:"humidity_unit,omitempty"`
	TemperatureValue       float32 `json:"temperature_primary_value,omitempty"`
	TemperatureUnit        string  `json:"temperature_primary_unit,omitempty"`
	PressureValue          float32 `json:"pressure_value,omitempty"`
	PressureUnit           string  `json:"pressure_unit,omitempty"`
	TemperatureBackupValue float32 `json:"temperature_secondary_value,omitempty"`
	TemperatureBackupUnit  string  `json:"temperature_secondary_unit,omitempty"`
	AltitudeValue          float32 `json:"altitude_value,omitempty"`
	AltitudeUnit           string  `json:"altitude_unit,omitempty"`
	IlluminanceValue       float32 `json:"illuminance_value,omitempty"`
	IlluminanceUnit        string  `json:"illuminance_unit,omitempty"`
	Timestamp              int64   `json:"timestamp,omitempty"`
}

type ArduinoReader struct {
	serialPort *serial.Port
}

func NewArduinoReader(devicePath string) *ArduinoReader {
	log.Printf("Reader: Attempting to connect Arduino device...")
	c := &serial.Config{Name: devicePath, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("READER: Waiting for Arduino external sensors to warm up")
	ar := &ArduinoReader{serialPort: s}
	ar.GetSparkFunWeatherShieldData()
	time.Sleep(5 * time.Second)
	ar.GetSparkFunWeatherShieldData()
	time.Sleep(5 * time.Second)
	return ar
}

func (ar *ArduinoReader) GetSparkFunWeatherShieldData() *TimeSeriesData {
	n, err := ar.serialPort.Write([]byte(RX_BYTE))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1028)
	n, err = ar.serialPort.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	var tsd TimeSeriesData
	err = json.Unmarshal(buf[:n], &tsd)
	if err != nil {
		return nil
	}
	tsd.Timestamp = time.Now().Unix()
	return &tsd
}
