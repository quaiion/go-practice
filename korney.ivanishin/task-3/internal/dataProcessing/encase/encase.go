package encase

import (
        "github.com/quaiion/go-practice/convertation/internal/dataProcessing/currency"
        "encoding/json"
        "errors"
        "sort"
        "os"
)

func EncaseJsonData(outFilePath string, currList currency.CurrencyList) error {
        err := prepareOutputEnv(outFilePath)
        if err != nil {
                return errors.Errorf("failed preparing the output environment // %w",
                                     err)
        }

        currList = transformData(currList)
        
        data, err := json.MarshalIdent(currList, ``, ` `)
        if err != nil {
                return errors.Errorf("failed encoding (marshalling) output data // %w",
                                     err)
        }

        err = writeJsonData(data)
        if err != nil {
                return errors.Errorf("failed writing output json data // %w",
                                     err)
        }

        return nil
}

/** separate function just for flexibility */
func prepareOutputEnv(filePath string) error {
        err := os.MkdirAll(filepath.Dir(filePath), 0644)
        if err != nil {
                return nil, errors.Errorf("failed to create specified directories // %w",
                                          err)
        }

        return nil
}

/** separate function just for flexibility */
func transformData(currList currency.CurrencyList) currency.CurrencyList {
        currList.Sort()
        return currList
}

/** separate function just for flexibility */
func writeJsonData(currList currency.CurrencyList) error {
        err = os.WriteFile(outFilePath, data, 0644)
        if err != nil {
                return errors.Errorf("failed writing output data to the file // %w",
                                     err)
        }
}
