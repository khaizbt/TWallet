package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	UploadAvatar(ID int, location string) (User, error)
	GetUserById(ID int) (User, error)
	GetBalance(ID int) (User, error)
}

type service struct { //Internal Struct
	repository Repository //Buat manggil struct repository melalui Interface repository
}

func NewService(repository Repository) *service { //
	return &service{repository} //mengirim ke interface harus pake pointer(ke alamat)
}

//Function Untuk Struct Service dengan tujuan mapping struct input ke struct user
func (s *service) RegisterUser(input RegisterUserInput) (User, error) { //Mengubah dari mapping struct input ke dalam struct user
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, err
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 { //Default dari golang jika user tidak ditemukan itu bukan nil, tapi IDnya 0
		return user, errors.New("User Tidak Ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil { //Email Not Available
		return false, err
	}

	if user.ID == 0 { //Email Available
		return true, nil
	}

	return false, nil //Nilai Default

}

/*
Upload Avatar
	- Dapatkan User berdasarkan ID
	- user Update attr avatar file name
	- simpan perubahan avatar
*/
func (s *service) UploadAvatar(ID int, location string) (User, error) {
	user, err := s.repository.FindByID(ID) //Mendapatkan ID User

	if err != nil {
		return user, err
	}

	user.AvatarImage = location

	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, err

}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID != ID {
		return user, errors.New("User Tidak Ditemukan")
	}

	return user, nil
}

func (s *service) GetBalance(ID int) (User, error) {
	user, err := s.repository.CekSaldo(ID)

	if err != nil {
		return user, err
	}

	return user, nil
}
