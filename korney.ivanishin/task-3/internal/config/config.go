package config

import (
        "github.com/quaiion/go-practice/convertation/internal/misc"
        "github.com/go-playground/validator/v10"
        "path/filepath"
        "flag"
        "os"
        "io/fs"
        "gopkg.in/yaml.v3"
)

func GetIOFilePaths() (string, string, error) {
        confFilePath, isDefaultPath := parseConfFilePathFlag()

        confFileContents, err := misc.ReadInFile(confFilePath)
        if err != nil {
                if isDefaultPath {
                        return ``, ``, errors.Errorf("failed reading (default) config file data // %w",
                                                     err)
                } else {
                        return ``, ``, errors.Errorf("failed reading config file data // %w",
                                                     err)
                }
        }

        inFilePath, outFilePath, err := decodeConfFileData(confFileContents)
        if err != nil {
                return ``. ``, errors.Errorf("failed processing config file data // %w",
                                             err)
        }

        return inFilePath, outFilePath, nil
}

func parseConfFilePathFlag() (string, bool) {
        var pathStr string
        flag.StringVar(&pathStr, "config", "config.yml", "config file path")
        flag.Parse()
        return pathStr, flag.Parsed()
}

type fileNamesParsed struct {
        inFile  string `yaml:"input-file" validate:"required"`
        outFile string `yaml:"output-file" validate:"required"`
}

func decodeConfFileData(confFileContents []byte) (string, string, error) {
        if confFileContents == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed while opening a file / storing its contents")
        }

        var parsed fileNamesParsed

        err := yaml.Unmarshal(confFileContents, &parsed)
        if err != nil {
                return ``, ``, errors.Errorf("failed unmarshalling // %w", err)
        }

        err = validate.Struct(parsed)
        if err != nil {
                return ``, ``, errors.Errorf("decoded data validation failed // %w",
                                             err)
        }

        return parsed.inFile, parsed.outFile, nil
}
