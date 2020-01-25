package logger

import (
	"github.com/golang/glog"
)

const (
	LOG_LEVEL string = "2"
)

func init() {

	glog.Infof("The default glog level is %s ", LOG_LEVEL)
	glog.V(1).Infof("Starting you application")

}
