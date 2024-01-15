package shipping

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"ruti-store/config"
)

type RajaOngkirShippingService struct {
	apiKey string
}

func NewShippingService() ShippingServiceInterface {
	return &RajaOngkirShippingService{
		apiKey: config.InitConfig().OngkirKey,
	}
}

func (s *RajaOngkirShippingService) GetAllShippingCost(request RajaOngkirRequest) (map[string]interface{}, error) {
	url := "https://api.rajaongkir.com/starter/cost"
	apiKey := config.InitConfig().OngkirKey
	couriers := []string{"jne", "pos", "tiki"}
	allResults := make(map[string]interface{})

	for _, courier := range couriers {
		requestData := map[string]interface{}{
			"origin":      "256",
			"destination": request.Destination,
			"weight":      request.Weight,
			"courier":     courier,
		}

		requestDataJSON, err := json.Marshal(requestData)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestDataJSON))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		allResults[courier] = result
	}

	return allResults, nil
}

func (s *RajaOngkirShippingService) GetProvince() (map[string]interface{}, error) {
	url := "https://api.rajaongkir.com/starter/province"
	apiKey := config.InitConfig().OngkirKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *RajaOngkirShippingService) GetCity(province string) (map[string]interface{}, error) {
	url := "https://api.rajaongkir.com/starter/city"
	apiKey := config.InitConfig().OngkirKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("province", province)
	req.URL.RawQuery = query.Encode()

	req.Header.Add("key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
