package settings

import (
	consts "github.com/STUD-IT-team/bauman-legends-backend/internal/app/consts"
	log "github.com/sirupsen/logrus"
)

func LogSetup() {
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = consts.LogDateFormat
	formatter.FullTimestamp = true
	formatter.DisableLevelTruncation = true
	log.SetFormatter(formatter)
}
