package conf

type cdn struct {
	AccessKey string
	SecretKey string
	Bucket    string
	UploadUrl string
	Url       string
	LocalPath string
}

var CDN = cdn{
	AccessKey: "SHFEDQXdwn4MO7-4OvW92Qr0LXzEeR-b7si5gpWW",
	SecretKey: "cbCu0I13O2fLIQq8DbLnhUNydDvNdyaSiAJSwnJj",
	Bucket:    "dousheng2022",
	UploadUrl: "https://rcnh4fg5p.hn-bkt.clouddn.com/",
	Url:       "https://cdn.nobugnolife.com/",
	LocalPath: "video/",
}
