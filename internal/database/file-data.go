package database

type FileMetadata struct {
	ID          int
	Name        string
	Size        int
	Type        string
	UploadDate  string
	UploaderID  int
	Permissions int
	Path        string
}

type FileContents struct {
	Data []byte
}
