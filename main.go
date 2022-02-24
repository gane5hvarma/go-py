// gp_py package
package hello

import "fmt"

func Hello(){
	words := []string{"hello","worlds"}
	printWords(words...)

}

func printWords(words ...string){
	wordCh:=make(chan string)
	go recieveFromCh(wordCh)
	for _,word := range words{
		fmt.Println(word)
		wordCh <- word
	}
}
func recieveFromCh(wordCh <- chan string) {
	for word := range wordCh {
		fmt.Println(word)

	}

}
