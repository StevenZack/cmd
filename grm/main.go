package main

import (
	"fmt"
	"github.com/StevenZack/cmd"
)

func main() {
	e := cmd.NewCmd("go", "run", "-tags", "gtk_3_18", "main.go").Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}
