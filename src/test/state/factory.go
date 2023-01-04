package state

import (
	"os"
)
import "github.com/ghodss/yaml"

type DataType int

const (
	Arithmetic DataType = iota
	Comparison
	Concatenation
	Combineslice
	ArithmeticYaml
)

func GetFactoryData(dataType DataType) *ExpressionData {
	switch dataType {
	case Arithmetic:
		return GetArithmeticData()
	case Comparison:
		return GetComparisonData()
	case Concatenation:
		return GetConcatenationData()
	case Combineslice:
		return GetCombineSliceData()
	case ArithmeticYaml:
		return GetArithmeticDataFromYaml()
	}
	return nil
}

func GetComparisonData() *ExpressionData {
	data := ExpressionData{
		InputMap: map[string]string{
			"First":  "3>2",
			"Second": "4<2",
			"Third":  "100<=(50+60)",
		},
		ExpectedOutputMap: map[string]string{
			"First":  "true",
			"Second": "false",
			"Third":  "true",
		},
		DeletedNames: []string{
			"First",
		},
	}
	return &data
}

func GetArithmeticData() *ExpressionData {
	data := ExpressionData{
		InputMap: map[string]string{
			"First":  "2+3",
			"Second": "10%3",
		},
		ExpectedOutputMap: map[string]string{
			"First":  "5",
			"Second": "1",
		},
		DeletedNames: []string{
			"First",
		},
	}
	return &data
}

func GetConcatenationData() *ExpressionData {

	data := ExpressionData{
		InputMap: map[string]string{
			"First":  "\"1\" + \"2\"",
			"Second": "\"India\" +  \"Bangalore\"",
		},
		ExpectedOutputMap: map[string]string{
			"First":  "12",
			"Second": "IndiaBangalore",
		},
		DeletedNames: []string{
			"First",
		},
	}
	return &data
}

func GetCombineSliceData() *ExpressionData {
	data := ExpressionData{
		InputMap: map[string]string{
			"First":  "[1, 2] +[3, 4]",
			"Second": "[\"bangalore\", \"chennai\"] + [\"mumbai\" , \"delhi\"]",
		},
		ExpectedOutputMap: map[string]string{
			"First":  "[1 2 3 4]",
			"Second": "[bangalore chennai mumbai delhi]",
		},
		DeletedNames: []string{
			"First",
		},
	}
	return &data
}
func GetArithmeticDataFromYaml() *ExpressionData {
	var data ExpressionData
	yfile, _ := os.ReadFile("arithmetic.yaml")
	yaml.Unmarshal(yfile, &data)
	return &data
}
