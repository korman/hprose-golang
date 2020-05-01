/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder_test.go                             |
|                                                          |
| LastModified: Apr 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeFloat32(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(math.MaxInt64)
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode([]byte{1})
	dec := NewDecoder(([]byte)(sb.String()))
	var f float32
	dec.Decode(&f)
	assert.Equal(t, float32(-1), f)
	dec.Decode(&f)
	assert.Equal(t, float32(0), f)
	dec.Decode(&f)
	assert.Equal(t, float32(1), f)
	dec.Decode(&f)
	assert.Equal(t, float32(123), f)
	dec.Decode(&f)
	assert.Equal(t, float32(math.MinInt64), f)
	dec.Decode(&f)
	assert.Equal(t, float32(-math.MaxInt64), f)
	dec.Decode(&f)
	assert.Equal(t, float32(math.MaxInt64), f)
	dec.Decode(&f)
	assert.Equal(t, float32(1), f)
	dec.Decode(&f)
	assert.Equal(t, float32(0), f)
	dec.Decode(&f)
	assert.Equal(t, float32(0), f)
	dec.Decode(&f)
	assert.Equal(t, float32(3.14), f)
	dec.Decode(&f)
	assert.True(t, math.IsNaN(float64(f)))
	dec.Decode(&f)
	assert.True(t, math.IsInf(float64(f), 1))
	dec.Decode(&f)
	assert.True(t, math.IsInf(float64(f), -1))
	dec.Decode(&f)
	assert.Equal(t, float32(0), f)
	dec.Decode(&f)
	assert.Equal(t, float32(1), f)
	dec.Decode(&f)
	assert.Equal(t, float32(123), f)
	assert.NoError(t, dec.Error)
	dec.Decode(&f)
	assert.EqualError(t, dec.Error, `strconv.ParseFloat: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&f)
	assert.True(t, math.IsNaN(float64(f)))
	dec.Decode(&f)
	assert.True(t, math.IsNaN(float64(f)))
}

func TestDecodeFloat64(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(math.MaxInt64)
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode([]byte{1})
	dec := NewDecoder(([]byte)(sb.String()))
	var f float64
	dec.Decode(&f)
	assert.Equal(t, float64(-1), f)
	dec.Decode(&f)
	assert.Equal(t, float64(0), f)
	dec.Decode(&f)
	assert.Equal(t, float64(1), f)
	dec.Decode(&f)
	assert.Equal(t, float64(123), f)
	dec.Decode(&f)
	assert.Equal(t, float64(math.MinInt64), f)
	dec.Decode(&f)
	assert.Equal(t, float64(-math.MaxInt64), f)
	dec.Decode(&f)
	assert.Equal(t, float64(math.MaxInt64), f)
	dec.Decode(&f)
	assert.Equal(t, float64(1), f)
	dec.Decode(&f)
	assert.Equal(t, float64(0), f)
	dec.Decode(&f)
	assert.Equal(t, float64(0), f)
	dec.Decode(&f)
	assert.Equal(t, float64(3.14), f)
	dec.Decode(&f)
	assert.True(t, math.IsNaN(f))
	dec.Decode(&f)
	assert.True(t, math.IsInf(f, 1))
	dec.Decode(&f)
	assert.True(t, math.IsInf(f, -1))
	dec.Decode(&f)
	assert.Equal(t, float64(0), f)
	dec.Decode(&f)
	assert.Equal(t, float64(1), f)
	dec.Decode(&f)
	assert.Equal(t, float64(123), f)
	assert.NoError(t, dec.Error)
	dec.Decode(&f)
	assert.EqualError(t, dec.Error, `strconv.ParseFloat: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&f)
	assert.True(t, math.IsNaN(f))
	dec.Decode(&f)
	assert.True(t, math.IsNaN(f))
}
