package logging

import (
	pb "async_logger/codegen"
	"io"

	"async_logger/internal/logger"
)

func GetLogs(
	server pb.Admin_LoggingServer,
	logger *logger.Logger,
) error {

	for {
		event := logger.GetLog()

		err := server.Send(&pb.Event{
			Timestamp: event.Timestamp,
			Consumer:  event.Consumer,
			Method:    event.Method,
			Host:      event.Host,
		})

		if err != nil {
			return err
		}
		if err == io.EOF {
			return nil
		}
	}

}
