package monitor

import (
	"proxclient/internal/logging"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/mgr"
)

const SERVICE_NAME = "ProxClientv5"

// 'IsOnline' checks if Prox Service is running
func IsOnline() (bool, error) {
	m, err := mgr.Connect()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Monitor Error")
		return false, err
	}

	s, err := m.OpenService(SERVICE_NAME)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Monitor Error")
		return false, err
	}
	status, err := s.Query()
	logging.Log.WithFields(logrus.Fields{
		"err": err,
	}).Error("Monitor Error")

	if status.State != 4 {
		return false, nil
	}

	return true, nil
}
