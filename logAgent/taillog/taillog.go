package taillog

import (
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		return err
	}
	return
}

func ReadChan() chan *tail.Line {
	return tailObj.Lines
}
