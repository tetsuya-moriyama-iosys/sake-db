package categoryPost

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Id                *int    `form:"id" validate:"omitempty,gt=1"`
	Name              string  `form:"name" validate:"required,max=100"`
	Parent            int     `form:"parent" validate:"required,gt=1"`
	Description       *string `form:"description" validate:"omitempty,max=5000"`
	VersionNo         *int    `form:"version_no" validate:"omitempty,gt=1"`
	SelectedVersionNo *int    `form:"selected_version_no" validate:"omitempty,gt=1"`
}
