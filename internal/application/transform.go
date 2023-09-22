package application

func TransformValues(values []int) []int {
	modifiedValues := make([]int, len(values))
	copy(modifiedValues, values)

	for i := 0; i < len(modifiedValues); i++ {
		modifiedValues[i] = modifiedValues[i] - 1
	}

	return modifiedValues
}
