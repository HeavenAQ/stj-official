// product status enum
enum ProductStatus {
  IN_STOCK = 'in-stock',
  OUT_OF_STOCK = 'out-of-stock',
  DISCONTINUED = 'discontinued'
}

// response type of listProducts
export type ProductWithInfo = {
  price: number
  id: string
  imageURLs: string[]
  is_hot: boolean
  status: ProductStatus
  quantity: number
  createdAt: string
  updatedAt: string
  category: string
  name: string
  introduction: string
  prize: string
  item_info: string
  recommendation: string
}
