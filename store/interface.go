package store

// 定义如何上传文件到bucket
// 做了抽象，方便以后切换到其他云存储
type Uploader interface {
	Upload(bucketName string, objectKey string, fileName string) error
}
