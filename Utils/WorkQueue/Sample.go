package WorkQueue

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"time"
)

func Sample() {

	JobQueue = make(chan Job, MaxQueue)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	Logger.SysLog.Infof("[Debug|Sample|WorkQueue] After Running")

	params := &JobParams{}
	params.AddString("OriginFullFilePath", "fff")
	params.AddString("ObjectPath", "ooo")
	work := Job{Payload: WorkTest, Data: params}
	Logger.SysLog.Infof("[Debug|Sample|WorkQueue] sending payload to work queue")

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

	Logger.SysLog.Infof("[Debug|Sample|WorkQueue] Sample Done")
}

func WorkTest(params IJobParams) error {
	time.Sleep(time.Duration(3) * time.Second)
	Logger.SysLog.Infof("[WorkQueue][WorkTest] Done")
	return nil
}
