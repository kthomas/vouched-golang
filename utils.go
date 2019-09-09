package vouched

func stringOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
