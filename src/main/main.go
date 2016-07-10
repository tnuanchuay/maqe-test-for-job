package main

import(
	"fmt"
	"crypto/sha1"
)

func main(){
	var codes []string

	h := sha1.New()

	for i := 0 ; i < 250 ; i++{
		h.Write([]byte(string(i)))
		codes = append(codes, fmt.Sprintf("MQ%x", h.Sum(nil) )[:10]);
	}

	
	for _,code := range codes{
		fmt.Println(code);
	}

}
