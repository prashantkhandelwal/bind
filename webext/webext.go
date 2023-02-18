package webext

import (
	"context"
	"log"
	"os"
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
	Title       string
	Description string
}

func Snap(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Creating timeout for 15 seconds
	ctx, cancel = context.WithTimeout(ctx, time.Second*Timeout)
	defer cancel()

	name := utils.GenRandStr(12)
	var buf []byte
	var file = Path + name + ".png"

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("Bindv1.0"),
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&buf),
	)

	if err != nil {
		log.Fatalf("ERROR:webext - Unable to extract meta(s) from the given URL - %s\n", url)
		return "", err
	}

	if err := os.WriteFile(file, buf, 0o644); err != nil {
		log.Fatalf("ERROR:webext - Cannot create snap for the link - %s\n", url)
		file = ""
		return "", err
	}

	return file, nil
}

func ExtractMeta(url string, webdata chan WebData) {
	start := time.Now()
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Creating timeout for 15 seconds
	ctx, cancel = context.WithTimeout(ctx, time.Second*Timeout)
	defer cancel()

	var pageTitle, description string

	var w = &WebData{}

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("Bindv1.0"),
		chromedp.Navigate(url),
		chromedp.Title(&pageTitle),
		chromedp.Evaluate(`document.querySelector("meta[name^='description' i]").getAttribute('content');`, &description),
	)

	if err != nil {
		log.Fatalf("ERROR:webext - Unable to extract meta(s) from the given URL - %s\n", url)
	}

	w.Title = pageTitle
	w.Description = description

	elapsed := time.Since(start)
	log.Printf("Time taken to fetch meta: %s", elapsed)

	webdata <- *w
}
