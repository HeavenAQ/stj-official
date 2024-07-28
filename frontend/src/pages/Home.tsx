import ItemCard from '../components/ItemCard'
import Loading from '../components/Loading'
import { useProductsQuery } from '../utils/query'

const Home: React.FC = () => {
  const { isPending, data: items } = useProductsQuery(10, 0)
  if (isPending) {
    return <Loading />
  }

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
        {items?.data.map((sake, i) => {
          return <ItemCard key={i} item={sake} />
        })}
      </div>
    </main>
  )
}
export default Home
