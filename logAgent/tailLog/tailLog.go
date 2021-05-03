package tailLog

import (
	"context"
	"fmt"
	"github.com/doudou215/LogCollection/logAgent/kafka"
	"github.com/hpcloud/tail"
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	tailObj.init()
	return tailObj
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
	if err != nil {
		fmt.Println("open file error ", err)
		return
	}
	fmt.Println("open file ", t.path)
	go t.run()
	return
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("%s_%s quit\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			fmt.Printf("%s\n", line.Text)
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
