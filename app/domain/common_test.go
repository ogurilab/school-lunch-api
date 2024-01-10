package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewResponse(t *testing.T) {

	testCases := []struct {
		name       string
		createStub func() interface{}
		check      func(*Response)
	}{
		{
			name: "OK",
			createStub: func() interface{} {
				return "data"
			},
			check: func(r *Response) {
				require.NotNil(t, r)
				require.Equal(t, "data", r.Data)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := tc.createStub()
			r := NewResponse(data)
			tc.check(r)
		})
	}

}
