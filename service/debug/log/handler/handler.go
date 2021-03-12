package handler

import (
	"context"
	"sync"

	"github.com/btccom/go-micro/v2/debug/log"
	"github.com/btccom/go-micro/v2/errors"
	pb "github.com/btccom/go-micro-platform/v2/service/debug/log/proto"
)

type Log struct {
	// per service log
	sync.RWMutex
	Logs map[string]log.Log

	// Ability to create new logger
	New func(string) log.Log
}

func (l *Log) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	if len(req.Service) == 0 {
		return errors.BadRequest("go.micro.debug.log", "Invalid service name")
	}

	l.Lock()
	defer l.Unlock()

	// get the service log
	serviceLog, ok := l.Logs[req.Service]
	if !ok {
		serviceLog = l.New(req.Service)
		l.Logs[req.Service] = serviceLog
	}

	// TODO: specify how many log records to read
	records, err := serviceLog.Read()
	if err != nil {
		return err
	}

	// append to records
	for _, rec := range records {
		rsp.Records = append(rsp.Records, &pb.Record{
			Timestamp: rec.Timestamp.Unix(),
			Metadata:  rec.Metadata,
			Message:   rec.Message.(string),
		})
	}

	return nil
}
