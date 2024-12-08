package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type BackupInfo struct {
	Name          string
	DownloadPath  string
	Timestamp     string
}

func newBackupInfo(fi os.FileInfo) BackupInfo {
	timestamp := fi.ModTime().Format("2006-01-02 15:04:05")
	dp := fmt.Sprintf("/download/backups/%s", fi.Name())
	
	return BackupInfo {
		Name: fi.Name(),
		DownloadPath: dp,
		Timestamp: timestamp,
	}
}

type PvwConfig struct {
	StaticPath  string
	AssetsPath  string
	DynResPath  string
	Port        int
	SslCertPath string
	SslKeyPath  string
}

// Returns the actual path of the resource inside of the static
// resources directory.
func (p PvwConfig) StaticFilePath(resource string) string {
	return filepath.Join(p.StaticPath, resource)
}

// Returns the actual path of the resource inside of the assets directory.
func (p PvwConfig) AssetFilePath(resource string) string {
	return filepath.Join(p.AssetsPath, resource)
}

// Returns the actual path of the resource inside of the dynamic resources
// directory.
func (p PvwConfig) DynResFilePath(resource string) string {
	return filepath.Join(p.DynResPath, resource)
}

// Returns a slice with the info of each backup.
func (p PvwConfig) GetBackupsInfo() ([]BackupInfo, error) {
	entries, err := os.ReadDir(filepath.Join(p.AssetsPath, "backups"))
	if err != nil {
		return nil, err
	}

	backupsInfo := make([]BackupInfo, 0)
	for _, entry := range entries {

		entryInfo, err := entry.Info()
		if err != nil {
			return nil, err
		}

		backupsInfo = append(backupsInfo, newBackupInfo(entryInfo))
	}

	return backupsInfo, nil
}

