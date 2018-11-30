package spider

import "fmt"

type spider struct {
}

func (s spider) Fetch(url string) {
	fmt.Println("fetch content from url")
}
