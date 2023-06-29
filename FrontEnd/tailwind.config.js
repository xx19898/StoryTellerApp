
module.exports = {
  content: ["./src/**/*.{html,js,tsx,ts}"],
  theme: {
    colors:{
      primary: "#D81E5B",
      secondary: "#3F3F37",
      secondPrimary:"#EB5E55",
      base:"#FFEBE7",
      special:"#2A1E5C",
      secSpecial:"#3B429F"
    },
    extend: {
      fontFamily:{
        'belanosima': ['Belanosima', 'sans-serif']
      }
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