package state

type DataType int

const (
	Arithematic DataType = iota
	Comparison
	Concatenation
	Combineslice
)

func GetFactoryData(dataType DataType) *ExpressionData {
	switch dataType {
	case Arithematic:
		return GetArithematicData()
	case Comparison:
		return GetComparisonData()
	case Concatenation:
		return GetConcatenationData()
	case Combineslice:
		return GetCombineSliceData()
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

func GetArithematicData() *ExpressionData {
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
