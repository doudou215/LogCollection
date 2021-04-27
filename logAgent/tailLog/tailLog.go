package tailLog

import (
	"fmt"
	"github.com/doudou215/LogCollection/logAgent/kafka"
	"github.com/hpcloud/tail"
)

type TailTak struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTak) {
	tailObj = &TailTak{
		path:  path,
		topic: topic,
	}
	tailObj.Init()
	return
}

func (t *TailTak) Init() {
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

func (t *TailTak) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			kafka.SentToKafka(t.topic, line.Text)
		}
	}
}
