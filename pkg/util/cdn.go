package util

import (
	"context"
	"douyin/pkg/constants"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadCdn(filename string) error {
	localFile := constants.CDN.LocalPath + filename
	key := filename

	putPolicy := storage.PutPolicy{
		Scope: constants.CDN.Bucket,
	}
	mac := qbox.NewMac(constants.CDN.AccessKey, constants.CDN.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = constants.CDN.Zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
