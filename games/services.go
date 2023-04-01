package games

import (
    "strings"
)

func removeDuplicateCharacters(str string) string {
    result := ""
    for _, char := range str {
        if !strings.Contains(result, string(char)) {
            result += string(char)
        }
    }
    return result
}

func (service *Service)ComputeFlames(name1 string, name2 string) string {
    // convert both names to lower case
    name1 = strings.ToLower(name1)
    name2 = strings.ToLower(name2)

    // remove all spaces
    name1 = strings.Replace(name1, " ", "", -1)
    name2 = strings.Replace(name2, " ", "", -1)

    // remove all duplicate characters
    name1 = removeDuplicateCharacters(name1)
    name2 = removeDuplicateCharacters(name2)

    // count the number of characters in each name
    count1 := len(name1)
    count2 := len(name2)

    // count the number of characters in both names
    count := count1 + count2

    // count the number of characters in both names that are common
    for _, char := range name1 {
        if strings.Contains(name2, string(char)) {
            count -= 2
        }
    }

    // compute the flames
    flames := []string{"F", "L", "A", "M", "E", "S"}
    for len(flames) > 1 {
        index := count % len(flames)
        if index == 0 {
            index = len(flames)
        }
        flames = append(flames[:index-1], flames[index:]...)
    }

    return flames[0]
}
