/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./template/*.html"],
  theme: {
    extend: {
      colors: {
        'murky-blue': '#a7b3c2',
        'tn-white': {
          100: '#d5d6db',
          200: '#9699a3'
        }
      }
    },
  },
  plugins: [],
}

