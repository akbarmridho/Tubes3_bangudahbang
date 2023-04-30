/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      primary: {
        base: '#2977EF',
        dark: '#1363DF',
      },
      secondary: {
        base: '#111825',
        dark: '#010101',
      },
    },
    extend: {
      colors: {
        white: '#FBFEFF',
      },
    },
  },
  plugins: [require("daisyui")],
}