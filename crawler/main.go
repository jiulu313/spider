package main

import (
	"spider/crawler/engine"
	"spider/crawler/zhenai/parse"
	"spider/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseCityList,
	})
}
