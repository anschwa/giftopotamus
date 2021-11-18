module.exports = {
  purge: [
    './templates/**/*.html.tmpl',
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {
      cursor: ['disabled'],
      opacity: ['disabled'],
      backgroundColor: ['odd', 'even'],
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
};
