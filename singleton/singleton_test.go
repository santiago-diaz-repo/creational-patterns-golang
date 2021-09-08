package singleton

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func Test_GetInstance(t *testing.T) {

	dataTable := []struct {
		name string
		want string
	}{
		{"1", "Creating instance\n"},
		{"2", ""},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := captureOutput(func() {
				GetInstance()
			})
			fmt.Println("****** " + got)
			if got != v.want {
				t.Errorf("Error, wanted: %s**, got: %s**", v.want, got)
			}
		})
	}

}

// captureOutput is explained in https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
