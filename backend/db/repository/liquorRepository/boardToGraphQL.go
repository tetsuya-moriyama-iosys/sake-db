/**
  GraphQL変換ロジックが煩雑になったので分ける
*/

package liquorRepository

import "backend/graph/graphModel"

func (m *BoardListResponse) ToGraphQL() *graphModel.UserEvaluateList {
	//resentPostsの変換
	var recentPosts []*graphModel.UserLiquor
	for _, post := range m.RecentPosts {
		recentPosts = append(recentPosts, post.ToGraphQL())
	}

	var rate5Liquors []*graphModel.UserLiquor
	var rate4Liquors []*graphModel.UserLiquor
	var rate3Liquors []*graphModel.UserLiquor
	var rate2Liquors []*graphModel.UserLiquor
	var rate1Liquors []*graphModel.UserLiquor
	var noRateLiquors []*graphModel.UserLiquor

	rateMap := map[int]*[]*graphModel.UserLiquor{
		5: &rate5Liquors,
		4: &rate4Liquors,
		3: &rate3Liquors,
		2: &rate2Liquors,
		1: &rate1Liquors,
	}

	for _, group := range m.GroupedByRate {
		// group.Rate が nil の場合は noRateLiquors へ
		if group.Rate == nil {
			for _, post := range group.Posts {
				noRateLiquors = append(noRateLiquors, post.ToGraphQL())
			}
			continue
		}

		// 対応する rate のスライスに追加
		if targetSlice, exists := rateMap[*group.Rate]; exists {
			for _, post := range group.Posts {
				*targetSlice = append(*targetSlice, post.ToGraphQL())
			}
		}
	}

	return &graphModel.UserEvaluateList{
		RecentComments: recentPosts,
		Rate5Liquors:   rate5Liquors,
		Rate4Liquors:   rate4Liquors,
		Rate3Liquors:   rate3Liquors,
		Rate2Liquors:   rate2Liquors,
		Rate1Liquors:   rate1Liquors,
		NoRateLiquors:  noRateLiquors,
	}

}

func (m *LiquorDetail) ToGraphQL() *graphModel.UserLiquor {
	return &graphModel.UserLiquor{
		ID:           m.ID.Hex(),
		Name:         m.Name,
		CategoryID:   m.CategoryID,
		CategoryName: m.CategoryName,
		ImageBase64:  m.ImageBase64,
	}

}

func (m *Post) ToGraphQL() *graphModel.UserLiquor {
	comment := m.Text
	return &graphModel.UserLiquor{
		ID:           m.ID.Hex(),
		LiquorID:     m.Liquor.ID.Hex(),
		Name:         m.Liquor.Name,
		CategoryID:   m.Liquor.CategoryID,
		CategoryName: m.Liquor.CategoryName,
		ImageBase64:  m.Liquor.ImageBase64,
		Comment:      &comment,
		Rate:         m.Rate,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (m *BoardModel) ToGraphQL() *graphModel.BoardPost {
	//userはnilの可能性があり、そのままObjectIDを変換して*stringに代入できないので変換
	var userId *string
	if m.UserId != nil {
		id := m.UserId.Hex()
		userId = &id
	}
	return &graphModel.BoardPost{
		ID:        m.ID.Hex(),
		UserID:    userId,
		LiquorID:  m.LiquorID.Hex(),
		Text:      m.Text,
		Rate:      m.Rate,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *BoardModelWithRelation) ToGraphQL() *graphModel.BoardPost {
	//userはnilの可能性があり、そのままObjectIDを変換して*stringに代入できないので変換
	var userId *string
	if m.UserId != nil {
		id := m.UserId.Hex()
		userId = &id
	}
	return &graphModel.BoardPost{
		ID:           m.ID.Hex(),
		UserName:     m.UserName,
		UserID:       userId,
		CategoryID:   m.CategoryID,
		CategoryName: m.CategoryName,
		LiquorID:     m.LiquorID.Hex(),
		LiquorName:   m.LiquorName,
		Text:         m.Text,
		Rate:         m.Rate,
		UpdatedAt:    m.UpdatedAt,
	}
}
