package minox

type CollectionLinks struct {
}

type Currency struct {
	ID string `json:"id"`
}

type Administrations []Administration

type Administration struct {
	// id: integer
	// example: 99998
	// minimum: 1
	// maximum: 99999
	ID int `json:"id"`

	// name: string
	// example: Breekstaal B.V.
	Name string `json:"name"`

	// number_of_periods: integer
	// example: 12
	NumberOfPeriods NumberOfPeriods `json:"number_of_periods"`

	VatOnSub bool     `json:"vat_on_sub"`
	Currency Currency `json:"currency"`

	// chamber_of_commerce: string
	// example: 30070111
	ChamberOfCommerce string `json:"chamber_of_commerce"`

	// country_code: string
	// example: NL
	// An ISO-3166-1-alpha2 country code. See:
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
	CountryCode string `json:"country_code"`

	// [administration_fiscal_year]
	FiscalYears FiscalYears `json:"fiscal_years"`

	LedgerAccounts LedgerAccounts `json:"ledger_accounts"`
}

// Enum: [ 1, 4, 12, 13 ]
type NumberOfPeriods int

type FiscalYears []FiscalYear

type FiscalYear struct {
	// number_of_periods: integer
	// example: 12
	NumberOfPeriods NumberOfPeriods `json:"number_of_periods"`

	// fiscal_year: integer
	// example: 2018
	FiscalYear int `json:"fiscal_year"`

	// start_date: string ($date)
	// example: 2018-01-01
	StartDate Date `json:"start_date"`

	// end_date: string ($date)
	// example: 2018-12-31
	EndDate Date `json:"end_date"`
}

type LedgerAccounts struct {
	ExchangeRate        AccountIDResponse `json:"exchange_rate"`
	TransactionRounding AccountIDResponse `json:"transaction_rounding"`
	AutomaticPayments   AccountIDResponse `json:"automatic_payments"`
	DebtCollection      AccountIDResponse `json:"debt_collection"`
	ProfitLoss          AccountIDResponse `json:"profit_loss"`
}

type AccountIDResponse struct {
	// id: account_idstring *
	// example: 4010
	// Value range: 1 to 999999. Must be a valid id from a general ledger
	// account, customer or supplier.
	ID string `json:"id"`

	// scheme: account_scheme string *
	Scheme AccountScheme `json:"scheme"`

	// type: account_type string *
	Type AccountType `json:"type"`

	// description: description string
	// Descriptive text defined by the user or the system.
	Description string `json:"description"`
}

// Enum: [ MINOX ]
type AccountScheme string

// Enum: [ ledger_account, supplier, customer ]
type AccountType string
