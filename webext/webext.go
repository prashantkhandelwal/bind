package webext

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/prashantkhandelwal/bind/utils"
)

const (
	UserAgentName = "Bindv1.0"
	Path          = "data\\images\\"
	Timeout       = 15
)

type WebData struct {
	Title string
	Snap  string
}

func extractSnapshot(name, url string, filename chan string) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte
	var file = Path + name + ".png"

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&buf),
	}); err != nil {
		log.Fatal(err)
		file = ""
	}

	if err := os.WriteFile(file, buf, 0o644); err != nil {
		log.Fatal(err)
		file = ""
	}

	filename <- file
}

func extractTitle(url string, title chan string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Creating timeout for 15 seconds
	ctx, cancel = context.WithTimeout(ctx, time.Second*Timeout)
	defer cancel()

	var pageTitle string

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("Bindv1.0"),
		chromedp.Navigate(url),
		chromedp.Title(&pageTitle),
	)
	if err != nil {
		log.Fatalf("ERROR:webext - Unable to extract title from the given URL - %s\n", url)
		pageTitle = ""
	}

	title <- pageTitle
}

func Extract(url string) (*WebData, error) {

	start := time.Now()
	if len(strings.TrimSpace(url)) == 0 {
		log.Fatalf("ERROR:webext - url cannot be empty.")
		return nil, errors.New("ERROR:webext - url cannot be empty")
	}

	title := make(chan string)
	name := make(chan string)

	var d = &WebData{}
	filename := utils.GenRandStr(12)

	go extractTitle(url, title)
	t := <-title

	go extractSnapshot(filename, url, name)
	n := <-name

	if len(strings.TrimSpace(t)) > 0 {
		d.Title = t
	}

	if len(strings.TrimSpace(n)) > 0 {
		d.Snap = n
	}

	close(title)
	close(name)
	elapsed := time.Since(start)
	log.Printf("Time taken to extract web data: %s", elapsed)
	return d, nil
}
