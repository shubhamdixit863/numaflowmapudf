package main

import (
	"context"
	"github.com/numaproj/numaflow-go/pkg/mapper"
	"log"
	"strconv"
)

type MultiplyByTen struct {
}

func (ml MultiplyByTen) Map(_ context.Context, keys []string, d mapper.Datum) mapper.Messages {
	msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	// Split the msg into an array with comma.
	num, err := strconv.Atoi(string(msg))

	if err != nil {
		return mapper.MessagesBuilder().Append(mapper.MessageToDrop())
	}
	return mapper.MessagesBuilder().Append(mapper.NewMessage([]byte(strconv.Itoa(num * 10))))

}

func main() {
	//It can take a function or Mapper interface type (a struct that is implementing the Mapper struct)
	err := mapper.NewServer(&MultiplyByTen{}).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map function server: ", err)
	}
}
