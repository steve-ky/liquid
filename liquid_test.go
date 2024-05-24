package liquid

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterationKeyedMap(t *testing.T) {
	vars := map[string]interface{}{
		"keyed_map": IterationKeyedMap(map[string]interface{}{"a": 1, "b": 2}),
	}
	engine := NewEngine()
	tpl, err := engine.ParseTemplate([]byte(`{% for k in keyed_map %}{{ k }}={{ keyed_map[k] }}.{% endfor %}`))
	require.NoError(t, err)
	out, err := tpl.RenderString(vars)
	require.NoError(t, err)
	require.Equal(t, "a=1.b=2.", out)
}

func ExampleIterationKeyedMap() {
	vars := map[string]interface{}{
		"map":       map[string]interface{}{"a": 1},
		"keyed_map": IterationKeyedMap(map[string]interface{}{"a": 1}),
	}
	engine := NewEngine()
	out, err := engine.ParseAndRenderString(
		`{% for k in map %}{{ k[0] }}={{ k[1] }}.{% endfor %}`, vars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	out, err = engine.ParseAndRenderString(
		`{% for k in keyed_map %}{{ k }}={{ keyed_map[k] }}.{% endfor %}`, vars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Output: a=1.
	// a=1.
}

func TestStringUnescape(t *testing.T) {
	vars := map[string]interface{}{}
	engine := NewEngine()

	out, err := engine.ParseAndRenderString(`{{ 'ab\nc' }}`, vars)
	require.NoError(t, err)
	require.Equal(t, "ab\\nc", out)

	out, err = engine.ParseAndRenderString(`{{ "ab\nc" }}`, vars)
	require.NoError(t, err)
	require.Equal(t, "ab\nc", out)

	out, err = engine.ParseAndRenderString(`{{ "ab\tc" }}`, vars)
	require.NoError(t, err)
	require.Equal(t, "ab\tc", out)

	_, err = engine.ParseAndRenderString(`{{ "ab\xc" }}`, vars)
	require.Error(t, err)

	out, err = engine.ParseAndRenderString(`{{ 'ab\xc' }}`, vars)
	require.NoError(t, err)
	require.Equal(t, "ab\\xc", out)
}
