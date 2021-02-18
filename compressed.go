package gomime

import (
	"strings"
)

var usualCompressedTypes = []string{
	"7z",
	"rar",
	"gzip",
	"zip",
	"tar",
}

// IsCompressed return determine whether the file is a compressed package
func IsCompressed(extension string) bool {
	for _, usualCompressedType := range usualCompressedTypes {
		if strings.ToLower(extension) == usualCompressedType {
			return true
		}
	}
	contentType := TypeByExtension(extension)
	if strings.HasSuffix(contentType, "compressed") {
		return true
	}
	return false
}
