import React from 'react'

export const OrderContext = React.createContext({
  order: [],
  setOrder: (order: any) => {}
})
