package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Format("2006-01-02 01:01:01")
	commitMsg := fmt.Sprint("changes at %s", timestamp)

	//add all tracked files

}
