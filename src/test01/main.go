package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
启动
-port=8080 -target=http://127.0.0.1:9411
*/

var (
	targetURL  *url.URL
	targetAddr *string //要代理的服务的http地址

)

func printMsg(r *http.Request) error {

	//gzip压缩格式
	if r.Header.Get("Content-Encoding") == "gzip" {
		fmt.Println("Content-Encoding: " + r.Header.Get("Content-Encoding"))
		reader, e := gzip.NewReader(r.Body)
		if e != nil {
			fmt.Printf("create gzip reader err: %s \n", e.Error())
			return e
		} else {
			//ioutil.ReadAll() 读取之后会关闭流
			//这里读出之后再新建一个流然后重新赋值
			bodyBytes, e := ioutil.ReadAll(reader)
			if e != nil {
				fmt.Printf("read request body err: %s \n", e.Error())
				return e
			} else {
				if len(bodyBytes) > 0 {
					fmt.Println(string(bodyBytes))
				}
			}

			//需要将数据再以gzip格式压缩，否则和header中的content-length对不上
			var zBuf bytes.Buffer
			zw := gzip.NewWriter(&zBuf)
			if _, err := zw.Write(bodyBytes); err != nil {
				return err
			}
			zw.Close()

			r.Body = ioutil.NopCloser(&zBuf)
		}
	} else {
		//其它压缩格式
		bodyBytes, e := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		if e != nil {
			fmt.Printf("read request body err: %s \n", e.Error())
			return e
		} else {
			if len(bodyBytes) > 0 {
				fmt.Println(string(bodyBytes))
			}
		}
	}
	return nil
}

func proxy(w http.ResponseWriter, r *http.Request) {
	if e := printMsg(r); e != nil {
		_, _ = fmt.Fprint(w, e.Error())
		return
	}

	o := new(http.Request)
	*o = *r

	o.Host = targetURL.Host
	o.URL.Scheme = targetURL.Scheme
	o.URL.Host = targetURL.Host
	o.URL.Path = r.URL.Path
	o.URL.RawQuery = r.URL.RawQuery
	o.Proto = r.Proto
	o.ProtoMajor = 1
	o.ProtoMinor = 1
	o.Close = false

	res, err := http.DefaultTransport.RoundTrip(o)
	if err != nil {
		fmt.Printf("http: proxy error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	hdr := w.Header()
	for k, vv := range res.Header {
		for _, v := range vv {
			hdr.Add(k, v)
		}
	}
	for _, c := range res.Cookies() {
		w.Header().Add("Set-Cookie", c.Raw)
	}

	w.WriteHeader(res.StatusCode)
	if res.Body != nil {
		_, _ = io.Copy(w, res.Body)
	}
}

func main() {

	//代理自身的端口号
	port := flag.Int("port", 0, "Http proxy liston port")
	//后端服务的地址
	targetAddr = flag.String("target", "", "The server address. Format: http://ip:port")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	fmt.Println("proxy addr: " + addr)
	fmt.Println("target addr: " + *targetAddr)

	if *port <= 0 {
		panic("proxy port is invalid!")
	}

	if len(*targetAddr) == 0 {
		panic("target addr is empty! Format: http://ip:port")
	}

	var e error
	targetURL, e = url.Parse(*targetAddr)
	if e != nil {
		panic(e.Error())
	}

	http.HandleFunc("/", proxy)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err.Error())
	}
}
