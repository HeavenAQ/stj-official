import { useState } from 'react'
import { toast } from 'react-hot-toast'
import FormLayout from '../layout/Form'

export default function PasswordReset() {
  const [email, setEmail] = useState<string>('')

  const registerUser = (e: React.FormEvent) => {
    e.preventDefault()
  }

  return (
    <FormLayout title="密碼重置">
      <form onSubmit={e => registerUser(e)}>
        <div className="flex flex-col mt-5">
          <input
            placeholder="電子信箱"
            type="email"
            className="py-2 px-4 rounded-xl"
            onChange={ev => setEmail(ev.target.value)}
            required
          />
        </div>
        <button className="p-2 mt-5 w-full text-white rounded-xl duration-300 bg-zinc-900 hover:bg-zinc-500 active:bg-zinc-400">
          密碼重置
        </button>
      </form>
    </FormLayout>
  )
}
