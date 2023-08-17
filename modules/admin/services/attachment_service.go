package services

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"quince/global"
	"quince/modules/admin/models"
	"quince/utils"
	"strconv"
	"strings"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/google/uuid"
)

// AttachmentService struct
type attachmentService struct {
	BaseService
}

//AttachmentService - instantiate de IModel filter
func NewAttachmentService() attachmentService {
	var cs attachmentService
	c := models.Attachment{}

	cs.IModel = &c
	return cs
}

// Upload
func (as *attachmentService) Upload(ctx *context.Context, name string, adminUserId int64, userId int64) (*models.Attachment, error) {
	file, h, err := ctx.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Custom file verification
	err = validateForAttachment(h)
	if err != nil {
		return nil, err
	}

	//Data table write
	saveName := uuid.New().String()
	//Suffix . (.png)
	fileExt := path.Ext(h.Filename)
	savePath := global.BA_CONFIG.Attachment.Path + saveName + fileExt
	saveRealDir := filepath.ToSlash(web.AppPath + "/" + global.BA_CONFIG.Attachment.Path)

	_, err = os.Stat(saveRealDir)
	if err != nil {
		_ = os.MkdirAll(saveRealDir, os.ModePerm)
	}

	saveURL := "/" + global.BA_CONFIG.Attachment.Url + saveName + fileExt

	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	io.Copy(f, file)

	attachmentInfo := models.Attachment{
		LoginUser:    &models.LoginUser{Base: models.Base{Id: adminUserId}},
		User:         &models.User{Base: models.Base{Id: userId}},
		OriginalName: h.Filename,
		SaveName:     saveName,
		SavePath:     saveRealDir + saveName + fileExt,
		Url:          saveURL,
		Extension:    strings.TrimLeft(fileExt, "."),
		Mime:         h.Header.Get("Content-Type"),
		Size:         h.Size,
		Md5:          utils.GetMd5String(saveName),
		Sha1:         utils.GetSha1String(saveName),
	}

	insertID, err := as.DataManipulation().Insert(&attachmentInfo)
	if err == nil {
		attachmentInfo.Id = insertID
		return &attachmentInfo, nil
	}
	return nil, err
}

// UploadMulti
func (as *attachmentService) UploadMulti(ctx *context.Context, name string, adminUserId int64, userId int64) ([]*models.Attachment, error) {
	var result []*models.Attachment
	//GetFiles return multi-upload files
	files, ok := ctx.Request.MultipartForm.File[name]
	if !ok {
		return nil, http.ErrMissingFile
	}

	for i := range files {
		h := files[i]
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()
		////create destination file making sure the path is writeable.
		//dst, err := os.Create("upload/" + files[i].Filename)
		//defer dst.Close()
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		////copy the uploaded file to the destination file
		//if _, err := io.Copy(dst, file); err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		//
		err = validateForAttachment(h)
		if err != nil {
			return nil, err
		}

		//Data table write
		saveName := uuid.New().String()
		//Suffix. (.png)
		fileExt := path.Ext(h.Filename)
		savePath := global.BA_CONFIG.Attachment.Path + saveName + fileExt
		saveRealDir := filepath.ToSlash(web.AppPath + "/" + global.BA_CONFIG.Attachment.Path)

		_, err = os.Stat(saveRealDir)
		if err != nil {
			_ = os.MkdirAll(saveRealDir, os.ModePerm)
		}

		saveURL := "/" + global.BA_CONFIG.Attachment.Url + saveName + fileExt

		f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		io.Copy(f, file)

		attachmentInfo := models.Attachment{
			LoginUser:    &models.LoginUser{Base: models.Base{Id: adminUserId}},
			User:         &models.User{Base: models.Base{Id: userId}},
			OriginalName: h.Filename,
			SaveName:     saveName,
			SavePath:     saveRealDir + saveName + fileExt,
			Url:          saveURL,
			Extension:    strings.TrimLeft(fileExt, "."),
			Mime:         h.Header.Get("Content-Type"),
			Size:         h.Size,
			Md5:          utils.GetMd5String(saveName),
			Sha1:         utils.GetSha1String(saveName),
		}

		insertID, err := as.DataManipulation().Insert(&attachmentInfo)
		if err == nil {
			attachmentInfo.Id = insertID
			//return &attachmentInfo, nil
			result = append(result, &attachmentInfo)
		} else {
			return nil, err
		}
	}

	//Return the object information of the uploaded file
	if result != nil {
		return result, nil
	}
	return nil, errors.New("error.file_access")

}

// validateForAttachment
func validateForAttachment(h *multipart.FileHeader) error {
	validateSize, _ := strconv.Atoi(global.BA_CONFIG.Attachment.ValidateSize)
	validateExt := global.BA_CONFIG.Attachment.ValidateExt
	if int(h.Size) > validateSize {
		return errors.New("error.file_too_big")
	}

	if !strings.Contains(validateExt, strings.ToLower(strings.TrimLeft(path.Ext(h.Filename), "."))) {
		return errors.New("error.file_type_unsupported")
	}

	return nil
}
