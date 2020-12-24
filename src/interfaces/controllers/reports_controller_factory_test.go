package controllers

import (
	"os"
	"reflect"
	"testing"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/gateways"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/presenters"
	"github.com/stretchr/testify/assert"
)

/*
 * NewReportsControllerFactory()
 */

func Test_NewReportsControllerFactory(t *testing.T) {
	f := NewReportsControllerFactory()

	assert.NotNil(t, f)
	assert.Equal(t, reflect.ValueOf(os.Getenv).Pointer(), reflect.ValueOf(f.osGetenv).Pointer())
}

/*
 * ReportsControllerFactory.Create()
 */

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
	assert.IsType(t, &gateways.ItemsRepository{}, c.itemsRepository)
	assert.NotNil(t, c.reportsRepository)
	assert.IsType(t, &gateways.ReportsRepository{}, c.reportsRepository)
	assert.NotNil(t, c.reportsPresenter)
	assert.IsType(t, &presenters.ReportsPresenter{}, c.reportsPresenter)
	mos.AssertNumberOfCalls(t, "Getenv", 1)
}
