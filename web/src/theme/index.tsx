import { extendTheme } from '@chakra-ui/react'

const colors = {
  brand: {
    900: '#1a365d',
    800: '#153e75',
    700: '#2a69ac',
  },
}

const shadows = {
  xs: '0 0 0 1px rgba(0, 0, 0, 0.15)'
};


export const theme = extendTheme({ colors, shadows })
