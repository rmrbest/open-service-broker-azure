package cosmosdb

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (s *serviceManager) Unbind(
	_ service.Instance,
	_ service.BindingDetails,
) error {
	return nil
}
