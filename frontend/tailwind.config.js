/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',
    'node_modules/flowbite-react/lib/esm/**/*.js'
  ],
  theme: {
    extend: {
      keyframes: {
        'skeleton-loading': {
          '0%': { 'background-color': 'hsl(200, 20%, 70%)' },
          '100%': { 'background-color': 'hsl(200, 20%, 95%)' }
        }
      },
      animation: {
        skeleton: 'skeleton-loading 1s linear infinite alternative'
      }
    }
  },
  plugins: [
    require('flowbite/plugin')({
      forms: false
    }),
    require('tailwindcss-animated')
  ]
}
