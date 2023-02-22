package workers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gocraft/work"
)

type Context struct{}

var JobEnqueuer *work.Enqueuer

func init() {

	namespace, _ := beego.AppConfig.String("redisnamespace")

	workerPool := work.NewWorkerPool(Context{}, 10, namespace, RedisPool)

	// Add middleware that will be executed for each job
	workerPool.Middleware((*Context).Log)

	workerPool.Job("keyword_crawler", (*Context).CrawlKeywork)

	// Start processing jobs
	workerPool.Start()

	JobEnqueuer = work.NewEnqueuer(namespace, RedisPool)
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	logs.Info("Starting job: ", job.Name)
	return next()
}
