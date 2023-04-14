package common

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	// 启动web服务
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
		if err != nil {
			fmt.Println("启动失败")
		}
		fmt.Printf("监控成功, 端口为:%d", port)
	}()
}
