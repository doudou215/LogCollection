package tailLog

import (
	"fmt"
	"github.com/doudou215/LogCollection/logAgent/kafka"
	"github.com/hpcloud/tail"
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init()
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	var err error
	t.instance, err = tail.TailFile(t.path, config)
	fmt.Println("open file ", t.path)
	if err != nil {
		fmt.Println("open file error ", err)
		return
	}
	go t.run()
	return
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
