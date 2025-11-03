package logging

import (
	pb "async_logger/codegen"
	"io"
)

func GetLogs(
	_ *pb.Nothing,
	server pb.Admin_LoggingServer,
) error {

	for {
		err := server.Send(&pb.Event{
			Timestamp: 0,
			Consumer:  "",
			Method:    "",
			Host:      "",
		})

		if err != nil {
			return err
		}
		if err == io.EOF {
			return nil
		}
	}

}
