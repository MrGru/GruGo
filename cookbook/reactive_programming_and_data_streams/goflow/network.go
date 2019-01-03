package goflow

import (
	flow "github.com/trustmaster/goflow"
)

func NewEncodingApp() *flow.Graph {
	e := flow.NewGraph()
	e.Add("encoder", new(Encoder))
	e.Add("printer", new(Printer))
	e.Connect("encoder", "Res", "printer", "Line")
	e.MapInPort("In", "encoder", "Val")
	return e
}
