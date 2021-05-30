package model

import "github.com/jinzhu/gorm"

type File struct {
	ID            uint32 `gorm:"primary_key" json:"id"`
	Filename      string `json:"filename"`
	Desc          string `json:"desc"`
	FileAccessUrl string `json:"file_access_url"`
	Type          int    `json:"type"`

	CreatedTime  int64  `json:"created_time"`
	CreatedBy    string `json:"created_by"`
	ModifiedTime int64  `json:"modified_time"`
	ModifiedBy   string `json:"modified_by"`
	DeletedTime  int64  `json:"deleted_time"`
	Status       uint   `json:"status"`
}

func (f File) TableName() string {
	return "files"
}

func (f File) GetFileByName(db *gorm.DB) ([]*File, error) {
	var files []*File
	err := db.Where("filename = ?", f.Filename).Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (f File) GetFileByAuthor(db *gorm.DB) ([]*File, error) {
	var files []*File
	err := db.Where("created_by = ?", f.CreatedBy).Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Create 文件上传记录写数据库
func (f File) Create(db *gorm.DB) error {
	return db.Create(&f).Error
}

// Update 更新文件表
func (f File) Update(db *gorm.DB, values interface{}) error {
	// 根据条件和 model 的值进行更新
	// db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
	err := db.Model(f).Where("id = ? AND status = ?", f.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (f File) Delete(db *gorm.DB) error {
	err := db.Where("id = ? AND status = ?", f.ID, 1).Delete(&f).Error
	if err != nil {
		return err
	}
	return nil
}
