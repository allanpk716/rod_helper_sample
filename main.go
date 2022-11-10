package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/WQGroup/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/pkg/errors"
	"gopkg.in/elazarl/goproxy.v1"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	go startHttpServer()
	go startLocalHttpProxyServer()
	// 等待服务器启动
	time.Sleep(2 * time.Second)

	//nowLancher := launcher.New().Headless(false)
	//purl := nowLancher.MustLaunch()
	//browser := rod.New().ControlURL(purl).MustConnect()
	//
	//logger.Infoln("LoadBody == true")
	//testProcessor(browser, true)
	//
	//logger.Infoln("LoadBody == false")
	//testProcessor(browser, false)

	select {}

	logger.Infoln("All Done")
}

func testProcessor(browser *rod.Browser, loadBody bool) {

	var err error
	{
		// 加载远程的，测试函数正常能够执行
		// 加载一个测试页面，不使用代理
		logger.Infoln("Try Load Remote Page Without Proxy")
		err = loadPage(browser, true)
		if err != nil {
			logger.Errorln(err)
		}
		logger.Infoln("Load Remote Page Without Proxy Success")
		// 加载一个测试页面，使用代理
		logger.Infoln("Try Load Remote Page With Proxy")
		err = loadPageWithProxy(browser, true, loadBody)
		if err != nil {
			logger.Errorln(err)
		}
		logger.Infoln("Load Remote Page With Proxy Success")
	}
	{
		// 加载本地的，且使用代理
		// 加载一个测试页面，不使用代理
		logger.Infoln("Try Load Local Page Without Proxy")
		err = loadPage(browser, false)
		if err != nil {
			logger.Errorln(err)
		}
		logger.Infoln("Load Local Page Without Proxy Success")
		// 加载一个测试页面，使用代理
		logger.Infoln("Try Load Local Page With Proxy")
		err = loadPageWithProxy(browser, false, loadBody)
		if err != nil {
			logger.Errorln(err)
		}
		logger.Infoln("Load Local Page With Proxy Success")
	}
}

// loadPage 加载一个测试页面，不使用代理
func loadPage(browser *rod.Browser, remoteOrLocalUrl bool) error {

	targetUrl := ""
	if remoteOrLocalUrl == true {
		targetUrl = remoteTargetUrl
	} else {
		targetUrl = localTargetUrl
	}

	page, e, err := NewPageNavigate(browser, targetUrl, 15*time.Second)
	defer func() {
		if page != nil {
			_ = page.Close()
		}
	}()
	if err != nil {
		// 这里可能会出现超时，但是实际上是成功的，所以这里不需要返回错误
		if errors.Is(err, context.DeadlineExceeded) == false {
			// 不是超时错误，那么就返回错误，跳过
			return err
		}
	}

	if e != nil && e.Response != nil {
		logger.Infoln(targetUrl, "No Proxy", "Response Status Code", e.Response.Status)
	}

	return nil
}

// loadPageWithProxy 加载一个测试页面，使用代理
func loadPageWithProxy(browser *rod.Browser, remoteOrLocalUrl, loadBody bool) error {

	targetUrl := ""
	if remoteOrLocalUrl == true {
		targetUrl = remoteTargetUrl
	} else {
		targetUrl = localTargetUrl

	}
	page, e, err := NewPageNavigateWithProxy(browser, localHttpProxyUrl, targetUrl, 15*time.Second, loadBody)
	defer func() {
		if page != nil {
			_ = page.Close()
		}
	}()
	if err != nil {
		// 这里可能会出现超时，但是实际上是成功的，所以这里不需要返回错误
		if errors.Is(err, context.DeadlineExceeded) == false {
			// 不是超时错误，那么就返回错误，跳过
			return err
		}
	}

	if e != nil && e.Response != nil {
		logger.Infoln(targetUrl, "With Proxy", "Response Status Code", e.Response.Status)
	}

	return err
}

func startHttpServer() {
	var srv *http.Server
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	engine := gin.Default()
	// 如果是 DebugMode 那么开启性能监控
	engine.GET("/test_page", func(c *gin.Context) {

		c.JSON(403, gin.H{"message": "Forbidden"})
	})
	engine.GET("/wait_long_time", func(c *gin.Context) {

		time.Sleep(120 * time.Second)
		c.JSON(200, gin.H{"message": "haha"})
	})
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", 19101),
		Handler: engine,
	}

	logger.Infoln("Try Start Http Server At Port", 19101)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Panicln("Start Server Error:", err)
	}
}

func startLocalHttpProxyServer() {

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	logger.Infoln("Try Start Http Proxy Server At Port", 19102)
	logger.Panicln(http.ListenAndServe("0.0.0.0:19102", proxy))
}

func NewPageNavigate(browser *rod.Browser, desURL string, timeOut time.Duration) (*rod.Page, *proto.NetworkResponseReceived, error) {

	page, err := newPage(browser)
	if err != nil {
		return nil, nil, err
	}

	return PageNavigate(page, desURL, timeOut)
}

func PageNavigate(page *rod.Page, desURL string, timeOut time.Duration) (*rod.Page, *proto.NetworkResponseReceived, error) {

	var e proto.NetworkResponseReceived
	wait := page.WaitEvent(&e)
	err := rod.Try(func() {
		page.Timeout(timeOut).MustNavigate(desURL).MustWaitLoad()
		wait()
	})
	if err != nil {
		return page, &e, err
	}
	if page == nil {
		return nil, nil, errors.New("page is nil")
	}

	return page, &e, nil
}

func NewPageNavigateWithProxy(browser *rod.Browser, proxyUrl string, desURL string, timeOut time.Duration, loadBody bool) (*rod.Page, *proto.NetworkResponseReceived, error) {

	page, err := newPage(browser)
	if err != nil {
		return nil, nil, err
	}

	return PageNavigateWithProxy(page, proxyUrl, desURL, timeOut, loadBody)
}

func PageNavigateWithProxy(page *rod.Page, proxyUrl string, desURL string, timeOut time.Duration, loadBody bool) (*rod.Page, *proto.NetworkResponseReceived, error) {

	router := page.HijackRequests()
	defer router.Stop()

	router.MustAdd("*", func(ctx *rod.Hijack) {
		px, _ := url.Parse(proxyUrl)
		err := ctx.LoadResponse(&http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyURL(px),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}, loadBody)
		if err != nil {
			return
		}
	})
	go router.Run()

	var e proto.NetworkResponseReceived
	wait := page.WaitEvent(&e)
	err := rod.Try(func() {
		page.Timeout(timeOut).MustNavigate(desURL).MustWaitLoad()
		wait()
	})
	if err != nil {
		return page, &e, err
	}
	if page == nil {
		return nil, nil, errors.New("page is nil")
	}

	return page, &e, nil
}

func newPage(browser *rod.Browser) (*rod.Page, error) {
	page, err := browser.Page(proto.TargetCreateTarget{URL: ""})
	if err != nil {
		return nil, err
	}
	return page, err
}

const localHttpProxyUrl = "http://127.0.0.1:19102"
const remoteTargetUrl = "https://baidu.com"
const localTargetUrl = "http://127.0.0.1:19101/test_page"
