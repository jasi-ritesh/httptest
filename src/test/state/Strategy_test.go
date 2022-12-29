package state

import (
	"testing"
)

func TestArithematicStrategyEngine(t *testing.T) {

	data := GetFactoryData(Arithematic)
	se := StateMachineEngine{t: t, data: data}
	se.RunStrategy()
}

func TestComparisonStrategyEngine(t *testing.T) {

	data := GetFactoryData(Comparison)
	se := StateMachineEngine{t: t, data: data}
	se.RunStrategy()
}

func TestStateMachineEngine_Concatenation(t *testing.T) {
	data := GetFactoryData(Concatenation)
	se := StateMachineEngine{t: t, data: data}
	se.RunStrategy()

}
func TestGetCombineSliceData(t *testing.T) {
	data := GetFactoryData(Combineslice)
	se := StateMachineEngine{t: t, data: data}
	se.RunStrategy()
}
