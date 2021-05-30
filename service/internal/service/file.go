package service

import (
	"errors"
	"mime/multipart"
	"os"
	"service/global"
	"service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// type UploadFileRequest struct {
// 	Filename string `json:"filename"`
// 	Desc string `json:"desc"`
// 	FileAccessUrl string `json:"file_access_url"`
// 	CreatedTime int64 `json:"created_time"`
// 	CreatedBy string `json:"created_by"`
// }

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	// fileName := upload.GetFileName(fileHeader.Filename)
	fileName := fileHeader.Filename // 不对文件名加密
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	// 将文件上传记录写入数据库Î
	// TODO: 文件上传写入数据库信息待完善
	// TODO: 上传重名文件
	err := svc.dao.UploadFile(fileName, "", "hyh", accessUrl, int(fileType))
	if err != nil {
		return nil, err
	}

	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
