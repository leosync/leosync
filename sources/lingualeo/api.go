package lingualeo

type linguaResp struct {
	ErrorMsg string     `json:"error_msg"`
	Count    uint       `json:"count_words"`
	ShowMore bool       `json:"show_more"`
	Userdict []Userdict `json:"userdict3"`
}

const (
	linguaDictUrl      = "http://lingualeo.com/userdict/json"
	linguaLoginUrl     = "http://api.lingualeo.com/api/login"
	linguaTranslateUrl = "http://api.lingualeo.com/gettranslates"
	linguaAddWordUrl   = "http://api.lingualeo.com/addword"
	httpTimeout        = 15
)

var leoClient linguaLeoClient

func GetClient() linguaLeoClient {
	if (leoClient == linguaLeoClient{}) {
		connectionConfig := *createConfig()
		leoClient = linguaLeoClient{
			connection: createConnection(&connectionConfig),
		}
	}
	return leoClient
}
