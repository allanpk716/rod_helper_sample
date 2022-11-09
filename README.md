# 测试项目介绍

本项目的测试代码是根据 [memory leak by use HijackRequests() · Issue #748 · go-rod/rod (github.com)](https://github.com/go-rod/rod/issues/748) 的要求实现的。

那么这个 issues 提到了 `hijack` 的 `LoadResponse` 设置从原有的 `true` 改为 `false`，遇到访问 403 页面不会马上返回的问题。因为使用了我自己的封装库 [allanpk716/rod_helper: go-rod 的封装，适用于爬虫任务 (github.com)](https://github.com/allanpk716/rod_helper)

* `LoadResponse` 设置 `true` ，对应的库版本是 `v0.0.38`
* `LoadResponse` 设置 `false` ，对应的库版本是 `v0.0.39`

## 还原依赖

请先克隆本项目，然后在本项目的目录中执行

```go
go mod tidy
```

## 测试的流程

找到 `main.go` 文件，执行 `go run main.go` 即可开始测试。

完整日志如下：

```shell
2022-11-09 14:46:11 - [INFO]: Try Start Http Proxy Server At Port 19102
2022-11-09 14:46:11 - [INFO]: Try Start Http Server At Port 19101
2022-11-09 14:46:15 - [INFO]: LoadBody == true
2022-11-09 14:46:15 - [INFO]: Try Load Remote Page Without Proxy
2022-11-09 14:46:17 - [INFO]: https://baidu.com No Proxy Response Status Code 200
2022-11-09 14:46:17 - [INFO]: Load Remote Page Without Proxy Success
2022-11-09 14:46:17 - [INFO]: Try Load Remote Page With Proxy       
2022/11/09 14:46:18 [002] WARN: Error copying to client: readfrom tcp 127.0.0.1:19102->127.0.0.1:50573: write tcp 127.0.0.1:19102->127.0.0.1:50573: wsasend: An existing connection was forcibly closed by the remote host.
2022/11/09 14:46:18 [002] WARN: Error copying to client: readfrom tcp 10.2.0.1:50574->14.215.177.38:443: read tcp 127.0.0.1:19102->127.0.0.1:50573: wsarecv: An existing connection was forcibly closed by the remote host.
2022-11-09 14:46:32 - [ERROR]: error value: context.deadlineExceededError{}                    
goroutine 1 [running]:                                                                         
runtime/debug.Stack()                                                                          
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65                      
github.com/go-rod/rod.Try.func1()                                                              
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45         
panic({0x9f3440, 0x110aa00})                                                                   
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207                          
github.com/go-rod/rod/lib/utils.glob..func2({0x9f3440?, 0x110aa00?})                           
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc00068cb90?, 0xa88e6b?, 0x11?})                            
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a           
github.com/go-rod/rod.(*Page).MustNavigate(0xc00055a160, {0xa88e6b?, 0xc000608e40?})           
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:226 +0x8b          
main.PageNavigateWithProxy.func2()                                                             
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:230 +0x4b     
github.com/go-rod/rod.Try(0xc00055a000?)                                                       
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:217 +0x62         
main.PageNavigateWithProxy(0xc00055a000, {0xa8da10, 0x16}, {0xa88e6b, 0x11}, 0x37e11d600, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:229 +0x2c5    
main.NewPageNavigateWithProxy(0x0?, {0xa8da10, 0x16}, {0xa88e6b, 0x11}, 0xc000494000?, 0xc0?)  
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:205 +0xb7     
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)                                              
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c     
main.testProcessor(0xc000140380?, 0x4?)                                                        
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:54 +0x172     
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c
main.testProcessor(0xc000140380?, 0x4?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:71 +0x31a
main.main()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:32 +0xf4

2022-11-09 14:46:48 - [INFO]: Load Local Page With Proxy Success
2022-11-09 14:46:48 - [INFO]: LoadBody == false
2022-11-09 14:46:48 - [INFO]: Try Load Remote Page Without Proxy
2022-11-09 14:46:49 - [INFO]: https://baidu.com No Proxy Response Status Code 200
2022-11-09 14:46:49 - [INFO]: Load Remote Page Without Proxy Success
2022-11-09 14:46:49 - [INFO]: Try Load Remote Page With Proxy
2022/11/09 14:46:50 [005] WARN: Error copying to client: readfrom tcp 10.2.0.1:50634->14.215.177.38:443: read tcp 127.0.0.1:19102->127.0.0.1:50633: wsarecv: An existing connection was forcibly closed by the remote host.
2022/11/09 14:46:50 [005] WARN: Error copying to client: readfrom tcp 127.0.0.1:19102->127.0.0.1:50633: write tcp 127.0.0.1:19102->127.0.0.1:50633: wsasend: An existing connection was forcibly closed by the remote host.
2022-11-09 14:47:04 - [ERROR]: error value: context.deadlineExceededError{}           
goroutine 1 [running]:                                                                
runtime/debug.Stack()                                                                 
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65             
github.com/go-rod/rod.Try.func1()                                                     
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45
panic({0x9f3440, 0x110aa00})                                                          
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207                 
github.com/go-rod/rod/lib/utils.glob..func2({0x9f3440?, 0x110aa00?})                  
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc00068d720?, 0xa88e6b?, 0x11?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a
(base) PS C:\WorkSpace\Go2Hell\src\github.com\allanpk716\rod_helper_sample> go run main.go
2022-11-09 14:47:25 - [INFO]: Try Start Http Proxy Server At Port 19102
2022-11-09 14:47:25 - [INFO]: Try Start Http Server At Port 19101
2022-11-09 14:47:29 - [INFO]: LoadBody == true
2022-11-09 14:47:29 - [INFO]: Try Load Remote Page Without Proxy
2022-11-09 14:47:31 - [INFO]: https://baidu.com No Proxy Response Status Code 200
2022-11-09 14:47:31 - [INFO]: Load Remote Page Without Proxy Success
2022-11-09 14:47:31 - [INFO]: Try Load Remote Page With Proxy       
2022-11-09 14:47:33 - [INFO]: https://baidu.com With Proxy Response Status Code 200
2022-11-09 14:47:33 - [INFO]: Load Remote Page With Proxy Success
2022-11-09 14:47:33 - [INFO]: Try Load Local Page Without Proxy  
2022/11/09 14:47:33 [046] WARN: Error copying to client: readfrom tcp 10.2.0.1:50841->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50838: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [050] WARN: Error copying to client: readfrom tcp 10.2.0.1:50847->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50846: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [047] WARN: Error copying to client: readfrom tcp 10.2.0.1:50843->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50840: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [049] WARN: Error copying to client: readfrom tcp 10.2.0.1:50848->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50844: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [052] WARN: Error copying to client: readfrom tcp 10.2.0.1:50852->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50851: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [045] WARN: Error copying to client: readfrom tcp 10.2.0.1:50839->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50837: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [051] WARN: Error copying to client: readfrom tcp 10.2.0.1:50850->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50849: wsarecv: An established connection was aborted by the software in your host machine.
2022/11/09 14:47:33 [048] WARN: Error copying to client: readfrom tcp 10.2.0.1:50845->60.188.66.33:443: read tcp 127.0.0.1:19102->127.0.0.1:50842: wsarecv: An established connection was aborted by the software in your host machine.
2022-11-09 14:47:34 - [INFO]: http://127.0.0.1:19101/test_page No Proxy Response Status Code 403
2022-11-09 14:47:34 - [INFO]: Load Local Page Without Proxy Success
2022-11-09 14:47:34 - [INFO]: Try Load Local Page With Proxy       
2022-11-09 14:47:34 - [INFO]: http://127.0.0.1:19101/test_page With Proxy Response Status Code 403
2022-11-09 14:47:35 - [INFO]: Load Local Page With Proxy Success
2022-11-09 14:47:35 - [INFO]: LoadBody == false
2022-11-09 14:47:35 - [INFO]: Try Load Remote Page Without Proxy
2022-11-09 14:47:36 - [INFO]: https://baidu.com No Proxy Response Status Code 200
2022-11-09 14:47:36 - [INFO]: Load Remote Page Without Proxy Success
2022-11-09 14:47:36 - [INFO]: Try Load Remote Page With Proxy
2022/11/09 14:47:37 [055] WARN: Error copying to client: readfrom tcp 127.0.0.1:19102->127.0.0.1:50868: write tcp 127.0.0.1:19102->127.0.0.1:50868: wsasend: An existing connection was forcibly closed by the remote host.
2022/11/09 14:47:37 [055] WARN: Error copying to client: readfrom tcp 10.2.0.1:50869->14.215.177.38:443: read tcp 127.0.0.1:19102->127.0.0.1:50868: wsarecv: An existing connection was forcibly closed by the remote host.
2022-11-09 14:47:51 - [ERROR]: error value: context.deadlineExceededError{}
goroutine 1 [running]:
runtime/debug.Stack()
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65
github.com/go-rod/rod.Try.func1()
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45
panic({0xb833c0, 0x129aa00})
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207
github.com/go-rod/rod/lib/utils.glob..func2({0xb833c0?, 0x129aa00?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc000238310?, 0xc18e8b?, 0x11?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a
github.com/go-rod/rod.(*Page).MustNavigate(0xc00014c000, {0xc18e8b?, 0xc00048cba0?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:226 +0x8b
main.PageNavigateWithProxy.func2()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:230 +0x4b
github.com/go-rod/rod.Try(0xc00014a0b0?)
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:217 +0x62
main.PageNavigateWithProxy(0xc00014a0b0, {0xc1da30, 0x16}, {0xc18e8b, 0x11}, 0x37e11d600, 0x0)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:229 +0x2db
main.NewPageNavigateWithProxy(0x0?, {0xc1da30, 0x16}, {0xc18e8b, 0x11}, 0xc000360000?, 0x50?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:205 +0xb7
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c
main.testProcessor(0xc000140380?, 0x4?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:54 +0x172
main.main()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:35 +0x13c

2022-11-09 14:47:51 - [INFO]: Load Remote Page With Proxy Success
2022-11-09 14:47:51 - [INFO]: Try Load Local Page Without Proxy
2022-11-09 14:47:52 - [INFO]: http://127.0.0.1:19101/test_page No Proxy Response Status Code 403
2022-11-09 14:47:52 - [INFO]: Load Local Page Without Proxy Success
2022-11-09 14:47:52 - [INFO]: Try Load Local Page With Proxy
2022-11-09 14:48:07 - [ERROR]: error value: context.deadlineExceededError{}
goroutine 1 [running]:
runtime/debug.Stack()
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65
github.com/go-rod/rod.Try.func1()
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45
panic({0xb833c0, 0x129aa00})
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207
github.com/go-rod/rod/lib/utils.glob..func2({0xb833c0?, 0x129aa00?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc0001492c0?, 0xc2812a?, 0x20?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a
github.com/go-rod/rod.(*Page).MustNavigate(0xc00014c4d0, {0xc2812a?, 0xc000642a20?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:226 +0x8b
main.PageNavigateWithProxy.func2()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:230 +0x4b
github.com/go-rod/rod.Try(0xc00014c420?)
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:217 +0x62
main.PageNavigateWithProxy(0xc00014c420, {0xc1da30, 0x16}, {0xc2812a, 0x20}, 0x37e11d600, 0x0)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:229 +0x2db
main.NewPageNavigateWithProxy(0x0?, {0xc1da30, 0x16}, {0xc2812a, 0x20}, 0xc000361c00?, 0xd0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:205 +0xb7
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c
main.testProcessor(0xc000140380?, 0x4?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:71 +0x31a
main.main()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:35 +0x13c

2022-11-09 14:48:07 - [INFO]: Load Local Page With Proxy Success
2022-11-09 14:48:07 - [INFO]: All Done

```



可以看到，如果设置`true`都是很快就返回，哪怕设置的超时时间是 15s，都是瞬间返回。

```shell
2022-11-09 14:46:15 - [INFO]: LoadBody == true
2022-11-09 14:46:15 - [INFO]: Try Load Remote Page Without Proxy
2022-11-09 14:46:17 - [INFO]: https://baidu.com No Proxy Response Status Code 200
2022-11-09 14:46:17 - [INFO]: Load Remote Page Without Proxy Success
2022-11-09 14:46:17 - [INFO]: Try Load Remote Page With Proxy       
2022/11/09 14:46:18 [002] WARN: Error copying to client: readfrom tcp 127.0.0.1:19102->127.0.0.1:50573: write tcp 127.0.0.1:19102->127.0.0.1:50573: wsasend: An existing connection was forcibly closed by the remote host.
2022/11/09 14:46:18 [002] WARN: Error copying to client: readfrom tcp 10.2.0.1:50574->14.215.177.38:443: read tcp 127.0.0.1:19102->127.0.0.1:50573: wsarecv: An existing connection was forcibly closed by the remote host.
2022-11-09 14:46:32 - [ERROR]: error value: context.deadlineExceededError{}                    
goroutine 1 [running]:                                                                         
runtime/debug.Stack()                                                                          
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65                      
github.com/go-rod/rod.Try.func1()                                                              
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45         
panic({0x9f3440, 0x110aa00})                                                                   
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207                          
github.com/go-rod/rod/lib/utils.glob..func2({0x9f3440?, 0x110aa00?})                           
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc00068cb90?, 0xa88e6b?, 0x11?})                            
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a           
github.com/go-rod/rod.(*Page).MustNavigate(0xc00055a160, {0xa88e6b?, 0xc000608e40?})           
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:226 +0x8b          
main.PageNavigateWithProxy.func2()                                                             
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:230 +0x4b     
github.com/go-rod/rod.Try(0xc00055a000?)                                                       
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:217 +0x62         
main.PageNavigateWithProxy(0xc00055a000, {0xa8da10, 0x16}, {0xa88e6b, 0x11}, 0x37e11d600, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:229 +0x2c5    
main.NewPageNavigateWithProxy(0x0?, {0xa8da10, 0x16}, {0xa88e6b, 0x11}, 0xc000494000?, 0xc0?)  
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:205 +0xb7     
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)                                              
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c     
main.testProcessor(0xc000140380?, 0x4?)                                                        
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:54 +0x172     
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c
main.testProcessor(0xc000140380?, 0x4?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:71 +0x31a
main.main()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:32 +0xf4
```

可以看到，如果设置`false`403 情况下，都需要等待超时完毕才会继续。等够了 15s。

```shell
2022-11-09 14:47:35 - [INFO]: Load Local Page With Proxy Success
2022-11-09 14:47:35 - [INFO]: LoadBody == false
2022-11-09 14:47:35 - [INFO]: Try Load Remote Page Without Proxy
2022-11-09 14:47:36 - [INFO]: https://baidu.com No Proxy Response Status Code 200
2022-11-09 14:47:36 - [INFO]: Load Remote Page Without Proxy Success
2022-11-09 14:47:36 - [INFO]: Try Load Remote Page With Proxy
2022/11/09 14:47:37 [055] WARN: Error copying to client: readfrom tcp 127.0.0.1:19102->127.0.0.1:50868: write tcp 127.0.0.1:19102->127.0.0.1:50868: wsasend: An existing connection was forcibly closed by the remote host.
2022/11/09 14:47:37 [055] WARN: Error copying to client: readfrom tcp 10.2.0.1:50869->14.215.177.38:443: read tcp 127.0.0.1:19102->127.0.0.1:50868: wsarecv: An existing connection was forcibly closed by the remote host.
2022-11-09 14:47:51 - [ERROR]: error value: context.deadlineExceededError{}
goroutine 1 [running]:
runtime/debug.Stack()
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65
github.com/go-rod/rod.Try.func1()
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45
panic({0xb833c0, 0x129aa00})
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207
github.com/go-rod/rod/lib/utils.glob..func2({0xb833c0?, 0x129aa00?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc000238310?, 0xc18e8b?, 0x11?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a
github.com/go-rod/rod.(*Page).MustNavigate(0xc00014c000, {0xc18e8b?, 0xc00048cba0?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:226 +0x8b
main.PageNavigateWithProxy.func2()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:230 +0x4b
github.com/go-rod/rod.Try(0xc00014a0b0?)
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:217 +0x62
main.PageNavigateWithProxy(0xc00014a0b0, {0xc1da30, 0x16}, {0xc18e8b, 0x11}, 0x37e11d600, 0x0)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:229 +0x2db
main.NewPageNavigateWithProxy(0x0?, {0xc1da30, 0x16}, {0xc18e8b, 0x11}, 0xc000360000?, 0x50?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:205 +0xb7
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c
main.testProcessor(0xc000140380?, 0x4?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:54 +0x172
main.main()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:35 +0x13c

2022-11-09 14:47:51 - [INFO]: Load Remote Page With Proxy Success
2022-11-09 14:47:51 - [INFO]: Try Load Local Page Without Proxy
2022-11-09 14:47:52 - [INFO]: http://127.0.0.1:19101/test_page No Proxy Response Status Code 403
2022-11-09 14:47:52 - [INFO]: Load Local Page Without Proxy Success
2022-11-09 14:47:52 - [INFO]: Try Load Local Page With Proxy
2022-11-09 14:48:07 - [ERROR]: error value: context.deadlineExceededError{}
goroutine 1 [running]:
runtime/debug.Stack()
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/debug/stack.go:24 +0x65
github.com/go-rod/rod.Try.func1()
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:213 +0x45
panic({0xb833c0, 0x129aa00})
        C:/WorkSpace/Go2Hell/go1.18.6/src/runtime/panic.go:838 +0x207
github.com/go-rod/rod/lib/utils.glob..func2({0xb833c0?, 0x129aa00?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/lib/utils/utils.go:60 +0x25
github.com/go-rod/rod.genE.func1({0xc0001492c0?, 0xc2812a?, 0x20?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:36 +0x5a
github.com/go-rod/rod.(*Page).MustNavigate(0xc00014c4d0, {0xc2812a?, 0xc000642a20?})
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/must.go:226 +0x8b
main.PageNavigateWithProxy.func2()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:230 +0x4b
github.com/go-rod/rod.Try(0xc00014c420?)
        C:/WorkSpace/Go2Hell/pkg/mod/github.com/go-rod/rod@v0.111.0/utils.go:217 +0x62
main.PageNavigateWithProxy(0xc00014c420, {0xc1da30, 0x16}, {0xc2812a, 0x20}, 0x37e11d600, 0x0)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:229 +0x2db
main.NewPageNavigateWithProxy(0x0?, {0xc1da30, 0x16}, {0xc2812a, 0x20}, 0xc000361c00?, 0xd0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:205 +0xb7
main.loadPageWithProxy(0xc000140380?, 0x4?, 0x0?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:120 +0x8c
main.testProcessor(0xc000140380?, 0x4?)
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:71 +0x31a
main.main()
        C:/WorkSpace/Go2Hell/src/github.com/allanpk716/rod_helper_sample/main.go:35 +0x13c

2022-11-09 14:48:07 - [INFO]: Load Local Page With Proxy Success
```



## 测试思路

下面会解释代码测试的思路。

首先，会开启两个 Server（日志如下）：

* 一个是模拟本地的 403 返回值的网站
* 一个是模拟本地的 http 代理服务器

```shell
2022-11-09 14:13:27 - [INFO]: Try Start Http Proxy Server At Port 19102
2022-11-09 14:13:27 - [INFO]: Try Start Http Server At Port 19101
```

然后，会先设置 loadbody = true

通过访问 `baidu.com` 测试正常网站的执行效果。

再通过访问上面开启的 本地HTTP 服务器，看执行效果。



然后，会先设置 loadbody = false

通过访问 `baidu.com` 测试正常网站的执行效果。

再通过访问上面开启的 本地HTTP 服务器，看执行效果。
