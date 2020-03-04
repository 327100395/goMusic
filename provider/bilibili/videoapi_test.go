package bilibili

import "testing"

func TestGetURL(t *testing.T) {
	var (
		input = "135463335"
		output = "https://interface.bilibili.com/v2/playurl?appkey=iVGUTjsxvpLeuDCf" +
		"&cid=135463335&otype=json&platform=flash&qn=80&quality=80&type=&sign=2b90191093ff9712e448b54ee0749007"
	)
	url := GetVideoURL(input)
	if url != output {
		t.Error("GetURL error", url, output)
	}
}
