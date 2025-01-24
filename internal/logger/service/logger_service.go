package service

import (
	"github.com/Ajulll22/belajar-microservice/internal/logger/config"
	"github.com/Ajulll22/belajar-microservice/internal/logger/model"
	"github.com/Ajulll22/belajar-microservice/internal/logger/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/broker"
)

type LoggerService interface {
	IndexApplicationLog(m *model.ApplicationLog) error
	IndexAuditLog(m *model.AuditLog) error
	IndexPerformLog(m *model.PerformLog) error
	IndexErrorLog(m *model.ErrorLog) error
}

func NewLoggerService(cfg config.Config, rmq broker.RabbitMQ, loggerRepository repository.LoggerRepository) LoggerService {
	return &loggerService{cfg, rmq, loggerRepository}
}

type loggerService struct {
	cfg              config.Config
	rmq              broker.RabbitMQ
	loggerRepository repository.LoggerRepository
}

func (s *loggerService) IndexApplicationLog(m *model.ApplicationLog) error {
	return s.loggerRepository.Index(s.cfg.APPLICATION_LOG_INDEX, m)
}

func (s *loggerService) IndexAuditLog(m *model.AuditLog) error {
	return s.loggerRepository.Index(s.cfg.AUDIT_LOG_INDEX, m)
}

func (s *loggerService) IndexPerformLog(m *model.PerformLog) error {
	return s.loggerRepository.Index(s.cfg.PERFORM_LOG_INDEX, m)
}

func (s *loggerService) IndexErrorLog(m *model.ErrorLog) error {
	return s.loggerRepository.Index(s.cfg.ERROR_LOG_INDEX, m)
}
