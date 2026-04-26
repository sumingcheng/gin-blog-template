import { extendTheme } from '@chakra-ui/react';

const theme = extendTheme({
  fonts: {
    heading: `-apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans SC', sans-serif`,
    body: `-apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans SC', sans-serif`,
  },
  styles: {
    global: {
      body: {
        bg: '#fff',
        color: '#111',
      },
    },
  },
});

export { theme };
