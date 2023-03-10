package main

import (
	"encoding/json"

	"github.com/opensourceways/software-package-server/softwarepkg/app"
)

type server struct {
	service app.SoftwarePkgMessageService
}

func (s *server) handleCIChecking(data []byte) error {
	msg := new(msgToHandleCIChecking)

	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	cmd, err := msg.toCmd()
	if err != nil {
		return err
	}

	return s.service.HandleCIChecking(cmd)
}

func (s *server) handleRepoCreated(data []byte) error {
	msg := new(msgToHandleRepoCreated)

	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	cmd, err := msg.toCmd()
	if err != nil {
		return err
	}

	return s.service.HandleRepoCreated(cmd)
}

func (s *server) handlePkgRejected(data []byte) error {
	msg := new(msgToHandlePkgRejected)

	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	return s.service.HandlePkgRejected(msg.toCmd())
}
