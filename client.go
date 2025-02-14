package herschel

import (
	"context"

	"github.com/pkg/errors"

	sheets "google.golang.org/api/sheets/v4"
)

// Client provides methods to manipulate spreadsheets.
type Client struct {
	service *sheets.Service
}

// NewClient returns a new instance
func NewClient() (*Client, error) {
	ctx := context.Background()
	service, err := sheets.NewService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create service with client")
	}

	return &Client{service: service}, nil
}

/*
 * Low-level Spreadsheet api calls
 */
func (c Client) updateCellValues(spreadsheetID string, sheetName string, values [][]interface{}) error {
	if c.service == nil {
		return errors.New("service not initiallized")
	}
	if _, err := c.service.Spreadsheets.Values.Update(spreadsheetID, sheetName, &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         values,
	}).ValueInputOption("USER_ENTERED").Do(); err != nil {
		return err
	}

	return nil
}

func (c Client) batchUpdate(spreadsheetID string, requests []*sheets.Request) error {
	if c.service == nil {
		return errors.New("service not initiallized")
	}
	if len(requests) == 0 {
		return nil
	}

	if _, err := c.service.Spreadsheets.BatchUpdate(spreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}).Do(); err != nil {
		return err
	}
	return nil
}
