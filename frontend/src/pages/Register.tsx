import { HttpStatusCode } from 'axios'
import { useCallback, useState } from 'react'
import { toast } from 'react-hot-toast'
import { useNavigate } from 'react-router-dom'
import { registerUser } from '../api/user'
import Loading from '../components/Loading'
import FormLayout from '../layout/Form'

export default function Register() {
  const navigate = useNavigate()
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [confirmPassword, setConfirmPassword] = useState<string>('')
  const [isLoading, setIsLoading] = useState<boolean>(false)

  const isInvalidPassword = (condition: boolean, message: string) => {
    if (condition) {
      toast.error(message)
      setPassword('')
      setConfirmPassword('')
      setIsLoading(false)
      return true
    }
    return false
  }

  const onSubmit = useCallback(
    async (e: React.FormEvent) => {
      e.preventDefault()
      setIsLoading(true)

      // check if password and confirm password match
      const chk1 = isInvalidPassword(password !== confirmPassword, '密碼不相符')
      const chk2 = isInvalidPassword(password.length < 8, '密碼長度不足')
      if (chk1 || chk2) return

      // create user and show toast message
      registerUser(email, password)
        .then(res => {
          switch (res.status) {
            case HttpStatusCode.Ok:
              toast.success('註冊成功')
              break
            default:
              toast.error('註冊失敗，請再試一次')
          }
          // succeeded, redirect to login page
          setIsLoading(false)
          navigate('/login')
        })
        .catch(err => {
          err.response?.status === HttpStatusCode.Conflict
            ? toast.error('此電子信箱已被註冊')
            : toast.error('註冊失敗，請再試一次')
          setIsLoading(false)
        })
    },
    [email, password, confirmPassword]
  )

  return (
    <FormLayout title="會員註冊">
      {isLoading && (
        <div className="flex justify-center items-center mx-auto mt-20 w-32 h-32">
          <Loading />
        </div>
      )}
      <form
        className={`${isLoading ? 'hidden' : ''}`}
        onSubmit={e => onSubmit(e)}
      >
        <div className="flex flex-col mt-5">
          <input
            placeholder="電子信箱"
            type="email"
            value={email}
            className="py-2 px-4 rounded-xl"
            onChange={ev => setEmail(ev.target.value)}
            autoComplete="email"
            required
          />
        </div>
        <div className="flex flex-col mt-5">
          <input
            placeholder="密碼"
            type="password"
            className="py-2 px-4 rounded-xl"
            value={password}
            onChange={ev => setPassword(ev.target.value)}
            autoComplete="new-password"
            required
          />
        </div>
        <div className="flex flex-col mt-5">
          <input
            placeholder="確認密碼"
            type="password"
            className="py-2 px-4 rounded-xl"
            value={confirmPassword}
            onChange={ev => setConfirmPassword(ev.target.value)}
            autoComplete="new-password"
            required
          />
        </div>
        <button className="p-2 mt-16 w-full text-white rounded-xl duration-300 bg-zinc-900 hover:bg-zinc-500 active:bg-zinc-400">
          註冊
        </button>
      </form>
    </FormLayout>
  )
}
