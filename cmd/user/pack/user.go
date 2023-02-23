// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
)

// User pack user info
func User(u *repository.User) *common.User {
	if u == nil {
		return nil
	}
	return &common.User{
		Id:              int64(u.ID),
		Name:            u.UserName,
		Avatar:          "",
		BackgroundImage: "https://cdn.cdnjson.com/wx3.sinaimg.cn/large/87c01ec7gy1fsnqqm9l6pj21kw0w0aos.jpg",
		Signature:       "",
	}
}

// Users pack list of user info
func Users(us []*repository.User) []*common.User {
	users := make([]*common.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
