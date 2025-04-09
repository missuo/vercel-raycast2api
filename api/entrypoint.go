/*
 * @Author: Vincent Yang
 * @Date: 2025-04-09 16:00:16
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2025-04-09 16:07:09
 * @FilePath: /vercel-raycast2api/api/entrypoint.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2025 by Vincent, All Rights Reserved.
 */

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	raycast2api "github.com/missuo/raycast2api/service"
)

var (
	cfg *raycast2api.Config
	app *gin.Engine
)

func init() {
	cfg = raycast2api.InitConfig()
	app = raycast2api.Router(cfg)
}

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

var _ http.HandlerFunc = Entrypoint
