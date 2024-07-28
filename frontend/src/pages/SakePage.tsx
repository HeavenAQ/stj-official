import { FC, useEffect, useState } from 'react'
import { redirect, useParams } from 'react-router-dom'
import Markdown from 'react-markdown'
import { ProductWithInfo } from '../types/api/product'
import { useProductQuery } from '../utils/query'
import Loading from '../components/Loading'

interface SakeDescriptionProps {
  introduction?: string
  itemInfo?: string
  recommendation?: string
  prize?: string
}

const SakeDescriptionTabs: FC<SakeDescriptionProps> = ({
  introduction,
  itemInfo,
  recommendation,
  prize
}) => {
  const [activeTab, setActiveTab] = useState(0)
  const activeTabStyle = () => {
    switch (activeTab) {
      case 0:
        return 'after:ml-0'
      case 1:
        return 'after:ml-[calc(100%/3)]'
      case 2:
        return 'after:ml-[calc(100%/3*2)]'
    }
  }

  const activeContentStyle = () => {
    switch (activeTab) {
      case 0:
        return 'translate-x-0'
      case 1:
        return '-translate-x-[calc(100%/3)]'
      case 2:
        return '-translate-x-[calc(100%/3*2)]'
    }
  }
  return (
    <div className="overflow-x-hidden lg:mt-10">
      <div className="text-sm font-semibold text-center text-gray-500 border-b-2">
        <div
          className={`relative grid grid-cols-3 mx-auto rounded-lg after:absolute after:inset-0 after:w-[calc(100%/3)] after:bg-gray-200 after:rounded-tl-lg after:rounded-tr-lg after:-z-10 isolate ${activeTabStyle()} after:duration-300`}
        >
          <button
            className="inline-block p-2 rounded-t-lg border-b-2 border-transparent duration-300 cursor-pointer md:hover:bg-zinc-100"
            onClick={() => setActiveTab(0)}
          >
            商品描述
          </button>
          <button
            className="inline-block p-2 rounded-t-lg border-b-2 border-transparent duration-300 cursor-pointer md:hover:bg-gray-100"
            aria-current="page"
            onClick={() => setActiveTab(1)}
          >
            商品詳情
          </button>
          <button
            className="inline-block p-2 rounded-t-lg border-b-2 border-transparent duration-300 cursor-pointer md:hover:bg-gray-100"
            onClick={() => setActiveTab(2)}
          >
            其他資訊
          </button>
        </div>
      </div>
      <div
        className={`py-4 w-[300%] grid grid-cols-3 transition-transform duration-300 ease-in-out ${activeContentStyle()}`}
      >
        <div>
          {prize !== undefined ? (
            <h1 className="mb-3 font-bold text-md">{prize}</h1>
          ) : null}
          <p>{introduction}</p>
        </div>
        <div>
          <Markdown
            components={{
              ul: ({ ...props }) => (
                <ul className="ml-8 list-disc">{props.children}</ul>
              )
            }}
          >
            {itemInfo}
          </Markdown>
        </div>
        <div>
          <Markdown
            components={{
              ul: ({ ...props }) => (
                <ul className="ml-8 list-disc">{props.children}</ul>
              )
            }}
          >
            {recommendation}
          </Markdown>
        </div>
      </div>
    </div>
  )
}

interface SakeCardProps {
  sake: ProductWithInfo
}
const SakeCard: FC<SakeCardProps> = ({ sake }) => {
  useEffect(() => {
    const hash = window.location.hash
    if (hash) {
      const id = hash.replace('#', '')
      const element = document.getElementById(id)
      element?.scrollIntoView({ behavior: 'smooth' })
    }
  }, [])

  return (
    <div id={sake.id}>
      <div className="gap-8 sm:gap-10">
        <div className="grid grid-rows-2 h-full">
          <div className="mx-auto h-auto text-center max-w-96">
            <img
              className="object-cover w-full h-full rounded-lg animate-skeleton"
              width="100%"
              height="100%"
              src={sake.imageURLs[0]}
              alt="service"
              loading="lazy"
            />
          </div>
          <div className="grid mt-6 list-disc">
            <SakeDescriptionTabs
              introduction={sake.introduction}
              itemInfo={sake.item_info}
              recommendation={sake.recommendation}
              prize={sake.prize}
            />
          </div>
        </div>
      </div>
    </div>
  )
}

const Sake: FC = () => {
  const { sakeId } = useParams()
  if (sakeId === undefined) {
    redirect('/')
  }

  const { isPending, data: sakeInfo } = useProductQuery(sakeId as string)
  if (isPending || !sakeInfo) {
    return (
      <div className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
        <Loading />
      </div>
    )
  }

  return (
    <div className="px-4 mx-auto w-full sm:px-8 md:px-14 max-w-[1000px] animate-fade">
      <div
        key={sakeInfo.data.name}
        id={sakeInfo.data.name}
        className="pt-8 pb-16 lg:items-center lg:pt-20 lg:pb-0"
      >
        <SakeCard sake={sakeInfo.data} />
      </div>
    </div>
  )
}
export default Sake
