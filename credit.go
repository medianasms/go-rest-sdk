package medianasms

import "encoding/json"

// getCreditResType get credit response type
type getCreditResType struct {
	Credit float64 `json:"credit"`
}

// GetCredit get credit for user
func (sms *MedianaSMS) GetCredit() (float64, error) {
	_res, err := sms.get("/credit", nil)
	if err != nil {
		return 0, err
	}

	res := &getCreditResType{}
	if err = json.Unmarshal(_res.Data, res); err != nil {
		return 0, err
	}

	return res.Credit, nil
}
