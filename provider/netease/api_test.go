package netease

import (
	"fmt"
	"testing"
)

// func TestRsaEncrypt(t *testing.T) {
// 	rs := RsaEncrypt([]byte("aaabaaabaaabaaab"))
// 	fmt.Println("rs=" + rs)
// 	a := big.NewInt(0)
// 	a.SetString("010001", 16)
// 	fmt.Printf("%v", a)
// }

func TestWeapi(t *testing.T) {
	// http://music.163.com/weapi/song/enhance/player/url?csrf_token=
	testInput := `{"ids":["28528420","347230"],"br":320000,"csrf_token":""}`
	param := weapi(testInput)
	fmt.Println(param)
}

func TestLinuxApi(t *testing.T) {
	// testInput := `{"ids":"[33911781]","br":999000}`
	// https://music.163.com/api/linux/forward
	testInput := `{"method":"POST","url":"https://music.163.com/api/song/enhance/player/url","params":{"ids":"[28445807,474567613]","br":999000}}`
	// param := linuxApi(testInput)
	// fmt.Println(param)
	fmt.Println(testInput)
	fmt.Println(linuxAPI(testInput))
}

// func TestQpow(t *testing.T) {
// 	rs := qpow(*big.NewInt(111111111), *big.NewInt(111111111), *big.NewInt(int64(1e9 + 7)))
// 	fmt.Printf("%v\n", &rs)
// }

// func TestRaEncrypt(t *testing.T) {
// 	RsaEncrypt([]byte("10000000000000"))
// }

// func TestMD5(t *testing.T) {
// 	fmt.Println(MD5("a"))
// }

func TestDoHttp(t *testing.T) {
	testInput := `{"method":"POST","url":"https://music.163.com/api/song/enhance/player/url","params":{"ids":"[28445807,509098750,28528420]","br":999000}}`
	rs := doHTTP("POST", "https://music.163.com/api/linux/forward", testInput, "linux")
	fmt.Println(rs)
}

func TestGetSongUrl(t *testing.T) {
	rs := GetSongURL([]string{"22730174"}, 320)
	fmt.Println(rs)
}

func TestExp(t *testing.T) {
	// rs := doHttp("GET", "https://music.163.com/song?id=28445807", "", "")
	// reg, _ := regexp.Compile(`"title": "[^"]+`)
	// fmt.Println(reg.FindAllString(rs, 1))

}

func TestSearch(t *testing.T) {
	rs := Search("トゥルーエンド プレイヤー")
	fmt.Println(rs)
	// var val songBean.SongData = rs[0]
	// fmt.Println(val.GetUrl(320))
}

// func TestAesEncryptCBC(t *testing.T) {
// 	fmt.Println(aesEncryptCBC("aaabaaabaaabaaab"))
// }
