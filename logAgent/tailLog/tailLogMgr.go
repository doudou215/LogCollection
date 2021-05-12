package tailLog

import (
	"fmt"
	"github.com/doudou215/LogCollection/logAgent/etcd"
	"time"
)

type TailMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

var tailLogMgr *TailMgr

func Init(logEntryConf []*etcd.LogEntry) {
	tailLogMgr = &TailMgr{
		logEntry:    logEntryConf,
		tskMap:      make(map[string]*TailTask, 10),
		newConfChan: make(chan []*etcd.LogEntry), // 没有缓冲区的通道
	}
	for _, logEntry := range logEntryConf {
		tailTask := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tailLogMgr.tskMap[mk] = tailTask
	}

	go tailLogMgr.run()
}

func (t *TailMgr) run() {
	for {
		select {
		case newConfs := <-t.newConfChan:
			fmt.Println("new configuration has been arried")
			for _, newConf := range newConfs {
				mk := fmt.Sprintf("%s_%s", newConf.Path, newConf.Topic)
				_, ok := t.tskMap[mk]
				if !ok {
					tailobj := NewTailTask(newConf.Path, newConf.Topic)
					t.tskMap[mk] = tailobj
				}
			}

			for _, oldConf := range t.logEntry {
				isDelete := true
				for _, newConf := range newConfs {
					if newConf.Topic == oldConf.Topic && newConf.Path == oldConf.Path {
						isDelete = false
						break
					}
				}

				if isDelete {
					mk := fmt.Sprintf("%s_%s", oldConf.Path, oldConf.Topic)
					fmt.Printf("%s is delete\n", mk)
					t.tskMap[mk].cancel()
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

// chan<- 只写的通道
func GetNewConfChan() chan<- []*etcd.LogEntry {
	return tailLogMgr.newConfChan
}
