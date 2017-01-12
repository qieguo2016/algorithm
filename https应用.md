**https应用中的一些注意点**

1. 首次http请求的重定向

用户直接在浏览器地址栏输入网址时，浏览器默认会使用http进行访问；另外你的链接被别人分享的时候，也很有可能被写成http网址，
所以很多服务器都在http页面上使用30x重定向跳转来使用https协议访问。但是如果你的http页面被劫持了，那么用户根本到不了你的https页面，
也就意味着你精细准备的https根本派不上用场就被人截胡了。这时候你可以启动HSTS(HTTP Strict Transport Security)来让浏览器强制使用https访问。
HSTS可以让浏览器帮你做30x跳转，省一次HTTP请求。另外，浏览器本地替换可以保证只会发送HTTPS请求，避免被劫持。

要使用HSTS，只需要在你的HTTPS网站响应头中，加入下面这行：

> strict-transport-security: max-age=16070400; includeSubDomains

includeSubDomains是可选的，用来指定是否作用于子域名。支持HSTS的浏览器遇到这个响应头，会把当前网站加入HSTS列表，
然后在max-age指定的秒数内，当前网站所有请求都会被重定向为https。即使用户主动输入http://或者不输入协议部分，都将重定向到https://地址。

2. 全站资源https化

决定要使用https之时，往往需要全站升级到https，避免出现Mixed Content的情况（在https站点中加载http资源）。因为加载的http资源如果是不安全的，那么启用https也没意义了。
另一方面，浏览器也越来越严格，Mixed Content资源在现代浏览器中可能已经加载不了了（尤其是js这种危险系数极高的资源）。
如果站点比较大，在全站往HTTPS迁移的过程中，工作量往往非常巨大，尤其是将所有资源都替换为 HTTPS 这一步，很容易产生疏漏。
这时候可以通过 upgrade-insecure-requests 这个 CSP 指令，可以让浏览器帮忙做这个转换。启用这个策略后，有两个变化：

* 页面所有 HTTP 资源，会被替换为 HTTPS 地址再发起请求；
* 页面所有站内链接，点击后会被替换为 HTTPS 地址再跳转；
