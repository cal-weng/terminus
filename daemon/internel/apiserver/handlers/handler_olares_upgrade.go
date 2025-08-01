package handlers

import (
	"github.com/beclab/Olares/daemon/pkg/cluster/state"
	"net/http"

	"github.com/beclab/Olares/daemon/pkg/commands"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

func (h *Handlers) RequestOlaresUpgrade(ctx *fiber.Ctx, cmd commands.Interface) error {
	var req state.UpgradeTarget
	if err := h.ParseBody(ctx, &req); err != nil {
		klog.Error("parse request error, ", err)
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}
	if err := req.IsValidRequest(); err != nil {
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}

	if _, err := cmd.Execute(ctx.Context(), req); err != nil {
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}

	return h.OkJSON(ctx, "successfully created upgrade target")
}

func (h *Handlers) CancelOlaresUpgrade(ctx *fiber.Ctx, cmd commands.Interface) error {
	if _, err := cmd.Execute(ctx.Context(), nil); err != nil {
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}

	return h.OkJSON(ctx, "successfully cancelled upgrade/download")
}

func (h *Handlers) ConfirmOlaresUpgrade(ctx *fiber.Ctx) error {
	target, err := state.GetOlaresUpgradeTarget()
	if err != nil {
		return h.ErrJSON(ctx, http.StatusInternalServerError, err.Error())
	}
	if target == nil {
		return h.ErrJSON(ctx, http.StatusNotFound, "upgrade target not found")
	}
	if target.DownloadOnly {
		target.DownloadOnly = false
		err = target.Save()
		if err != nil {
			return h.ErrJSON(ctx, http.StatusInternalServerError, err.Error())
		}
		if target.Downloaded {
			err = state.CheckCurrentStatus(ctx.Context())
			if err != nil {
				klog.Warning("failed to refresh current status immediately after confirmation, ", err)
			}
		}
	}

	return h.OkJSON(ctx, "successfully confirmed upgrade")
}
