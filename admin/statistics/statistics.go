package statistics

import (
	pb "async_logger/codegen"
	"fmt"
	"io"
	"time"
)

type Statistics struct {
	IntervalSeconds int32
	ticker          time.Ticker
}

func GetStatistics(
	statInterval *pb.StatInterval,
	server pb.Admin_StatisticsServer,
) error {
	interval := statInterval.GetIntervalSeconds()
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	fmt.Println("interval:", interval)

	for {
		err := server.Send(&pb.Stat{
			Timestamp:  0,
			ByMethod:   nil,
			ByConsumer: nil,
		})

		if err != nil {
			return err
		}
		if err == io.EOF {
			return nil
		}
	}

	/*

		for {
			inWord, err := inStream.Recv()
			if err == io. EOF {
				return nil
			}
			if err != nil {
				return err
			}

			out := &translit.Word{
			Word: tr.Translit(inWord.Word),
			fmt.Println(inWord.Word, "->" , out.Word)
			inStream.Send(out)

		}

		return nil

	*/

}
