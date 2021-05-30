package dao

import (
	"service/internal/model"
	"time"
)

func (d *Dao) UploadFile(filename, desc, createdBy, accessUrl string, filetype int) error {
	file := model.File{
		Filename:      filename,
		Desc:          desc,
		CreatedBy:     createdBy,
		CreatedTime:   time.Now().Unix(),
		FileAccessUrl: accessUrl,
		Type:          filetype,
	}
	return file.Create(d.engine)
}
