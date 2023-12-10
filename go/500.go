import (
    "strings"
    "fmt"
)

func findWords(words []string) []string {
    row1 := "qwertyuiop"
    row2 := "asdfghjkl"
    row3 := "zxcvbnm"

    result := make([]string, 0)

    for _,word := range words {
        fmt.Println(word)
        temp := strings.ToLower(word)
        char1 := string(temp[0])
        fmt.Println(char1)


        addWord := true

        if strings.Contains(row1, string(char1)) {

            for _, char := range temp {
                if !strings.Contains(row1, string(char)) {
                    addWord = false
                    break
                }
            }

        }else if strings.Contains(row2, string(char1)) {

            for _, char := range temp {
                if !strings.Contains(row2, string(char)) {
                    addWord = false
                    break
                }
            }
        }else {

            for _, char := range temp {
                if !strings.Contains(row3, string(char)) {
                    addWord = false
                    break
                }
            }
        }
        if addWord {
            result = append(result, word)
        }
    }
    return result
}
