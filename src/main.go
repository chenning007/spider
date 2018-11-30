package main

import "fmt"
import "launcher"

func main() {
	fmt.Println("hello spider, start from here")

	launcher.RunLauncher()

	launcher.Stop()

	fmt.Println("spider stopped")
}
