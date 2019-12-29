package cli

func normalize(source []string) ([]string, error) {
	if len(source) == 0 {
		return []string{"."}, nil
	} else {
		return source, nil
	}
}
