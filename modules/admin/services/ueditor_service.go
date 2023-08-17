package services

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"quince/global"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/google/uuid"
)

// UeditorService struct
type ueditorService struct {
}

func NewUeditorService() ueditorService {
	var cs ueditorService

	return cs
}

// GetConfig
func (*ueditorService) GetConfig() map[string]interface{} {
	ueditorConfig := global.BA_CONFIG.Ueditor
	result := map[string]interface{}{
		"catcherActionName":       ueditorConfig.CatcherActionName,
		"catcherAllowFiles":       strings.Split(ueditorConfig.CatcherAllowFiles, "|"),
		"catcherFieldName":        ueditorConfig.CatcherFieldName,
		"catcherLocalDomain":      strings.Split(ueditorConfig.CatcherLocalDomain, "|"),
		"catcherMaxSize":          ueditorConfig.CatcherMaxSize,
		"catcherPathFormat":       ueditorConfig.CatcherPathFormat,
		"catcherUrlPrefix":        ueditorConfig.CatcherUrlPrefix,
		"fileActionName":          ueditorConfig.FileActionName,
		"fileAllowFiles":          strings.Split(ueditorConfig.FileAllowFiles, "|"),
		"fileFieldName":           ueditorConfig.FileFieldName,
		"fileManagerActionName":   ueditorConfig.FileManagerActionName,
		"fileManagerAllowFiles":   strings.Split(ueditorConfig.FileManagerAllowFiles, "|"),
		"fileManagerListPath":     ueditorConfig.FileManagerListPath,
		"fileManagerListSize":     ueditorConfig.FileManagerListSize,
		"fileManagerUrlPrefix":    ueditorConfig.FileManagerUrlPrefix,
		"fileMaxSize":             ueditorConfig.FileMaxSize,
		"filePathFormat":          ueditorConfig.FilePathFormat,
		"fileUrlPrefix":           ueditorConfig.FileUrlPrefix,
		"imageActionName":         ueditorConfig.ImageActionName,
		"imageAllowFiles":         strings.Split(ueditorConfig.ImageAllowFiles, "|"),
		"imageCompressBorder":     ueditorConfig.ImageCompressBorder,
		"imageCompressEnable":     ueditorConfig.ImageCompressEnable,
		"imageFieldName":          ueditorConfig.ImageFieldName,
		"imageInsertAlign":        ueditorConfig.ImageInsertAlign,
		"imageManagerActionName":  ueditorConfig.ImageManagerActionName,
		"imageManagerAllowFiles":  strings.Split(ueditorConfig.ImageManagerAllowFiles, "|"),
		"imageManagerInsertAlign": ueditorConfig.ImageManagerInsertAlign,
		"imageManagerListPath":    ueditorConfig.ImageManagerListPath,
		"imageManagerListSize":    ueditorConfig.ImageManagerListSize,
		"imageManagerUrlPrefix":   ueditorConfig.ImageManagerUrlPrefix,
		"imageMaxSize":            ueditorConfig.ImageMaxSize,
		"imagePathFormat":         ueditorConfig.ImagePathFormat,
		"imageUrlPrefix":          ueditorConfig.ImageUrlPrefix,
		"scrawlActionName":        ueditorConfig.ScrawlActionName,
		"scrawlAllowFiles":        strings.Split(ueditorConfig.ScrawlAllowFiles, "|"),
		"scrawlFieldName":         ueditorConfig.ScrawlFieldName,
		"scrawlInsertAlign":       ueditorConfig.ScrawlInsertAlign,
		"scrawlMaxSize":           ueditorConfig.ScrawlMaxSize,
		"scrawlPathFormat":        ueditorConfig.ScrawlPathFormat,
		"scrawlUrlPrefix":         ueditorConfig.ScrawlUrlPrefix,
		"snapscreenActionName":    ueditorConfig.SnapscreenActionName,
		"snapscreenInsertAlign":   ueditorConfig.SnapscreenInsertAlign,
		"snapscreenPathFormat":    ueditorConfig.SnapscreenPathFormat,
		"snapscreenUrlPrefix":     ueditorConfig.SnapscreenUrlPrefix,
		"videoActionName":         ueditorConfig.VideoActionName,
		"videoAllowFiles":         strings.Split(ueditorConfig.VideoAllowFiles, "|"),
		"videoFieldName":          ueditorConfig.VideoFieldName,
		"videoMaxSize":            ueditorConfig.VideoMaxSize,
		"videoPathFormat":         ueditorConfig.VideoPathFormat,
		"videoUrlPrefix":          ueditorConfig.VideoUrlPrefix,
	}

	return result
}

