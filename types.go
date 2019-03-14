package minox

import uuid "github.com/satori/go.uuid"

type CollectionLinks struct {
}

type Currency struct {
	ID           string  `json:"id"`
	ExchangeRate float64 `json:"exchange_rate"`
	Description  string  `json:"description"`
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

	VATOnSub bool     `json:"vat_on_sub"`
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

	LedgerAccounts struct {
		ExchangeRate        AccountIDResponse `json:"exchange_rate"`
		TransactionRounding AccountIDResponse `json:"transaction_rounding"`
		AutomaticPayments   AccountIDResponse `json:"automatic_payments"`
		DebtCollection      AccountIDResponse `json:"debt_collection"`
		ProfitLoss          AccountIDResponse `json:"profit_loss"`
	} `json:"ledger_accounts"`
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

type TransactionLineBatchPut []TransactionLinePut

type TransactionLinePut struct {
	Account struct {
		ID string `json:"id"`
	} `json:"account"`
	CostCenter struct {
		ID int `json:"id,omitempty"`
	} `json:"cost_center"`
	Credit        float64 `json:"credit"`
	Debit         float64 `json:"debit"`
	DocumentDate  *Date   `json:"document_date,omitempty"`
	DueDate       *Date   `json:"due_date,omitempty"`
	EntryNumber   int     `json:"entry_number,omitempty"`
	InvoiceNumber int     `json:"invoice_number,omitempty"`
	Journal       struct {
		ID string `json:"id"`
	} `json:"journal"`
	PaymentTerm struct {
		ID int `json:"id,omitempty"`
	} `json:"payment_term"`
	Period struct {
		ID         int `json:"id"`
		FiscalYear int `json:"fiscal_year"`
	} `json:"period"`
	Description string `json:"description"`
	VAT         struct {
		ID     int     `json:"id,omitempty"`
		Amount float64 `json:"amount,omitempty"`
	} `json:"vat"`
}

type TransactionLineBatchGet []TransactionLineGet

type TransactionLineGet struct {
	Account struct {
		ID          string `json:"id"`
		Scheme      string `json:"scheme"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"account"`
	BatchID    string `json:"batch_id"`
	CostCenter struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"cost_center"`
	CostType struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"cost_type"`
	Credit   float64 `json:"credit"`
	Currency struct {
		ID           string `json:"id"`
		ExchangeRate int    `json:"exchange_rate"`
		Description  string `json:"description"`
	} `json:"currency"`
	Debit         float64 `json:"debit"`
	DocumentDate  Date    `json:"document_date"`
	DueDate       *Date   `json:"due_date,omitempty"`
	EntryNumber   int     `json:"entry_number"`
	ID            string  `json:"id"`
	InvoiceNumber int     `json:"invoice_number"`
	Journal       struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Type        string `json:"type"`
		LineNumber  int    `json:"line_number"`
		PageNumber  int    `json:"page_number"`
	} `json:"journal"`
	Messages []struct {
		MessageCode      string `json:"message_code"`
		MessageType      string `json:"message_type"`
		Message          string `json:"message"`
		Field            string `json:"field"`
		TenantID         int    `json:"tenant_id"`
		AdministrationID int    `json:"administration_id"`
	} `json:"messages"`
	PaymentTerm struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		Search      []struct {
			ID    string `json:"id"`
			Value string `json:"value"`
		} `json:"search"`
		Days    int    `json:"days"`
		Mandate string `json:"mandate"`
		Start   string `json:"start"`
		Terms   struct {
			Type       string `json:"type"`
			Days       int    `json:"days"`
			Percentage int    `json:"percentage"`
		} `json:"terms"`
	} `json:"payment_term"`
	Period struct {
		ID          int    `json:"id"`
		FiscalYear  int    `json:"fiscal_year"`
		Description string `json:"description"`
	} `json:"period"`
	Description string    `json:"description"`
	VAT         VAT       `json:"vat"`
	DocumentID  uuid.UUID `json:"document_id"`
}

