// handler/grpc_handler.go
package handler

import (
	"context"
	"errors"
	"time"

	"github.com/hariszaki17/library-management/book-service/config"
	"github.com/hariszaki17/library-management/book-service/handler/dto"
	"github.com/hariszaki17/library-management/book-service/models"
	"github.com/hariszaki17/library-management/book-service/usecase"
	"github.com/hariszaki17/library-management/proto/constants"
	pb "github.com/hariszaki17/library-management/proto/gen/book/proto"
	"github.com/hariszaki17/library-management/proto/logging"
	"github.com/hariszaki17/library-management/proto/utils"
	"github.com/sirupsen/logrus"
)

func NewRPC(useCase usecase.BookUsecase) pb.BookServiceServer {
	return &rpc{
		bookUsecase: useCase,
	}
}

type rpc struct {
	bookUsecase                       usecase.BookUsecase
	pb.UnimplementedBookServiceServer // Embed the unimplemented server
}

func (r *rpc) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetBooks",
		"userID":      userID,
	}).Info("Invoke RPC - GetBooks")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)
	books, err := r.bookUsecase.GetBooks(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.GetBooks")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "GetBooks",
		"userID":      userID,
	}).Info("Finished RPC - GetBooks")
	return dto.ToGetBooksResponse(books), nil
}

func (r *rpc) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "CreateBook",
		"userID":      userID,
	}).Info("Invoke RPC - CreateBook")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	// Parse the string into a time.Time object
	parsedDate, err := time.Parse(constants.FormatDate, req.PublishedAt)
	if err != nil {
		logger.WithError(err).Error("Error parse PublishedAt")
		return nil, err
	}
	_, err = r.bookUsecase.CreateBook(ctx, &models.Book{
		Title:       req.Title,
		AuthorID:    uint(req.AuthorId),
		CategoryID:  uint(req.CategoryId),
		ISBN:        req.Isbn,
		PublishedAt: parsedDate,
	})
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.CreateBook")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "CreateBook",
		"userID":      userID,
	}).Info("Finished RPC - CreateBook")
	return dto.ToCreateBookResponse("successfully create a book"), nil
}

func (r *rpc) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UpdateBook",
		"userID":      userID,
	}).Info("Invoke RPC - UpdateBook")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method bookUsecase.UpdateBook, id must be > 0")
		return nil, errors.New("id must be > 0")
	}
	data := req.GetData().AsMap() // Convert Struct to map[string]interface{}

	_, err := r.bookUsecase.UpdateBook(ctx, uint(req.Id), data)
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.UpdateBook")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "UpdateBook",
		"userID":      userID,
	}).Info("Finished RPC - UpdateBook")
	return dto.ToUpdateBookResponse("successfully update a book"), nil
}

func (r *rpc) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "DeleteBook",
		"userID":      userID,
	}).Info("Invoke RPC - DeleteBook")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method bookUsecase.DeleteBook, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.bookUsecase.DeleteBook(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.DeleteBook")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "DeleteBook",
		"userID":      userID,
	}).Info("Finished RPC - DeleteBook")
	return dto.ToDeleteBookResponse("successfully delete a book"), nil
}

func (r *rpc) BorrowBookByID(ctx context.Context, req *pb.BorrowBookByIDRequest) (*pb.BorrowBookByIDResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "BorrowBookByID",
		"userID":      userID,
	}).Info("Invoke RPC - BorrowBookByID")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method bookUsecase.BorrowBookByID, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.bookUsecase.BorrowBookByID(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.BorrowBookByID")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "BorrowBookByID",
		"userID":      userID,
	}).Info("Finished RPC - BorrowBookByID")
	return dto.ToBorrowBookByIDResponse("successfully borrow a book"), nil
}

func (r *rpc) ReturnBookByID(ctx context.Context, req *pb.ReturnBookByIDRequest) (*pb.ReturnBookByIDResponse, error) {
	requestID := utils.ExtractRequestID(ctx)
	userID := utils.ExtractUserID(ctx)
	logger := logging.Logger.WithField("requestID", requestID)

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "ReturnBookByID",
		"userID":      userID,
	}).Info("Invoke RPC - ReturnBookByID")
	ctx = context.WithValue(ctx, constants.RequestIDKeyCtx, requestID)

	if req.Id < 1 {
		logger.Error("Error while calling method bookUsecase.ReturnBookByID, id must be > 0")
		return nil, errors.New("id must be > 0")
	}

	err := r.bookUsecase.ReturnBookByID(ctx, uint(req.Id))
	if err != nil {
		logger.WithError(err).Error("Error while calling method bookUsecase.ReturnBookByID")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"serviceName": config.Data.ServiceName,
		"rpc":         "ReturnBookByID",
		"userID":      userID,
	}).Info("Finished RPC - ReturnBookByID")
	return dto.ToReturnBookByIDResponse("successfully return a book"), nil
}
