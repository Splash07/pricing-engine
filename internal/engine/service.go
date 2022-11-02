package engine

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

type PriceManagementService interface {
	UploadPriceChangeRequestsByCSV(context.Context, *os.File) (string, error)
	GetUploadStatus(context.Context, string) error
}

type priceManagementService struct {
	uploadJobRepo          UploadJobRepository
	priceChangeRequestRepo PriceChangeRequestRepository
	logger                 *zap.SugaredLogger
}

func NewPriceManagementService(
	uploadJobRepo UploadJobRepository,
	priceChangeRequestRepo PriceChangeRequestRepository,
	logger *zap.SugaredLogger,
) *priceManagementService {
	return &priceManagementService{
		uploadJobRepo:          uploadJobRepo,
		priceChangeRequestRepo: priceChangeRequestRepo,
		logger:                 logger,
	}
}

func (svc priceManagementService) UploadPriceChangeRequestsByCSV(ctx context.Context, file *os.File) (string, error) {
	fcsv := csv.NewReader(file)

	var (
		requests               = []PriceChangeRequest{}
		requestsDuplicationMap = make(map[[2]string]bool)
		err                    error
		// uploadJob              UploadJob
	)

	for {
		rStr, readErr := fcsv.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			svc.logger.Errorf("error when parsing csv file: ", readErr.Error())
			err = readErr
			break
		}

		request := svc.parsePriceChangeRequest(rStr)
		if _, exists := requestsDuplicationMap[[2]string{request.SkuID, string(request.PriceType)}]; exists {
			err = fmt.Errorf("cannot upload file, duplicate sky and price type: %s, %s", request.SkuID, request.PriceType)
			break
		}
		requests = append(requests, request)
		requestsDuplicationMap[[2]string{request.SkuID, string(request.PriceType)}] = true
	}

	return "", err

}

func (svc priceManagementService) parsePriceChangeRequest(str []string) PriceChangeRequest {
	return PriceChangeRequest{}
}
