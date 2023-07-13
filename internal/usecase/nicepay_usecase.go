package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/json"

	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/doffy007/golang-nicepay/config"
	"github.com/doffy007/golang-nicepay/internal/constants"
	"github.com/doffy007/golang-nicepay/internal/entity"
	"github.com/doffy007/golang-nicepay/internal/repository"
	"github.com/doffy007/golang-nicepay/internal/request"
	"github.com/doffy007/golang-nicepay/internal/response"
	"github.com/joho/godotenv"
)

type NicepayUsecase interface {
	NicepayRegistration(request.RegistrationRequest) (bool, response.RegistrationResponse)
	NicepayStatus(request.StatusRequest) (bool, response.StatusResponse)
	NicepayPayment(request.PaymentRequest) (bool, response.PaymentResponse)
}

type nicepayUsecase struct {
	ctx               context.Context
	config            *config.Config
	nicepayRepository repository.NicepayRepository
}

func (uc usecase) NicepayHandler() NicepayUsecase {
	return &nicepayUsecase{
		ctx:               uc.ctx,
		config:            uc.config,
		nicepayRepository: uc.repository.NicepayRepository(),
	}
}

// NicepayRegistration implements.
func (uc nicepayUsecase) NicepayRegistration(param request.RegistrationRequest) (bool, response.RegistrationResponse) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	// Get nicepay key
	merchantKey := os.Getenv("NICEPAY_MERCHANT_ID")
	merchantId := os.Getenv("NICEPAY_MERCHANT_KEY")

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowTIme := time.Now().In(loc)
	now := nowTIme.Format("2006-01-02 15:04:05")
	timeStamp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(now, "-", "-"), " ", ""), ":", "")
	amt := strconv.FormatFloat(float64(param.Amount), 'f', 0, 64)

	dataToken := []byte(timeStamp + merchantId + param.ReferenceNo + amt + merchantKey)
	merchantToken := fmt.Sprintf("%x", sha256.Sum256(dataToken))

	if param.Timestamp == "" {
		param.Timestamp = timeStamp
	}

	if param.MerchantId == "" {
		param.MerchantId = merchantId
	}

	if param.MerchantToken == "" {
		param.MerchantToken = merchantToken
	}

	if param.PayMethod == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "PayMethod must be defined",
		}
	}

	if param.PayMethod == string(constants.MethodCreditCard) {
		if param.InstallmentType == 0 {
			return false, response.RegistrationResponse{
				ResultCode:    http.StatusBadRequest,
				ResultMessage: "InstallmentType must be defined",
			}
		}

		if param.InstallmentMonth == 0 {
			return false, response.RegistrationResponse{
				ResultCode:    http.StatusBadRequest,
				ResultMessage: "InstallmentMonth must be defined",
			}
		}
	}

	if param.Currency == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Currency must be defined",
		}
	}

	if param.Amount == 0 {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Amount Cannot 0",
		}
	}

	if param.ReferenceNo == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "ReferenceNo must be defined",
		}
	}

	if param.GoodsName == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "GoodsName must be defined",
		}
	}

	if param.BillingName == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "BillingName must be defined",
		}
	}

	if param.BillingPhone == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "BillingPhone must be defined",
		}
	}

	if param.BillingEmail == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "BillingEmail must be defined",
		}
	}

	if param.DeliveryName == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "DeliveryName must be defined",
		}
	}

	if param.DeliveryPhone == "" {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "DeliveryPhone must be defined",
		}
	}

	var payload entity.Nicepay
	rec, _ := json.Marshal(param)
	json.Unmarshal(rec, &payload)

	err = uc.nicepayRepository.Register(payload)

	if err != nil {
		return false, response.RegistrationResponse{
			ResultCode:    http.StatusInternalServerError,
			ResultMessage: err.Error(),
		}
	}

	return true, response.RegistrationResponse{
		ResultCode:    http.StatusCreated,
		ResultMessage: "Success Register",
	}
}

