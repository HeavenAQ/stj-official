import { SakeInfo } from '../data/sakes'
import React, { useState } from 'react'

interface ItemCardProps {
  item: SakeInfo
  href: string
}

const ItemCard: React.FC<ItemCardProps> = ({ item, href }) => {
  const [quantity, setQuantity] = useState(0)
  const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      entry.isIntersecting
        ? entry.target.classList.add('animate-fade-up')
        : entry.target.classList.remove('animate-fade-up')
    })
  })
  return (
    <div
      className="relative mx-auto w-full max-w-sm bg-white rounded-lg border border-gray-200 shadow-md duration-300 dark:bg-gray-800 dark:border-gray-700"
      ref={node => {
        node && observer.observe(node)
      }}
    >
      {item.isHot && (
        <span className="inline-flex absolute top-2 left-2 items-center py-0.5 px-2.5 mb-2 text-xs font-medium text-red-800 bg-red-100 rounded-full dark:text-red-300 dark:bg-red-900">
          <span className="w-2 h-2 bg-red-500 rounded-full me-1"></span>
          熱銷
        </span>
      )}
      <a href={`/sakes/${href}`}>
        <img className="p-4 rounded-t-lg" src={item.images[0]} alt="product" />
      </a>
      <a href={`/sakes/${href}`}>
        <div className="w-1/2 mx-auto mb-4 py-2 px-2 font-medium leading-none text-center rounded-full animate-pulse  text-md text-zinc-600 bg-zinc-200 dark:text-zinc-200 dark:bg-zinc-900">
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
            href="/"
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
