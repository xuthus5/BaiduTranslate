package BaiduTranslate

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	bi := BaiduInfo{
		AppID:"XXXXX",
		Salt:Salt(5),
		SecretKey:"XXXXX",
		From:"auto",
		To:"en",
	}
	bi.Text = "你好,世界"
	bi.Translate()
}