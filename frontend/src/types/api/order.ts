enum OrderStatus {
  PENDING = 'pending',
  PROCESSING = 'processing',
  SHIPPED = 'shipped',
  DELIVERED = 'delivered',
  CANCELLED = 'cancelled'
}

type Item = {
  productId: string
  quantity: number
  name: string
}

export type CreateOrderRequest = {
  ShippingAddress: string
  Phone: string
  Email: string
  Items: Item[]
}

export type CreateOrderResponse = {
  created_at: string
  order_status: OrderStatus
  items: Item[]
  order_id: string
}
