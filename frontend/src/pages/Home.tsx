import { sakes } from '../data/sakes'
import ItemCard from '../components/ItemCard'

const Home: React.FC = () => {
  const allSakes = sakes.getSakeInfoList()
  return (
    <main className="px-3 mx-auto mt-28 h-full font-noto-serif animate-fade-down ">
      <header className="mb-14 w-full rounded-lg shadow-lg h-[350px] shadow-gray-500 md:h-[600px]">
        <img
          className="object-cover w-full h-full rounded-lg"
          src="/images/misc/Hero.webp"
          alt="home"
        />
      </header>
      <h1 className="mb-10 text-3xl font-bold text-center text-zinc-700">
        所有商品
      </h1>
      <hr className="my-8 w-full bg-gray-200 border-0 dark:bg-gray-700 h-[2px]" />
      <div className="grid grid-cols-1 gap-y-10 mx-auto mb-32 w-full sm:grid-cols-2 sm:gap-x-10 lg:grid-cols-3">
        {allSakes.map((sake, i) => {
          return <ItemCard key={i} item={sake} href={String(i)} />
        })}
      </div>
    </main>
  )
}
export default Home
