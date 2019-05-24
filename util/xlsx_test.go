package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadTasksFromXlsx(t *testing.T) {
	dir := os.Getenv("DATA_HOME")
	LoadTasksFromXlsx("all", filepath.Join(dir, "repo.xlsx"))
}
