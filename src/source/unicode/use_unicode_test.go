package unicode

import (
	"testing"
	"unicode"
	"fmt"
)

// unicode & ascii . ascii is only the first 0 ~ 255 of unicode . ascii is a small part of unicode

func TestFuncUnicode(t *testing.T){

	const mixed = "\b5Ὂg̀9! ℃ᾭG"

	for _ ,c := range mixed {
		// check if the c is a rune control character
		if unicode.IsControl(c) {
			fmt.Println("is control rune")
		}
		// check if character is a letter
		if unicode.IsLetter(c) {
			fmt.Println("is Letter rune")
		}
		// check if c is a decimal digital
		if unicode.IsDigit(c){
			fmt.Println("is digit rune")
		}
		// check if c is a graphic rune
		if unicode.IsGraphic(c){
			fmt.Println("is Graphic rune")
		}
	}


}