// Code generated by "./generator ./session/org.deepin.dde.xsettings1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package xsettings1

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type XSettings interface {
	xSettings // interface org.deepin.dde.XSettings1
	proxy.Object
}

type objectXSettings struct {
	interfaceXSettings // interface org.deepin.dde.XSettings1
	proxy.ImplObject
}

func NewXSettings(conn *dbus.Conn) XSettings {
	obj := new(objectXSettings)
	obj.ImplObject.Init_(conn, "org.deepin.dde.XSettings1", "/org/deepin/dde/XSettings1")
	return obj
}

type xSettings interface {
	GoGetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetColor(flags dbus.Flags, arg0 string) ([]int16, error)
	GoGetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetInteger(flags dbus.Flags, arg0 string) (int32, error)
	GoGetScaleFactor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetScaleFactor(flags dbus.Flags) (float64, error)
	GoGetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetScreenScaleFactors(flags dbus.Flags) (map[string]float64, error)
	GoGetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetString(flags dbus.Flags, arg0 string) (string, error)
	GoListProps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListProps(flags dbus.Flags) (string, error)
	GoSetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call
	SetColor(flags dbus.Flags, arg0 string, arg1 []int16) error
	GoSetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call
	SetInteger(flags dbus.Flags, arg0 string, arg1 int32) error
	GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call
	SetScaleFactor(flags dbus.Flags, arg0 float64) error
	GoSetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]float64) *dbus.Call
	SetScreenScaleFactors(flags dbus.Flags, arg0 map[string]float64) error
	GoSetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call
	SetString(flags dbus.Flags, arg0 string, arg1 string) error
	ConnectSetScaleFactorStarted(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectSetScaleFactorDone(cb func()) (dbusutil.SignalHandlerId, error)
}

type interfaceXSettings struct{}

func (v *interfaceXSettings) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceXSettings) GetInterfaceName_() string {
	return "org.deepin.dde.XSettings1"
}

// method GetColor

func (v *interfaceXSettings) GoGetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetColor", flags, ch, arg0)
}

func (*interfaceXSettings) StoreGetColor(call *dbus.Call) (arg1 []int16, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceXSettings) GetColor(flags dbus.Flags, arg0 string) ([]int16, error) {
	return v.StoreGetColor(
		<-v.GoGetColor(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetInteger

func (v *interfaceXSettings) GoGetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetInteger", flags, ch, arg0)
}

func (*interfaceXSettings) StoreGetInteger(call *dbus.Call) (arg1 int32, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceXSettings) GetInteger(flags dbus.Flags, arg0 string) (int32, error) {
	return v.StoreGetInteger(
		<-v.GoGetInteger(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetScaleFactor

func (v *interfaceXSettings) GoGetScaleFactor(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScaleFactor", flags, ch)
}

func (*interfaceXSettings) StoreGetScaleFactor(call *dbus.Call) (arg0 float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceXSettings) GetScaleFactor(flags dbus.Flags) (float64, error) {
	return v.StoreGetScaleFactor(
		<-v.GoGetScaleFactor(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetScreenScaleFactors

func (v *interfaceXSettings) GoGetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetScreenScaleFactors", flags, ch)
}

func (*interfaceXSettings) StoreGetScreenScaleFactors(call *dbus.Call) (arg0 map[string]float64, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceXSettings) GetScreenScaleFactors(flags dbus.Flags) (map[string]float64, error) {
	return v.StoreGetScreenScaleFactors(
		<-v.GoGetScreenScaleFactors(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetString

func (v *interfaceXSettings) GoGetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetString", flags, ch, arg0)
}

func (*interfaceXSettings) StoreGetString(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceXSettings) GetString(flags dbus.Flags, arg0 string) (string, error) {
	return v.StoreGetString(
		<-v.GoGetString(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListProps

func (v *interfaceXSettings) GoListProps(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListProps", flags, ch)
}

func (*interfaceXSettings) StoreListProps(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceXSettings) ListProps(flags dbus.Flags) (string, error) {
	return v.StoreListProps(
		<-v.GoListProps(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetColor

func (v *interfaceXSettings) GoSetColor(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 []int16) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetColor", flags, ch, arg0, arg1)
}

func (v *interfaceXSettings) SetColor(flags dbus.Flags, arg0 string, arg1 []int16) error {
	return (<-v.GoSetColor(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetInteger

func (v *interfaceXSettings) GoSetInteger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetInteger", flags, ch, arg0, arg1)
}

func (v *interfaceXSettings) SetInteger(flags dbus.Flags, arg0 string, arg1 int32) error {
	return (<-v.GoSetInteger(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetScaleFactor

func (v *interfaceXSettings) GoSetScaleFactor(flags dbus.Flags, ch chan *dbus.Call, arg0 float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScaleFactor", flags, ch, arg0)
}

func (v *interfaceXSettings) SetScaleFactor(flags dbus.Flags, arg0 float64) error {
	return (<-v.GoSetScaleFactor(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetScreenScaleFactors

func (v *interfaceXSettings) GoSetScreenScaleFactors(flags dbus.Flags, ch chan *dbus.Call, arg0 map[string]float64) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetScreenScaleFactors", flags, ch, arg0)
}

func (v *interfaceXSettings) SetScreenScaleFactors(flags dbus.Flags, arg0 map[string]float64) error {
	return (<-v.GoSetScreenScaleFactors(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetString

func (v *interfaceXSettings) GoSetString(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetString", flags, ch, arg0, arg1)
}

func (v *interfaceXSettings) SetString(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoSetString(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// signal SetScaleFactorStarted

func (v *interfaceXSettings) ConnectSetScaleFactorStarted(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SetScaleFactorStarted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SetScaleFactorStarted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SetScaleFactorDone

func (v *interfaceXSettings) ConnectSetScaleFactorDone(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SetScaleFactorDone", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SetScaleFactorDone",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
