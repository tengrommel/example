package meta

import (
	myDB "awesomeProject/filestore-server/db"
	"sort"
)

// FileMeta：文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fMeta FileMeta) {
	fileMetas[fMeta.FileSha1] = fMeta
}

// UpdateFileMetaDB 新增/更新文件到mysql中
func UpdateFileMetaDB(fMeta FileMeta) bool {
	return myDB.OnFileUploadFinished(fMeta.FileSha1, fMeta.FileName, fMeta.FileSize, fMeta.Location)
}

func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tFile, err := myDB.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{}, err
	}
	fMeta := FileMeta{
		FileSha1: tFile.FileHash,
		FileName: tFile.FileName.String,
		FileSize: tFile.FileSize.Int64,
		Location: tFile.FileAddr.String,
	}
	return fMeta, nil
}

// GetFileMeta: 通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetLastFileMetas：获取批量的文件元信息列表
func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}
	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

// RemoveFileMeta：删除文件信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
