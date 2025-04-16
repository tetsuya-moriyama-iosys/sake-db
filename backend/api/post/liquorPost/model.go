package liquorPost

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Id                *string `form:"id" binding:"omitempty,len=24"`
	Name              string  `form:"name" binding:"required,max=100"`
	CategoryID        int     `form:"category" binding:"required,gt=1"`
	Description       string  `form:"description" binding:"omitempty,max=5000"`
	Youtube           string  `form:"youtube" binding:"omitempty,youtube"` //TODO: youtubeはvalidatorの拡張なのでたぶん適用されなさそう
	VersionNo         *int    `form:"version_no" binding:"omitempty,gte=1"`
	SelectedVersionNo *int    `form:"selected_version_no" binding:"omitempty,gte=1"`
}
