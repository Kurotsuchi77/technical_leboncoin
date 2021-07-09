package fizzbuzz

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

// Request - Format of a request to fizzbuzz Service
type Request struct {
	gorm.Model
	Int1  int64
	Int2  int64
	Limit int64
	Str1  string
	Str2  string
}

// Result - Format of a result from fizzbuzz Service
type Result []string

// Service - fizzbuzz Service structure, containing data used by fizzbuzz
type Service struct {
	Db *gorm.DB
}

// NewService - Returns a pointer to a new fizzbuzz Service
func NewService(db *gorm.DB) *Service {
	return &Service{
		Db: db,
	}
}

// CreateRequest - Creates a fizzbuzz Request object from a map of parameters
func (service *Service) CreateRequest(parameters map[string]string) (*Request, error) {
	var int1, int2, limit int64
	var err error

	// Convert string parameters to int64
	if int1, err = strconv.ParseInt(parameters["int1"], 10, 64); err != nil {
		return nil, err
	}
	if int2, err = strconv.ParseInt(parameters["int2"], 10, 64); err != nil {
		return nil, err
	}
	if limit, err = strconv.ParseInt(parameters["limit"], 10, 64); err != nil {
		return nil, err
	}

	return &Request{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  parameters["str1"],
		Str2:  parameters["str2"],
	}, nil
}

// GetFizzBuzz - Performs fizzbuzz with given Request, returns a fizzbuzz Result
func (service *Service) GetFizzBuzz(request *Request) (*Result, error) {
	res := make(Result, 0, request.Limit)

	for i := int64(1); i < request.Limit+1; i += 1 {
		if i%request.Int1 == 0 && i%request.Int2 == 0 {
			res = append(res, request.Str1+request.Str2)
		} else if i%request.Int1 == 0 {
			res = append(res, request.Str1)
		} else if i%request.Int2 == 0 {
			res = append(res, request.Str2)
		} else {
			res = append(res, strconv.FormatInt(i, 10))
		}
	}

	// Add to DB
	if err := service.Db.Save(request); err.Error != nil {
		return nil, err.Error
	}

	return &res, nil
}

// GetMostRequested - 
func (service *Service) GetMostRequested() (*Request, int64) {
	var req Request
	var count int64

	// Query most requested parameters
	service.Db.Model(&Request{}).Select("int1, int2, \"limit\", str1, str2").Group("int1, int2, \"limit\", str1, str2").Order("COUNT(*) DESC").Limit(1).
		Scan(&req)

	// Get number of queries
	service.Db.Model(&Request{}).Where(&req).Count(&count)
	if count == 0 {
		return nil, 0
	}

	return &req, count
}
