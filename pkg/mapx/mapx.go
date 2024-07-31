package mapx

import "fmt"

func RunMap() {
	umap := make(map[string]int)
	umap["2"] = 2

	if value, ok := umap["2"]; ok {
		fmt.Println(value)
	}
}
