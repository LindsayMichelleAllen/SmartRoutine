package model_test

import (
	"api/services/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func deviceSetup() *model.Device {
	dvc := &model.Device{}
	dvc.SetId("976431852")
	dvc.SetName("DevName")
	dvc.SetUserId("123456789")
	return dvc
}

func TestDeviceMethods(t *testing.T) {
	t.Run("Getters", func(t *testing.T) {
		t.Run("should be able to get the device id", func(t *testing.T) {
			dvc := deviceSetup()
			assert.Equal(t, "976431852", dvc.GetId())
		})
		t.Run("should be able to get the device name", func(t *testing.T) {
			dvc := deviceSetup()
			assert.Equal(t, "DevName", dvc.GetName())
		})
		t.Run("should be able to get the device userid", func(t *testing.T) {
			dvc := deviceSetup()
			assert.Equal(t, "123456789", dvc.GetUserId())
		})
	})
	t.Run("Setters", func(t *testing.T) {
		t.Run("should be able to set the device id", func(t *testing.T) {
			dvc := deviceSetup()
			assert.Equal(t, "976431852", dvc.GetId())
			dvc.SetId("NewDevID")
			assert.Equal(t, "NewDevID", dvc.GetId())
		})
		t.Run("should be able to set the device name", func(t *testing.T) {
			dvc := deviceSetup()
			assert.Equal(t, "DevName", dvc.GetName())
			dvc.SetName("NewDevName")
			assert.Equal(t, "NewDevName", dvc.GetName())
		})
		t.Run("should be able to set the device userid", func(t *testing.T) {
			dvc := deviceSetup()
			assert.Equal(t, "123456789", dvc.GetUserId())
			dvc.SetUserId("987654321")
			assert.Equal(t, "987654321", dvc.GetUserId())
		})
	})
}
