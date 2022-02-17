package main

import (
	"fmt"
)

func main() {
	dd := &downloadFromDisk{
		password: &mobileTokenDynamics{mobileNumber: "13427495733"},
		filePath: "module 6",
	}
	dd.DownloadFile()
}

type DynamicPassword interface {
	GetPassword() string
}

type mobileTokenDynamics struct {
	mobileNumber string
}

func (mbd *mobileTokenDynamics) GetPassword() string {
	return "1234567"
}

type downloadFromDisk struct {
	password DynamicPassword
	filePath string
}

func (dd *downloadFromDisk) DownloadFile() {
	if err := dd.loginCheck(); err != nil {
		// re-login if password is wrong
	}
	dd.downloadFromAliYun(dd.filePath)
}

func (dd *downloadFromDisk) loginCheck() error {
	err := dd.checkPassword(dd.password.GetPassword())
	if err != nil {
		return err
	}
	return nil
}

func (dd *downloadFromDisk) downloadFromAliYun(file string) error {
	// todo download file from filepath
	return nil
}

func (dd *downloadFromDisk) checkPassword(pw string) error {
	if pw != "1234567" {
		fmt.Errorf("incorrect password")
	}
	return nil
}
