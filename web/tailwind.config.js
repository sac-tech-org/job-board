/** @type {import('tailwindcss').Config} */

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
  safelist: [
    // for the Button component
    'bg-green-500',
    'hover:bg-green-500/90',
    'text-zinc-50',
  ]
};
