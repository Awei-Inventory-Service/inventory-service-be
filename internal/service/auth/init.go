package auth

func NewUserService(userRepository userRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
