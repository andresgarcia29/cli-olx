package device

type DeviceFileType string

const (
	File   DeviceFileType = "file"
	Folder DeviceFileType = "folder"
)

type DeviceFiles struct {
	Path              string
	Type              DeviceFileType
	HasIgnoreField    bool
	GenerateFunction  func(string)
	ExecutionFunction func(string) error
}
