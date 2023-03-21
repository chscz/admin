package main

import (
	"admin/internal/route"
)

func main() {
	// Gin 라우터 생성

	router := route.SetRouter()
	// 서버 시작
	router.Run(":8080")
}
