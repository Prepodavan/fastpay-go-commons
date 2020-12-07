package requests

import "fmt"

type SetAvailablePlatformsRequest struct {
	TechnicalSignRequest
	AvailablePlatforms string `yaml:"availablePlatforms" json:"availablePlatforms" valid:"required"`
}

func (request *SetAvailablePlatformsRequest) String() string {
	return fmt.Sprintf("setAvailablePlatforms%s", request.AvailablePlatforms)
}
