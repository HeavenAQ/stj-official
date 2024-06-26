export default function NotFound404() {
  return (
    <main className="flex flex-col justify-center items-center w-full absolute top-1/2 right-1/2 translate-x-1/2 -translate-y-3/4">
      <h1 className="text-9xl font-extrabold tracking-widest text-zinc-600">
        404
      </h1>
      <div className="absolute px-2 text-sm bg-red-400 rounded rotate-12">
        Page Not Found
      </div>
      <button className="mt-5">
        <a
          href="/"
          className="inline-block relative text-sm font-medium text-red-400 focus:ring focus:outline-none active:text-orange-500 group"
        >
          <p>
            <span className="block relative py-3 px-8 rounded-md border border-current duration-300 hover:text-white hover:bg-zinc-600">
              Go back home
            </span>
          </p>
        </a>
      </button>
    </main>
  )
}
