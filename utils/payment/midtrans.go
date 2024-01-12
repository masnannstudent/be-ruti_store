package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"ruti-store/config"
)

func InitSnapMidtrans(config config.Config) snap.Client {
	var snapClient snap.Client
	snapClient.New(config.ServerKey, midtrans.Sandbox)
	return snapClient
}

func InitCoreMidtrans(config config.Config) coreapi.Client {
	var coreClient coreapi.Client
	coreClient.New(config.ServerKey, midtrans.Sandbox)
	return coreClient
}