// UploadImage
func (us *ueditorService) UploadImage(ctx *context.Context) map[string]interface{} {
	fieldName := global.BA_CONFIG.Ueditor.ImageFieldName
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::imageFieldName.",
		}
	}
	return us.upFile(fieldName, ctx)
}

// upFile
func (*ueditorService) upFile(fieldName string, ctx *context.Context) map[string]interface{} {
	result := make(map[string]interface{})
	file, h, err := ctx.Request.FormFile(fieldName)
	if err != nil {
		result["state"] = err.Error()
		return result
	}
	defer file.Close()

	//Custom file verification
	err = validateForAttachment(h)
	if err != nil {
		result["state"] = err.Error()
		return result
	}

	//
	saveName := uuid.New().String()
	//Suffix. (.png)
	fileExt := path.Ext(h.Filename)
	savePath := "static/uploads/ueditor/" + saveName + fileExt
	saveRealDir := filepath.ToSlash(web.AppPath + "/static/uploads/ueditor/")

	_, err = os.Stat(saveRealDir)
	if err != nil {
		_ = os.MkdirAll(saveRealDir, os.ModePerm)
	}

	saveURL := "/static/uploads/ueditor/" + saveName + fileExt

	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		result["state"] = err.Error()
		return result
	}
	defer f.Close()
	_, err = io.Copy(f, file)

	if err != nil {
		result["state"] = err.Error()
		return result
	}

	result = map[string]interface{}{
		"state":    "SUCCESS",
		"url":      saveURL,
		"title":    saveName + fileExt,
		"original": h.Filename,
		"type":     strings.TrimLeft(fileExt, "."),
		"size":     h.Size,
	}

	return result
}

// ListImage
func (us *ueditorService) ListImage(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	allowFiles := global.BA_CONFIG.Ueditor.ImageManagerAllowFiles
	//ext
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	listSize := global.BA_CONFIG.Ueditor.ImageManagerListSize
	if allowFiles == "" || listSize == "" || len(get) <= 0 {
		result["state"] = "config params error."
	}

	listSizeInt, _ := strconv.Atoi(listSize)
	result = us.fileList(allowFiles, listSizeInt, get)

	return result
}

// fileList
func (us *ueditorService) fileList(allowFiles string, listSize int, get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	_ = result
	var sizeInt, startInt int

	dir := "/static/uploads/ueditor/"

	//
	size := get.Get("size")
	if size == "" {
		sizeInt = listSize
	} else {
		sizeInt, _ = strconv.Atoi(size)
	}

	start := get.Get("start")
	if start == "" {
		startInt = 0
	} else {
		startInt, _ = strconv.Atoi(start)
	}

	end := startInt + sizeInt

	//
	files := us.getFiles(dir, allowFiles)
	if files == nil || len(files) <= 0 {
		result = map[string]interface{}{
			"state": "no match file",
			"list":  map[string]interface{}{},
			"start": startInt,
			"total": 0,
		}

		return result
	}

	//
	lenFiles := len(files)

	result = map[string]interface{}{
		"state": "SUCCESS",
		"start": startInt,
		"total": lenFiles,
	}

	if startInt > lenFiles || end < 0 || startInt > end {
		result["list"] = map[string]interface{}{}
	}

	endInt := 0

	if end > lenFiles {
		endInt = lenFiles
	} else {
		endInt = end
	}

	result["list"] = files[startInt:endInt]

	return result
}

