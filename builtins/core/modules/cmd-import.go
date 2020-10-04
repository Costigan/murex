package modules

import (
	"bytes"
	"errors"
	"os"
	"strings"

	"github.com/lmorg/murex/builtins/core/httpclient"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/readline"
)

func packageDirExists(pack string) error {
	_, err := os.Stat(pack)
	if os.IsNotExist(err) {
		return nil
	}

	return errors.New("A file or directory already exists with that package name")
}

func importModules(p *lang.Process) error {
	path, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	if path == profile.ModulePath+profile.PackagesFile {
		return errors.New("You cannot import the same file as the master packages.json file")
	}

	if utils.IsURL(path) {
		resp, err := httpclient.Request(p.Context, "GET", path, nil, p.Config, true)
		if err != nil {
			return err
		}

		f, err := utils.NewTempFile(resp.Body, "_package.json")
		if err != nil {
			return err
		}

		path = f.FileName

		defer f.Close()
	}

	importDb, err := readPackagesFile(path)
	if err != nil {
		return err
	}

	db, err := readPackagesFile(profile.ModulePath + profile.PackagesFile)
	if err != nil {
		return err
	}

	for i := range importDb {
		err = p.SetPwd(profile.ModulePath)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			continue
		}

		p.Stderr.Writeln(bytes.Repeat([]byte{'-'}, readline.GetTermWidth()))
		p.Stderr.Writeln([]byte("Importing `" + importDb[i].Package + "`...."))
		err = packageDirExists(importDb[i].Package)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			continue
		}

		importDb[i].Package, _, err = getPackage(p, importDb[i].URI)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			continue
		}

		db = append(db, importDb[i])

		_, err = profile.LoadPackage(profile.ModulePath+importDb[i].Package, true)
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
		}
	}

	var message string

	err = writePackagesFile(&db)
	if err != nil {
		message += err.Error() + utils.NewLineString
	}

	if message != "" {
		return errors.New(strings.TrimSpace(message))
	}

	return nil
}
