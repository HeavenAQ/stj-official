import { useState } from "react";
import { toast } from "react-hot-toast";

export default function PasswordReset() {
  const [email, setEmail] = useState<string>("");

  const registerUser = (e: React.FormEvent) => {
    e.preventDefault();
  };

  return (
    <div className="w-[90%] sm:w-[80%] md:w-[70%] max-w-[500px] max-h-[700px] h-[60%] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-xl bg-zinc-600 flex flex-col px-10 shadow-xl shadow-gray-500 justify-center animate-fade">
      <h1 className="mb-3 text-3xl text-center text-white">會員註冊</h1>
      <form onSubmit={(e) => registerUser(e)}>
        <div className="flex flex-col mt-5">
          <input
            placeholder="電子信箱"
            type="email"
            className="py-2 px-4 rounded-xl"
            onChange={(ev) => setEmail(ev.target.value)}
            required
          />
        </div>
        <button className="p-2 mt-5 w-full text-white rounded-xl duration-300 bg-zinc-900 hover:bg-zinc-500 active:bg-zinc-400">
          密碼重置
        </button>
      </form>
    </div>
  );
}
