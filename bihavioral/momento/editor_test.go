package momento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditor(t *testing.T) {
	editor := Editor{}
	editor.Type("Hi, guys!")
	mementoEditor := editor.Save()

	editorNew := Editor{}
	editorNew.Restore(mementoEditor)

	assert.Equal(t, editor.text, editorNew.text)
}
