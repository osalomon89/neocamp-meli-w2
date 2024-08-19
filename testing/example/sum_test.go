package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumaSucceso(t *testing.T) {
	resultado := Suma(2, 3)
	esperado := 5
	if resultado != esperado {
		t.Errorf("Suma(2, 3) = %d; esperado %d", resultado, esperado)
	}
}

func TestSubtracao(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		resultado, err := Subtracao(5, 3)
		assert.Nil(t, err)
		assert.Equal(t, 2, resultado)
	})

	t.Run("error", func(t *testing.T) {
		resultado, err := Subtracao(3, 5)
		assert.NotNil(t, err)
		assert.Equal(t, resultado, 0)
		assert.Equal(t, err.Error(), "error in subtracao")
	})
}

func TestSubtracaoError(t *testing.T) {
	resultado, err := Subtracao(3, 5)
	assert.NotNil(t, err)
	assert.Equal(t, resultado, 0)
	assert.Equal(t, err.Error(), "error in subtracao")
}
