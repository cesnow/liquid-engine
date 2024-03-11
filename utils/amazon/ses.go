package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/settings"
)

func SesSendMail(Recipient string, Subject, Msg string) bool {

	Sender := settings.AmazonSesSettings.Sender
	CharSet := settings.AmazonSesSettings.Charset

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(settings.AmazonSesSettings.Region),
		Credentials: credentials.NewStaticCredentials(
			settings.AmazonSesSettings.AccessKey,
			settings.AmazonSesSettings.SecretKey,
			settings.AmazonSesSettings.Token,
		),
	})
	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(Msg),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case ses.ErrCodeMessageRejected:
				logger.SysLog.Error(ses.ErrCodeMessageRejected, awsErr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				logger.SysLog.Error(ses.ErrCodeMailFromDomainNotVerifiedException, awsErr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				logger.SysLog.Error(ses.ErrCodeConfigurationSetDoesNotExistException, awsErr.Error())
			default:
				logger.SysLog.Error(awsErr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			logger.SysLog.Error(err.Error())
		}

		return false
	}

	logger.SysLog.Info("Email Sent to address: " + Recipient)
	logger.SysLog.Info(result)
	return true

	//from := Settings.SmtpSettings.FromMail
	//// Set up authentication information.
	//mailPlainAuth := smtp.PlainAuth("",
	//	Settings.SmtpSettings.Account, Settings.SmtpSettings.Password, Settings.SmtpSettings.Host)
	//// Connect to the server, authenticate, set the sender and recipient,
	//// and send the email all in one step.
	//addr := fmt.Sprintf("%s:%d", Settings.SmtpSettings.Host, Settings.SmtpSettings.Port)
	//err := smtp.SendMail(addr, mailPlainAuth, from, []string{to}, msg)
	//if err != nil {
	//	logger.SysLog.Warnf("%s", err.Error())
	//	return false
	//}
	//return true
}
