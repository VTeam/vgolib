package filewatch

import (
	"context"
	"os"
	"sync"
	"time"

	"log"
)

type toWatch []string
type updateHandler func([]string)

func NewFileWatcher(ctx context.Context, toWatch toWatch,
	handler updateHandler, interval time.Duration) *FileWatcher {
	return &FileWatcher{
		ctx:         ctx,
		interval:    interval,
		handler:     handler,
		watchRecord: newWatchRecord(toWatch),
	}
}

type FileWatcher struct {
	ctx         context.Context
	handler     updateHandler
	interval    time.Duration
	watchRecord watchRecord
	once        sync.Once
}

func (fw *FileWatcher) Run() {
	fw.once.Do(func() {
		ticker := time.NewTicker(fw.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if updated := fw.watchRecord.Watch(); len(updated) > 0 {
					fw.handler(updated)
				}
			case <-fw.ctx.Done():
				return
			}
		}

	})
}

type watchRecord map[string]time.Time

func (wr watchRecord) Watch() []string {
	var updated []string

	for name, prev := range wr {
		stat, err := os.Stat(name)
		if err != nil {
			switch {
			case os.IsNotExist(err) && !prev.IsZero():
				// file has been deleted
				updated = append(updated, name)
				wr[name] = time.Time{}
			case os.IsNotExist(err):
				// file not exists, ignore
			default:
				log.Fatalf("While get state of file %s error: %s", name, err)
			}
			continue
		}
		if prev != stat.ModTime() {
			updated = append(updated, name)
			wr[name] = stat.ModTime()
		}
	}
	return updated
}

func newWatchRecord(files toWatch) watchRecord {
	var watchRecord = watchRecord(map[string]time.Time{})
	for _, f := range files {
		watchRecord[f] = time.Time{}
	}
	watchRecord.Watch()

	return watchRecord
}
