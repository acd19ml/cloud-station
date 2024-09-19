package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	accessKey    = os.Getenv("ALI_AK")
	accessSecret = os.Getenv("ALI_SK")
	endpoint     = os.Getenv("ALI_OSS_ENDPOINT")
	bucketName   = os.Getenv("ALI_BUCKET_NAME")
	uploadFile   = ""
	help         = false
)

// 实现文件上传
func upload(file_path string) error {
	// Create a new client
	client, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return err
	}

	// Get the bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// Upload a file to the bucket
	if err := bucket.PutObjectFromFile(file_path, file_path); err != nil {
		return err
	}

	// Print a link to download the file
	downloadURL, err := bucket.SignURL(file_path, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("下载链接: %s\n", downloadURL)
	fmt.Println("请在1天之内下载")

	return nil
}

func validate() error {
	if endpoint == "" || accessKey == "" || accessSecret == "" {
		return fmt.Errorf("endpoint, accessKey, accessSecret has one empty")
	}

	if uploadFile == "" {
		return fmt.Errorf("upload file path required")
	}

	return nil
}

/*
&uploadFile：这是一个指向 uploadFile 变量的指针，表示解析命令行参数时将值存储到 uploadFile 中。
"f"：这是命令行标志的名称，表示当用户在命令行中输入 -f 标志时，接下来的内容会作为文件名解析。
""：这是标志的默认值，如果用户没有提供 -f 标志，uploadFile 的值将保持为空字符串。
"上传文件的名称"：这是命令行标志的帮助信息，当用户请求帮助（例如使用 -h 或 --help 标志）时，将显示这个描述信息。
*/
func loadParams() {
	flag.BoolVar(&help, "h", false, "帮助信息")
	flag.StringVar(&uploadFile, "f", "", "上传文件的名称")
	flag.Parse() // 解析命令行参数

	// 判断Cli是否需要打印帮助信息
	if help {
		usage()
		os.Exit(0)
	}
}

func usage() {
	// 打印帮助信息
	fmt.Fprintf(os.Stderr, `cloudstation version: 1.0.0
Usage: cloudstation [-h] -f <upload_file_path>
Options:
`)
	flag.PrintDefaults() // 打印所有的命令行参数
}

func main() {
	// 参数加载
	loadParams()

	// 参数校验
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常：%s\n", err)
		usage()
		os.Exit(1)
	}

	// 文件上传
	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件异常：%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("上传文件成功: %s\n", uploadFile)
}
