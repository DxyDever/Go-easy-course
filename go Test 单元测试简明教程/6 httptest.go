package go_Test_单元测试简明教程

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)

	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")

	handleError(t, err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world , but get", string(body))
	}
}

/*
net.Listen("tcp", "127.0.0.1:0")：监听一个未被占用的端口，并返回 Listener。
调用 http.Serve(ln, nil) 启动 http 服务。
使用 http.Get 发起一个 Get 请求，检查返回值是否正确。
尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景。


*/