// getFiles
func (us *ueditorService) getFiles(dir, allowFiles string) []map[string]string {
	path := filepath.ToSlash(web.AppPath + dir)
	var filesArr []map[string]string

	if !strings.HasPrefix(path, "/") {
		path = path + "/"
	}

	_, err := os.Stat(path)
	if err != nil {
		return nil
	}

	//files, err := ioutil.ReadDir(path)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			childFilesArr := us.getFiles(filepath.ToSlash(path+file.Name()+"/"), allowFiles)
			if len(childFilesArr) > 0 {
				filesArr = append(filesArr, childFilesArr...)
			}
		} else {
			if !strings.Contains(allowFiles, strings.ToLower(strings.TrimLeft(filepath.Ext(file.Name()), "."))) {
				continue
			}
			fileInfo, _ := file.Info()
			filesArr = append(filesArr, map[string]string{
				"url":   dir + file.Name(),
				"mtime": strconv.Itoa(int(fileInfo.ModTime().Unix())),
			})
		}
	}

	return filesArr
}

// UploadVideo
func (us *ueditorService) UploadVideo(ctx *context.Context) map[string]interface{} {
	fieldName := global.BA_CONFIG.Ueditor.VideoFieldName
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::videoFieldName.",
		}
	}

	return us.upFile(fieldName, ctx)
}

// UploadFile
func (us *ueditorService) UploadFile(ctx *context.Context) map[string]interface{} {
	fieldName := global.BA_CONFIG.Ueditor.FileFieldName
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::fileFieldName.",
		}
	}

	return us.upFile(fieldName, ctx)
}

// ListFile
func (us *ueditorService) ListFile(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	allowFiles := global.BA_CONFIG.Ueditor.FileManagerAllowFiles
	//ext
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	listSize := global.BA_CONFIG.Ueditor.FileManagerListSize
	if allowFiles == "" || listSize == 0 || len(get) <= 0 {
		result["state"] = "config params error."
	}

	listSizeInt := listSize
	result = us.fileList(allowFiles, listSizeInt, get)

	return result
}

// UploadScrawl
func (us *ueditorService) UploadScrawl(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	pathFormat := global.BA_CONFIG.Ueditor.ScrawlPathFormat
	maxSize := global.BA_CONFIG.Ueditor.ScrawlMaxSize
	allowFiles := global.BA_CONFIG.Ueditor.ScrawlAllowFiles
	//ext
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	oriName := global.BA_CONFIG.Ueditor.ScrawlFieldName

	if pathFormat == "" || maxSize == 0 || allowFiles == "" || oriName == "" {
		result["state"] = "config params error."
		return result
	}

	config := map[string]string{
		"pathFormat": pathFormat,
		"maxSize":    strconv.Itoa(maxSize),
		"allowFiles": allowFiles,
		"oriName":    oriName,
	}

	base64Data := get.Get(oriName)
	return us.upBase64(config, base64Data)
}

// upBase64
func (us *ueditorService) upBase64(config map[string]string, base64Data string) map[string]interface{} {
	result := make(map[string]interface{})
	imgByte, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		result["state"] = "picture content get error. err:" + err.Error()
		return result
	}

	path := "/static/uploads/ueditor/scrawl/"
	dirName := filepath.ToSlash(web.AppPath + path)
	file := make(map[string]string)
	file["filesize"] = strconv.Itoa(len(string(imgByte)))
	file["oriName"] = config["oriName"]
	file["ext"] = ".png"
	file["name"] = uuid.New().String() + file["ext"]
	file["fullName"] = dirName + file["name"]
	file["urlName"] = path + file["name"]

	fullName := file["fullName"]

	//Check if the file size exceeds the limit
	fileSizeInt, _ := strconv.Atoi(file["filesize"])
	maxSizeInt, _ := strconv.Atoi(config["maxSize"])
	if fileSizeInt >= maxSizeInt {
		result["state"] = "error.file_too_big"
		return result
	}

	//Failed to create directory
	_, err = os.Stat(dirName)
	if err != nil {
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			result["state"] = "error.create_folder"
			return result
		}
	}

	//
	f, err := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		result["state"] = "Error writing file content,err :" + err.Error()
		return result
	}
	defer f.Close()
	_, err = f.Write(imgByte)
	if err != nil {
		result["state"] = "Failed to write file,err :" + err.Error()
		return result
	}

	result = map[string]interface{}{
		"state":    "SUCCESS",
		"url":      file["urlName"],
		"title":    file["name"],
		"original": file["oriName"],
		"type":     file["ext"],
		"size":     file["filesize"],
	}

	return result
}

