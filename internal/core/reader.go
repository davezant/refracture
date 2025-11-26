package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"refracture/pkg/refr_strings"
	"strings"
)

func (r *Reader) ReplaceInFiles() error {
	for _, s := range r.filePathNew {
		if filepath.Ext(s) == refr_strings.ContainsShow(IterableReadings, filepath.Ext(s)) {
			var content = ""
			file, err := os.Open(s)
			if err != nil {
				return err
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			content = string(data)
			for i := range r.resOld {
				if strings.Contains(content, r.resOld[i]) {
					fmt.Println(s + " Has: " + r.resOld[i])
					fmt.Println("Changing by :" + r.resNew[i])
					content = strings.ReplaceAll(content, r.resOld[i], r.resNew[i])
				}
			}
			for i := range r.filePathOld {
				if strings.Contains(content, r.filePathOld[i]) {
					fmt.Println(s + " Has: " + r.filePathOld[i])
					fmt.Println("Changing by : " + r.filePathNew[i])
					content = strings.ReplaceAll(content, r.filePathOld[i], r.filePathNew[i])
				}
			}
			err = os.WriteFile(s, []byte(content), 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
