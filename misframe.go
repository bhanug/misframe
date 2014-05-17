package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/feeds"
	"github.com/russross/blackfriday"
)

type Post struct {
	Title         string
	Date          time.Time
	FormattedDate string
	Url           string
	Content       string // This is parsed Markdown.
}

const (
	FileTimeForm = "2006-01-02 15:04:05"
	WebTimeForm  = "Jan 02 2006"
)

var (
	ContentDir   = "./posts"
	ReloadUrl    = "RELOAD_POSTS" // Make sure this is something weird!
	Listen       = ":8080"
	TemplateFile = "template.html"

	Posts = []*Post{}
	Urls  = map[string]*Post{}

	templ *template.Template = nil
)

func readPostFile(filename string) *Post {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Failed to open file:", filename)
		return nil
	}
	defer f.Close()

	r := bufio.NewReader(f)

	meta := map[string]string{
		"title": ".",
		"date":  ".",
		"url":   ".",
	}

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}

		parts := strings.Split(string(line), ": ")
		if len(parts) == 2 {
			if meta[parts[0]] == "." {
				meta[parts[0]] = parts[1]
			} else {
				continue
			}
		} else {
			break
		}
	}

	if meta["title"] == "." || meta["date"] == "." || meta["url"] == "." {
		return nil
	}

	created, err := time.Parse(FileTimeForm, meta["date"])
	if err != nil {
		log.Println(filename, "had an invalid date")
		return nil
	}

	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Failed to read content from", filename)
		return nil
	}

	post := Post{
		Title:         meta["title"],
		Date:          created,
		FormattedDate: created.Format(WebTimeForm),
		Url:           meta["url"],
		Content:       string(md(content)),
	}

	return &post
}

func loadPosts() {
	var err error

	templ, err = template.New("template").ParseFiles(TemplateFile)
	if err != nil {
		log.Fatal("Failed to parse template")
	}

	fileInfos, err := ioutil.ReadDir(ContentDir)
	if err != nil {
		log.Fatalf("Failed to read content directory: %v", err)
	}

	for _, fileInfo := range fileInfos {
		post := readPostFile(ContentDir + "/" + fileInfo.Name())
		if post != nil {
			Posts = append(Posts, post)
			Urls[post.Url] = post
		}
	}

	// reverse the order
	for i := 0; i < len(Posts)/2; i++ {
		tmp := Posts[i]
		Posts[i] = Posts[len(Posts)-i-1]
		Posts[len(Posts)-i-1] = tmp
	}
}

func md(input []byte) []byte {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_GITHUB_BLOCKCODE
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_HEADER_IDS

	return blackfriday.Markdown(input, renderer, extensions)
}

func main() {
	flag.StringVar(&ContentDir, "content", ContentDir, "Directory with the blog content")
	flag.StringVar(&ReloadUrl, "reload-url", ReloadUrl, "URL that reloads posts from the disk")
	flag.StringVar(&Listen, "listen", Listen, "Listen address")
	flag.StringVar(&TemplateFile, "template", TemplateFile, "Template file used to render HTML")
	flag.Parse()

	loadPosts()

	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		feed := &feeds.Feed{
			Title: "Misframe",
			Link:  &feeds.Link{Href: "http://misfra.me/"},
		}
		for i := 0; i < 5; i++ {
			feed.Items = append(feed.Items, &feeds.Item{
				Title:   Posts[i].Title,
				Link:    &feeds.Link{Href: "http://misfra.me/" + Posts[i].Url},
				Created: Posts[i].Date,
				Author:  &feeds.Author{"Preetam Jinka", ""},
			})
		}
		atom, _ := feed.ToAtom()
		fmt.Fprintln(w, atom)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimLeft(r.URL.Path, "/")
		if post, present := Urls[url]; present {
			buffer := &bytes.Buffer{}
			templ.ExecuteTemplate(buffer, "Single", post)

			content := buffer.String()
			buffer.Reset()
			templ.ExecuteTemplate(w, "Page", struct {
				Title   string
				Content string
			}{
				Title:   post.Title,
				Content: content,
			})
		} else {
			buffer := &bytes.Buffer{}
			templ.ExecuteTemplate(buffer, "All", Posts)
			content := buffer.String()
			buffer.Reset()
			templ.ExecuteTemplate(w, "Page", struct {
				Title   string
				Content string
			}{
				Title:   "Misframe",
				Content: content,
			})
		}
	})

	http.HandleFunc("/"+ReloadUrl, func(w http.ResponseWriter, r *http.Request) {
		Posts = Posts[:0]
		Urls = map[string]*Post{}

		loadPosts()

		fmt.Fprintln(w, "Updated!")
	})

	panic(http.ListenAndServe(Listen, nil))
}
