package work_queue

import (
	"github.com/cesnow/liquid-engine/logger"
	"time"
)

func Sample() {

	JobQueue = make(chan Job, MaxQueue)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	logger.SysLog.Infof("[Debug|Sample|WorkQueue] After Running")

	params := &JobParams{}
	params.AddString("OriginFullFilePath", "fff")
	params.AddString("ObjectPath", "ooo")
	work := Job{Payload: WorkTest, Data: params}
	logger.SysLog.Infof("[Debug|Sample|WorkQueue] sending payload to work queue")

	// Push the work onto the queue.
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work
	JobQueue <- work

	logger.SysLog.Infof("[Debug|Sample|WorkQueue] Sample Done")
}

func WorkTest(params IJobParams) error {
	time.Sleep(time.Duration(3) * time.Second)
	logger.SysLog.Infof("[WorkQueue][WorkTest] Done")
	return nil
}
