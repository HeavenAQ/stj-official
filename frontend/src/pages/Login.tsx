export default function Login() {
  return (
    <div className="max-h-[700px] w-[90%] sm:w-[80%] md:w-[70%] max-w-[500px] h-[60%] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-xl bg-zinc-600 flex flex-col px-10 shadow-xl shadow-gray-500 justify-center">
      <h1 className="text-3xl text-center text-white">會員登入</h1>
      <form>
        <div className="flex flex-col mt-5">
          <input
            placeholder="電子信箱"
            type="email"
            className="py-2 px-4 rounded-xl"
          />
        </div>
        <div className="flex flex-col mt-5">
          <input
            placeholder="密碼"
            type="password"
            className="py-2 px-4 rounded-xl"
          />
        </div>
        <div className="inline-flex justify-start items-center mt-2 w-full">
          <input type="checkbox" className="m-2 cursor-pointer" />
          <label className="text-white">記住我</label>
          <a href="#" className="ml-auto text-white cursor-pointer">
            <label className="duration-300 cursor-pointer hover:text-zinc-400">
              忘記密碼
            </label>
          </a>
        </div>
        <button className="p-2 mt-5 w-full text-white rounded-xl duration-300 bg-zinc-500 hover:bg-zinc-900 active:bg-zinc-500">
          登入
        </button>

        <div className="inline-flex justify-center items-center w-full">
          <hr className="my-8 w-full h-px bg-gray-200 border-0 dark:bg-gray-700" />
          <span className="absolute left-1/2 px-3 font-medium -translate-x-1/2 bg-zinc-600 text-slate-300">
            or
          </span>
        </div>
        <button className="p-2 w-full text-white rounded-xl duration-300 bg-zinc-900 hover:bg-zinc-500 active:bg-zinc-400">
          註冊
        </button>
      </form>
    </div>
  );
}
