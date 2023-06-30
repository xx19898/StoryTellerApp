
module.exports = {
  content: ["./src/**/*.{html,js,tsx,ts}"],
  theme: {
    screens: {
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1280px',
      '2xl': '1536px',
    },
    colors:{
      primary: "#D81E5B",
      secondary: "#3F3F37",
      secondPrimary:"#EB5E55",
      base:"#FFEBE7",
      special:"#2A1E5C",
      secSpecial:"#3B429F",
      white: "#FFFFFF",
      black: "#000000"

    },
    extend: {
      fontFamily:{
        'belanosima': ['Belanosima', 'sans-serif']
      },
      boxShadow: {
        '3xl': '3px 3px 6px 0px rgba(0,0,0,0.47)',
        '4xl': [
            '0 35px 35px rgba(0, 0, 0, 0.25)',
            '0 45px 65px rgba(0, 0, 0, 0.15)'
        ]
      },
    },
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          "primary": "#D81E5B",
          "secondary": "#3F3F37",
          "accent": "#EB5E55",
          "neutral": "#FFEBE7",
          "base-100": "#ebe4f1",
          "info": "#3a9dcb",
          "success": "#19a95f",
          "warning": "#fac533",
          "error": "#dd4072",
        },
      },
    ],
  },
  plugins: [require("daisyui")],
}