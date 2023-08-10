package facebook

type ChromiumAccountFbHoldAds []accountFbHoldAds

type accountFbHoldAds struct {
	Id               string
	Name             string
	Coookie          string
	TotalAccountAds  int
	DetailAccountAds []adsAcount
	TotalPage        int
	DetailPage       []page
	Business         []business
}

func (c *ChromiumAccountFbHoldAds) Name() string {
	return "account_facebook_hold_ads"
}

func (c *ChromiumAccountFbHoldAds) Len() int {
	return len(*c)
}

func (c *ChromiumAccountFbHoldAds) Parse(masterKey []byte) error {
	return nil
}
