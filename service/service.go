package service

import (
	"time"

	"github.com/EricTusk/template-http-grpc/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type service struct {
	number   float32
	sentence string
}

func NewService() (api.TemplateHTTPGRPCServiceServer, error) {
	s := &service{
		sentence: "default",
	}
	return s, nil
}

func (s *service) Echo(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {
	Count.WithLabelValues("echo").Add(1)

	logrus.Infof("number: %v, sentence: %s", req.GetEcho().GetNumber(), req.GetEcho().GetSentence())

	timer := prometheus.NewTimer(prometheus.ObserverFunc(Latency.WithLabelValues("echo").Observe))
	// do something
	time.Sleep(2 * time.Second)
	lastEcho := &api.EchoInfo{
		Number:   s.number,
		Sentence: s.sentence,
	}
	s.number = req.GetEcho().GetNumber()
	s.sentence = req.GetEcho().GetSentence()
	timer.ObserveDuration()

	res := &api.EchoResponse{
		LastEcho:    lastEcho,
		CurrentEcho: req.GetEcho(),
	}

	return res, nil
}

func (s *service) GetSystemInfo(ctx context.Context, req *api.GetSystemInfoRequest) (*api.GetSystemInfoResponse, error) {
	Count.WithLabelValues("get_system_info").Add(1)

	logrus.Infof("I am ok....")

	timer := prometheus.NewTimer(prometheus.ObserverFunc(Latency.WithLabelValues("get_system_info").Observe))
	// do something
	time.Sleep(2 * time.Second)
	timer.ObserveDuration()

	res := &api.GetSystemInfoResponse{
		Info: "I am ok....",
	}

	return res, nil
}
