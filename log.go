package mdserver

import (
	"log"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
)

func initLog() *rotatelogs.RotateLogs {

	rl := rotatelogs.NewRotateLogs(
		"./access_log.%Y%m%d",
	)

	// Optional fields must be set afterwards
	rl.LinkName = "./access_log"
	rl.RotationTime = 86400 * time.Second
	rl.MaxAge = 86400 * time.Second * 2
	rl.Offset = 0

	log.SetOutput(rl)

	return rl
}
