package extract

import (
        "github.com/quaiion/go-practice/convertation/internal/misc"
        "github.com/quaiion/go-practice/convertation/internal/dataProcessing/currency"
        "encoding/xml"
        "strings"
        "strconv"
        "errors"
        "golang.org/x/net/html/charset"
        "github.com/go-playground/validator/v10"
)

func ExtractXmlData(inFilePath string) (currency.CurrencyList, error) {
        fileData, err := misc.ReadInFile(inFilePath)
        if err != nil {
                return nil, errors.Errorf("failed reading input file data // %w",
                                          err)
        }

        data, err := decodeXmlFileData(fileData)
        if err != nil {
                return nil, errors.Errorf("failed decoding input file data // %w",
                                          err)
        }

        return data, nil
}

func decodeXmlFileData(fileData []byte) (currency.CurrencyList, error) {
        /** for now encoding is hardcoded based on the input files' format */
        charsetReader, err := charset.NewReaderLabel(`windows-1251`, fileData)
        if err != nil {
                return nil, errors.Errorf("failed creating a charset reader // %w",
                                          err)
        }

        currList := []currency.Currency

        decoder := xml.NewDecoder(charsetReader)
        for token, err := decoder.Token() ; token != nil ; token, err = decoder.Token() {
                if err != nil {
                        return nil, errors.Errorf("failed parsing a token from input data // %w",
                                                  err)
                }

                tokenType := token.(type)
                if tokenType == xml.StartElement && tokenType.Name.Local == `Valute` {
                        var curr currency.Currency
                        err = decoder.DecodeElement(&curr, &tokenType)
                        if err != nil {
                                return nil, errors.Errorf("failed decoding an xml currency record // %w",
                                                          err)
                        }

                        err = validate.Struct(curr)
                        if err != nil {
                                /**
                                 * we only validate the "required" condition and
                                 * just discard elements that dont satisfy it,
                                 * so there is no need to pass this error upwards
                                 */
                                 continue
                        }

                        curr.ValueStr = strings.ReplaceAll(curr.ValueStr, `,`, `.`)
                        curr.Value, err = strconv.ParseFloat(curr.ValueStr, 64)
                        if err != nil {
                                return nil, errors.Errorf("failed converting a 'Value' record to float // %w",
                                                          err)
                        }

                        currList = append(currList, curr)
                }
        }

        return currList, nil
}
