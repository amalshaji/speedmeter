package main

import (
	"fmt"
	"time"

	"github.com/amalshaji/speedmeter/utils"
)

func main() {
	rInit, tInit := utils.GetBytes()
	for {
		fmt.Print("\033[u\033[K")
		time.Sleep(1 * time.Second)
		go func() {
			r, t := utils.GetBytes()
			rDiff, tDiff := r-rInit, t-tInit
			rInit, tInit = r, t
			fmt.Printf("Download: %d KBps | Upload: %d KBps", rDiff/1024, tDiff/1024)
		}()
	}
}
