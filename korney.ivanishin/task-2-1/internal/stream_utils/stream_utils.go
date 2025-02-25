package stream_utils

import "fmt"

func ScanBoundReq() (string, int32, error) {
	var (
		boundReqTypeSymb string
		boundReq int32
	)

	_, err := fmt.Scanf("%s %d", &boundReqTypeSymb, &boundReq)
	return boundReqTypeSymb, boundReq, err
}

func ScanInt32() (int32, error) {
	var val int32
	_, err := fmt.Scan(&val)
	return val, err
}

func FlushInput() {
        var flushStr string
        for nFlushed := 1 ; nFlushed != 0 ; nFlushed, _ = fmt.Scanln(&flushStr) {}
}
