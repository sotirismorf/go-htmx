const colors = require('tailwindcss/colors')

module.exports = {
  content: ["./**/*.{html,templ,js}"],
  theme: {
    fontFamily: {
      'sans': ['Fira Sans', 'sans-serif'],
    },
    extend: {
      boxShadow: {
        'solid': '6px 6px rgba(0, 0, 0, 0.3)',
      },
      colors: {
        dark: {
          1: colors.neutral[700],
          2: colors.neutral[800],
          3: colors.neutral[900],
        },
        light: {
          1: '#FBF1C7'
        }
      },
    },
  },
  plugins: [],
}
