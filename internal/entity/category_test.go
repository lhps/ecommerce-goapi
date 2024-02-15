package entity_test

import (
	"testing"

	"github.com/lhps/ecommerce-goapi/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestEntity_NewCategory(t *testing.T) {
	categoryName := "Category 1"
	category := entity.NewCategory(categoryName)
	require.NotNil(t, category)
	require.NotEmpty(t, category.ID)
}
