package accessory

import (
	"github.com/gumpyoung/hc/service"
)

type Doorbell struct {
	*Accessory
	Doorbell *service.Doorbell
}

// NewDoorbell returns a doorbell which implements model.Doorbell.
func NewDoorbell(info Info) *Doorbell {
	acc := Doorbell{}
	acc.Accessory = New(info, TypeVideoDoorbell)
	acc.Doorbell = service.NewDoorbell()
	acc.AddService(acc.Doorbell.Service)

	return &acc
}
