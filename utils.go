package flatfile

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func FlatfileString(value string) *string {
	if value == "" {
		return nil
	}

	return &value
}

func FlatfileBool(value bool) *bool {
	return &value
}
