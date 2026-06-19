package utils

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
	"video_feed/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// OSSClient 阿里云OSS客户端
type OSSClient struct {
	client     *oss.Client
	bucket     *oss.Bucket
	bucketName string
	domain     string
}

var ossClient *OSSClient

// InitOSS 初始化OSS客户端
func InitOSS(cfg *config.OSSConfig) error {
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return fmt.Errorf("初始化OSS客户端失败: %v", err)
	}

	bucket, err := client.Bucket(cfg.BucketName)
	if err != nil {
		return fmt.Errorf("获取Bucket失败: %v", err)
	}

	ossClient = &OSSClient{
		client:     client,
		bucket:     bucket,
		bucketName: cfg.BucketName,
		domain:     cfg.Domain,
	}

	log.Println("阿里云OSS初始化成功")
	return nil
}

// GetOSSClient 获取OSS客户端
func GetOSSClient() *OSSClient {
	return ossClient
}

// UploadFile 上传文件到OSS
func (o *OSSClient) UploadFile(file *multipart.FileHeader, objectKey string) (string, error) {
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close()

	// 上传文件
	err = o.bucket.PutObject(objectKey, src)
	if err != nil {
		return "", fmt.Errorf("上传文件失败: %v", err)
	}

	// 返回访问URL
	return o.GetFileURL(objectKey), nil
}

// UploadVideo 上传视频文件
func (o *OSSClient) UploadVideo(file *multipart.FileHeader) (string, error) {
	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	objectKey := fmt.Sprintf("videos/%d%s", time.Now().UnixNano(), ext)

	return o.UploadFile(file, objectKey)
}

// UploadCover 上传封面图片
func (o *OSSClient) UploadCover(file *multipart.FileHeader) (string, error) {
	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	objectKey := fmt.Sprintf("covers/%d%s", time.Now().UnixNano(), ext)

	return o.UploadFile(file, objectKey)
}

// GetFileURL 获取文件访问URL
func (o *OSSClient) GetFileURL(objectKey string) string {
	// 确保域名以 / 结尾
	domain := o.domain
	if !strings.HasSuffix(domain, "/") {
		domain += "/"
	}
	return domain + objectKey
}

// DeleteFile 删除文件
func (o *OSSClient) DeleteFile(objectKey string) error {
	return o.bucket.DeleteObject(objectKey)
}

// GetObjectKeyFromURL 从URL中提取ObjectKey
func (o *OSSClient) GetObjectKeyFromURL(url string) string {
	domain := o.domain
	if !strings.HasSuffix(domain, "/") {
		domain += "/"
	}
	return strings.TrimPrefix(url, domain)
}

// IsVideoFile 检查是否为视频文件
func IsVideoFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	videoExts := []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv"}
	for _, v := range videoExts {
		if ext == v {
			return true
		}
	}
	return false
}

// IsImageFile 检查是否为图片文件
func IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	for _, v := range imageExts {
		if ext == v {
			return true
		}
	}
	return false
}

// ReadFile 读取文件内容
func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	return io.ReadAll(src)
}
