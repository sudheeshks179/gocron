// Tests for gocron
package gocron

import (
	"fmt"
	"testing"
	"time"
)

var err = 1

func task() {
	fmt.Println("I am a running job.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func TestSecond(*testing.T) {
	var id1, id2 int64 = 1000, 2000
	defaultScheduler.Every(1).Second().Do(task)
	defaultScheduler.Every(1).Second().DoWithId(id1, taskWithParams, int(id1), "Running ")
	defaultScheduler.Every(1).Second().DoWithId(id2, taskWithParams, int(id2), "Running ")
	fmt.Println("Id1:", id1, "Id2:", id2)
	defaultScheduler.Start()
	fmt.Println("Timer for 4 secs starts")
	time.Sleep(4 * time.Second)
	fmt.Println("Stoping Id :", id1)
	defaultScheduler.RemoveById(id1)
	defaultScheduler.RemoveById(id1)
	fmt.Println("Timer for 10  secs starts")
	time.Sleep(10 * time.Second)
}
