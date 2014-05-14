package main

import (
	"github.com/tamaxyo/fswatcher/config"
	"github.com/tamaxyo/fswatcher/watcher"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	FILENAME          = ".fswatcher"
	PATTERN_SEPARATOR = ","
)

func main() {
	input, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatal("error: ", err)
	}

	configs, err := config.Parse(input)
	if err != nil {
		log.Fatal("parse error: ", err)
	}

	for _, c := range configs {
		execute(c)
	}

	done := make(chan bool)
	<-done
}

func execute(c config.Config) {
	rw, err := watcher.NewWatcher(c.Path, c.Recursive)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-rw.Event:
				name := filepath.Base(ev.Name)
				if match(name, c.Pattern) && !match(name, c.Ignore) {
					log.Println("event: ", ev)

					cmd := setupCommand(c.Path, c.Command)
					if err := cmd.Run(); err != nil {
						log.Println(err)
					}
				}
			case err := <-rw.Error:
				log.Println("error: ", err)
			}
		}
	}()
}

func match(name string, pattern string) bool {
	for _, p := range strings.Split(pattern, PATTERN_SEPARATOR) {
		if match, _ := filepath.Match(p, name); match {
			return true
		}
	}
	return false
}

func setupCommand(workdir string, command string) *exec.Cmd {
	commands := strings.Split(command, " ")

	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Dir = workdir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
