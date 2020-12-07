package responses

import "github.com/SolarLabRU/fastpay-go-commons/models"

type RoutesItem struct {
	Commission *models.CurrencyContractRoutingItem
	Middle     *models.CurrencyContractRoutingItem
	Input      *models.CurrencyContractRoutingItem
}

func (ri *RoutesItem) ListContracts() models.CurrencyContractRoutingItemSlice {
	return models.CurrencyContractRoutingItemSlice{
		ri.GetInput(),
		ri.GetMiddle(),
		ri.GetCommission(),
	}
}
