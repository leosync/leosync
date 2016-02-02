package lingualeo

type Word struct {
	Id         uint   `json:"word_id"`
	Value      string `json:"word_value"`
	Transcript string `json:"transcription"`
	//	Created     time.Time       `json:"created_at"`
	//	LastUpdated time.Time       `json:"last_updated_at"`
	Translates []UserTranslate `json:"translate"`
	SoundUrl   string          `json:"sound_url"`
	PictureUrl string          `json:"pic_url"`
}

type UserTranslate struct {
	Value      string `json:"value"`
	PictureUrl string `json:"pic_url"`
}

type Userdict struct {
	Name  string `json:"name"`
	Count uint   `json:"count"`
	Words []Word `json:"words"`
}
