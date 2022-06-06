package api_manager

type Video struct {
	Title string
	Id    string
	Link  string
}

func Get(url string) (*[]Video, error) {

}

func CreateUrl(search string, engine *ApiEngine) string {

}
