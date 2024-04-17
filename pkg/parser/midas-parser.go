package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sefikcan/hisse-crawler/internal/asset"
	"io"
	"net/http"
	"time"
)

const (
	assetApiUrl = "https://www.getmidas.com/canli-borsa/"
)

var client = &http.Client{Timeout: 30 * time.Second}

func ParseMidas(symbol string) (*asset.Asset, error) {
	req, err := http.NewRequest(http.MethodGet, assetApiUrl+symbol, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, handleStatusCodeError(res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	result := &asset.Asset{
		Name:        asset.String{Value: findText(doc, ".plan-comparison-title h1")},
		Price:       asset.String{Value: findText(doc, ".detail-card-container .detail-cards p.val:first-child")},
		DailyVolume: asset.String{Value: findText(doc, ".detail-card-container .detail-cards p.val:last-child")},
		DailyChange: asset.String{Value: findText(doc, ".buy-block .info.daily p.val")},
	}

	return result, nil
}

func handleStatusCodeError(code int, status string) error {
	return fmt.Errorf("unexpected status code: %d %s", code, status)
}

func findText(doc *goquery.Document, selector string) string {
	text := doc.Find(selector).Text()
	return text
}
