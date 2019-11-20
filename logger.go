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

// Print format & print log
func (d *GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		d.WithFields(logrus.Fields{
			"module":  "gorm",
			"type":    "sql",
			"query":   v[3],
			"values":  v[4],
			"latency": v[2],
			"rows":    v[5],
			"src_ref": v[1],
		}).
			Debug("sql query")
	case "log":
		d.WithFields(logrus.Fields{"module": "gorm", "type": "log", "query": v[2]}).Debug("query log")
	}
}
