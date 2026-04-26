import ReactDOM from 'react-dom/client';
import App from './App';
import './assets/index.css';
import { ChakraProvider } from '@chakra-ui/react';
import { theme } from './theme';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <ChakraProvider theme={theme}>
    <App />
  </ChakraProvider>
);
