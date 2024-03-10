package amazon

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/settings"
	"strconv"
	"strings"
	"time"
)

func CloudFrontInvalidation(path ...string) {
	logger.SysLog.Infof("[CreateInvalidation] SetInvalidation Path: %+v", path)
	// Set Invalidation
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(settings.AmazonS3Settings.Region),
		Credentials: credentials.NewStaticCredentials(
			settings.AmazonS3Settings.AccessKey,
			settings.AmazonS3Settings.SecretKey,
			"",
		),
	}))
	var finalPath []string
	for _, p := range path {
		if !strings.HasPrefix(p, "/") {
			p = fmt.Sprintf("/%s", p)
		}
		finalPath = append(finalPath, p)
	}
	cloudFrontSvc := cloudfront.New(mySession)
	caller := strconv.FormatInt(time.Now().UnixNano(), 10)
	InvalidationInput := &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(settings.AmazonS3Settings.CloudFrontId),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(caller),
			Paths: &cloudfront.Paths{
				Items:    aws.StringSlice(finalPath),
				Quantity: aws.Int64(int64(len(finalPath))),
			},
		},
	}
	req, resp := cloudFrontSvc.CreateInvalidationRequest(InvalidationInput)
	err := req.Send()
	if err == nil {
		logger.SysLog.Infof("[CreateInvalidation] %+v", resp)
	} else {
		logger.SysLog.Errorf("[CreateInvalidation] Failed, %+v", err)
	}
}
