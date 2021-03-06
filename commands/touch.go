package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"errors"
)

func Touch(currDir entities.Dir, args ...string) error {
	if len(args) == 0 {
		return errors.New("touch: missing file operand")
	}

	for _, path := range args {
		var dest entities.Dir
		var err error

		dirPath, filename := utils.GetDest(path)
		dest, err = utils.Unpath(dirPath, currDir)
		if err != nil {
			return err
		}
		_, err = entities.NewFile(filename, dest)
		if err != nil {
			return err
		}
	}

	return nil
}
