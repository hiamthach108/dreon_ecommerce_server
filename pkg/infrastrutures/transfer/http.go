package transfer

import (
	"dreon_ecommerce_server/shared/errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type httpTransfer struct {
	client http.Client
}

func NewHttpTransfer() *httpTransfer {
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	return &httpTransfer{
		client: c,
	}
}

func (h *httpTransfer) PerformGet(url string) (data []byte, err error) {
	resp, err := h.client.Get(url)
	if err != nil {
		return
	}
	defer func() {
		resp.Body.Close()
		h.client.CloseIdleConnections()
	}()
	if resp.StatusCode != http.StatusOK {
		err = errors.NewNotFound(fmt.Errorf("status code %d", resp.StatusCode), "Cannot get data from url")
		return
	}
	data, err = io.ReadAll(resp.Body)
	return
}
