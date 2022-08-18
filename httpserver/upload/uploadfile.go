package upload

import (
	"awesomeProject/config/global"
	"awesomeProject/pkg/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type Filetype int

// 获取文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

// 获取文件路径
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 上传文件路径
func GetSavePath() string {
	return global.UploadSetting.UploadSavePath
}

// 检查上传文件中是否有相同文件
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查文件大小
func CheckMaxSize(f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	if size >= global.UploadSetting.UploadImageMaxSize*1024*1024 {
		return true
	}
	return false
}

// 检查文件权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建保存路径
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// 保存文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
