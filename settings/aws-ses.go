package settings

type amazonSesSettings struct {
	Region    string
	Sender    string
	Charset   string
	AccessKey string
	SecretKey string
	Token     string
}

var AmazonSesSettings *amazonSesSettings
