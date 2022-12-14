package cards

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusId int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	//create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	msg := ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your card is expired"
	default:
		msg = "Your card was declined"
	}

	//test card number : 4242 4242 4242 4242
	//test card date : 12/34
	//test card cvv : 567

	return msg

}
