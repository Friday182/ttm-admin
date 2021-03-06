package websocket

import (
	"net/http"
	"testing"

	"github.com/gogf/gf/net/ghttp"
)

func Benchmark(t *testing.B) {
	for i := 0; i < 100; i++ {
		r := ghttp.Request{}
		Handler(&r)
	}
}

func TestWebsocketHandler(t *testing.T) {
	//创建一个请求
	req, err := http.NewRequest("GET", "/api/ws", nil)
	if err != nil {
			t.Fatal(err)
	}

	// 我们创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
	//rr := httptest.NewRecorder()

	Handler(req)
}