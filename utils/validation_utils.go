package utils

import (
    "strings"
)

func IsValidExtension(ext string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".png", ".pdf"} // Add more extensions as needed
	for _, validExt := range validExtensions {
		if strings.ToLower(ext) == validExt {
			return true
		}
	}
	return false
}
