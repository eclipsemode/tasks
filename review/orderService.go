package review

import (
	"log"
	"time"
)

type OrderService struct {
	BookingService BookingService
	UserService    UserService
}

type UserService interface {
	LockUser(User) error
	UnlockUser(User) error
}

type BookingService interface {
	BookFlight() (string, *BookingServiceError)
}

type User struct {
	ID string
}

// Исправил BookedAt на тип time.Time
type Receipt struct {
	ID          string
	BookingCode string
	BookedAt    time.Time
}

type BookingServiceError struct {
	error
	TryAgain bool
}

func (s *OrderService) HandleBookingOrder(user User) *Receipt {
	receipt := Receipt{ID: uuid.New().String()}

	if err := s.UserService.LockUser(user); err != nil {
		log.Logger.Err(err)
		return nil
	}

	// Добавил defer для UnlockUser чтобы функция отрабатывала в любом случае. Возврат значения при ошибке nil происходит в конце.
	defer func() {
		if err := s.UserService.UnlockUser(user); err != nil {
			log.Logger.Err(err)
		}
	}()

	// Добавил метку для доступа из внутреннего switch case
outerLoop:
	for {
		bookingCode, err := s.BookingService.BookFlight()

		switch {
		case err == nil:
			receipt.BookedAt = time.Now()
			receipt.BookingCode = bookingCode
			return &receipt
		case err.TryAgain:
			// Добавил continue в TryAgain чтобы пропустить итерацию и на следущей заного пробовать забронировать.
			continue
		default:
			log.Logger.Err(err)
			break outerLoop
		}
	}

	// Поставил nil вместо &receipt так как данный return сработает только в плохом сценарии обработки всей функции.
	return nil
}
