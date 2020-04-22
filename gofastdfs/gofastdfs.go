package gofastdfs

import (
	`fmt`
	`github.com/astaxie/beego/httplib`
	`mime/multipart`
)

func UploadFiles(file *multipart.FileHeader)  {
		var obj interface{}
		req := httplib.Post("http://49.234.235.249:3666/group1/upload")

		req.PostFile("file", file.Filename) //注意不是全路径
		req.Param("output", "json")
		req.Param("scene", "")
		req.Param("path", "")
		req.ToJSON(&obj)
		fmt.Println(obj)
}