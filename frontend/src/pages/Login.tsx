import { useMutation, useQueryClient } from '@tanstack/react-query'
import { AxiosError, HttpStatusCode } from 'axios'
import { useCallback, useState } from 'react'
import { toast } from 'react-hot-toast'
import { useNavigate } from 'react-router-dom'
import { loginUser } from '../api/user'
import Loading from '../components/Loading'
import FormLayout from '../layout/Form'

interface LoginParams {
  email: string
  password: string
}

const useLoginMutation = () => {
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  const mutation = useMutation({
    mutationKey: ['login'],
    mutationFn: async ({ email, password }: LoginParams) =>
      loginUser(email, password),
    onSuccess: res => {
      toast.success('登入成功')
      queryClient.setQueryData(['user'], () => ({ data: res.data.user }))
      sessionStorage.setItem('access_token', res.data.access_token)
      sessionStorage.setItem('refresh_token', res.data.refresh_token)
      sessionStorage.setItem('session_id', res.data.session_id)
      navigate('/')
    },
    onError: error => {
      const err = error as AxiosError
      if (err.response?.status === HttpStatusCode.Unauthorized) {
        toast.error('電子信箱或密碼錯誤')
      } else {
        toast.error('登入失敗，請再試一次')
      }
      console.error(error)
    }
  })

  return mutation
}

export default function Login() {
  const navigate = useNavigate()
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [rememberMe, setRememberMe] = useState<boolean>(false)
  const mutation = useLoginMutation()

  // on submit login
  const onSubmit = useCallback(
    async (e: React.FormEvent) => {
      e.preventDefault()

      // ensure email and password are valid
      if (!email || !password) {
        return
      } else if (password.length < 8) {
        toast.error('密碼長度不足')
      }

      // call login mutation
      mutation.mutate({ email, password })
    },
    [email, password, mutation]
  )

  return (
    <FormLayout title="會員登入">
      {mutation.isPending && (
        <div className="flex justify-center items-center mx-auto mt-20 w-32 h-32">
          <Loading />
        </div>
      )}
      <form
        onSubmit={onSubmit}
        className={`${mutation.isPending ? 'hidden' : ''}`}
      >
        <div className="flex flex-col mt-5">
          <input
            placeholder="電子信箱"
            type="email"
            value={email}
            onChange={ev => setEmail(ev.target.value)}
            className="py-2 px-4 rounded-xl"
            autoComplete="email"
          />
        </div>
        <div className="flex flex-col mt-5">
          <input
            placeholder="密碼"
            value={password}
            onChange={ev => setPassword(ev.target.value)}
            type="password"
            className="py-2 px-4 rounded-xl"
            autoComplete="current-password"
          />
        </div>
        <div className="inline-flex justify-start items-center mt-2 w-full">
          <input
            type="checkbox"
            className="m-2 cursor-pointer"
            onClick={() => setRememberMe(!rememberMe)}
          />
          <label className="text-white">記住我</label>
          <a
            href="/password-reset"
            className="ml-auto text-white cursor-pointer"
          >
            <label className="duration-300 cursor-pointer hover:text-zinc-400">
              忘記密碼?
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
        <button
          onClick={e => {
            e.preventDefault()
            navigate('/register')
          }}
          className="p-2 w-full text-white rounded-xl duration-300 bg-zinc-900 hover:bg-zinc-500 active:bg-zinc-400"
        >
          註冊
        </button>
      </form>
    </FormLayout>
  )
}
