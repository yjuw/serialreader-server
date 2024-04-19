package internal // github.com/bartmika/serialreader-server/internal/server_impl.go

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/yjuw/serialreader-server/proto"
)

type SerialReaderServerImpl struct {
	arduinoReader *ArduinoReader
	pb.SerialReaderServer
}

func (s *SerialReaderServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *SerialReaderServerImpl) GetSparkFunWeatherShieldData(ctx context.Context, in *pb.GetTimeSeriesData) (*pb.SparkFunWeatherShieldTimeSeriesData, error) {
	datum := s.arduinoReader.GetSparkFunWeatherShieldData()
	return &pb.SparkFunWeatherShieldTimeSeriesData{
		Status:                 true,
		Timestamp:              ptypes.TimestampNow(), // Note: https://godoc.org/github.com/golang/protobuf/ptypes#Timestamp
		HumidityValue:          datum.HumidityValue,
		HumidityUnit:           datum.HumidityUnit,
		TemperatureValue:       datum.TemperatureValue,
		TemperatureUnit:        datum.TemperatureUnit,
		PressureValue:          datum.PressureValue,
		PressureUnit:           datum.PressureUnit,
		TemperatureBackupValue: datum.TemperatureBackupValue,
		TemperatureBackupUnit:  datum.TemperatureBackupUnit,
		AltitudeValue:          datum.AltitudeValue,
		AltitudeUnit:           datum.AltitudeUnit,
		IlluminanceValue:       datum.IlluminanceValue,
		IlluminanceUnit:        datum.IlluminanceUnit,
	}, nil
}
