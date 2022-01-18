package service

import (
	"fmt"
	"strconv"
	"time"
	"wineSong/entity"
	"wineSong/middleware"
	"wineSong/models"
	"wineSong/repository"
	"wineSong/util"
)

func Login(user models.User) entity.Result {
	var res entity.Result
	userSearch := repository.SearchUserByName(user.UserName)
	if userSearch.Id == 0 {
		res = util.ErrReturn(604, nil, "用户不存在")
	} else {
		if util.GetMd5(user.Password, "wineSong") == userSearch.Password {
			res = SuccessLogin(userSearch)
		} else {
			res = util.ErrReturn(601, nil, "密码错误")
		}
	}
	return res
}

func SuccessLogin(u models.User) entity.Result {
	var userInfo models.UserInfo
	var loginLog models.LoginLog
	var res entity.Result
	token, err := middleware.NewToken(u.Id)
	if err == nil {
		fmt.Println(err)
	}
	userInfo = models.UserInfo{
		Id:       u.Id,
		UserName: u.UserName,
		Role:     u.Role,
		Phone:    u.Phone,
		Token:    token,
		Qq:       u.Qq,
	}
	res.Data = userInfo
	res.Code = 200
	res.Message = "成功"
	loginLog.Time = time.Now().Format("2006-01-02 15:04:05")
	loginLog.UserId = u.Id
	err = repository.CreateLoginLog(loginLog)
	fmt.Println(err)
	return res
}

func GetUserInfo(userId int) entity.Result {
	var res entity.Result
	user := repository.SearchUserById(userId)
	if user.Id == 0 {
		res = util.ErrReturn(604, nil, "用户不存在")
	} else {
		token, err := middleware.NewToken(user.Id)
		if err == nil {
			fmt.Println(err)
		}
		userInfo := models.UserInfo{
			Id:       user.Id,
			UserName: user.UserName,
			Role:     user.Role,
			Phone:    user.Phone,
			Token:    token,
			Qq:       user.Qq,
		}
		res.Code = 200
		res.Data = userInfo
	}
	return res
}

func Register(user models.User) entity.Result {
	var res entity.Result
	userAccount := repository.SearchUserByName(user.UserName)
	if userAccount.Id == 0 {
		user.Password = util.GetMd5(user.Password, "wineSong")
		user.Role = "user"
		if err := repository.Register(user); err != nil {
			res = util.ErrReturn(603, err, "sql错误")
		} else {
			res.Code = 200
			res.Message = "成功"
		}
	} else {
		res = util.ErrReturn(602, nil, "账号已注册")
	}
	return res
}

func SearchAuctioneer(userId string) entity.Result {
	var res entity.Result
	if userId == "" {
		auctioneer := repository.SearchAuctioneer()
		res.Code = 200
		res.Data = auctioneer
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			res = util.ErrReturn(605, err, "系统错误")
		} else {
			res.Code = 200
			res.Data = repository.SearchAuctioneerByUserId(id)
		}
	}

	return res
}

func UpdateUser(user models.User) entity.Result {
	var res entity.Result
	searchUser := repository.SearchUserById(user.Id)
	if searchUser.UserName != user.UserName {
		searchUserName := repository.SearchUserByName(user.UserName)
		if searchUserName.Id != 0 {
			res = util.ErrReturn(602, nil, "用户名重复")
			return res
		}
	}
	if searchUser.Phone != user.Phone {
		searchUserName := repository.SearchUserByPhone(user.Phone)
		if searchUserName.Id != 0 {
			res = util.ErrReturn(602, nil, "手机号已注册重复")
			return res
		}
	}
	if searchUser.Qq != user.Qq {
		searchUserName := repository.SearchUserByQq(user.Qq)
		if searchUserName.Id != 0 {
			res = util.ErrReturn(602, nil, "qq号已被绑定")
			return res
		}
	}
	if err := repository.UpdateUser(user); err != nil {
		res = util.ErrReturn(603, err, "sql错误")
	} else {
		res.Code = 200
		res.Message = "成功"
	}
	return res
}

func SearchAllUser() entity.Result {
	var res entity.Result
	res.Code = 200
	res.Data = repository.SearchAllUser()
	return res
}
