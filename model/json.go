/**
 * @Author: Hardews
 * @Date: 2023/10/8 23:45
 * @Description:
**/

package model

import "time"

type JsonBody struct {
	Repository struct {
		Name    string `json:"name"`    // 仓库名
		Private bool   `json:"private"` // 是否是公开的
		Owner   struct {
			Name string `json:"name"`
		} `json:"owner"`
		CreatedAt int       `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		PushedAt  int       `json:"pushed_at"`
	} `json:"repository"`
	HeadCommit struct {
		Id        string    `json:"id"`
		TreeId    string    `json:"tree_id"`
		Message   string    `json:"message"` // commit message
		Timestamp time.Time `json:"timestamp"`
	} `json:"head_commit"`
}
