package shipping

type ShippingServiceInterface interface {
	GetAllShippingCost(request RajaOngkirRequest) (map[string]interface{}, error)
	GetProvince() (map[string]interface{}, error)
	GetCity(province string) (map[string]interface{}, error)
}
