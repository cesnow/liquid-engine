package settings

type amazonS3Settings struct {
	Region           string
	Bucket           string
	CloudFrontDomain string
	CloudFrontUri    string
	BucketUri        string
	AccessKey        string
	SecretKey        string
	CloudFrontId     string
}

var AmazonS3Settings *amazonS3Settings
