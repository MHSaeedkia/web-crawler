package log

import (
	"project-root/sys-modules/log/Lib/FireLog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	err := FireLog.Log("telebot info", "telebot", "INFO")
	assert.Nil(t, err)
	err = FireLog.Log("telebot warning", "telebot", "WARNING")
	assert.Nil(t, err)
	err = FireLog.Log("telebot error", "telebot", "ERROR")
	assert.Nil(t, err)
	err = FireLog.Log("crawler info 1 ", "webcrawler", "INFO")
	assert.Nil(t, err)
	err = FireLog.Log("crawler info 2 ", "webcrawler", "INFO")
	assert.Nil(t, err)
	err = FireLog.Log("crawler warning", "webcrawler", "WARNING")
	assert.Nil(t, err)
	err = FireLog.Log("crawler error", "webcrawler", "ERROR")
	assert.Nil(t, err)

}
