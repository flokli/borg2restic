package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type BorgRepo struct {
	Archives []*BorgArchive `json:"archives"`

	mountPoint string
}

func (br *BorgRepo) LoadBorgArchives() error {
	// obtain a listing of the repo
	cmd := exec.Command("borg", "list", "--json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("unable to run borg list: %w", err)
	}

	err = json.Unmarshal(out.Bytes(), &br)
	if err != nil {
		return fmt.Errorf("unable to serialize borg list output: %w", err)
	}

	for _, borgArchive := range br.Archives {
		err := borgArchive.ParseTimestamps()
		if err != nil {
			return fmt.Errorf("unable to parse timestamps for archive %v: %w", borgArchive.ID, err)
		}
	}

	return nil
}

// Mount mounts an repo at the chosen destination path
// archiveName can be left to the empty string, in that case,
// a listing of all archives is provided at the root of the mount
func (br *BorgRepo) Mount(dest string) error {
	if br.mountPoint != "" {
		return fmt.Errorf("already mounted at %v", br.mountPoint)
	}

	args := []string{"mount", "-o", "ignore_permissions", "::", dest}
	cmd := exec.Command("borg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("%+v", args)

	br.mountPoint = dest

	return cmd.Run()
}

// Unmount does unmount the repo.
func (br *BorgRepo) Unmount() error {
	if br.mountPoint == "" {
		return fmt.Errorf("nothing mounted")
	}

	cmd := exec.Command("fusermount", "-u", br.mountPoint)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
