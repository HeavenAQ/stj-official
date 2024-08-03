import React, { FC } from 'react'
import { useCart } from '../contexts/CartContext'

interface OrderSummaryStepProps {
  firstName: string
  lastName: string
  email: string
  phone: any
  address: string
  birthYear: number
}

const OrderSummaryStep: FC<OrderSummaryStepProps> = ({
  firstName,
  lastName,
  email,
  phone,
  address,
  birthYear
}) => {
  const today = new Date()
  const age = today.getFullYear() - birthYear
  const cartContext = useCart()

  return (
    <section className=" bg-zinc-600 rounded-xl mx-auto md:px-0 px-6 mb-14 pb-9 animate-fade w-[95%] sm:w-[85%] md:w-[80%] lg:w-[75%] text-white">
      <div className="pt-10 mx-auto w-full md:w-4/5 max-w-[1000px] space-y-6 md:text-xl sm:text-lg text-md">
        <h1 className="my-8 text-xl sm:text-2xl md:text-3xl text-center font-bold text-white lg:mb-12 lg:text-3xl tracking-[20px]">
          訂單內容
        </h1>
        <hr className="my-4 w-full bg-gray-200 border-0 dark:bg-gray-700 h-[2px]" />
        <p className="font-bold mt-10">購買人資訊：</p>
        <p className="grid grid-cols-2">
          <span>姓名：</span>
          <span>{lastName + ' ' + firstName}</span>
        </p>
        <p className="grid grid-cols-2">
          <span>年齡：</span>
          <span>{age} 歲</span>
        </p>
        <p className="grid grid-cols-2">
          <span>信箱：</span>
          <span className="break-words">{email}</span>
        </p>
        <p className="grid grid-cols-2">
          <span>電話：</span>
          <span>{phone}</span>
        </p>
        <p className="grid grid-cols-2">
          <span>寄送地址：</span>
          <span>{address}</span>
        </p>

        <hr className="my-2 w-full bg-gray-200 border-0 dark:bg-gray-700 h-[1px]" />
        <p className="font-bold mt-10">購買商品：</p>
        {cartContext.cart.map((item, idx) => (
          <p key={idx} className="grid grid-cols-2">
            <span>{item.product.name}</span>
            <span>{item.quantity} 支</span>
          </p>
        ))}
      </div>
    </section>
  )
}

export default OrderSummaryStep
