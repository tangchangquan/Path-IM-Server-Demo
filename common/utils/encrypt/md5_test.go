package encrypt

import "testing"

func TestMd5(t *testing.T) {
	md5 := Md5("123456")
	t.Log(md5)
}
