package sliceutils

import (
	"github.com/decentralisedkev/Neo-Go-API/models"
)

// TODO: rmeove dependency on models.host
func RemoveDuplicates(elements []models.Host) []models.Host {
	// Use map to record duplicates as we find them.
	encountered := map[models.Host]bool{}
	result := []models.Host{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
