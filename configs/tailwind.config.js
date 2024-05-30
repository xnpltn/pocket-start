/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './internal/templates/**/*.templ',
    './internal/handler/**/*.go'
  ],
  theme: {
    extend: {
      fontFamily: {
        'inter': ['Inter', 'sans-serif'],
      },
      colors: {
        'light-gray': '#D8DEE9',
        'dark-gray': '#2E3440',
        'light-b': '#ECEFF4'
      }
    }
  },
};

