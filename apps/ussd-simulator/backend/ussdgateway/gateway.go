package ussdgateway

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/monime-lab/gok/syserr"
	"github.com/monime-lab/gwater/envutil"
	"github.com/monime-lab/gwater/typeutil"
	"github.com/monimesl/monime-cli/pkg/utils/monimeapis"
)

type Gateway struct {
	ctx            context.Context
	exchangeUrl    string
	exchangeMsisdn string
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Initialize(ctx context.Context) {
	g.ctx = ctx
	exchangeUrl := strings.TrimSpace(envutil.GetOptionalValue("MONIME_CLI_USSD_GATEWAY_EXCHANGE_URL", ""))
	exchangeMsisdn := strings.TrimSpace(envutil.GetOptionalValue("MONIME_CLI_USSD_GATEWAY_EXCHANGE_MSISDN", ""))
	if exchangeUrl != "" {
		u, err := url.ParseRequestURI(exchangeUrl)
		if err != nil {
			log.Fatalf("Invalid exchange URL: %s", exchangeUrl)
		}
		exchangeUrl = u.String()
	} else {
		exchangeMsisdn = ""
		exchangeUrl = "/ussd-exchanges"
	}
	g.exchangeUrl = exchangeUrl
	g.exchangeMsisdn = exchangeMsisdn
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
			Network     string  `json:"network"`
			Msisdn      *string `json:"msisdn"`
			InitialCode string  `json:"initialCode"`
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
			Network     string  `json:"network"`
			Msisdn      *string `json:"msisdn"`
			InitialCode string  `json:"initialCode"`
		}{
			Network:     request.NetworkId,
			Msisdn:      typeutil.StringPtr(g.exchangeMsisdn),
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
	result, err := monimeapis.ApiRequest[ExchangeReq, Session](ctx, nil, http.MethodPost, g.exchangeUrl, req)
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
