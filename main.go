package dogulib

import (
	"bufio"
	"log"
	"os"
)

// credits : https://gist.github.com/pkulak/93336af9bb9c7207d592
func ReadFileConcurrently(concurrency int, filename string, f func(queue chan string, complete chan bool)) {
	workQueue := make(chan string)

	complete := make(chan bool)

	go func() {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			workQueue <- scanner.Text()
		}

		close(workQueue)
	}()

	for i := 0; i < concurrency; i++ {
		go f(workQueue, complete)
	}

	for i := 0; i < concurrency; i++ {
		<-complete
	}
}
