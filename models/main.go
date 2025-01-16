package models

func UpdateUserFields(user *User, updateUser User) {
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Password != nil {
		user.Password = updateUser.Password
	}
	if updateUser.InviteCodeID != 0 {
		user.InviteCodeID = updateUser.InviteCodeID
	}
	// if updateUser.InviteCode != (InviteCode{}) {
	// 	user.InviteCode = updateUser.InviteCode
	// }
	if updateUser.IsAdmin != user.IsAdmin {
		user.IsAdmin = updateUser.IsAdmin
	}
}
