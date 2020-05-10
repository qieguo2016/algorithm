/*
	ring buff，an implement of cache
*/

package array

import (
	"errors"
	"io"
)

var ErrOutOfRange = errors.New("out of range")


// ring buffer底层是一个数组，使用单调递增的左右下标来标识虚拟的存储空间，写到数组末尾之后再从头部开始写
type RingBuf struct {
	begin int64 // 虚拟存储空间左边界
	end   int64 // 虚拟存储空间右边界
	data  []byte
	index int  // 数组可写下标，另外一种常见的优化是控制数组长度为2的n次方，然后用位操作代替取模
}

func NewRingBuf(size int) (rb RingBuf) {
	rb.data = make([]byte, size)
	rb.begin = 0
	rb.end = 0
	rb.index = 0
	return
}

func (rb *RingBuf) Size() int64 {
	return int64(len(rb.data))
}

func (rb *RingBuf) Begin() int64 {
	return rb.begin
}

func (rb *RingBuf) End() int64 {
	return rb.end
}

// 从尾部写入，写满数组后覆盖写头部，达到ring效果
func (rb *RingBuf) Write(p []byte) (n int, err error) {
	if len(p) > len(rb.data) {
		err = ErrOutOfRange
		return
	}
	for n < len(p) {
		written := copy(rb.data[rb.index:], p[n:])
		rb.end += int64(written)
		n += written
		rb.index += written
		if rb.index >= len(rb.data) {  // 回到头部
			rb.index -= len(rb.data)
		}
	}
	if int(rb.end-rb.begin) > len(rb.data) {
		rb.begin = rb.end - int64(len(rb.data))
	}
	return
}

// 从指定位置写入，写满数组后覆盖写头部，达到ring效果
func (rb *RingBuf) WriteAt(p []byte, offset int64) (n int, err error) {
	if offset+int64(len(p)) > rb.end || offset < rb.begin {
		err = ErrOutOfRange
		return
	}
	var writeOff int
	if rb.end-rb.begin < int64(len(rb.data)) {
		writeOff = int(offset - rb.begin)
	} else {
		writeOff = rb.index + int(offset-rb.begin)
	}
	if writeOff > len(rb.data) {
		writeOff -= len(rb.data)
	}
	writeEnd := writeOff + int(rb.end-offset)
	if writeEnd <= len(rb.data) {
		n = copy(rb.data[writeOff:writeEnd], p)
	} else {
		n = copy(rb.data[writeOff:], p)
		if n < len(p) {
			n += copy(rb.data[:writeEnd-len(rb.data)], p[n:])
		}
	}
	return
}

func (rb *RingBuf) ReadAt(p []byte, offset int64) (n int, err error) {
	if offset > rb.end || offset < rb.begin {
		err = ErrOutOfRange
		return
	}
	var readOff int
	if rb.end-rb.begin < int64(len(rb.data)) {
		readOff = int(offset - rb.begin)
	} else {
		readOff = rb.index + int(offset-rb.begin)
	}
	if readOff >= len(rb.data) {
		readOff -= len(rb.data)
	}
	readEnd := readOff + int(rb.end-offset)
	if readEnd <= len(rb.data) {
		n = copy(p, rb.data[readOff:readEnd])
	} else {
		n = copy(p, rb.data[readOff:])
		if n < len(p) {
			n += copy(p[n:], rb.data[:readEnd-len(rb.data)])
		}
	}
	if n < len(p) {
		err = io.EOF
	}
	return
}

func (rb *RingBuf) Evacuate(offset int64, length int) (newOff int64) {
	if offset+int64(length) > rb.end || offset < rb.begin {
		return -1
	}
	var readOff int
	if rb.end-rb.begin < int64(len(rb.data)) {
		readOff = int(offset - rb.begin)
	} else {
		readOff = rb.index + int(offset-rb.begin)
	}
	if readOff >= len(rb.data) {
		readOff -= len(rb.data)
	}

	if readOff == rb.index {
		// no copy evacuate
		rb.index += length
		if rb.index >= len(rb.data) {
			rb.index -= len(rb.data)
		}
	} else if readOff < rb.index {
		var n = copy(rb.data[rb.index:], rb.data[readOff:readOff+length])
		rb.index += n
		if rb.index == len(rb.data) {
			rb.index = copy(rb.data, rb.data[readOff+n:readOff+length])
		}
	} else {
		var readEnd = readOff + length
		var n int
		if readEnd <= len(rb.data) {
			n = copy(rb.data[rb.index:], rb.data[readOff:readEnd])
			rb.index += n
		} else {
			n = copy(rb.data[rb.index:], rb.data[readOff:])
			rb.index += n
			var tail = length - n
			n = copy(rb.data[rb.index:], rb.data[:tail])
			rb.index += n
			if rb.index == len(rb.data) {
				rb.index = copy(rb.data, rb.data[n:tail])
			}
		}
	}
	newOff = rb.end
	rb.end += int64(length)
	if rb.begin < rb.end-int64(len(rb.data)) {
		rb.begin = rb.end - int64(len(rb.data))
	}
	return
}

func (rb *RingBuf) Resize(newSize int) {
	if len(rb.data) == newSize {
		return
	}
	newData := make([]byte, newSize)
	var offset int
	if rb.end-rb.begin == int64(len(rb.data)) {
		offset = rb.index
	}
	if int(rb.end-rb.begin) > newSize {
		discard := int(rb.end-rb.begin) - newSize
		offset = (offset + discard) % len(rb.data)
		rb.begin = rb.end - int64(newSize)
	}
	n := copy(newData, rb.data[offset:])
	if n < newSize {
		copy(newData[n:], rb.data[:offset])
	}
	rb.data = newData
	rb.index = 0
}

func (rb *RingBuf) Skip(length int64) {
	rb.end += length
	rb.index += int(length)
	for rb.index >= len(rb.data) {
		rb.index -= len(rb.data)
	}
	if int(rb.end-rb.begin) > len(rb.data) {
		rb.begin = rb.end - int64(len(rb.data))
	}
}