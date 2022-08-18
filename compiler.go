package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/skip2/go-qrcode"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

func main() {
	dir_html := os.Args[1]

	os.Chdir(dir_html)

	//fmt.Println(dir_html)

	if _, err := os.Stat("index.html"); err != nil {
		fmt.Println("index.html not found in dir !")
		os.Exit(-1)
	}

	compile("index.html")
}

func compile(path string) {
	file, err := os.ReadFile(path)

	if err != nil {
		println("error !")
	}
	//fmt.Println(file)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(file)))

	var styles []string
	doc.Find("link[rel=\"stylesheet\"]").Each(func(i int, s *goquery.Selection) {
		if attr, ex := s.Attr("href"); ex {
			fmt.Println("css found: " + attr)
			file, err := os.ReadFile(attr)
			if err != nil {
				println("error read " + attr + " !")
			}

			file_string := string(file)
			fmt.Print(len(file_string))
			minified := do_min("text/css", file_string)

			s.Remove()

			styles = append(styles, minified)
			fmt.Print(" -> ")
			fmt.Println(len(minified))
		}
	})

	doc.Find("head").AppendHtml("<style>" + do_min("text/css", strings.Join(styles, "\n")) + "</style>")

	var scripts []string

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		fmt.Print("js found: ")

		var js_file string

		if attr, ex := s.Attr("src"); ex {
			fmt.Println(attr)
			file, err := os.ReadFile(attr)
			if err != nil {
				println("error read " + attr + " !")
			}

			js_file = string(file)
		} else {
			js_file = s.Text()
		}

		s.Remove()

		fmt.Print(len(js_file))
		minified := do_min("text/javascript", js_file)
		fmt.Print(" -> ")
		fmt.Println(len(minified))
		scripts = append(scripts, minified)
	})

	doc.Find("body").AppendHtml("<script>" + do_min("text/javascript", strings.Join(scripts, "\n")) + "</script>")
	packed_html, _ := doc.Html()
	fmt.Print("minifing index.html: " + strconv.Itoa(len(packed_html)) + " -> ")
	minified := do_min("text/html", packed_html)
	fmt.Println(len(minified))

	os.WriteFile("compiled.html", []byte(minified), 0644)

	generate_qr(minified)
}

func do_min(mediatype string, text string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	s, err := m.String(mediatype, text)
	if err != nil {
		fmt.Println(mediatype)
		panic(err)
	}

	return s
}

func generate_qr(text string) {
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))

	qr_text := "data:text/html;charset=utf-8;base64," + encodedText

	fmt.Println("qr length: " + strconv.Itoa(len(qr_text)))

	fmt.Println(qr_text)

	err := qrcode.WriteFile(qr_text, qrcode.Highest, 256, "qr.png")

	if err != nil {
		fmt.Println(err)
	}
}
