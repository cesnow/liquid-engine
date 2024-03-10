package amazon

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/settings"
	"github.com/cesnow/liquid-engine/utils/pathutil"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const (
	S3AclPublic  = "public-read"
	S3AclPrivate = "private"
)

var S3AclSetting = S3AclPublic

func SetACLIsPublic() {
	S3AclSetting = S3AclPublic
}

func SetACLIsPrivate() {
	S3AclSetting = S3AclPrivate
}

func getAwsSession() (*session.Session, error) {
	config := &aws.Config{
		Region: aws.String(settings.AmazonS3Settings.Region),
		Credentials: credentials.NewStaticCredentials(
			settings.AmazonS3Settings.AccessKey,
			settings.AmazonS3Settings.SecretKey,
			"",
		),
	}
	//if Settings.AmazonS3Settings.Endpoint != "" {
	//	config.Endpoint = aws.String(Settings.AmazonS3Settings.Endpoint)
	//}
	awsSession, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	return awsSession, nil
}

func GetFullUrl(src string) string {
	dirPart := strings.Split(src, string(os.PathSeparator))
	return pathutil.UrlJoin(settings.AmazonS3Settings.CloudFrontUri, dirPart...)
}

func DownloadObject(objectPath string) ([]byte, error) {
	awsSession, _ := getAwsSession()
	downloader := s3manager.NewDownloader(awsSession, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 4
		d.BufferProvider = s3manager.NewPooledBufferedWriterReadFromProvider(1024 * 1024 * 8)
	})

	buff := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(settings.AmazonS3Settings.Bucket),
		Key:    aws.String(objectPath),
	})
	data := buff.Bytes()

	if err != nil {
		logger.SysLog.Errorf("[S3][DownloadObject] Failed, %s", err)
		return nil, err
	}
	return data, nil
}

func MoveObject(objectPath string, toObjectPath string) error {
	awsSession, _ := getAwsSession()
	svc := s3.New(awsSession)
	logger.SysLog.Infof("[MoveObject] Copy %s to %s", path.Join(settings.AmazonS3Settings.Bucket, objectPath), toObjectPath)
	copyObjectInput := &s3.CopyObjectInput{
		Key:        aws.String(toObjectPath),
		Bucket:     aws.String(settings.AmazonS3Settings.Bucket),
		CopySource: aws.String(path.Join(settings.AmazonS3Settings.Bucket, objectPath)),
	}
	deleteObjectInput := &s3.DeleteObjectInput{
		Bucket: aws.String(settings.AmazonS3Settings.Bucket),
		Key:    aws.String(objectPath),
	}
	cpOutput, err := svc.CopyObject(copyObjectInput)
	_, _ = svc.DeleteObject(deleteObjectInput)
	if err != nil {
		logger.SysLog.Warnf("[S3MoveObject] Move Failed, %s", err.Error())
		return err
	}
	logger.SysLog.Infof("%+v", cpOutput)
	return nil
}

func DeleteFolder(folderPath string) error {

	if strings.HasPrefix(folderPath, "/") {
		folderPath = folderPath[1:]
	}

	awsSession, _ := getAwsSession()
	svc := s3.New(awsSession)
	listObjectsInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(settings.AmazonS3Settings.Bucket),
		Prefix: aws.String(folderPath),
	}
	listObjectsOutput, err := svc.ListObjectsV2(listObjectsInput)
	if err != nil {
		logger.SysLog.Warnf("[S3DeleteFolder] Failed to list objects in folder %s, %s", folderPath, err.Error())
		return err
	}
	for _, object := range listObjectsOutput.Contents {
		deleteObjectInput := &s3.DeleteObjectInput{
			Bucket: aws.String(settings.AmazonS3Settings.Bucket),
			Key:    object.Key,
		}
		go func(inp *s3.DeleteObjectInput) {
			_, _ = svc.DeleteObject(inp)
		}(deleteObjectInput)
		//_, err := svc.DeleteObject(deleteObjectInput)
		//if err != nil {
		//	logger.SysLog.Warnf("[S3DeleteFolder] Failed to delete object %s, %s", *object.Key, err.Error())
		//}
	}
	//logger.SysLog.Infof("[S3DeleteFolder] Deleted all objects under folder %s", folderPath)
	return nil
}

func DeleteObject(objectPath string) error {
	awsSession, _ := getAwsSession()
	svc := s3.New(awsSession)
	logger.SysLog.Infof("[DeleteObject] Delete %s", path.Join(settings.AmazonS3Settings.Bucket, objectPath))
	deleteObjectInput := &s3.DeleteObjectInput{
		Bucket: aws.String(settings.AmazonS3Settings.Bucket),
		Key:    aws.String(objectPath),
	}
	deleteObjOutput, err := svc.DeleteObject(deleteObjectInput)
	if err != nil {
		logger.SysLog.Warnf("[S3DeleteObject] Failed, %s", err.Error())
		return err
	}
	if deleteObjOutput.DeleteMarker != nil && deleteObjOutput.VersionId != nil {
		logger.SysLog.Infof("[S3DeleteObject] DeleteMarker: %t, VersionId: %s",
			*deleteObjOutput.DeleteMarker, *deleteObjOutput.VersionId)
	}
	return err
}

func CheckObject(objectPath string) (bool, error) {

	awsSession, _ := getAwsSession()

	downloader := s3.New(awsSession)
	_, err := downloader.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(settings.AmazonS3Settings.Bucket),
		Key:    aws.String(objectPath),
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func UploadObjectTTLFromIO(ioReader io.Reader, objectPath string, expires *time.Time) error {

	awsSession, _ := getAwsSession()
	uploader := s3manager.NewUploader(awsSession)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(settings.AmazonS3Settings.Bucket),
		Key:                aws.String(objectPath),
		Body:               ioReader,
		ACL:                aws.String(S3AclSetting),
		ContentDisposition: aws.String("attachment"),
		Expires:            expires,
	})

	if err != nil {
		logger.SysLog.Errorf("[S3][UploadObjectFromPath] Failed, %s", err)
		return err
	}
	return nil
}

func UploadObjectWithData(data []byte, objectPath string, contentType string) error {
	reader := bytes.NewReader(data)
	awsSession, _ := getAwsSession()
	uploader := s3manager.NewUploader(awsSession)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(settings.AmazonS3Settings.Bucket),
		Key:         aws.String(objectPath),
		Body:        reader,
		ACL:         aws.String(S3AclSetting),
		ContentType: aws.String(contentType),
		// ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		logger.SysLog.Errorf("[S3][UploadObjectFromPath] Failed, %s", err)
		return err
	}
	return nil
}

func UploadObjectFromPath(originFilePath string, objectPath string) error {
	// Open File into Buffer
	file, err := os.Open(originFilePath)
	if err != nil {
		logger.SysLog.Errorf("[S3][UploadObjectFromPath] Failed, %s", err)
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)

	awsSession, err := getAwsSession()
	if err != nil {
		logger.SysLog.Errorf("[S3][UploadObjectFromPath] Failed, %s", err)
		return err
	}

	uploader := s3manager.NewUploader(awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(settings.AmazonS3Settings.Bucket),
		Key:                aws.String(objectPath),
		Body:               bytes.NewReader(buffer),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ACL:                aws.String(S3AclSetting),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		logger.SysLog.Errorf("[S3][UploadObjectFromPath] Failed, %s", err)
		return err
	}
	return nil
}
