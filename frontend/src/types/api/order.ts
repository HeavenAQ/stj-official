type Item = {
  productId: string
  quantity: number
  name: string
}

export type CreateOrderRequest = {
  ShippingAddress: string
  Phone: string
  Email: string
  Items: Array<Item>
}
