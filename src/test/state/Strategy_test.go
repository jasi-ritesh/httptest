package state

import (
	"testing"
)

func TestArithmeticStrategyEngine(t *testing.T) {

	data := GetFactoryData(Arithmetic)
	se := StateMachineEngine{t: t, engineName: "Arithmetic", data: data}
	se.RunStrategy()
}

func TestComparisonStrategyEngine(t *testing.T) {

	data := GetFactoryData(Comparison)
	se := StateMachineEngine{t: t, engineName: "Comparison", data: data}
	se.RunStrategy()
}

func TestStateMachineEngine_Concatenation(t *testing.T) {
	data := GetFactoryData(Concatenation)
	se := StateMachineEngine{t: t, engineName: "Concatenation", data: data}
	se.RunStrategy()

}
func TestGetCombineSliceData(t *testing.T) {
	data := GetFactoryData(Combineslice)
	se := StateMachineEngine{t: t, engineName: "Combineslice", data: data}
	se.RunStrategy()
}

func TestConcurrentStateMachineEngine_RunStrategy(t *testing.T) {

	cse := ConcurrentStateMachineEngine{engines: []*StateMachineEngine{
		&StateMachineEngine{t: t, engineName: "Arithmetic", data: GetFactoryData(Arithmetic)},
		&StateMachineEngine{t: t, engineName: "Comparison", data: GetFactoryData(Comparison)},
		&StateMachineEngine{t: t, engineName: "Concatenation", data: GetFactoryData(Concatenation)},
		&StateMachineEngine{t: t, engineName: "Combineslice", data: GetFactoryData(Combineslice)},
	}}

	cse.RunStrategy()

}

func TestArithmeticFromYaml(t *testing.T) {
	data := GetFactoryData(ArithmeticYaml)
	se := StateMachineEngine{t: t, engineName: "ArithmeticYaml", data: data}
	se.RunStrategy()
}
