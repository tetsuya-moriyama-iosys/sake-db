package liquorPost

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Id                *string `form:"id" validate:"omitempty,len=24"`
	Name              string  `form:"name" validate:"required,max=100"`
	CategoryID        int     `form:"category" validate:"required,gt=1"`
	Description       *string `form:"description" validate:"omitempty,max=5000"`
	Youtube           *string `form:"youtube" validate:"omitempty,url"`
	VersionNo         *int    `form:"version_no" validate:"omitempty,gt=1"`
	SelectedVersionNo *int    `form:"selected_version_no" validate:"omitempty,gt=1"`
}
