package aliyun_test

import (
	"os"
	"testing"

	"acd19ml/cloud-station/store"
	"acd19ml/cloud-station/store/aliyun"

	"github.com/stretchr/testify/assert"
)

var (
	uploader store.Uploader
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

func TestUpload(t *testing.T) {
	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		t.Log("Upload success")
	}
	if err != nil {
		t.Log(err)
	}
}

func TestUploadError(t *testing.T) {
	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_testtt.go")
	should.Error(err, "open store_testtt.go: no such file or directory")
}

// 通过 init 编写 uploader 实例化逻辑
func init() {
	// 初始化AliOssStore
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}

	// AliOssStore实现了Uploader接口，传给全局变量
	uploader = ali
}
