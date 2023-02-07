package main

import (
	"fmt"
	"net/url"
)

func main() {

	// Percent Encoding
	text := "This is a trial for Percent Encoding 🙊"
	fmt.Println(url.QueryEscape(text))

	//Query String/Form Post
	values := url.Values{}
	values.Add("Nama", "Febryan")
	values.Add("Umur", "28 tahun")
	values.Add("Hobi", "Sepak Bola")
	values.Add("Emoji", "🦁")
	fmt.Println(values.Encode())

	//Path Encoding
	path := "test?test++"
	fmt.Println(url.PathEscape(path))

	//Test Membuat URL Lengkap
	baseUrl, err := url.Parse("https://contohurl.com")
	if err != nil {
		panic("cannot pare url")
	}

	pathtest := "some random path"

	baseUrl.Path += url.PathEscape(pathtest)

	qs := url.Values{}
	qs.Add("handle", "@febryanzz")
	qs.Add("seq", "716386138")
	qs.Add("another", "justaddthisrandomstring")

	baseUrl.RawQuery = qs.Encode()

	fmt.Println(baseUrl.String())

}
