package global

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()
var WorkerWs = workerWsInit()
