package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	File, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer File.Close()

	info, err := File.Stat()
	if err != nil {
		return err
	}

	if info.Size() < offset {
		return ErrOffsetExceedsFileSize
	}
	if info.Size() == 0 || (info.IsDir()) {
		return ErrUnsupportedFile
	}
	if (limit > info.Size()-offset) || (limit == 0) {
		limit = info.Size() - offset
	}

	_, err = File.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		log.Println("ошибка создания файла", err)
		return err
	}
	defer outFile.Close()

	bar := pb.Full.Start64(limit)
	defer bar.Finish()
	barReader := bar.NewProxyReader(File)

	_, err = io.CopyN(outFile, barReader, limit)
	if err != nil {
		return err
	}
	return nil
}
