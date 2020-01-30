package taillog

import "awesomeProject/logAgent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	//tskMap map[string]*TailTask
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{logEntry: logConf}
	for _, logEntry := range logConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
