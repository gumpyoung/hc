package accessory

import (
	"github.com/gumpyoung/hc/service"
)

//Button StatefulSwitch or StatelessSwitch
type Button struct {
	*Accessory
	StatefulSwitch  *service.StatefulProgrammableSwitch
	StatelessSwitch *service.StatelessProgrammableSwitch
}

// NewButton returns a switch which implements model.ProgrammableSwitch.
func NewButton(info Info, stateful bool) *Button {
	acc := Button{}
	acc.Accessory = New(info, TypeProgrammableSwitch)

	if stateful {
		acc.StatefulSwitch = service.NewStatefulProgrammableSwitch()
		acc.AddService(acc.StatefulSwitch.Service)
	} else {
		acc.StatelessSwitch = service.NewStatelessProgrammableSwitch()
		acc.AddService(acc.StatelessSwitch.Service)
	}
	return &acc
}
