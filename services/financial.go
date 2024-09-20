package services

import (
	"os"

	"github.com/IvanARodriguez/payme/models"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/account"
)

func CreateStripeBusiness(b models.Business) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET")

	acctParams := &stripe.AccountParams{
		Type: stripe.String("express"),
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			Name: stripe.String(b.Name),
		},
		Country: stripe.String("US"),
		Email:   stripe.String(b.Email),
	}

	stripeAccount, err := account.New(acctParams)

	if err != nil {
		return "", err
	}

	return stripeAccount.ID, nil

}
