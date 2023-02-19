package constants

import "github.com/qiniu/go-sdk/v7/storage"

type cdn struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Url       string
	LocalPath string
	Zone      *storage.Zone
}

var CDN = cdn{
	AccessKey: "SHFEDQXdwn4MO7-4OvW92Qr0LXzEeR-b7si5gpWW",
	SecretKey: "cbCu0I13O2fLIQq8DbLnhUNydDvNdyaSiAJSwnJj",
	Bucket:    "douyin-6379",
	Url:       "http://cdn.nobugnolife.com/",
	LocalPath: "./tmp/videos/",
	Zone:      &storage.ZoneHuadong,
}
