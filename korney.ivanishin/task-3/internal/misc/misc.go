package misc

import (
        "errors"
)

func ReadInFile(filePath string) ([]byte, error) {
        fileData, err := os.ReadFile(filePath)
        if err != nil {
                var pathErr *fs.PathErr
                if errors.As(err, &pathErr) {
                        err = errors.Errorf("failed to open config file path: %w",
                                            pathErr.Path)
                }

                return nil, err
        }

        return fileData, nil
}
