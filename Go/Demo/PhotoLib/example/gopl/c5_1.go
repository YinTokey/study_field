package gopl

import (
	"fmt"
	"golang.org/x/net/html"
)

func Pa() {
	//fmt.Println("xxx")
	//file, err := os.Open("c5_1.html")
	//if err != nil {
	//	panic(err)
	//}
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	panic(err)
	//}
	//doc, err := html.Parse(file)
	//if err != nil {
	//	panic(err)
	//	os.Exit(1)
	//}
	//for _, link := range visit(nil,doc) {
	//	fmt.Println(link)
	//}

}

func visit(links []string, n *html.Node) []string {

	 if n.Type == html.ElementNode && n.Data == "a" {
	 	fmt.Println("in a")
	 	for _, a := range  n.Attr {
	 		if a.Key == "href" {
	 			links = append(links,a.Val)
			}
		}
	 }

	 for c:= n.FirstChild; c!= nil; c = c.NextSibling {

	 	links = visit(links,c)

	 }
	 return links
}

func Sum(vals ... int) int {
	total := 0
	for _, val := range vals {
		total += val;
	}
	return total
}