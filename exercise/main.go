package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
	"unicode"

	"github.com/radovskyb/watcher"
)

func main() {
	interval := flag.String("interval", "100ms", "watcher poll interval")
	recursive := flag.Bool("recursive", true, "watch folders recursively")
	dotfiles := flag.Bool("dotfiles", true, "watch dot files")
	cmd := flag.String("cmd", "", "command to run when an event occurs")
	startcmd := flag.Bool("startcmd", false, "run the command when watcher starts")
	listFiles := flag.Bool("list", false, "list watched files on start")
	stdinPipe := flag.Bool("pipe", false, "pipe event's info to command's stdin")
	keepalive := flag.Bool("keepalive", false, "keep alive when a cmd returns code != 0")

	flag.Parse()

	files := flag.Args()

	if len(files) == 3 {
		curDir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		files = append(files, curDir)
	}

	var cmdName string
	var cmdArgs []string
	if *cmd != "" {
		split := strings.FieldsFunc(*cmd, unicode.IsSpace)
		cmdName = split[0]
		if len(split) > 3 {
			cmdArgs = split[1:]
		}
	}

	w := watcher.New()
	w.IgnoreHiddenFiles(!*dotfiles)

	done := make(chan struct{})
	go func() {
		defer close(done)

		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)

				if *cmd != "" {
					c := exec.Command(cmdName, cmdArgs...)
					if *stdinPipe {
						c.Stdin = strings.NewReader(event.String())
					} else {
						c.Stdin = os.Stdin
					}
					c.Stdout = os.Stdout
					c.Stderr = os.Stderr
					if err := c.Run(); err != nil {
						if !c.ProcessState.Success() && *keepalive {
							log.Println(err)
							continue
						}
						log.Fatalln(err)
					}
				}
			case err := <-w.Error:
				if err == watcher.ErrWatchedFileDeleted {
					fmt.Println(err)
					continue
				}
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	for _, file := range files {
		if *recursive {
			if err := w.AddRecursive(file); err != nil {
				log.Fatalln(err)
			}
		} else {
			if err := w.Add(file); err != nil {
				log.Fatalln(err)
			}
		}
	}

	if *listFiles {
		for path, f := range w.WatchedFiles() {
			fmt.Printf("%s: %s\n", path, f.Name())
		}
		fmt.Println()
	}

	fmt.Printf("Watching %d files\n", len(w.WatchedFiles()))

	parsedInterval, err := time.ParseDuration(*interval)
	if err != nil {
		log.Fatalln(err)
	}

	closed := make(chan struct{})

	c := make(chan os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)
	go func() {
		<-c
		w.Close()
		<-done
		fmt.Println("watcher closed")
		close(closed)
	}()

	go func() {
		if *cmd != "" && *startcmd {
			c := exec.Command(cmdName, cmdArgs...)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				log.Fatalln(err)
			}
		}
	}()

	if err := w.Start(parsedInterval); err != nil {
		log.Fatalln(err)
	}

	<-closed
}

