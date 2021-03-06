package catalog

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Catalog struct {
	ID          uint64 `json:"Id"`
	Name        string `json:"Name"`
	Bundle      string `json:"Bundle" gorm:"size:65532"`
	Description string `json:"Description" gorm:"size:65532"`
	IconData    string `json:"IconData" gorm:"size:65532"`
	AccountId   uint64 `json:"AccountId"`
	Type        uint8  `json:"Type"`
}

func ImageHandle(request *http.Request) (string, error) {
	var buf []byte
	var err error

	icon, _, err := request.FormFile("icon")
	if err != nil {
		return "", nil
	} else {
		buf, err = ioutil.ReadAll(icon)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("data:%s;base64,%s",
		http.DetectContentType(buf),
		base64.StdEncoding.EncodeToString(buf)), nil
}
