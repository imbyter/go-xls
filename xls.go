package xls

import (
	"io"
	"os"

	"github.com/imbyter/ole2"
)

//Open one xls file
func Open(file string, charset string) (*WorkBook, error) {
	if fi, err := os.Open(file); err == nil {
		return OpenReader(fi, charset)
	} else {
		return nil, err
	}
}

//Open one xls file and return the closer
func OpenWithCloser(file string, charset string) (*WorkBook, io.Closer, error) {
	if fi, err := os.Open(file); err == nil {
		wb, err := OpenReader(fi, charset)
		return wb, fi, err
	} else {
		return nil, nil, err
	}
}

//Open xls file from reader
func OpenReader(reader io.ReadSeeker, charset string) (wb *WorkBook, err error) {
	var ole *ole2.Ole
	if ole, err = ole2.Open(reader, charset); err == nil {
		var dir []*ole2.File
		if dir, err = ole.ListDir(); err == nil {
			var book *ole2.File
			var root *ole2.File
			for _, file := range dir {
				name := file.Name()
				if name == "Workbook" {
					if book == nil {
						book = file
					}
					//book = file
					// break
				}
				if name == "Book" {
					book = file
					// break
				}
				if name == "Root Entry" {
					root = file
				}
			}
			if book != nil {
				wb = newWorkBookFromOle2(ole.OpenFile(book, root))
				return
			}
		}
	}
	return
}
