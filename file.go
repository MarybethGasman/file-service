package main

import (
	"douyin-file-service/config"
	"github.com/kataras/iris/v12"
	"io"
	"os"
)

var (
	FilePath = config.AppConfig.GetString("video.filePath")
)

type FileController struct {
}

func (fc *FileController) PostUpload(ctx iris.Context) {
	fname := ctx.URLParamDefault("name", "")
	r := ctx.Request()
	file, _, err := r.FormFile("data")
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 1,
			"status_msg":  "没有找到data参数,err: " + err.Error(),
		})
		return
	}
	defer file.Close()

	if b, _ := isHasDir(FilePath); !b {
		err = os.MkdirAll(FilePath, 0777)
		if err != nil {
			ctx.JSON(map[string]interface{}{
				"status_code": 5,
				"status_msg":  "create folder failed,err: " + err.Error(),
			})
			return
		}
	}

	fileName := fname //strconv.FormatInt(time.Now().UnixNano(), 10) + head.Filename

	fw, err := os.Create(FilePath + fileName)

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 2,
			"status_msg":  "create file failed,err: " + err.Error(),
		})
		return
	}
	defer fw.Close()
	_, err = io.Copy(fw, file)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 3,
			"status_msg":  "copy file failed,err: " + err.Error(),
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"status_code": 0,
		"status_msg":  "save file success",
		"file_url":    FilePath + fileName,
	})
}
func isHasDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}
