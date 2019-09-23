package wxpay

import (
	"io/ioutil"
	"log"
)

const (
	AccountTypeNormal          = "normal"
	AccountTypeServiceProvider = "service_provider"
)

type Account struct {
	acctType  string
	appID     string
	subAppId  string
	mchID     string
	subMchId  string
	apiKey    string
	certData  []byte
	isSandbox bool
}

// 创建微信支付账号
func NewAccount(appID string, mchID string, apiKey string, isSanbox bool) *Account {
	return &Account{
		acctType:  AccountTypeNormal,
		appID:     appID,
		subAppId:  "",
		mchID:     mchID,
		subMchId:  "",
		apiKey:    apiKey,
		isSandbox: isSanbox,
	}
}

// 创建特约商户微信支付账号
func NewServiceProviderAccount(appID string, subAppId string, mchID string, subMchId string, apiKey string, isSanbox bool) *Account {
	return &Account{
		acctType:  AccountTypeServiceProvider,
		appID:     appID,
		subAppId:  subAppId,
		mchID:     mchID,
		subMchId:  subMchId,
		apiKey:    apiKey,
		isSandbox: isSanbox,
	}
}

// 设置证书
func (a *Account) SetCertData(certPath string) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Println("读取证书失败")
		return
	}
	a.certData = certData
}
