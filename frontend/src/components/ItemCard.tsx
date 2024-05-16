import React, { useState } from 'react'
import { SakeInfo } from '../data/sakes'

interface ItemCardProps {
  item: SakeInfo
}

const ItemCard: React.FC<ItemCardProps> = ({ item }) => {
  const [quantity, setQuantity] = useState(0)
  return (
    <div className="relative mx-auto w-full max-w-sm bg-white rounded-lg border border-gray-200 shadow-md duration-300 dark:bg-gray-800 dark:border-gray-700">
      {item.isHot && (
        <span className="inline-flex absolute top-2 left-2 items-center py-0.5 px-2.5 mb-2 text-xs font-medium text-red-800 bg-red-100 rounded-full dark:text-red-300 dark:bg-red-900">
          <span className="w-2 h-2 bg-red-500 rounded-full me-1"></span>
          熱銷
        </span>
      )}
      <a href="#">
        <img
          className="p-4 rounded-t-lg"
          src={item.images[0]}
          alt="product image"
        />
        <div className="absolute top-[65%] lg:top-64 left-1/2 py-2 px-3 font-medium leading-none text-center rounded-full animate-pulse -translate-x-1/2 text-md text-zinc-600 bg-zinc-200 dark:text-zinc-200 dark:bg-zinc-900">
          點擊看更多
        </div>
      </a>

      <div className="inline-flex justify-evenly items-center mb-5 w-full">
        <button
          className="w-8 h-8 text-zinc-700"
          type="button"
          onClick={() => {
            if (quantity === 0) return
            setQuantity(quantity - 1)
          }}
        >
          -
        </button>
        <div className="w-3">{quantity}</div>
        <button
          className="w-8 h-8 text-zinc-700"
          type="button"
          onClick={() => setQuantity(quantity + 1)}
        >
          +
        </button>
      </div>
      <div className="px-5 pb-5 mx-auto w-44">
        <div className="flex justify-center items-center">
          <a
            href="#"
            className="py-2.5 px-5 text-sm font-medium text-center text-white rounded-lg duration-300 focus:ring-4 focus:outline-none bg-zinc-700 dark:bg-zinc-600 dark:hover:bg-zinc-400 dark:focus:ring-zinc-800 hover:bg-zinc-500 focus:ring-zinc-300"
          >
            加入詢價單
          </a>
        </div>
      </div>
    </div>
  )
}

export default ItemCard
