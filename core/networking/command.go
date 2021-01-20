package networking

import "go_rts/core/geometry"

const (
	DeselectUnitsCommand  = "deselect units"
	SelectUnitsCommand    = "select_units"
	SetDestinationCommand = "set destination"
)

type NetworkCommand struct {
	Name string
	Data interface{}
}

func NewDeselectUnitsCommand(p geometry.Point) NetworkCommand {
	return NetworkCommand{
		Name: DeselectUnitsCommand,
		Data: p,
	}
}

func NewSelectUnitsCommand(r geometry.Rectangle) NetworkCommand {
	return NetworkCommand{
		Name: SelectUnitsCommand,
		Data: r,
	}
}

func NewSetDestinationCommand(p geometry.Point) NetworkCommand {
	return NetworkCommand{
		Name: SetDestinationCommand,
		Data: p,
	}
}
