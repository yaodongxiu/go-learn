package do

import (
	"testing"
)

func Test_sendData(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data := fetchQuestionNil(1002)

		if data != nil {
			sendData(data)
		}
	})
}
