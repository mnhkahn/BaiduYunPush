package BaiduYunPush

type URLParamMap map[string][]string

type BaiduResponseSuccess struct {
	RequestId      string `json:"request_id"`
	ResponseParams struct {
		ChannelId    float64 `json:"channel_id"`
		ChannelToken string  `json:"channel_token"`
	} `json:"response_params"`
}

type BaiduResponseError struct {
	RequestId    string `json:"request_id"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}
