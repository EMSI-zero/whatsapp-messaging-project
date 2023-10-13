package services

type ServiceIniterFunc func() error

var ServiceIniters []ServiceIniterFunc

func InitServices() error {
	for _, service := range ServiceIniters {
		if err := service(); err != nil {
			return err
		}
	}

	return nil
}

func RegisterService(service ServiceIniterFunc) {
	ServiceIniters = append(ServiceIniters, service)
}
