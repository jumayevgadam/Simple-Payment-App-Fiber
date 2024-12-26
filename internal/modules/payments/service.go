package payments

import (
	"context"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/models/statistics"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Service interface {
	AddPayment(ctx *fiber.Ctx, checkPhoto *multipart.FileHeader, studentID int, paymentRequest paymentModel.Request) (int, error)
	GetPayment(ctx context.Context, studentID, paymentID int) (*paymentModel.AllPaymentDTO, error)
	StudentUpdatePayment(ctx *fiber.Ctx, studentID, paymentID int, updateRequest paymentModel.UpdatePaymentRequest) (
		string, error,
	)
	StudentDeletePayment(ctx context.Context, studentID, paymentID int) error

	ListPaymentsByStudent(ctx context.Context, studentID int, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*paymentModel.AllPaymentDTO], userModel.StudentNameAndSurname, error,
	)

	AdminUpdatePaymentOfStudent(ctx context.Context, studentID, paymentID int, paymentStatus string) (string, error)
	AdminDeleteStudentPayment(ctx context.Context, studentID, paymentID int) error

	// ------------------ STATISTICS -----------------------//
	AdminGetStatisticsAboutYear(ctx context.Context, timeModel timeModel.AcademicYearRequest) (statistics.StatisticsAboutUniversity, error)
	AdminGetStatisticsAboutFaculty(ctx context.Context, facultyID int, timeModel timeModel.AcademicYearRequest) (
		statistics.StatisticsAboutUniversity, error)
}
