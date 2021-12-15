package main

import (
	"fmt"
	"mysite/repositories/filesystem"
)

func main() {
	f := &filesystem.UserFileRepository{}
	u := f.GetByEmail("e")
	fmt.Print(u)
}
