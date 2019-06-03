package timer

import "time"

const StartSetInfoTime = time.Second * 10

func getDuration() time.Duration {
	t := time.Now().Add(time.Hour)
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	return t.Sub(time.Now())
}

func watchFunc(fn func()) {
	// 每个一个小时去更新一下企业信息，并且是整点的时候更新，保证所有服务同步
	go func() {
		t := time.NewTimer(StartSetInfoTime)
		defer t.Stop()
		for {
			<-t.C
			fn()
			t.Reset(getDuration())
		}
	}()
}
