/* This file contains the service definitions for the flames package. */
package games

import (
    "strings"
)

// removeDuplicateCharacters removes all duplicate characters from a string
// str: the string to remove duplicate characters from
// returns: the string with all duplicate characters removed
func removeDuplicateCharacters(str string) string {
    result := ""
    for _, char := range str {
        if !strings.Contains(result, string(char)) {
            result += string(char)
        }
    }
    return result
}

// ComputeFlames computes the flames between two names
// name1: the first name
// name2: the second name
// returns: the flames between the two names
func (service *Service) ComputeFlames(name1 string, name2 string) string {
	// Remove all duplicate characters from the names
	name1 = removeDuplicateCharacters(name1)
	name2 = removeDuplicateCharacters(name2)

	// Convert the names to lowercase
	name1 = strings.ToLower(name1)
	name2 = strings.ToLower(name2)

	// Count the number of characters in each name
	count1 := len(name1)
	count2 := len(name2)

	// Count the number of characters in common between the two names
	common := 0
	for _, char := range name1 {
		if strings.Contains(name2, string(char)) {
			common++
		}
	}

	// Compute the flames
	flames := (count1 + count2) - (2 * common)
	flames %= 6

	switch flames {
	case 0:
		return "F"
	case 1:
		return "L"
	case 2:
		return "A"
	case 3:
		return "M"
	case 4:
		return "E"
	case 5:
		return "S"
	}
	return ""
}
