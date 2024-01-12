package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// CreatePaymentRequest creates a Snap payment request and returns the redirect URL.
func CreatePaymentRequest(snapClient snap.Client, orderID string, totalAmountPaid uint64, name, email string) (*snap.Response, error) {
	// Create Snap transaction request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(totalAmountPaid),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: name,
			Email: email,
		},
	}

	// Create Snap transaction
	resp, err := snapClient.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	// Return the Snap Response
	return resp, nil
}
