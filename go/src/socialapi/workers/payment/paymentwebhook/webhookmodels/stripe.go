package webhookmodels

type StripePlan struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Interval string `json:"interval"`
}

type StripeSubscription struct {
	ID         string     `json:"id"`
	CustomerId string     `json:"customer"`
	Plan       StripePlan `json:"plan"`
}

type StripePeriod struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

type StripeInvoiceData struct {
	SubscriptionId string       `json:"id"`
	Period         StripePeriod `json:"period"`
	Plan           StripePlan   `json:"plan"`
}

type StripeInvoiceLines struct {
	Data  []StripeInvoiceData `json:"data"`
	Count int                 `json:"count"`
}

type StripeInvoice struct {
	ID         string             `json:"id"`
	CustomerId string             `json:"customer"`
	AmountDue  float64            `json:"amount_due"`
	Currency   string             `json:"currency"`
	Lines      StripeInvoiceLines `json:"lines"`
}

type StripeCard struct {
	Id      string `json:"id"`
	ExpYear string `json:"exp_year"`
	Last4   string `json:"last4"`
	Brand   string `json:"brand"`
}

type StripeCharge struct {
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
	CustomerId string  `json:"customer"`
}

type StripeCustomer struct {
	ID string `json:"id"`
}
