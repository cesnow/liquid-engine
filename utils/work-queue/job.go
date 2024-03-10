package work_queue

import "github.com/cesnow/liquid-engine/logger"

type IJobParams interface {
	AddString(string, string)
	String(string) string
}

type JobParams struct {
	dictString map[string]string
}

func (j *JobParams) AddString(k string, v string) {
	if j.dictString == nil {
		j.dictString = make(map[string]string)
	}
	j.dictString[k] = v
}

func (j *JobParams) String(k string) string {
	if v, f := j.dictString[k]; f {
		return v
	}
	logger.SysLog.Warnf("[JobParams][GetString] Can't find key `%s`", k)
	return ""
}

// Job represents the job to be run
type Job struct {
	Payload func(IJobParams) error
	Data    *JobParams
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job
