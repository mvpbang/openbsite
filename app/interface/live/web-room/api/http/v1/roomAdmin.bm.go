// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api/http/v1/roomAdmin.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api/http/v1/roomAdmin.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

// ===================
// RoomAdmin Interface
// ===================

// History 相关服务
type RoomAdmin interface {
	// 获取主播拥有的的所有房管, 无需登录态
	// `method:"GET"
	GetByRoom(ctx context.Context, req *RoomAdminGetByRoomReq) (resp *RoomAdminGetByRoomResp, err error)
}

var v1RoomAdminSvc RoomAdmin

// @params RoomAdminGetByRoomReq
// @router GET /xlive/web-room/v1/roomAdmin/get_by_room
// @response RoomAdminGetByRoomResp
func roomAdminGetByRoom(c *bm.Context) {
	p := new(RoomAdminGetByRoomReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomAdminSvc.GetByRoom(c, p)
	c.JSON(resp, err)
}

// RegisterV1RoomAdminService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1RoomAdminService(e *bm.Engine, svc RoomAdmin, midMap map[string]bm.HandlerFunc) {
	v1RoomAdminSvc = svc
	e.GET("/xlive/web-room/v1/roomAdmin/get_by_room", roomAdminGetByRoom)
}
