
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
      black: "#000000",
      darkerPrimary:"#AC1848",
      darkestPrimary:"#6C0F2D",
      darkerSecondary: "#32322C",
      darkestSecondary: "#1F1F1B",
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
          "darker-primary":"#AC1848",
          "darkest-primary":"#6C0F2D",
          "secondary": "#3F3F37",
          "darker-secondary": "#32322C",
          "darkest-secondary": "#1F1F1B",
          "accent": "#EB5E55",
          "neutral": "#FFFFFF",
          "base-100": "#ebe4f1",
          "info": "#3a9dcb",
          "success": "#19a95f",
          "warning": "#fac533",
          "error": "#dd4072",
          "white": "#FFFFFF",
          "black": "#000000"
        },
      },
    ],
  },
  plugins: [require("daisyui")],
}