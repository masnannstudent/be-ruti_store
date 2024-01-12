package payment

import (
	"github.com/midtrans/midtrans-go/coreapi"
	"ruti-store/module/feature/order/domain"
)

func TransactionStatus(transactionStatusResp *coreapi.TransactionStatusResponse) domain.Status {

	var status domain.Status
	if transactionStatusResp.TransactionStatus == "capture" {
		if transactionStatusResp.FraudStatus == "challenge" {
			status.PaymentStatus = "challenge"
			status.OrderStatus = "challenge"
		} else if transactionStatusResp.FraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			status.PaymentStatus = "Konfirmasi"
			status.OrderStatus = "Proses"
		}
	} else if transactionStatusResp.TransactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		status.PaymentStatus = "Konfirmasi"
		status.OrderStatus = "Proses"
	} else if transactionStatusResp.TransactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
	} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		status.PaymentStatus = "Gagal"
		status.OrderStatus = "Gagal"
	} else if transactionStatusResp.TransactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		status.PaymentStatus = "Menunggu Konfirmasi"
		status.OrderStatus = "Menunggu Konfirmasi"
	}

	return status
}
