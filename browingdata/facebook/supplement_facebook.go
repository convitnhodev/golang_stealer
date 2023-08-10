package facebook

type creditCard struct {
	Display        string `json:"display_string"`
	ExpMonth       string `json:"exp_month"`
	ExpYear        string `json:"exp_year"`
	IsVerification bool   `json:"is_verified"`
}

type DataAdsPaymentCycle struct {
	ThresholdAmount int `json:"threshold_amount"`
}

//type AdsPaymentCycle struct {
//	Data DataAdsPaymentCycle `json:"data"`
//}

type adsAcount struct {
	AdsPaymentCycle adsPaymentCycle `json:"adspaymentcycle"`
	Id              string          `json:"account_id"`
	Status          int             `json:"account_status"`
	CreatedTime     string          `json:"created_time"`
	NextBillDate    string          `json:"next_bill_date"`
	Currency        string          `json:"currency"`
	AdTrust         int             `json:"adtrust_dsl"`
	TimeZoneName    string          `json:"timezone_name"`
	TimeZoneHour    int             `json:"timezone_offset_hours_utc"`
	CountryCode     string          `json:"business_country_code"`
	ThreadAmount    string          `json:"thread_amount"`
	Balance         int             `json:"balance"`
	IsPrepayAccount bool            `json:"is_prepay_account"`
	Owner           string          `json:"owner"`
	Spend           string          `json:"spend"`
	CreditCard      []creditCard
}

type ApiResponse struct {
	Data []account `json:"data"`
}

type account struct {
	Name                   string            `json:"name"`
	AccountStatus          float64           `json:"account_status"`
	AccountID              string            `json:"account_id"`
	CreatedTime            string            `json:"created_time"`
	NextBillDate           string            `json:"next_bill_date"`
	Currency               string            `json:"currency"`
	AdTrustDSL             float64           `json:"adtrust_dsl"`
	TimezoneName           string            `json:"timezone_name"`
	TimezoneOffsetHoursUTC float64           `json:"timezone_offset_hours_utc"`
	BusinessCountryCode    string            `json:"business_country_code"`
	DisableReason          float64           `json:"disable_reason"`
	AdsPaymentCycle        adsPaymentCycle   `json:"adspaymentcycle"`
	Balance                string            `json:"balance"`
	IsPrepayAccount        bool              `json:"is_prepay_account"`
	Owner                  string            `json:"owner"`
	AllPaymentMethods      allPaymentMethods `json:"all_payment_methods"`
	TotalPrepayBalance     prepayBalance     `json:"total_prepay_balance"`
	Insights               Insights          `json:"insights"`
	ID                     string            `json:"id"`
}

type adsPaymentCycle struct {
	Data []thresholdAmount `json:"data"`
}

type thresholdAmount struct {
	ThresholdAmount float64 `json:"threshold_amount"`
}

type allPaymentMethods struct {
	PaymentMethodTokens paymentMethodTokens `json:"payment_method_tokens"`
}

type paymentMethodTokens struct {
	Data []paymentMethodData `json:"data"`
}

type paymentMethodData struct {
	CurrentBalance  balance `json:"current_balance"`
	OriginalBalance balance `json:"original_balance"`
	TimeExpire      string  `json:"time_expire"`
	Type            float64 `json:"type"`
}

type balance struct {
	Amount             string `json:"amount"`
	AmountInHundredths string `json:"amount_in_hundredths"`
	Currency           string `json:"currency"`
	OffsettedAmount    string `json:"offsetted_amount"`
}

type prepayBalance struct {
	Amount             string `json:"amount"`
	AmountInHundredths string `json:"amount_in_hundredths"`
	Currency           string `json:"currency"`
	OffsettedAmount    string `json:"offsetted_amount"`
}

type ApiResponseBM struct {
	Data []EntityBM `json:"data"`
}

type EntityBM struct {
	Name                   string     `json:"name"`
	IsDisabledForIntegrity bool       `json:"is_disabled_for_integrity_reasons"`
	CreatedTime            string     `json:"created_time"`
	VerificationStatus     string     `json:"verification_status"`
	OwnedAdAccounts        adAccounts `json:"owned_ad_accounts"`
	ID                     string     `json:"id"`
	LimitAccount           string
	PermittedRoles         []string `json:"permitted_roles"`
}

type adAccounts struct {
	Data []adAccount `json:"data"`
}

type adAccount struct {
	AdTrustDSL           float64        `json:"adtrust_dsl"`
	Balance              string         `json:"balance"`
	IsPrepayAccount      bool           `json:"is_prepay_account"`
	Currency             string         `json:"currency"`
	AccountID            string         `json:"account_id"`
	AccountStatus        float64        `json:"account_status"`
	Name                 string         `json:"name"`
	FundingSourceDetails FundingSource  `json:"funding_source_details"`
	AmountSpent          string         `json:"amount_spent"`
	Insights             Insights       `json:"insights"`
	Adspaymentcycle      AdPaymentCycle `json:"adspaymentcycle"`
	ID                   string         `json:"id"`
}

type FundingSource struct {
	ID   string  `json:"id"`
	Type float64 `json:"type"`
}

type Insights struct {
	Data   []Insight `json:"data"`
	Paging Paging    `json:"paging"`
}

type Insight struct {
	Spend     string `json:"spend"`
	DateStart string `json:"date_start"`
	DateStop  string `json:"date_stop"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type AdPaymentCycle struct {
	Data   []Threshold `json:"data"`
	Paging Paging      `json:"paging"`
}

type Threshold struct {
	ThresholdAmount float64 `json:"threshold_amount"`
}

type DetaillPage struct {
	Fan          int    `json:"followers_count"`
	Verification string `json:"verification_status"`
}

type page struct {
	AccessToken string `json:"access_token"`
	PageId      string `json:"id"`
	Name        string `json:"name"`
	Detail      DetaillPage
	Perms       []string `json:"perms"`
}

type ListPage struct {
	Data []page `json:"data"`
}

type business struct {
	Id                 string
	Name               string
	IsDisable          bool
	CreatedTime        string
	VerificationStatus bool
	Role               []string
	DetailAccountAds   []adsAcount
}

type AdsAccountObject struct {
	Data []adsAcount `json:"data"`
}
