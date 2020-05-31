package main

import (
	"github.com/radovskyb/watcher"
	"regexp"
	"time"
)

func watch_sources() {
	// Watch go files.
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write, watcher.Create)
	w.AddFilterHook(watcher.RegexFilterHook(regexp.MustCompile("(?i).go$"), false))
	if err := w.AddRecursive("."); err != nil {
		println(err.Error())
		w.Close()
	} else {
		// Fire up the watcher!
		go func() {
			for {
				select {
				case <-w.Event:
					println("TODO!")
				case err := <-w.Error:
					println(err.Error())
					return
				case <-w.Closed:
					return
				}
			}
		}()
		go func() {
			if err := w.Start(time.Millisecond * 200); err != nil {
				println(err.Error())
				w.Close()
			}
		}()
	}
}
