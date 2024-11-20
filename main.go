package main

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
	"time"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func buildGo(cCtx *cli.Context) error {
	args := cCtx.Args()
	if args.Len() < 1 {
		return cli.Exit("Usage: go-resource buld resource_dir [output.go]", 1)
	}

	directory, err := filepath.Abs(args.Get(0))
	if err != nil {
		return cli.Exit(err, 3)
	}

	output := os.Stdout
	defer output.Close()

	if args.Len() > 1 {
		ohandle, err := os.OpenFile(args.Get(1), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return cli.Exit(err, 2)
		}
		output = ohandle
	}

	type kv struct {
		Name   string
		Value  string
		BValue []byte
	}

	type cfg struct {
		Package string   `yaml:"package"`
		Ignore  []string `yaml:"ignore"`
		Error   bool     `yaml:"use_errors"`
	}

	type pkg struct {
		Package string
		Fields  []kv
		Config  cfg
	}

	p := pkg{
		Package: "main",
		Fields:  make([]kv, 0, 32),
		Config:  cfg{},
	}

	configName := filepath.Join(directory, "resources.yaml")
	configData, err := os.ReadFile(configName)
	if err == nil {
		err = yaml.Unmarshal(configData, &p.Config)
		if err != nil {
			return cli.Exit(err, 6)
		}
	} else {
		p.Config.Package = "main"
		p.Config.Ignore = []string{
			"resources.yaml",
		}
	}

	err = filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		name := string(path[len(directory)+1:])
		value, err := os.ReadFile(path)
		if err != nil {
			return cli.Exit(err, 3)
		}

		dst := make([]byte, base64.StdEncoding.EncodedLen(len(value)))
		base64.StdEncoding.Encode(dst, value)

		for _, re := range p.Config.Ignore {
			matched, err := regexp.Match(re, []byte(name))
			if err != nil {
				return cli.Exit(err, 7)
			}
			if matched {
				return nil
			}
		}

		p.Fields = append(p.Fields, kv{Name: name, Value: string(dst), BValue: value})
		return nil
	})
	if err != nil {
		return cli.Exit(err, 3)
	}

	t := template.New("gofile")
	t = t.Funcs(template.FuncMap{"isWrap": func(x int) bool { return x%12 == 11 }})

	t, err = t.Parse(Rs("pkg.go.tpl"))
	if err != nil {
		return cli.Exit(err, 4)
	}

	err = t.Execute(output, p)
	if err != nil {
		return cli.Exit(err, 5)
	}

	return nil
}

func createDirectory(cCtx *cli.Context) error {
	args := cCtx.Args()
	if args.Len() != 1 {
		return cli.Exit("Usage: go-resource create <directory>", 1)
	}

	dir := args.Get(0)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return cli.Exit(err, 2)
	}

	name := filepath.Join(dir, "resources.yaml")
	err = os.WriteFile(name, R("resources.yaml.tpl"), 0644)
	if err != nil {
		return cli.Exit(err, 3)
	}

	log.Printf("Directory `%s` created.\n", dir)
	log.Printf("Config `%s` created.\n", name)
	log.Printf("Please edit the file (%s) and place your resources into the directory.\n", name)
	log.Printf("After that run: go-resource build %s <file.go>.\n", dir)
	log.Printf("Note: You can use go:generate tags in Your files.\n")
	log.Printf("example: //go:generate go-resource build %s resources.go\n", dir)
	return nil
}

func run(args []string) {
	app := &cli.App{
		Name:     "go-resource",
		Version:  "0.0.2",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Dmitry E. Oboukhov",
				Email: "unera@uvw.ru",
			},
		},
		Copyright: "(c) 2024",
		Usage:     "golang resource compiler",
		Commands: []*cli.Command{
			&cli.Command{
				Name:      "build",
				Aliases:   []string{"b"},
				Usage:     "build go file contains resources",
				ArgsUsage: "directory filename.go",
				Action:    buildGo,
			},
			&cli.Command{
				Name:      "create",
				Aliases:   []string{"c"},
				Usage:     "create resource infrastructure",
				ArgsUsage: "directory",
				Action:    createDirectory,
			},
			&cli.Command{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "print version",
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("%s\n", cCtx.App.Version)
					return nil
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		log.Fatal(err)
	}

}

func main() {
	run(os.Args)
}

//go:generate go-resource build templates resources.go
