package qrng

import (
	"math/big"
)

type datatype string

const (
	u8  datatype = "uint8"
	u16 datatype = "uint16"
)

func init() {

}

func FetchSeed() (string, error) {
	result := ""
	numbers, err := queryDefaultSize(u16)
	if err != nil {
		return "", err
	}
	for _, s := range numbers {

		res := big.NewInt(int64(s)).Text(16)
		if len(res) == 3 {
			res = "0" + res
		}
		result = result + res
	}
	return result, nil
}

func FetchE2EEKey() (string, error) {
	result := ""
	numbers, err := queryDefaultSize(u8)
	if err != nil {
		return "", err
	}
	for i, s := range numbers {
		if i%4 == 0 {
			res := big.NewInt(int64(s)).Text(16)
			if len(res) == 3 {
				res = "0" + res
			}
			result = result + res
		}
	}
	return result, nil
}
