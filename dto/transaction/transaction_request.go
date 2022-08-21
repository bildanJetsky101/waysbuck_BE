package transactionsdto

//import category "waysbuck/models"

type CreateTransactionRequest struct {
	Price     int   `json:"price" form:"price" gorm:"type: int"`
	BuyerID   int   `json:"buyer_id" gorm:"type: int"`
	ProductID []int `json:"product_id" gorm:"type: int"`
}
