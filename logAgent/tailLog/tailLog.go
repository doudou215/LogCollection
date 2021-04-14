package tailLog

import (
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)
func Init(filename string) error {
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 0},
		MustExist: false,
		Poll: true,
	}

	tailObj, err := tail.TailFile(filename, config)
	if err != nil {
		return err
	}
	return nil
}

func GetTailChan() <-chan *tail.Line {
	return tailObj.Lines
}