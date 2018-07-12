 package main

 import (
         "fmt"
         "strings"
 )

 func snakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
         //snake_case to camelCase

         isToUpper := false

         for k, v := range inputUnderScoreStr {
                 if k == 0 {
                         camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
                 } else {
                         if isToUpper {
                                 camelCase += strings.ToUpper(string(v))
                                 isToUpper = false
                         } else {
                                 if v == '_' {
                                         isToUpper = true
                                 } else {
                                         camelCase += string(v)
                                 }
                         }
                 }
         }
         return

 }

 func main() {
         snakeCase := "this_is_a_statement_with_underscore_which_is_also_known_as_Snake_Case"
         result := snakeCaseToCamelCase(snakeCase)

         fmt.Println(snakeCase)
         fmt.Println(result)
 }