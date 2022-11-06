// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	aCtx, aCancel := chromedp.NewExecAllocator(context.Background(), []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		//chromedp.Flag("blink-settings", "imagesEnabled=false"), // 禁用图片加载
		chromedp.WindowSize(1920, 1080),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"),
		chromedp.ProxyServer("196.51.137.209:8800"),
		chromedp.ExecPath("/headless-shell/headless-shell"),
	}...)
	defer aCancel()
	ctx, cancel := chromedp.NewContext(
		aCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.amazon.co.jp`),
	)
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			chromedp.Run(ctx, chromedp.Sleep(3*time.Second))
			html := ""
			chromedp.Run(ctx, chromedp.OuterHTML(`html`, &html))
			fmt.Println(html)
			return nil
		}),
	})
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
