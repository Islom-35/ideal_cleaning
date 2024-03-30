package app

import (
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/domain"
	"example.com/m/internal/user/errors"
	"example.com/m/pkg/utils"
	"fmt"
	"log"
)

type userUseCase struct {
	userRepo    domain.UserRepository
	userFactory domain.UserFactory
	pb.UnimplementedUserServiceServer
}

type UserUseCase interface {
	SignUpAdmin(user *pb.NewUser) error
	SignUpUser(user *pb.NewUser) error
	SignInUser(email, password string) (bool, error)
	GetUsers(*pb.UserRequest) (*pb.UsersResponse, error)
	GetUser(id *pb.UserId) (*pb.UserResponse, error)
	Update(user *pb.UserUpdate) error
	Delete(id *pb.UserId) error
}

func NewUserUseCase(userRepo domain.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) SignUpAdmin(admin *pb.NewUser) error {
	adminUser := u.userFactory.CreateAdmin(admin)

	err := utils.ValidateUserInfoForSignUp(
		adminUser.FullName,
		adminUser.Email,
		adminUser.Phone,
		adminUser.Password,
	)

	if err != nil {
		fmt.Sprintf("failed sign up admin")
		return err
	}

	exists, err := u.userRepo.UserExistByEmail(adminUser.Email)
	if err != nil {
		log.Println("internal error: " + err.Error())
		return err
	}

	if exists {
		return errors.ErrEmailExist
	} else {
		err := u.userRepo.Save(adminUser)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u *userUseCase) SignUpUser(user *pb.NewUser) error {
	userSignUp := u.userFactory.CreateUser(user)

	err := utils.ValidateUserInfoForSignUp(
		userSignUp.FullName,
		userSignUp.Email,
		userSignUp.Phone,
		userSignUp.Password,
	)

	if err != nil {
		fmt.Sprintf("failed sign up user")
		return err
	}

	exists, err := u.userRepo.UserExistByEmail(userSignUp.Email)
	if err != nil {
		log.Println("internal error: " + err.Error())
		return err
	}

	if exists {
		return errors.ErrEmailExist
	} else {
		err := u.userRepo.Save(userSignUp)
		if err != nil {
			return err
		}
		return nil
	}
}

func (u *userUseCase) SignInUser(email, password string) (bool, error) {
	err := utils.ValidateUserInfoForSignIn(email, password)

	if err != nil {
		return false, nil
	}

	userExists, err := u.userRepo.UserExistByEmail(email)
	if err != nil {
		return false, err
	}

	if !userExists {
		return false, errors.ErrUserNotFound
	}

	hashedPassword, err := u.userRepo.GetHashedPasswordByEmail(email)
	if err != nil {
		return false, err
	}

	err = utils.CheckPassword(password, hashedPassword)
	if err != nil {
		return false, errors.ErrBadCredentials
	}

	return true, nil
}

func (u *userUseCase) GetUsers(request *pb.UserRequest) (*pb.UsersResponse, error) {
	userList, err := u.userRepo.GetUsers(request)

	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return userList, nil
}

func (u *userUseCase) GetUser(userId *pb.UserId) (*pb.UserResponse, error) {
	user, err := u.userRepo.FindById(userId)

	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}

func (u *userUseCase) Update(userUpdate *pb.UserUpdate) error {
	var id *pb.UserId
	existUser, err := u.userRepo.FindById(id)

	if err != nil {
		return errors.ErrUserNotFound
	}
	existUser.FullName = userUpdate.FullName
	existUser.Phone = userUpdate.Phone

	err = u.userRepo.UpdateUser(userUpdate)
	if err != nil {
		return errors.ErrUpdateFailed
	}

	return nil
}

func (u *userUseCase) Delete(userId *pb.UserId) error {
	err := u.userRepo.DeleteUser(userId)

	if err != nil {
		return errors.ErrUserDeleteFailed
	}

	return nil
}
