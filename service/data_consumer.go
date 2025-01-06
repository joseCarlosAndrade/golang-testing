package service

func ConsumeData(provider DataProvider, id int) (string, error) {
	randomNumber, err := provider.GetRandomData(id)
	if err != nil {
		return "", err
	}

	result := checkParity(randomNumber)
	return result, nil
}

func checkParity(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}