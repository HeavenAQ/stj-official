import { useState } from "react";
import { toast } from "react-hot-toast";

export default function Register() {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [confirmPassword, setConfirmPassword] = useState<string>("");

  const registerUser = (e: React.FormEvent) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      toast.error("密碼不相符");
      return;
    }
  };

  return (
    <div className="w-[90%] sm:w-[80%] md:w-[70%] max-w-[500px] max-h-[700px] h-[60%] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-xl bg-zinc-600 flex flex-col px-10 shadow-xl shadow-gray-500 justify-center">
      <h1 className="text-3xl text-center text-white">會員註冊</h1>
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
        <div className="flex flex-col mt-5">
          <input
            placeholder="密碼"
            type="password"
            className="py-2 px-4 rounded-xl"
            onChange={(ev) => setPassword(ev.target.value)}
            required
          />
        </div>
        <div className="flex flex-col mt-5">
          <input
            placeholder="確認密碼"
            type="password"
            className="py-2 px-4 rounded-xl"
            onChange={(ev) => setConfirmPassword(ev.target.value)}
            required
          />
        </div>
        <button className="p-2 mt-16 w-full text-white rounded-xl duration-300 bg-zinc-900 hover:bg-zinc-500 active:bg-zinc-400">
          註冊
        </button>
      </form>
    </div>
  );
}
