package users

func EntityToDto(userEntity User) *ReadUserResponse {
	return &ReadUserResponse{
		Id:        userEntity.Id,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Score:     userEntity.Score,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}
