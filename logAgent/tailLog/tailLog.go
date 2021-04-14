package tailLog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(filename string) error {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 0},
		MustExist: false,
		Poll:      true,
	}

	tailObj, err := tail.TailFile(filename, config)
	fmt.Println(tailObj.Filename)
	if err != nil {
		return err
	}
	return nil
}

func GetTailChan() <-chan *tail.Line {
	return tailObj.Lines
}
