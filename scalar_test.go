package githubv4

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Base64String(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := Base64String{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x Base64String
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}

func Test_BigInt(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x BigInt
			x.N.SetInt64(-3)
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`-3`), b)
			}
		})
		t.Run("Case2", func(t *testing.T) {
			var x BigInt
			x.N.SetInt64(100000)
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`100000`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x BigInt
			err := x.UnmarshalText([]byte(`-4`))
			if assert.NoError(t, err) {
				expected := big.NewInt(-4)
				assert.True(t, x.N.Cmp(expected) == 0)
			}
		})
		t.Run("Case2", func(t *testing.T) {
			var x BigInt
			err := x.UnmarshalText([]byte(`100009`))
			if assert.NoError(t, err) {
				expected := big.NewInt(100009)
				assert.True(t, x.N.Cmp(expected) == 0)
			}
		})
		t.Run("Case3", func(t *testing.T) {
			var x BigInt
			err := x.UnmarshalText([]byte(`+5`))
			if assert.NoError(t, err) {
				expected := big.NewInt(5)
				assert.True(t, x.N.Cmp(expected) == 0)
			}
		})
	})
}

func Test_Date(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := Date{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x Date
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}

func Test_DateTime(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := DateTime{time.Date(2023, 6, 1, 23, 59, 59, 0, time.UTC)}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`2023-06-01T23:59:59Z`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x DateTime
			err := x.UnmarshalText([]byte(`2023-06-01T23:59:59Z`))
			if assert.NoError(t, err) {
				expected := time.Date(2023, 6, 1, 23, 59, 59, 0, time.UTC)
				assert.True(t, expected.Equal(x.Time))
			}
		})
	})
}

func Test_GitObjectID(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := GitObjectID{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x GitObjectID
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}

func Test_GitTimestamp(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := GitTimestamp{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x GitTimestamp
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}

func Test_GitSSHRemote(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := GitSSHRemote{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x GitSSHRemote
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}

func Test_HTML(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := HTML{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x HTML
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}

func Test_ID(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		x := ID{S: "x"}
		b, err := json.Marshal(x)
		if assert.NoError(t, err) {
			assert.Equal(t, `"x"`, string(b))
		}
	})
	t.Run("Case2", func(t *testing.T) {
		var x ID
		err := json.Unmarshal([]byte(`"x"`), &x)
		if assert.NoError(t, err) {
			assert.Equal(t, "x", x.S)
		}
	})
}

func Test_PreciseDateTime(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := PreciseDateTime{time.Date(2023, 6, 1, 23, 59, 59, 123400000, time.UTC)}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`2023-06-01T23:59:59.1234Z`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x PreciseDateTime
			err := x.UnmarshalText([]byte(`2023-06-01T23:59:59.1234Z`))
			if assert.NoError(t, err) {
				expected := time.Date(2023, 6, 1, 23, 59, 59, 123400000, time.UTC)
				assert.True(t, expected.Equal(x.Time))
			}
		})
	})
}

func Test_URI(t *testing.T) {
	t.Run("MarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			x := URI{"x"}
			b, err := x.MarshalText()
			if assert.NoError(t, err) {
				assert.Equal(t, []byte(`x`), b)
			}
		})
	})
	t.Run("UnmarshalText", func(t *testing.T) {
		t.Run("Case1", func(t *testing.T) {
			var x URI
			err := x.UnmarshalText([]byte(`x`))
			if assert.NoError(t, err) {
				assert.Equal(t, "x", x.S)
			}
		})
	})
}
