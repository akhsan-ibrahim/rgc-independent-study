//  File to enable tailwindcss intellisense
const colors = require('tailwindcss/colors')

module.exports = {
    content: ['index.html'],
  theme: {
    fontFamily: {
        'display': 'Poppins',
    },
    container:{
        center: true,
        padding: '16px',
    },
    extend: {
        colors: {
            transparent: 'transparent',
            current: 'currentColor',
            fuchsia: colors.fuchsia,
            slate: colors.slate,      
        },
        screen: {
            '2xl': '1320px',
        },
    },    
  },
}