// CatchImage
func (us *ueditorService) CatchImage(ctx *context.Context) map[string]interface{} {
	result := make(map[string]interface{})
	pathFormat := global.BA_CONFIG.Ueditor.CatcherPathFormat
	maxSize := global.BA_CONFIG.Ueditor.CatcherMaxSize
	allowFiles := global.BA_CONFIG.Ueditor.CatcherAllowFiles
	//ext
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	oriName := "remote.png"

	if pathFormat == "" || maxSize == 0 || allowFiles == "" {
		result["state"] = "config params error."
		return result
	}

	config := map[string]string{
		"pathFormat": pathFormat,
		"maxSize":    strconv.Itoa(maxSize),
		"allowFiles": allowFiles,
		"oriName":    oriName,
	}

	fieldName := global.BA_CONFIG.Ueditor.CatcherFieldName

	source := make([]string, 0)
	ctx.Input.Bind(&source, fieldName)

	var list []map[string]string
	//
	if len(source) <= 0 {
		result = map[string]interface{}{
			"state": "ERROR",
			"list":  list,
		}
		return result
	}

	for _, imgURL := range source {
		info := us.saveRemote(config, imgURL)
		if info["state"] == "SUCCESS" {
			list = append(list, map[string]string{
				"state":    info["state"],
				"url":      info["url"],
				"size":     info["size"],
				"title":    info["title"],
				"original": info["original"],
				"source":   imgURL,
			})
		} else {
			list = append(list, map[string]string{
				"state":    info["state"],
				"url":      "",
				"size":     "",
				"title":    "",
				"original": "",
				"source":   imgURL,
			})
		}

	}

	result = map[string]interface{}{
		"state": "SUCCESS",
		"list":  list,
	}

	return result
}

// saveRemote
func (*ueditorService) saveRemote(config map[string]string, fieldName string) map[string]string {
	result := make(map[string]string)
	imgURL := strings.ReplaceAll(fieldName, "&amp;", "&")

	if imgURL == "" {
		result["state"] = "error.empty_link"
		return result
	}

	//http
	if !strings.HasPrefix(imgURL, "http") {
		result["state"] = "error.not_http_link"
		return result
	}

	//Get request headers and detect dead links
	response, err := http.Get(imgURL)

	if err != nil || response.StatusCode != 200 {
		result["state"] = "error.link_not_available"
		return result
	}
	defer response.Body.Close()
	//Format verification (extension verification and content-type verification)
	if !strings.Contains(response.Header.Get("Content-Type"), "image") {
		result["state"] = "error.link_content_type"
		return result
	}

	fileType := strings.TrimLeft(filepath.Ext(imgURL), ".")
	if fileType == "" || !strings.Contains(config["allowFiles"], fileType) {
		result["state"] = "error.link_suffix"
		return result
	}

	path := "/static/uploads/ueditor/remote/"
	dirName := filepath.ToSlash(web.AppPath + path)

	file := make(map[string]string)
	file["oriName"] = filepath.Ext(imgURL)
	file["filesize"] = "0"
	file["ext"] = file["oriName"]
	file["name"] = uuid.New().String() + file["ext"]
	file["fullName"] = dirName + file["name"]
	file["urlName"] = path + file["name"]

	//Check if the file size exceeds the limit
	fileSizeInt, _ := strconv.Atoi(file["filesize"])
	maxSizeInt, _ := strconv.Atoi(config["maxSize"])
	if fileSizeInt >= maxSizeInt {
		result["state"] = "error.file_too_big"
		return result
	}

	//
	_, err = os.Stat(dirName)
	if err != nil {
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			result["state"] = "error.create_folder"
			return result
		}
	}

	//写入文件
	img := response.Body
	f, err := os.Create(file["fullName"])

	if err != nil {
		result["state"] = "Failed to write file,err :" + err.Error()
		return result
	}

	w, err := io.Copy(f, img)
	if err != nil {
		result["state"] = "Failed to write file,err :" + err.Error()
		return result
	}

	file["filesize"] = strconv.Itoa(int(w))

	result = map[string]string{
		"state":    "SUCCESS",
		"url":      file["urlName"],
		"title":    file["name"],
		"original": file["oriName"],
		"type":     file["ext"],
		"size":     file["filesize"],
	}

	return result

}
