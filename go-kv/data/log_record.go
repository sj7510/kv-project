package data

// LogRecordPos 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id，表示数据存储到了那个文件当中
	Offset int64  // 偏移，表示数据存储到了文件的哪个位置
}
