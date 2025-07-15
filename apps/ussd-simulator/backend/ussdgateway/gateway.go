package ussdgateway

import (
	"context"
	"github.com/monime-lab/gok/syserr"
	"github.com/monimesl/monime-cli/pkg/utils/monimeapis"
	"net/http"
	"time"
)

type Gateway struct {
	ctx context.Context
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Initialize(ctx context.Context) {
	g.ctx = ctx
}

type ExchangeRequest struct {
	NetworkId       string `json:"networkId"`
	SessionId       string `json:"sessionId"`
	ReplyData       string `json:"replyData"`
	InitialUssdCode string `json:"initialUssdCode"`
}

type ExchangeResponse struct {
	SessionId       string `json:"sessionId"`
	Terminate       bool   `json:"terminate"`
	ResponseMessage string `json:"responseMessage"`
}

func (g *Gateway) Exchange(request ExchangeRequest) (*ExchangeResponse, error) {
	ctx, cancel := context.WithTimeout(g.ctx, 15*time.Second)
	defer cancel()
	resp, err := g.exchange(ctx, request)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (g *Gateway) exchange(ctx context.Context, request ExchangeRequest) (ExchangeResponse, error) {
	type ExchangeReq struct {
		Session *struct {
			Network     string `json:"network"`
			InitialCode string `json:"initialCode"`
		} `json:"session,omitempty"`
		Reply *struct {
			SessionId string `json:"sessionId"`
			Data      string `json:"data"`
		} `json:"reply,omitempty"`
	}
	type Session struct {
		SessionId   string `json:"sessionId"`
		NextMessage string `json:"nextMessage"`
		IsFinal     bool   `json:"isFinal"`
	}

	var req ExchangeReq
	if request.SessionId == "" {
		req.Session = &struct {
			Network     string `json:"network"`
			InitialCode string `json:"initialCode"`
		}{
			Network:     request.NetworkId,
			InitialCode: request.InitialUssdCode,
		}
	} else {
		req.Reply = &struct {
			SessionId string `json:"sessionId"`
			Data      string `json:"data"`
		}{
			SessionId: request.SessionId,
			Data:      request.ReplyData,
		}
	}
	result, err := monimeapis.ApiRequest[ExchangeReq, Session](ctx, nil, http.MethodPost, "/ussd-exchanges", req)
	if err != nil {
		return ExchangeResponse{}, err
	}
	if request.SessionId != "" && request.SessionId != result.Result.SessionId {
		return ExchangeResponse{}, syserr.ArgumentsInvalid(
			"Mismatch between request session ID and response session ID")
	}
	return ExchangeResponse{
		Terminate:       result.Result.IsFinal,
		SessionId:       result.Result.SessionId,
		ResponseMessage: result.Result.NextMessage,
	}, nil
}
