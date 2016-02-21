package storage

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type S3 struct {
	Bucket *s3.Bucket
}

func (s *S3) Init(auth aws.Auth, region aws.Region, bucketName string) error {
	client := s3.New(auth, region)

	s.Bucket = client.Bucket(bucketName)

	_, err := s.Bucket.GetBucketContents()
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchBucket" {
		err = s.Bucket.PutBucket(s3.BucketOwnerFull)
	}

	return err
}

func NewS3(auth aws.Auth, region aws.Region, bucketName string) (*S3, error) {
	s := new(S3)
	return s, s.Init(auth, region, bucketName)
}

func (s *S3) Save(url string) (string, error) {
	if url == "" {
		return "", ErrURLEmpty
	}

	var (
		code string
		err  error
	)

	for i := 0; i < 10; i++ {
		code = getRandomString(8)

		if _, err = s.Bucket.GetKey(code); err != nil {
			err = s.Bucket.Put(code, []byte(url), "text/plain", s3.BucketOwnerFull)
			if err == nil {
				return code, err
			}
		}
	}

	return "", ErrCodeNotSet
}

func (s *S3) SaveName(code string, url string) error {
	if code == "" {
		return ErrNameEmpty
	}
	if url == "" {
		return ErrURLEmpty
	}

	return s.Bucket.Put(code, []byte(url), "text/plain", s3.BucketOwnerFull)
}

func (s *S3) Load(code string) (string, error) {
	if code == "" {
		return "", ErrNameEmpty
	}

	url, err := s.Bucket.Get(code)
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchKey" {
		return "", ErrCodeNotSet
	}
	return string(url), err
}
