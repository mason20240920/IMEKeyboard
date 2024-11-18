package IMEKeyboard

import (
	"fmt"
	"testing"
)

func TestKeyboardPos_ChooseWord(t *testing.T) {
	model, err := KBPSharedInstance(BaiduBrand, "GooglePixel2XL", YamlFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	coordinate := model.LogoKeyboardAction()
	fmt.Println(coordinate.X)
	fmt.Println(coordinate.Y)
}
