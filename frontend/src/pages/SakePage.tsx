import { FC, useEffect, useState } from 'react'
import { sakes, SakeInfo, SakeDescription } from '../data/sakes'
import { useParams } from 'react-router-dom'
import Markdown from 'react-markdown'

interface SakeDescriptionProps {
  description: SakeDescription
}
const SakeDescriptionTabs: FC<SakeDescriptionProps> = ({ description }) => {
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
            className="inline-block p-2 rounded-t-lg border-b-2 border-transparent duration-300 cursor-pointer md:hover:border-gray-500"
            onClick={() => setActiveTab(0)}
          >
            商品描述
          </button>
          <button
            className="inline-block p-2 rounded-t-lg border-b-2 border-transparent duration-300 cursor-pointer md:hover:border-gray-500"
            aria-current="page"
            onClick={() => setActiveTab(1)}
          >
            商品詳情
          </button>
          <button
            className="inline-block p-2 rounded-t-lg border-b-2 border-transparent duration-300 cursor-pointer md:hover:border-gray-500"
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
          {description && description.title !== undefined ? (
            <h1 className="mb-3 font-bold text-md">{description?.title}</h1>
          ) : null}
          <p>{description?.introduction}</p>
        </div>
        <div>
          <Markdown
            components={{
              ul: ({ node, ...props }) => (
                <ul className="ml-8 list-disc">{props.children}</ul>
              )
            }}
          >
            {description?.itemInfo}
          </Markdown>
        </div>
        <div>
          <Markdown
            components={{
              ul: ({ node, ...props }) => (
                <ul className="ml-8 list-disc">{props.children}</ul>
              )
            }}
          >
            {description?.recommendation}
          </Markdown>
        </div>
      </div>
    </div>
  )
}

interface SakeCardProps {
  sake: SakeInfo
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
              className="object-cover w-full h-full rounded-lg"
              width="100%"
              height="100%"
              src={sake.images[0]}
              alt="service"
            />
            <h2 className="hidden text-lg font-medium font-m-plus tracking-[10px]">
              {sake.subtitle}
            </h2>
          </div>
          <div className="grid mt-6 list-disc">
            <SakeDescriptionTabs description={sake.description} />
          </div>
        </div>
      </div>
    </div>
  )
}

const sakeInfos = sakes.getSakeInfoList()
const Sake: FC = () => {
  const { sakeId } = useParams()
  const sakeIdNum = Number(sakeId as string)

  return (
    <div className="px-4 mx-auto w-full sm:px-8 md:px-14 max-w-[1000px] animate-fade-down">
      <div
        key={sakeInfos[sakeIdNum].title}
        id={sakeInfos[sakeIdNum].title}
        className="pt-8 pb-16 lg:items-center lg:pt-20 lg:pb-0"
      >
        <SakeCard sake={sakeInfos[sakeIdNum]} />
      </div>
    </div>
  )
}
export default Sake
