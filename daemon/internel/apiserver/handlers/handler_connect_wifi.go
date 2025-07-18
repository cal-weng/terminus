package handlers

import (
	"net/http"

	"github.com/beclab/Olares/daemon/pkg/commands"
	connectwifi "github.com/beclab/Olares/daemon/pkg/commands/connect_wifi"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

type ConnectWifiReq struct {
	Password string `json:"password"`
	SSID     string `json:"ssid"`
}

func (h *Handlers) PostConnectWifi(ctx *fiber.Ctx, cmd commands.Interface) error {
	var req ConnectWifiReq
	if err := h.ParseBody(ctx, &req); err != nil {
		klog.Error("parse request error, ", err)
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}

	if _, err := cmd.Execute(ctx.Context(), &connectwifi.Param{
		SSID:     req.SSID,
		Password: req.Password,
	}); err != nil {
		return h.ErrJSON(ctx, http.StatusBadRequest, err.Error())
	}

	return h.OkJSON(ctx, "success to connect wifi")
}
