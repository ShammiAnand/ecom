package cart

import (
	"errors"

	"github.com/shammianand/ecom/types"
)

func getProductIDsFromCartItems(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, errors.New("invalid quantity")
		}
		productIds = append(productIds, item.ProductId)
	}
	return productIds, nil
}

func (h *Handler) createOrder(
	ps []types.Product,
	items []types.CartItem,
	userID int,
) (int, float64, error) {

	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are in stock
	err := checkIfCartIsInStock(items, productMap)
	if err != nil {
		return 0, 0.0, err
	}
	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)
	// reduce the quantity of products
	for _, item := range items {
		product := productMap[item.ProductId]
		product.Quantity -= item.Quantity
		err := h.productStore.UpdateProduct(product)
		if err != nil {
			return 0, 0, err
		}
	}

	// create order and order items for each product
	orderId, err := h.store.CreateOrder(types.Order{
		UserId:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "India",
	})

	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductId].Price,
		})
	}

	return orderId, totalPrice, nil
}

func checkIfCartIsInStock(items []types.CartItem, productMap map[int]types.Product) error {
	if len(items) == 0 {
		return errors.New("cart is empty")
	}

	for _, item := range items {
		product, ok := productMap[item.ProductId]
		if !ok {
			return errors.New("product ID does not exist")
		}
		if product.Quantity < item.Quantity {
			return errors.New("not enought quantity")
		}
	}

	return nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	totalPrice := 0.0
	for _, item := range items {
		totalPrice += float64(item.Quantity) * productMap[item.ProductId].Price
	}
	return totalPrice
}
