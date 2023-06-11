package githubv4

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/jbrekelmans/go-graphql"
	"github.com/stretchr/testify/assert"
)

func Test_enhanceError(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		err := errors.New("some err")
		actual := enhanceError(err)
		assert.Same(t, err, actual)
	})
	t.Run("Case2", func(t *testing.T) {
		err1 := errors.New("some err")
		err2 := &graphql.Error{
			Err:     err1,
			Message: "msg",
		}
		err3 := enhanceError(err2)
		if assert.IsType(t, &Error{}, err3) {
			assert.Equal(t, &Error{
				Err:     err1,
				Message: "msg",
			}, err3)
			assert.Same(t, err1, err3.(*Error).Err)
		}
	})
	t.Run("Case3", func(t *testing.T) {
		base := &graphql.Error{
			Message: "msg1",
			Errors: []graphql.ErrorItem{
				{
					Message: "msg2",
				},
			},
		}
		actual := enhanceError(base)
		assert.Equal(t, &Error{
			Message: "msg1",
			Errors: []ErrorItem{
				{
					Message: "msg2",
				},
			},
		}, actual)
	})
}

func Test_enhanceErrorItem(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		raw := map[string]json.RawMessage{
			"extensions": json.RawMessage(`{"ext1":"ext1Val","ext2":["ext2Val"]}`),
			"type":       json.RawMessage(`"type123"`),
			"something":  json.RawMessage(`123`),
		}
		base := graphql.ErrorItem{
			Message: "msg",
			Raw:     raw,
		}
		actual := enhanceErrorItem(base)
		expected := ErrorItem{
			Extensions: map[string]any{
				"ext1": "ext1Val",
				"ext2": []any{"ext2Val"},
			},
			Message: "msg",
			Raw: map[string]json.RawMessage{
				"something": json.RawMessage(`123`),
			},
			Type: "type123",
		}
		assert.Equal(t, expected, actual)
	})
}

func Test_tryGet(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		s := tryGet[string](nil, "k")
		assert.Equal(t, "", s)
	})
	t.Run("Case2", func(t *testing.T) {
		raw := map[string]json.RawMessage{
			"k": json.RawMessage("3"),
		}
		s := tryGet[string](raw, "k")
		assert.Equal(t, "", s)
		assert.Len(t, raw, 1)
	})
	t.Run("Case3", func(t *testing.T) {
		raw := map[string]json.RawMessage{
			"k": json.RawMessage(`"v"`),
		}
		s := tryGet[string](raw, "k")
		assert.Equal(t, "v", s)
		assert.Len(t, raw, 0)
	})
	t.Run("Case4", func(t *testing.T) {
		raw := map[string]json.RawMessage{
			"k1": json.RawMessage(`"v1"`),
			"k2": json.RawMessage(`null`),
		}
		s := tryGet[string](raw, "k1")
		assert.Equal(t, "v1", s)
		assert.Equal(t, map[string]json.RawMessage{
			"k2": json.RawMessage(`null`),
		}, raw)
	})
}
