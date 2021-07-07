package route

import (
	"fmt"
	"src/model"
)

// Routing : 라우팅 설정
func Routing(ctx model.Context) {
	// 빌드된 프론트 앱 js, css, html 파일
	// 프론트 앱 진입점 (BASE URL)
	ctx.Echo.Static("/v", fmt.Sprintf("%v/app/www", ctx.RootDir))
	// asset
	ctx.Echo.Static("/assets", fmt.Sprintf("%v/app/www/assets", ctx.RootDir))
	// svg
	ctx.Echo.Static("/svg", fmt.Sprintf("%v/app/www/svg", ctx.RootDir))

	/*
		모두
	*/
	ctx.Echo.File("/favicon.ico/", fmt.Sprintf("%v/app/src/assets/icon/favicon.png", ctx.RootDir))
	ctx.Echo.File("/robots.txt/", fmt.Sprintf("%v/app/src/assets/text/robots.txt", ctx.RootDir))
	ctx.Echo.File("/*", fmt.Sprintf("%v/app/www/index.html", ctx.RootDir))
}
