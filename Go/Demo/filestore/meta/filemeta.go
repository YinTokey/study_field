package meta


// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	//  初始化 哈希表
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta) {
	
	fileMetas[fmeta.FileSha1] = fmeta
}

func GetFileMeta(fileSha1 string) FileMeta  {
	return fileMetas[fileSha1]
}