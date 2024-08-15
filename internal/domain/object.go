package domain

type Object struct {
	BucketName string
	ObjectName string
	TypeData   string
	Size       int64
	Data       []byte
}