// NicepayStatus implements.
func (uc nicepayUsecase) NicepayStatus(param request.StatusRequest) (bool, response.StatusResponse) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	// Get nicepay key
	merchantKey := os.Getenv("NICEPAY_MERCHANT_ID")
	merchantId := os.Getenv("NICEPAY_MERCHANT_KEY")

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowTIme := time.Now().In(loc)
	now := nowTIme.Format("2006-01-02 15:04:05")
	timeStamp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(now, "-", "-"), " ", ""), ":", "")
	amt := strconv.FormatFloat(float64(param.Amount), 'f', 0, 64)

	dataToken := []byte(timeStamp + merchantId + param.ReferenceNo + amt + merchantKey)
	merchantToken := fmt.Sprintf("%x", sha256.Sum256(dataToken))

	var filter entity.Status
	tId, _ := fmt.Println(rand.Intn(100))

	if param.Timestamp == "" {
		param.Timestamp = timeStamp
	}

	if param.TransactionId == "" {
		param.TransactionId = merchantKey + strconv.Itoa(tId)
	}

	if param.MerchantId == "" {
		param.MerchantId = merchantId
	}

	if param.ReferenceNo == "" {
		return false, response.StatusResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "ReferenceNo must be defined",
		}
	}

	if param.Amount == 0 {
		return false, response.StatusResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Amount Cannot 0",
		}
	}

	if param.MerchantToken == "" {
		param.MerchantToken = merchantToken
	}

	status, err := uc.nicepayRepository.Status(filter, []string{})
	if err != nil {
		return false, response.StatusResponse{
			ResultCode:    http.StatusInternalServerError,
			ResultMessage: err.Error(),
		}
	}

	if status == nil {
		return false, response.StatusResponse{
			ResultCode:    http.StatusInternalServerError,
			ResultMessage: err.Error(),
		}
	}

	return true, response.StatusResponse{
		ResultCode:    http.StatusCreated,
		TransactionId: merchantId + strconv.Itoa(tId),
		MerchantId:    merchantId,
		ReferenceNo:   param.ReferenceNo,
		Amount:        strconv.Itoa(param.Amount),
	}
}

func GenerateToken(s string) string {
	var payload request.RegistrationRequest
	// Get nicepay key
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	merchantKey := os.Getenv("NICEPAY_MERCHANT_ID")
	merchantId := os.Getenv("NICEPAY_MERCHANT_KEY")

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowTIme := time.Now().In(loc)
	now := nowTIme.Format("2006-01-02 15:04:05")
	timeStamp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(now, "-", "-"), " ", ""), ":", "")
	amt := strconv.FormatFloat(float64(payload.Amount), 'f', 0, 64)

	dataToken := []byte(timeStamp + merchantId + payload.ReferenceNo + amt + merchantKey)
	merchantToken := fmt.Sprintf("%x", sha256.Sum256(dataToken))

	return merchantToken
}

// NicepayPayment implements.
func (uc nicepayUsecase) NicepayPayment(params request.PaymentRequest) (bool, response.PaymentResponse) {
	tokenGenerate := GenerateToken(params.MerchantToken)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowTIme := time.Now().In(loc)
	now := nowTIme.Format("2006-01-02 15:04:05")
	timeStamp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(now, "-", "-"), " ", ""), ":", "")

	if params.Timestamp == "" {
		params.Timestamp = timeStamp
	}

	if params.TransactionId == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Transaction ID must be define",
		}
	}

	if params.CCNumber == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "CC Number must be define",
		}
	}

	if params.CardExpiry == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "CC Number must be define",
		}
	}

	if params.CardCVV == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Card CVV must be define",
		}
	}

	if params.CardHolderName == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Card Holder Name must be define",
		}
	}

	if params.RecurringToken == " " {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Recurring Token must be define",
		}
	}

	if params.PreAuthToken == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Pre-Auth Token must be define",
		}
	}

	if params.MerchantToken == "" {
		params.MerchantToken = tokenGenerate
	}

	if params.CallbackUrl == "" {
		return false, response.PaymentResponse{
			ResultCode:    http.StatusBadRequest,
			ResultMessage: "Callback Url must be define",
		}
	}

	return true, response.PaymentResponse{
		ResultCode:    http.StatusCreated,
		ResultMessage: "Success",
		TransactionID: params.TransactionId,
	}
}
