package storage

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type S3 struct {
	Bucket *s3.Bucket
}

func NewS3(auth aws.Auth, region aws.Region, bucketName string) (*S3, error) {
	s := &S3{
		Bucket: s3.New(auth, region).Bucket(bucketName),
	}

	_, err := s.Bucket.GetBucketContents()
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchBucket" {
		err = s.Bucket.PutBucket(s3.BucketOwnerFull)
	}

	return s, err
}

func (s *S3) Save(url string) (string, error) {
	if _, err := validateURL(url); err != nil {
		return "", err
	}

	var (
		short string
		err   error
	)

	for i := 0; i < 10; i++ {
		short = getRandomString(8)

		if _, err = s.Bucket.GetKey(short); err != nil {
			err = s.Bucket.Put(short, []byte(url), "text/plain", s3.BucketOwnerFull)
			if err == nil {
				return short, nil
			}
		}
	}

	return "", ErrShortNotSet
}

func (s *S3) SaveName(short string, url string) error {
	if err := validateShort(short); err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	return s.Bucket.Put(short, []byte(url), "text/plain", s3.BucketOwnerFull)
}

func (s *S3) Load(short string) (string, error) {
	if err := validateShort(short); err != nil {
		return "", err
	}

	url, err := s.Bucket.Get(short)
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchKey" {
		return "", ErrShortNotSet
	}
	return string(url), err
}
