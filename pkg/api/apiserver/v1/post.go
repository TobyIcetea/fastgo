package v1

import "time"

// Post 表示博客文章
type Post struct {
	// postID 表示博文 ID
	PostID string `json:"postID"`
	// userID 表示用户 ID
	UserID string `json:"userID"`
	// title 表示博客标题
	Title string `json:"title"`
	// content 表示博客内容
	Content string `json:"content"`
	// createAt 表示博客创建时间
	CreateAt time.Time `json:"createAt"`
	// updateAt 表示博客最后更新时间
	UpdateAt time.Time `json:"updateAt"`
}

// CreatePostRequest 表示创建文章请求
type CreatePostRequest struct {
	// title 表示博客标题
	Title string `json:"title"`
	// content 表示博客内容
	Content string `json:"content"`
}

// CreatePostResponse 表示创建文章响应
type CreatePostResponse struct {
	// postID 表示创建的文章 ID
	PostID string `json:"postID"`
}

// UpdatePostRequest 表示更新文章请求
type UpdatePostRequest struct {
	// postID 表示要更新的文章 ID，对应 {postID}
	PostID string `json:"postID" uri:"postID"`
	// title 表示更新后的博客标题
	Title *string `json:"title"`
	// content 表示更新后的博客内容
	Content *string `json:"content"`
}

// UpdatePostResponse 表示更新文章响应
type UpdatePostResponse struct {
}

// DeletePostRequest 表示删除文章请求
type DeletePostRequest struct {
	// postIDs 表示要删除的文章 ID 列表
	PostIDs []string `json:"postIDs"`
}

// DeletePostResponse 表示删除文章响应
type DeletePostResponse struct {
}

// GetPostRequest 表示获取文章请求
type GetPostRequest struct {
	// postID 表示要获取的文章 ID
	PostID string `json:"postID" uri:"postID"`
}

// GetPostResponse 表示获取文章响应
type GetPostResponse struct {
	// post 表示返回的文章信息
	Post *Post `json:"post"`
}

// ListPostRequest 表示获取文章列表请求
type ListPostRequest struct {
	// offset 表示偏移量
	Offset int64 `json:"offset"`
	// limit 表示每页数量
	Limit int64 `json:"limit"`
	// title 表示可选的标题过滤
	Title *string `json:"title"`
}

// ListPostResponse 表示获取文章列表响应
type ListPostResponse struct {
	// total_count 表示总文章数
	TotalCount int64 `json:"total_count"`
	// posts 表示文章列表
	Posts []*Post `json:"posts"`
}
