package service

import (
	"hmdp-go/internal/constant"
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
	"hmdp-go/pkg/redisutil"
	"hmdp-go/pkg/util"
	"log"
)

// BlogService 博客服务接口
type BlogService interface {
	QueryMyBlog(userID uint64, page *model.PaginationRequest) ([]*model.Blog, error)
	QueryHotBlog(userID uint64, page *model.PaginationRequest) ([]*model.Blog, error)
	QueryBlogById(userID uint64, id uint64) (*model.Blog, error)
}

// blogService 博客服务实现
type blogService struct {
	repo *repository.Repository
}

// NewBlogService 创建博客服务实例
func NewBlogService(repo *repository.Repository) BlogService {
	return &blogService{repo: repo}
}

// QueryMyBlog 查询当前登录用户的博客
func (s *blogService) QueryMyBlog(userID uint64, page *model.PaginationRequest) ([]*model.Blog, error) {
	return s.repo.BlogRepo.FindByUserIDWithPage(userID, page.GetOffset(), page.PageSize)
}

// QueryHotBlog 查询热门探店笔记，按liked降序
func (s *blogService) QueryHotBlog(userID uint64, page *model.PaginationRequest) ([]*model.Blog, error) {
	// 调用仓库层查询热门博客，按liked降序
	blogs, err := s.repo.BlogRepo.FindHotWithPage(page.GetOffset(), page.PageSize)
	if err != nil {
		return nil, err
	}
	// 遍历博客列表，查询关联的用户信息和是否被当前登录用户点赞
	for _, blog := range blogs {
		// 查询博客关联的用户信息
		s.queryBlogUser(blog)
		// 判断博客是否被当前登录用户点赞
		s.isBlogLiked(userID, blog)
	}

	return blogs, nil
}

// QueryBlogById 根据ID查询博客详情
func (s *blogService) QueryBlogById(userID uint64, id uint64) (*model.Blog, error) {
	blog, err := s.repo.BlogRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if blog == nil {
		return nil, nil
	}

	// 查询博客关联的用户信息
	s.queryBlogUser(blog)

	// 判断博客是否被当前登录用户点赞
	s.isBlogLiked(userID, blog)

	return blog, nil
}

// queryBlogUser 查询博客关联的用户信息
func (s *blogService) queryBlogUser(blog *model.Blog) {
	// 根据用户ID查询用户信息
	user, err := s.repo.UserRepo.FindByID(blog.UserID)
	if err != nil {
		log.Printf("查询用户信息失败,用户ID: %d,错误: %v\n", blog.UserID, err)
		return
	}

	// 设置用户昵称和头像
	blog.Name = user.NickName
	blog.Icon = user.Icon
}

// isBlogLiked 判断博客是否被当前登录用户点赞
func (s *blogService) isBlogLiked(userId uint64, blog *model.Blog) {
	userIdStr := util.Uint64ToString(userId)
	// 1.判断当前登录用户是否已经点赞
	key := constant.BLOG_LIKED_KEY + userIdStr
	score, err := redisutil.ZScore(key, userIdStr)
	if err != nil {
		log.Printf("查询Redis点赞记录失败,Key: %s,错误: %v\n", key, err)
		blog.IsLike = false
		return
	}
	// 2. 设置是否被点赞
	blog.IsLike = score != 0
}
