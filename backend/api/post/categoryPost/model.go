package categoryPost

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Id                *int    `form:"id" binding:"omitempty,gt=1"`
	Name              string  `form:"name" binding:"required,max=100"`
	Parent            int     `form:"parent" binding:"required,gt=1"`
	Description       *string `form:"description" binding:"omitempty,max=5000"`
	VersionNo         *int    `form:"version_no" binding:"omitempty,gt=1"`
	SelectedVersionNo *int    `form:"selected_version_no" binding:"omitempty,gt=1"`
}
