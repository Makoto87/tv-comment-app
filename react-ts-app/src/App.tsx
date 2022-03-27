import logo from './logo.svg';
import './App.css';
import { Button, ChakraProvider } from '@chakra-ui/react';
import theme from './theme/theme';
import { Home } from './components/pages/Home';

function App() {
return (
      <ChakraProvider theme={theme}>
        <Home />
      </ChakraProvider>
    )
}

export default App;
