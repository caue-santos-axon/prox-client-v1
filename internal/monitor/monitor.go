package monitor

import (
	"proxclient/internal/logging"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/mgr"
)

func IsOnline() (bool, error) {
	m, err := mgr.Connect()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't set mrg service monitor")
		return false, err
	}

	s, err := m.OpenService("ProxClientv5")
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't open service")
		return false, err
	}
	status, err := s.Query()
	logging.Log.WithFields(logrus.Fields{
		"err": err,
	}).Error("Coundn't get current status server")

	if status.State != 4 {
		return false, nil
	}

	return true, nil
}
