/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js,templ,go,gohtml}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}

