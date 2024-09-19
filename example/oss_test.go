package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
	client       *oss.Client
)

func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}
	// Upload a file to the bucket
	// objectKey: mydir/test.go, oss server will create a directory named mydir
	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

func init() {

	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	fmt.Println("EndPoint:", OssEndpoint)
	fmt.Println("AccessKey:", AccessKey)
	if err != nil {
		panic(err)
	}

	client = c
}
