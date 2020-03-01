package main

import (
	"archive/zip"
	"io"
	"os"
	"sync"
)

func main() {
	fileNames := os.Args[1:]
	sequenceArch(fileNames)
	conArch(fileNames)
}

func conArch(fileNames []string) {
	waitGroup := sync.WaitGroup{}
	for _, fileName := range fileNames {
		fName := fileName
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup, fileName string) {
			defer func() {
				waitGroup.Done()
			}()
			archiver(fName)
		}(&waitGroup, fileName)
	}
	waitGroup.Wait()
}

func sequenceArch(fileNames []string) {
	for _, fileName := range fileNames {
		archiver(fileName)
	}
}

func archiver(fileName string) {
	archiveFile, err := os.Create(fileName + ".zip")
	if err != nil {
		return

	}
	defer func() {
		err := archiveFile.Close()
		if err != nil {
			return
		}
	}()

	writer := zip.NewWriter(archiveFile)
	defer func() {
		writer.Close()
		if err != nil {
			return

		}
	}()

	file, err := os.Open(fileName)

	defer func() {
		err := file.Close()
		if err != nil {
			return

		}
	}()

	archive, err := writer.Create(fileName) // внутри архива
	if err != nil {
		return

	}
	_, err = io.Copy(archive, file) // справа читает, слева закидывать
	if err != nil {
		return
	}
}
