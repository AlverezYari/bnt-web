/** @type {import('tailwindcss').Config} */
module.exports = {
  content:  [
    './ui/**/*.html',
    './ui/**/*.tmpl',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'hero-pattern': "url('/static/img/bkg.png')",
      },
      colors: {
        skyCream: '#FEF9E3',
        mainText: '021534',
      },

  },
  plugins: [],
},
}

