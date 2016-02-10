package main

import (
	"archive/zip"
	"os"
	"path/filepath"
	"io"
	"flag"
)

const ptrSize = 32 << uintptr(^uintptr(0) >> 63)

func main() {
	target := flag.String("t", "", "Provide installation target directory. Mandatory.")
	installServer(*target)
}

func installServer(target string) {
	installDataSource(target)
}

func installDataSource(target string) {
	installPostgre(target)
}

func installPostgre(target string) {
	getOSType().accept(PostgreUnzipper{target: target})
}

func getOSType() osType {
	if ptrSize == 32 {
		return x86{}
	}
	return x64{}
}

type PostgreUnzipper struct {
	target string
}

func (provider PostgreUnzipper) visitx86() {
	unzip("vendor/postgresql-9.4.5-3-windows-binaries.zip", provider.getTarget())
}

func (provider PostgreUnzipper) visitx64() {
	unzip("vendor/postgresql-9.4.5-3-windows-x64-binaries.zip", provider.getTarget())
}

func (provider PostgreUnzipper) getTarget() string {
	return filepath.Join(provider.target, "lib")
}

func unzip(archiveName, target string) error {
	reader, err := zip.OpenReader(archiveName)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		err := writeArchiveElementOnTarget(file, target)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeArchiveElementOnTarget(file *zip.File, target string) error {
	path := filepath.Join(target, file.Name)

	if file.FileInfo().IsDir() {
		os.MkdirAll(path, file.Mode())
		return nil
	}

	return writeArchiveFileOnTarget(file, target, path)
}

func writeArchiveFileOnTarget(file *zip.File, target, path string) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	targetFile, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, fileReader); err != nil {
		return err
	}
	return nil
}

type osType interface {
	accept(visitor osTypeVisitor)
}

type x86 struct{}
type x64 struct{}

func (curOs x86) accept(visitor osTypeVisitor) {
	visitor.visitx86()
}

func (curOs x64) accept(visitor osTypeVisitor) {
	visitor.visitx64()
}

type osTypeVisitor interface {
	visitx86()
	visitx64()
}