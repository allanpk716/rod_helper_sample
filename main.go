package main

import (
	"context"
	"fmt"
	"github.com/WQGroup/logger"
	"github.com/allanpk716/rod_helper"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/pkg/errors"
	"gopkg.in/elazarl/goproxy.v1"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	go startHttpServer()
	go startLocalHttpProxyServer()
	// 等待服务器启动
	time.Sleep(2 * time.Second)

	var err error
	nowLancher := launcher.New().Headless(false)
	purl := nowLancher.MustLaunch()
	browser := rod.New().ControlURL(purl).MustConnect()

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
		err = loadPageWithProxy(browser, true)
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
		err = loadPageWithProxy(browser, false)
		if err != nil {
			logger.Errorln(err)
		}
		logger.Infoln("Load Local Page With Proxy Success")
	}

	logger.Infoln("All Done")
}

// loadPage 加载一个测试页面，不使用代理
func loadPage(browser *rod.Browser, remoteOrLocalUrl bool) error {

	targetUrl := ""
	if remoteOrLocalUrl == true {
		targetUrl = remoteTargetUrl
	} else {
		targetUrl = localTargetUrl
	}

	page, e, err := rod_helper.NewPageNavigate(browser, targetUrl, 15*time.Second)
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
func loadPageWithProxy(browser *rod.Browser, remoteOrLocalUrl bool) error {

	targetUrl := ""
	if remoteOrLocalUrl == true {
		targetUrl = remoteTargetUrl
	} else {
		targetUrl = localTargetUrl

	}
	page, e, err := rod_helper.NewPageNavigateWithProxy(browser, localHttpProxyUrl, targetUrl, 15*time.Second)
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

const localHttpProxyUrl = "http://127.0.0.1:19102"
const remoteTargetUrl = "https://baidu.com"
const localTargetUrl = "http://127.0.0.1:19101/test_page"
