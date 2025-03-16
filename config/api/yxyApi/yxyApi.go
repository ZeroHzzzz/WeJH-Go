package yxyApi

import "wejh-go/config/config"

var YxyHost = config.Config.GetString("yxy.host")

type YxyApi string

const (
	SecurityToken          YxyApi = "api/v1/login/security-token"
	SendVerificationCode   YxyApi = "api/v1/login/send-code"
	LoginByCode            YxyApi = "api/v1/login/code"
	SlientLogin            YxyApi = "api/v1/login/silent"
	CardBalance            YxyApi = "api/v1/card/balance"
	ConsumptionRecords     YxyApi = "api/v1/card/consumption-records"
	Auth                   YxyApi = "api/v1/electricity/auth"
	ElectricityBalance     YxyApi = "api/v1/electricity/surplus"
	RechargeRecords        YxyApi = "api/v1/electricity/recharge-records"
	ElectricityConsumption YxyApi = "api/v1/electricity/usage-records"
	BusAuth                YxyApi = "api/v1/bus/auth"
	BusInfo                YxyApi = "api/v1/bus/info"
	BusRecords             YxyApi = "api/v1/bus/records"
	BusQrCode              YxyApi = "api/v1/bus/qrcode"
	BusMessage             YxyApi = "api/v1/bus/message"
)
