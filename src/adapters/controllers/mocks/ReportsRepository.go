// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ReportsRepository is an autogenerated mock type for the ReportsRepository type
type ReportsRepository struct {
	mock.Mock
}

// UpdateWeekly provides a mock function with given fields: t, id, items
func (_m *ReportsRepository) UpdateWeekly(t time.Time, id string, items entities.Items) error {
	ret := _m.Called(t, id, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time, string, entities.Items) error); ok {
		r0 = rf(t, id, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWeeklyByTag provides a mock function with given fields: t, id, items, tag
func (_m *ReportsRepository) UpdateWeeklyByTag(t time.Time, id string, items entities.Items, tag string) error {
	ret := _m.Called(t, id, items, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time, string, entities.Items, string) error); ok {
		r0 = rf(t, id, items, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
