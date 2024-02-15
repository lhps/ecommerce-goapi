package entity_test

import (
	"testing"

	"github.com/lhps/ecommerce-goapi/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestEntity_NewProduct(t *testing.T) {
	categoryName := "Category 1"
	category := entity.NewCategory(categoryName)
	require.NotNil(t, category)
	require.NotEmpty(t, category.ID)

	productName := "Product 1"
	productDescription := "Product 1 description"
	productPrice := 100.00
	productImageUrl := "http://www.google.com"
	productCategoryId := category.ID

	product := entity.NewProduct(productName, productDescription, productCategoryId, productImageUrl, productPrice)

	require.NotNil(t, product)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.CategoryID, category.ID)
}
