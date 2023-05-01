/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      'primary-base': '#2977EF',
      'primary-dark': '#1363DF',
      'secondary-light': '#292F3B',
      'secondary-base': '#111825',
      'secondary-dark': '#010101',
    },
    extend: {
      colors: {
        white: '#FBFEFF',
      },
    },
  },
  plugins: [require("daisyui")],
}