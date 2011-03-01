package crypticmysql

import "testing"

type tester struct {
	key  []byte
	data []byte
}

var tests = []tester{
	{[]byte("abcdefghijklmnop"), []byte("brian")},
	{[]byte("abcdefghijklmnop"), []byte("")},
	{[]byte("abcdefghijklmnop"), []byte("0")},
	{[]byte("abcdefghijklmnop"), []byte("01")},
	{[]byte("abcdefghijklmnop"), []byte("012")},
	{[]byte("abcdefghijklmnop"), []byte("0123")},
	{[]byte("abcdefghijklmnop"), []byte("01234")},
	{[]byte("abcdefghijklmnop"), []byte("012345")},
	{[]byte("abcdefghijklmnop"), []byte("0123456")},
	{[]byte("abcdefghijklmnop"), []byte("01234567")},
	{[]byte("abcdefghijklmnop"), []byte("012345678")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789A")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789AB")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789ABC")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789ABCD")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789ABCDE")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789ABCDEF")},
	{[]byte("abcdefghijklmnop"), []byte("0123456789ABCDEFG")},
	{[]byte("abcdefghijklmnop"), []byte{0}},
	{[]byte("abcdefghijklmnop"), []byte{11, 11, 11, 11, 11}},
}

func testAES(t *testing.T, key, data []byte) {
	enc := AESEncrypt(data, key)
	dec := AESDecrypt(enc, key)
	if string(dec) != string(data) {
		t.Errorf("aes output should match %q is %q", string(dec),
string(data))
	}
}

func TestAES(t *testing.T) {
	for _, testcase := range tests {
		testAES(t, testcase.key, testcase.data)
	}
}

func BenchmarkAES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enc := AESEncrypt(tests[0].data, tests[0].key)
		dec := AESDecrypt(enc, tests[0].key)
		_ = dec
	}
}