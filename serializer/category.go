package serializer

import "gin_shop/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []*model.Category) (categorys []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categorys = append(categorys, category)
	}
	return categorys
}
