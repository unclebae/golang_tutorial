package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// var site_map = []byte(`
// <sitemapindex>
// 	<sitemap>
// 		<loc>http://hello.com</loc>
// 	</sitemap>
// 	<sitemap>
// 		<loc>http://hi.com</loc>
// 	</sitemap>
// 	<sitemap>
// 		<loc>http://site.com</loc>
// 	</sitemap>
// </sitemapindex>`)

// type SitemapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// }

// type SitemapIndexV2 struct {
// 	Locations []string `xml:"sitemap>loc"`
// }

// type Location struct {
// 	Loc string `xml:"loc"`
// }

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `<h1> Hey there </h1>
// 	<p>Go is fast</p>
// 	<p>.... and simple</p>`)
// }

// func main() {
// 	// resp, _ := http.Get("https://www.google.com/")
// 	// bytes, _ := ioutil.ReadAll(resp.Body)

// 	// string_body := string(bytes)
// 	// // fmt.Println(string_body)

// 	// resp.Body.Close()

// 	bytes := site_map
// 	var s SitemapIndexV2
// 	xml.Unmarshal(bytes, &s)
// 	fmt.Println(s.Locations)

// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	// x := 5
// 	// for {
// 	// 	fmt.Printf("Do stuff\n", x)
// 	// 	x += 3
// 	// 	if x >= 25 {
// 	// 		break
// 	// 	}
// 	// }

// 	for _, Location := range s.Locations {
// 		fmt.Printf("\n%s", Location)
// 	}

// 	// var grades1 map[string]float32
// 	grades := make(map[string]float32)

// 	grades["Timmy"] = 50
// 	grades["Kido"] = 90
// 	grades["Jess"] = 70
// 	grades["Sam"] = 60

// 	fmt.Println(grades)

// 	TimsGrade := grades["Timmy"]
// 	fmt.Println(TimsGrade)

// 	delete(grades, "Timmy")
// 	fmt.Println(grades)

// 	for k, v := range grades {
// 		fmt.Println("Key: ", k, " Value: ", v)
// 	}
// }

type NewsAggPage struct {
	Titles string
	News   string
}

func newAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Titles: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")

	fmt.Println(t.Execute(w, p))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newAggHandler)
	http.ListenAndServe(":8000", nil)
}
