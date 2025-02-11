package s3

type S3RequestSignUrl struct {
	Key string `json:"key"`
}

type S3ResponseSignUrl struct {
	Url string `json:"signedUrl"`
}
