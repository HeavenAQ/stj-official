import { CartItem, useCart } from '../contexts/CartContext'
import React, { useRef, useState } from 'react'
import { RxCrossCircled } from 'react-icons/rx'
import Markdown from 'react-markdown'
import Stepper from '../components/Stepper'
import DeliveryInfo from '../components/DeliveryInfo'

export const OrderPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState(0)
  const activeContentStyle = () => {
    switch (activeTab) {
      case 0:
        return 'translate-x-[calc(100%/3)]'
      case 1:
        return 'translate-x-0'
      case 2:
        return '-translate-x-[calc(100%/3)]'
    }
  }
  return (
    <div className="w-full h-full flex flex-col items-center justify-center mb-28 overflow-x-hidden">
      <Stepper activeIdx={activeTab as 0 | 1 | 2} />

      <div
        className={`py-4 w-[300%] grid grid-cols-3 transition-transform duration-300 ease-in-out ${activeContentStyle()}`}
      >
        <div>
          <CheckOrderStep />
        </div>
        <div>
          <DeliveryInfo />
        </div>
      </div>

      <div className="inline-flex w-full">
        <button
          className="w-28 py-3 text-center tracking-[5px] bg-zinc-400 rounded-lg text-white duration-300 font-semibold hover:bg-zinc-300"
          onClick={() => {
            setActiveTab(prevTab => (prevTab === 0 ? 0 : prevTab - 1))
          }}
        >
          上一步
        </button>
        <button
          className="w-28 py-3 text-center tracking-[5px] bg-zinc-400 rounded-lg text-white duration-300 font-semibold hover:bg-zinc-300 ml-auto"
          onClick={() => {
            setActiveTab(prevTab => (prevTab === 3 ? 2 : prevTab + 1))
          }}
        >
          下一步
        </button>
      </div>
    </div>
  )
}

const CheckOrderStep = () => {
  const cartContext = useCart()
  return (
    <table className=" w-full text-left bg-white rounded-lg overflow-hidden mb-10">
      <thead className="h-14 bg-zinc-100">
        <tr>
          <th className="sm:w-20 sm:table-cell hidden"></th>
          <th className="sm:pl-0 pl-3 min-w-32">商品</th>
          <th className="min-w-32">名稱</th>
          <th className="hidden lg:table-cell">資訊</th>
          <th className="text-center min-w-20">數量</th>
        </tr>
      </thead>
      <tbody className="bg-white">
        {cartContext.cart?.map(item => (
          <ProductRow key={item.product.id} item={item} />
        ))}
      </tbody>
    </table>
  )
}

interface Prop {
  item: CartItem
}

const createQuantityOptions = () => {
  return Array.from({ length: 99 }, (_, i) => i + 1).map(i => (
    <option key={i} value={i}>
      {i}
    </option>
  ))
}

const ProductRow: React.FC<Prop> = ({ item }) => {
  const rowRef = useRef<HTMLTableRowElement>(null)
  const cartContext = useCart()
  const quantityOptions = createQuantityOptions()

  // for calculating swiping position
  let downX: number
  let fromX = 0

  // when delete button is clicked
  const removeItem = () => {
    cartContext.removeFromCart(item.product.id)
  }

  // handle pointer move (for showing delete button in small screen)
  const onPointerMove = (e: PointerEvent) => {
    const newX = e.clientX
    if (rowRef && rowRef.current) {
      // move along with the pointer
      let toX = fromX + (newX - downX) * 2
      if (toX <= -80) {
        toX = -80
      } else if (toX >= 0) {
        toX = 0
      }
      rowRef.current.style.transform = `translate(${toX}px)`

      // sleep for a while and show/hide the delete button
      setTimeout(() => {
        if (toX < -20) {
          toX = -80
        } else {
          toX = 0
        }
        if (rowRef.current) {
          rowRef.current.style.transform = `translate(${toX}px)`
        }
      }, 100)

      fromX = toX
    }
  }

  //

  return (
    <tr
      className="mb-4 h-32 transition-all duration-700 ease-in-out cursor-pointer relative"
      ref={rowRef}
      onPointerDown={e => {
        downX = e.clientX
        rowRef.current?.addEventListener('pointermove', onPointerMove)
      }}
      onPointerUp={() => {
        rowRef.current?.removeEventListener('pointermove', onPointerMove)
      }}
    >
      <td
        className="z-10 rounded-tl-lg rounded-bl-lg duration-200 ease-in-out hover:bg-red-300 group/delete hidden sm:table-cell"
        onClick={removeItem}
      >
        <RxCrossCircled className="mx-auto text-xl text-red-500 group-hover/delete:text-white" />
      </td>
      <td>
        <img
          src={item.product.images[0]}
          width="100"
          height="100"
          alt="product"
        />
      </td>
      <td className="font-bold text-md">{item.product.title}</td>
      <td className="hidden lg:table-cell">
        <Markdown
          components={{
            ul: ({ node, ...props }) => (
              <ul className="ml-8 list-disc">{props.children}</ul>
            )
          }}
        >
          {item.product.description.itemInfo}
        </Markdown>
      </td>
      <td className="rounded-tr-lg rounded-br-lg text-center">
        <select
          className="cursor-pointer"
          defaultValue={item.quantity}
          onChange={e =>
            cartContext.updateCart(item.product.id, Number(e.target.value))
          }
        >
          {quantityOptions}
        </select>
      </td>
      <td
        className="w-20 h-full bg-red-500 absolute right-0 text-white flex items-center justify-center hover:bg-red-300 translate-x-20"
        onClick={removeItem}
      >
        刪除
      </td>
    </tr>
  )
}
export default OrderPage