type VAT struct {
	ID          int     `json:"id"`
	Type        VATType `json:"type"`
	Percentage  int     `json:"percentage"`
	AccountID   string  `json:"account_id"`
	Description string  `json:"description"`
	Amount      int     `json:"amount"`
}

// [
// 	Af te dragen (1a, 1b),
// 	Te vorderen (5b),
// 	Verlegd binnenland - levering (1e),
// 	Verlegd binnenland - verwerving (2a),
// 	Verlegd buiten EU - levering (3a),
// 	Verlegd binnen EU - levering (3b),
// 	Verlegd buiten EU - verwerving (4a),
// 	Verlegd binnen EU - verwerving (4b),
// 	Afstandsverkopen binnen de EU (3c),
// 	Diensten binnen de EU (ICP)
// ]
type VATType string

type Journals []Journal

type Journal struct {
	LedgerAccountID int    `json:"ledger_account_id"`
	IBAN            string `json:"iban"`
	Terms           []struct {
		Days          int    `json:"days"`
		TranslationID string `json:"translation_id"`
	} `json:"terms"`
	NextInvoiceNumber        int      `json:"next_invoice_number"`
	LedgerAccountDifferences int      `json:"ledger_account_differences"`
	CheckForNegativeCash     bool     `json:"check_for_negative_cash"`
	Currency                 Currency `json:"currency"`
	FreeEntryNumber          bool     `json:"free_entry_number"`
	EntryNumber              int      `json:"entry_number"`
	Type                     string   `json:"type"`
	CheckClosingBalance      bool     `json:"check_closing_balance"`
	ID                       string   `json:"id"`
	BIC                      string   `json:"bic"`
	Description              string   `json:"description"`
	Range                    struct {
		Start                   int  `json:"start"`
		End                     int  `json:"end"`
		IsAutomaticallyAssigned bool `json:"is_automatically_assigned"`
	}
	AllowLinkingOfDocuments bool `json:"allow_linking_of_documents"`
}

type LedgerAccounts []LedgerAccount

type LedgerAccount struct {
	Aggregation struct {
		AggregationLevel int    `json:"aggregation_level"`
		Description      string `json:"description"`
	} `json:"aggregation"`
	BudgetDivision struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"budget_division"`
	YearBudget float64 `json:"year_budget"`
	VAT        VAT     `json:"vat"`
	CostType   struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"cost_type"`
	Blocked struct {
		Active bool   `json:"active"`
		Reason string `json:"reason"`
	} `json:"blocked"`
	IsBalance          bool `json:"isBalance"`
	QuantityApplicable bool `json:"quantity_applicable"`
	// Whether the ledgeraccount needs to be compressed, also known as "verdichten".
	AggregateTransactions bool                  `json:"aggregate_transactions"`
	AllowFreeVATInput     bool                  `json:"allow_free_vat_input"`
	DebitCreditSuggestion DebitCreditSuggestion `json:"debit_credit_suggestion"`
	Memo                  struct {
		Active bool   `json:"active"`
		Text   string `json:"text"`
	} `json:"memo"`
	ID      int `json:"id"`
	Balance struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		IsGroup     bool   `json:"is_group"`
		Type        string `json:"type"`
		IsBalance   bool   `json:"isBalance"`
	} `json:"balance"`
	Description string `json:"description"`
	RGS         []struct {
		ReferenceCode    string `json:"reference_code"`
		Order            string `json:"order"`
		ReferenceNumber  string `json:"reference_number"`
		Description      string `json:"description"`
		DescriptionShort string `json:"description_short"`
		DebitCredit      string `json:"debit_credit"`
		Level            int    `json:"level"`
		Filters          struct {
			ZZP   bool `json:"ZZP"`
			AFREK bool `json:"AFREK"`
		} `json:"filters"`
	} `json:"rgs"`
}

// [
// 	debet,
// 	credit,
// 	none
// ]
type DebitCreditSuggestion string
