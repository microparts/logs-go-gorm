package gormLogger

import (
	"github.com/sirupsen/logrus"
)

type GormLogger struct {
	*logrus.Logger
}

//NewLogger initiates new gorm logger
func NewLogger(logger *logrus.Logger) *GormLogger {
	return &GormLogger{Logger: logger}
}

//Println accepts db logs ang uses logrus to log them
func (d *GormLogger) Println(v ...interface{}) {
	if v[0] == "sql" {
		d.WithFields(logrus.Fields{"query": v[3], "values": v[4], "rows": v[5], "latency": v[2]}).Info("Query sql")
	}
	if v[0] == "log" {
		d.WithFields(logrus.Fields{"query": v[2]}).Info("Query log")
	}
}

// Print format & print log
func (d *GormLogger) Print(values ...interface{}) {
	d.Println(values...)
}
