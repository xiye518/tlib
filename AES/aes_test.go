package AES

import (
	"encoding/base64"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	//t.Log(time.Now().Unix())
	//return
	{
		// token wem0Upqsl5MBD0Z39jWO/g==
		in := `{"account":"test","ip":"127.0.0.1","time":1562832202}`
		key := "0123456789abcdef"
		c := new(AESCrypt)
		result, err := c.Encrypt([]byte(in), []byte(key))
		if err != nil {
			t.Fatalf("加密失败：%s", err.Error())
		}
		t.Logf("加密后为：%s", result)
		baseStr := base64.StdEncoding.EncodeToString(result)
		t.Logf("加密后的base64为：%s", baseStr)

		bs, err := base64.StdEncoding.DecodeString(baseStr)
		data, err := c.Decrypt(bs, []byte(key))
		if err != nil {
			t.Fatalf("解密失败：%s", err.Error())
		}
		if string(data) != in {
			t.Fatalf("解密错误：%s", string(data))
		}
		t.Logf("解密后为：%s", data)

	}
}

func TestGetToken(t *testing.T) {
	in := `{"account":"test","ip":"127.0.0.1","time":0}`
	key := "0123456789abcdef"
	c := new(AESCrypt)
	result, err := c.Encrypt([]byte(in), []byte(key))
	if err != nil {
		t.Fatalf("加密失败：%s", err)
	}
	t.Logf("加密后为：%s", result)
	//5SOZ33f22XAIqPNiiy3aFVa7+nELOf/a/eJ27oN7q4i7WXvuFPragDPsCHiE5fXK
	baseStr := base64.StdEncoding.EncodeToString(result)
	t.Logf("加密后的base64为：%s", baseStr)
}

func TestDecode(t *testing.T) {
	key := "0123456789abcdef"
	in := `5SOZ33f22XAIqPNiiy3aFVa7+nELOf/a/eJ27oN7q4i7WXvuFPragDPsCHiE5fXK`
	content, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("content: %s", content)

	c := new(AESCrypt)
	result, err := c.Decrypt(content, []byte(key))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %s", result)

}
