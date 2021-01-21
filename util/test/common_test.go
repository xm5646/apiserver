/**
 * 功能描述:
 * @Date: 2021/1/15
 * @author: lixiaoming
 */
package test

import (
	"github.com/xm5646/log"
	"testing"
)

func TestGetTestFuncWithDB(t *testing.T) {
	GetTestFuncWithDB(func() {
		log.Infof("test ...")
	})
}
