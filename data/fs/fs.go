package fs

import (
	"log"
	"os"
	"skrive/data"
	"sort"
)

type FsStorage struct {
	path string
}

func (storage FsStorage) Append(dose data.Dose) error {
	file, err := os.OpenFile(storage.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if _, err := file.WriteString(encode(dose) + "\n"); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func (storage FsStorage) FetchAll() ([]data.Dose, error) {
	if bytes, err := os.ReadFile(storage.path); err != nil {
		return nil, err
	} else {
		raw := string(bytes)

		if doses, err := decode(raw); err != nil {
			return nil, err
		} else {
			sort.Slice(doses, func(i, j int) bool {
				return doses[i].Time.Unix() > doses[j].Time.Unix()
			})

			return doses, nil
		}
	}
}

func (storage FsStorage) overwrite(doses []data.Dose) error {
	file, err := os.CreateTemp("", "skrive-tmp")
	if err == nil {
		err = file.Chmod(0600)
	}

	if err == nil {
		for i := range doses {
			if _, err = file.WriteString(encode(doses[i]) + "\n"); err != nil {
				break
			}
		}
	}

	if err == nil {
		err = file.Sync()
	}

	closeErr := file.Close()
	if err == nil && closeErr != nil {
		return closeErr
	}

	if err == nil {
		err = os.Rename(file.Name(), storage.path)
	}

	return err
}

func (storage FsStorage) DeleteDose(id data.Id) error {
	original, err := storage.FetchAll()

	if err != nil {
		return err
	}

	result := make([]data.Dose, len(original)-1)

	for _, dose := range original {
		if dose.Id != id {
			result = append(result, dose)
		}
	}

	return storage.overwrite(result)
}
