package main

import (
	"github.com/opensourceways/software-package-server/softwarepkg/app"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

// msgToHandleCIChecking
type msgToHandleCIChecking struct {
	PkgId        string `json:"pkg_id"`
	RelevantPR   string `json:"relevant_pr"`
	FailedReason string `json:"failed_reason"`
}

func (msg *msgToHandleCIChecking) toCmd() (cmd app.CmdToHandleCIChecking, err error) {
	if cmd.RelevantPR, err = dp.NewURL(msg.RelevantPR); err != nil {
		return
	}

	cmd.PkgId = msg.PkgId
	cmd.FiledReason = msg.FailedReason

	return
}

// msgToHandleRepoCreated
type msgToHandleRepoCreated struct {
	PkgId        string `json:"pkg_id"`
	Platform     string `json:"platform"`
	RepoLink     string `json:"repo_link"`
	FailedReason string `json:"failed_reason"`
}

func (msg *msgToHandleRepoCreated) toCmd() (cmd app.CmdToHandleRepoCreated, err error) {
	cmd.PkgId = msg.PkgId
	cmd.FiledReason = msg.FailedReason

	if cmd.Platform, err = dp.NewPackagePlatform(msg.Platform); err != nil {
		return
	}

	if msg.RepoLink != "" {
		cmd.RepoLink, err = dp.NewURL(msg.RepoLink)
	}

	return
}

// msgToHandlePkgRejected
type msgToHandlePkgRejected struct {
	PkgId      string `json:"pkg_id"`
	Reason     string `json:"reason"`
	RejectedBy string `json:"rejected_by"`
}

func (msg *msgToHandlePkgRejected) toCmd() app.CmdToHandlePkgRejected {
	return app.CmdToHandlePkgRejected{
		PkgId:      msg.PkgId,
		Reason:     msg.Reason,
		RejectedBy: msg.RejectedBy,
	}
}
