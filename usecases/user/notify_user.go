package user

type UserService struct {
	// storage UserStorage
	// rateAPI RateAPI
}

// func NewUserService(storage UserStorage, rateAPI RateAPI) *UserService {
func NewUserService() *UserService {
	return &UserService{
		// storage: storage,
		// rateAPI: rateAPI,
	}
}

// SubscribeUser subscribes a user to daily rate updates
// func (us *UserService) SubscribeUser(userID int) error {
// return us.storage.AddUserSubscription(userID)
// }

// GetCurrentRate fetches the current rate from the rate API
func (us *UserService) GetCurrentRate() (string, error) {
	return "123", nil
}
