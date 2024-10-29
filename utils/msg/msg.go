package msg

const (
	// common
	SUCCESS = 200
	ERROR   = 500

	// 文章
	ERROR_ARTICLE_NOT_EXIST = 2001

	// 分类
	ERROR_CATEGORY_EXIST     = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	// common
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	// 文章
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",

	// 分类
	ERROR_CATEGORY_EXIST:     "该分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "该分类不存在",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
