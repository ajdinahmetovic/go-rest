package config

//AppConfig config struct
type appConfig struct {
	ElasticSearchURL string
	ItemServiceURL   string
}

//AppCfg instance
var AppCfg appConfig

//Load function to load ENV variables to config
func Load() {

	AppCfg = appConfig{
		ElasticSearchURL: "localhost:9200",
		ItemServiceURL:   "localhost:4040",
	}

	/*
		AppCfg = appConfig{
			ElasticSearchURL: os.Getenv("ES_HOST"),
			ItemServiceURL:   os.Getenv("ITEM_SERVICE_URL"),
		}
	*/
}
