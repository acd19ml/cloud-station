package aliyun

import (
	"acd19ml/cloud-station/store"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 对象是否实现了Uploader接口
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint, accessKey, accessSecret has one empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		Endpoint:     os.Getenv("ALI_OSS_ENDPOINT"),
		AccessKey:    os.Getenv("ALI_AK"),
		AccessSecret: os.Getenv("ALI_SK"),
	})
}

// AliOssStore 对象的构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	// Create a new client
	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}

	return &AliOssStore{
		client: c,
	}, nil
}

type AliOssStore struct {
	client *oss.Client
}

func (a *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// Get the bucket
	bucket, err := a.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// Upload a file to the bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	// Print a link to download the file
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("下载链接: %s\n", downloadURL)
	fmt.Println("请在1天之内下载")

	return nil
}
