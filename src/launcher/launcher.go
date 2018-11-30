package launcher

import (
	"fmt"
	"parser"
	"store"
)

func RunLauncher() {
	fmt.Println("Launch spider goroutines here")

	content := doFetchUrl(store.FetchNextUrl())
	parts := parser.Parse(content)

	fmt.Println(parts)

}

func Stop() {
	fmt.Println("Stop all spider goroutines")
}

/**
访问远程url获得内容
*/
func doFetchUrl(url string) string {

	fmt.Println("get content result")

	return ""
}
