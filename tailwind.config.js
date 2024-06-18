const colors = require('tailwindcss/colors')

module.exports = {
  content: ["./**/*.{html,templ,js}"],
  theme: {
    fontFamily: {
      'sans': ['Fira Sans', 'sans-serif'],
    },
    colors: {
      dark: {
        1: colors.neutral[700],
        2: colors.neutral[800],
        3: colors.neutral[900],
      }
    },
    extend: {},
  },
  plugins: [],
}
