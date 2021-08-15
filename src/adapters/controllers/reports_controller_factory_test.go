package controllers

import (
	"os"
	"reflect"
	"testing"

	itemsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/items"
	rptsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/reports"
	"github.com/stretchr/testify/assert"
)

func Test_NewReportsControllerFactory(t *testing.T) {
	f := NewReportsControllerFactory()

	assert.NotNil(t, f)
	assert.Equal(t, reflect.ValueOf(os.Getenv).Pointer(), reflect.ValueOf(f.osGetenv).Pointer())
}

func TestReportsControllerFactory_Create(t *testing.T) {
	mos := new(mockOS)
	mos.On("Getenv", "QIITA_ACCESS_TOKEN").Return("TOKEN")

	f := &ReportsControllerFactory{osGetenv: mos.Getenv}
	ic := f.Create()
	c, ok := ic.(*ReportsController)
	if !ok {
		t.Fatal()
	}

	assert.NotNil(t, c)
	assert.NotNil(t, c.itemsRepository)
	assert.IsType(t, &itemsrepo.Repository{}, c.itemsRepository)
	assert.NotNil(t, c.reportsRepository)
	assert.IsType(t, &rptsrepo.Repository{}, c.reportsRepository)
	mos.AssertNumberOfCalls(t, "Getenv", 1)
}
