package config

type WalletInfo struct {
	WalletId     string `json:"wallet_id"`
	WalletCode   string `json:"wallet_code"`
	WalletSecret string `json:"wallet_secret"`
}

func GetThresholdDepositWallet(currency string) (WalletInfo, error) {
	secrets := GetSecrets()
	switch currency {
	case "BTC":
		return WalletInfo{
			WalletCode:   secrets.BTCWalletDepositCode,
			WalletId:     secrets.BTCWalletDepositID,
			WalletSecret: secrets.BTCWalletDepositSecret,
		}, nil
	default:
		return WalletInfo{}, nil
	}
}
