import logo from './logo.svg';
import './App.css';
import { Button, ChakraProvider } from '@chakra-ui/react';
import theme from './theme/theme';
import { Home } from './components/pages/Home';
import { Program } from './components/pages/Program';
import { Comments } from './components/pages/Comments';
import { BrowserRouter } from 'react-router-dom';
import { Router } from './router/Router';

function App() {
return (
      <ChakraProvider theme={theme}>
            <BrowserRouter>
                  <Router />
            </BrowserRouter>
      </ChakraProvider>
    )
}

export default App;
