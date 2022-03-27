import logo from './logo.svg';
import './App.css';
import { Button, ChakraProvider } from '@chakra-ui/react';
import theme from './theme/theme';
import { Home } from './components/pages/Home';
import { Program } from './components/pages/Program';

function App() {
return (
      <ChakraProvider theme={theme}>
        <Program />
      </ChakraProvider>
    )
}

export default App;
