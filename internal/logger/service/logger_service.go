package service

type LoggerService interface {
	IndexApplicationLog() error
	IndexAuditLog() error
	IndexPerformLog() error
	IndexErrorLog() error
}
