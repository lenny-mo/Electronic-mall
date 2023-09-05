package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadAvatar 上传头像到本地
func UploadAvatar(file multipart.File, fileSize int64, userid uint, username string) (filepath string, err error) {
	bid := strconv.Itoa(int(userid))

	// 创建文件夹
	fileAddr := "./static/avatar/" + "user" + bid + "/"
	if !DirExistorNot(fileAddr) {
		err := CreateDir(fileAddr)
		if err != nil {
			return "", err
		}
	}

	avatarPath := fileAddr + username + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件失败")
		return "", err
	}

	// 创建文件
	err = os.WriteFile(avatarPath, content, os.ModePerm)
	if err != nil {
		fmt.Println("创建文件失败")
		return "", err
	}

	return avatarPath, nil
}

// DirExistorNot 判断文件夹路径是否存在
func DirExistorNot(fileAddr string) bool {
	// 使用 os.Stat() 获取文件或文件夹的信息
	s, err := os.Stat(fileAddr)
	// 如果出现错误，说明文件夹路径不存在
	if err != nil {
		return false
	}
	// 判断获取到的信息是否为文件夹
	return s.IsDir()
}

func CreateDir(fileAddr string) error {
	// 使用 MkdirAll() 创建文件夹
	err := os.MkdirAll(fileAddr, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
