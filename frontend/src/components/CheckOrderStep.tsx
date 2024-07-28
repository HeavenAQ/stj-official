import { CartItem, useCart } from '../contexts/CartContext'
import React, { useRef } from 'react'
import { RxCrossCircled } from 'react-icons/rx'
import Markdown from 'react-markdown'
import { QuantityOptions } from '../utils/options'

const CheckOrderStep = () => {
  const { cart, removeFromCart, updateCart } = useCart()

  return (
    <table className="w-full text-left bg-white rounded-lg overflow-hidden mb-10">
      <thead className="h-14 bg-zinc-600 text-white">
        <tr>
          <th className="sm:w-20 sm:table-cell hidden"></th>
          <th className="sm:pl-0 pl-3 min-w-32">
            <p className="ml-2">商品</p>
          </th>
          <th className="min-w-32">
            <p className="ml-2">名稱</p>
          </th>
          <th className="hidden lg:table-cell">
            <p className="ml-2">資訊</p>
          </th>
          <th className="text-center min-w-20">
            <p>數量</p>
          </th>
        </tr>
      </thead>
      <tbody className="bg-white">
        {cart?.map(item => (
          <ProductRow
            key={item.product.id}
            item={item}
            removeFromCart={removeFromCart}
            updateCart={updateCart}
          />
        ))}
      </tbody>
    </table>
  )
}

interface ProductRowProps {
  item: CartItem
  removeFromCart: (id: string) => void
  updateCart: (id: string, quantity: number) => void
}

const ProductRow: React.FC<ProductRowProps> = ({
  item,
  removeFromCart,
  updateCart
}) => {
  const rowRef = useRef<HTMLTableRowElement>(null)
  let downX: number
  let fromX = 0

  const onPointerMove = (e: PointerEvent | TouchEvent) => {
    const newX = 'touches' in e ? e.touches[0].clientX : e.clientX
    if (rowRef.current) {
      let toX = fromX + (newX - downX) * 2
      toX = Math.min(0, Math.max(toX, -80))
      rowRef.current.style.transform = `translate(${toX}px)`

      setTimeout(() => {
        if (rowRef.current) {
          rowRef.current.style.transform = `translate(${toX < -20 ? -80 : 0}px)`
        }
      }, 100)

      fromX = toX
    }
  }

  return (
    <tr
      className="mb-4 h-32 transition-all duration-700 ease-in-out cursor-pointer relative"
      ref={rowRef}
      onPointerDown={e => {
        downX = e.clientX
        rowRef.current?.addEventListener('pointermove', onPointerMove, {
          passive: true
        })
      }}
      onPointerUp={() => {
        rowRef.current?.removeEventListener('pointermove', onPointerMove)
      }}
      onTouchStart={e => {
        downX = e.touches[0].clientX
        rowRef.current?.addEventListener('touchmove', onPointerMove, {
          passive: true
        })
      }}
      onTouchEnd={() => {
        rowRef.current?.removeEventListener('touchmove', onPointerMove)
      }}
    >
      <td
        className="z-10 rounded-tl-lg rounded-bl-lg duration-200 ease-in-out hover:bg-red-300 group/delete hidden sm:table-cell"
        onClick={() => removeFromCart(item.product.id)}
      >
        <RxCrossCircled className="mx-auto text-xl text-red-500 group-hover/delete:text-white" />
      </td>
      <td>
        <img
          src={item.product.imageURLs[0]}
          width="100"
          height="100"
          alt="product"
        />
      </td>
      <td className="font-bold text-md">{item.product.name}</td>
      <td className="hidden lg:table-cell">
        <Markdown
          components={{
            ul: ({ node, ...props }) => (
              <ul className="ml-8 list-disc">{props.children}</ul>
            )
          }}
        >
          {item.product.item_info}
        </Markdown>
      </td>
      <td className="rounded-tr-lg rounded-br-lg text-center">
        <select
          id={'quantity' + item.product.id}
          className="cursor-pointer"
          defaultValue={item.quantity}
          onChange={e => updateCart(item.product.id, Number(e.target.value))}
        >
          <QuantityOptions />
        </select>
      </td>
      <td
        className="w-20 h-full bg-red-500 absolute right-0 text-white flex items-center justify-center hover:bg-red-300 translate-x-20"
        onClick={() => removeFromCart(item.product.id)}
      >
        刪除
      </td>
    </tr>
  )
}

export default CheckOrderStep
