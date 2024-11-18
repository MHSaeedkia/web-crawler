package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	err := Log("telebot info", "telebot", "INFO")
	assert.Nil(t, err)
	err = Log("telebot warning", "telebot", "WARNING")
	assert.Nil(t, err)
	err = Log("telebot error", "telebot", "ERROR")
	assert.Nil(t, err)
	err = Log("crawler info 1 ", "webcrawler", "INFO")
	assert.Nil(t, err)
	err = Log("crawler info 2 ", "webcrawler", "INFO")
	assert.Nil(t, err)
	err = Log("crawler warning", "webcrawler", "WARNING")
	assert.Nil(t, err)
	err = Log("crawler error", "webcrawler", "ERROR")
	assert.Nil(t, err)

}
