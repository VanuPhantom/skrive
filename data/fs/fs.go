package fs

import (
	"log"
	"os"
	"path"
	"skrive/data"
	"sort"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type FsStorage struct {
	Path string
}

func GetPath(fileArgument string) (*string, error) {
	if len(fileArgument) > 0 {
		return &fileArgument, nil
	}

	if value, isDefined := os.LookupEnv("SKRIVE_DOSES_PATH"); isDefined {
		return &value, nil
	}

	if dirname, err := os.UserHomeDir(); err != nil {
		return nil, homePathError{}
	} else {
		var p = path.Join(dirname, "doses.dat")
		return &p, nil
	}
}

type homePathError struct{}

func (e homePathError) Error() string {
	return "Could not find a home directory."
}

func (storage FsStorage) Append(dose data.Dose) error {
	if err := storage.probeAndMigrate(); err != nil {
		return err
	}

	file, err := os.OpenFile(storage.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err == nil && len(dose.Id) == 0 {
		var id string
		id, err = gonanoid.New()
		dose.Id = data.Id(id)
	}

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
	if err := storage.probeAndMigrate(); err != nil {
		return nil, err
	}

	if bytes, err := os.ReadFile(storage.Path); err != nil {
		return nil, err
	} else {
		raw := string(bytes)
		withoutHeader := strings.SplitN(raw, "Version:1\n", 2)[1]
		if doses, err := parseVersion1(withoutHeader); err != nil {
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
		file.WriteString("Version:1\n")
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
		err = os.Rename(file.Name(), storage.Path)
	}

	return err
}

func (storage FsStorage) DeleteDose(id data.Id) error {
	original, err := storage.FetchAll()

	if err != nil {
		return err
	}

	result := make([]data.Dose, 0)

	for _, dose := range original {
		if dose.Id != id {
			result = append(result, dose)
		}
	}

	return storage.overwrite(result)
}

func (storage FsStorage) probeAndMigrate() error {
	file, err := os.Open(storage.Path)

	var header string
	if err == nil {
		var buf = make([]byte, 10)
		_, err = file.Read(buf)
		header = string(buf)
	}

	if err != nil {
		return nil
	} else if !strings.HasPrefix(header, "Version:") {
		println("Migrating data from version 0 to version 1")
		var bytes []byte
		bytes, err = os.ReadFile(storage.Path)
		var doses []data.Dose

		if err == nil {
			doses, err = parseHeaderless(string(bytes))
		}

		if err == nil {
			err = storage.overwrite(doses)
		}

		return err
	} else if strings.HasPrefix("Version:1\n", header) {
		return nil
	} else {
		return DecodeError{Kind: UNKNOWN_HEADER, context: &strings.SplitN(header, "\n", 1)[0]}
	}
}
