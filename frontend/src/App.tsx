import { FC, useEffect, useState } from "react";
import { sakes, SakeInfo, SakeDescription } from "./data/sakes";
import Markdown from "react-markdown";

interface SakeDescriptionProps {
  description: SakeDescription;
}
const SakeDescriptionTabs: FC<SakeDescriptionProps> = ({ description }) => {
  const [activeTab, setActiveTab] = useState(0);
  const activeTabStyle = () => {
    switch (activeTab) {
      case 0:
        return "after:ml-0";
      case 1:
        return "after:ml-[calc(100%/3)]";
      case 2:
        return "after:ml-[calc(100%/3*2)]";
    }
  };

  const activeContentStyle = () => {
    switch (activeTab) {
      case 0:
        return "translate-x-0";
      case 1:
        return "-translate-x-[calc(100%/3)]";
      case 2:
        return "-translate-x-[calc(100%/3*2)]";
    }
  };
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
              ),
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
              ),
            }}
          >
            {description?.recommendation}
          </Markdown>
        </div>
      </div>
    </div>
  );
};

interface SakeCardProps {
  id: number;
  sake: SakeInfo;
}
const SakeCard: FC<SakeCardProps> = ({ id, sake }) => {
  const isImgLeft = id % 2 === 0;

  useEffect(() => {
    const hash = window.location.hash;
    if (hash) {
      const id = hash.replace("#", "");
      const element = document.getElementById(id);
      element?.scrollIntoView({ behavior: "smooth" });
    }
  }, []);

  return (
    <div id={sake.id}>
      <div className="gap-8 sm:gap-10 lg:inline-grid lg:grid-cols-2">
        {isImgLeft && (
          <div className="hidden list-disc lg:grid">
            <SakeDescriptionTabs description={sake.description} />
          </div>
        )}
        <div className="flex flex-col h-full">
          <div className="relative mx-auto w-full h-full text-center">
            <img
              className="object-cover w-full h-full rounded-lg"
              width="100%"
              height="100%"
              src={sake.images[0]}
              alt="service"
            />
            <h2 className="hidden text-lg font-medium lg:inline-block font-m-plus tracking-[10px]">
              {sake.subtitle}
            </h2>
          </div>
          <div className="grid mt-6 list-disc lg:hidden">
            <SakeDescriptionTabs description={sake.description} />
          </div>
        </div>
        {!isImgLeft && (
          <div className="hidden list-disc lg:grid">
            <SakeDescriptionTabs description={sake.description} />
          </div>
        )}
      </div>
    </div>
  );
};

const Sake: FC = () => {
  const servicesInfo = sakes.getSakeInfoList();

  return (
    <div className="overflow-scroll px-8 mx-auto w-full md:px-14 h-[100vh] snap-y snap-mandatory scroll-smooth">
      {servicesInfo.map((sake, i) => (
        <div
          key={sake.title}
          id={sake.title}
          className="flex justify-center pt-8 pb-16 lg:items-center lg:pt-0 lg:pb-0 h-[100vh] snap-start"
        >
          <SakeCard id={i} sake={sake} />
        </div>
      ))}
    </div>
  );
};

export default Sake;
