package validators

// Articles 内容表
type ArticleAdd struct {
	CategoryID      int64  `form:"category_id" json:"category_id" valid:"required~缺少cate_id,int,min(1)~cate_id必须大于0"` // 分类
	UserID          int64  `form:"user_id"  json:"user_id" valid:"required,int,min(1)"`                               // 作者
	Slug            string `form:"slug"  json:"slug" valid:"required,runelength(3|50)"`                               // 短链
	Title           string `form:"title"  json:"title" valid:"required,runelength(3|50)"`                             // 标题
	Subtitle        string `form:"subtitle"  json:"subtitle" valid:"optional"`                                        // 副标题
	Content         string `form:"content"  json:"content" valid:"required"`                                          // 内容
	PageImage       string `form:"page_image"  json:"page_image" valid:"required,url"`                                // 主图
	MetaDescription string `form:"meta_description"  json:"meta_description" valid:"optional"`                        // seo内容
	Recommend       int8   `form:"recommend" json:"recommend" valid:"required,range(0|1)"`                            // 是否推荐
	Sort            int    `form:"sort" json:"sort" valid:"required,int"`                                             // 排序
	//State           int8    `form:"state"  json:"state" valid:"required,range(0|1)"`                                   // 0默认草稿 1已发布
	ViewCount int `form:"view_count"  json:"view_count" valid:"required,int"` // 浏览量
}
