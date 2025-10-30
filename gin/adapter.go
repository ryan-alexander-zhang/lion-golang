package main

import "go.uber.org/zap"

type Service struct{ log *zap.Logger }

func NewService(base *zap.Logger) *Service {
	return &Service{log: base.With(zap.String("comp", "service"))}
}

func (s *Service) DoSomething() {
	s.log.Info("Service is doing something")
}
