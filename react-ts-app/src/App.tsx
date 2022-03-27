import logo from './logo.svg';
import './App.css';
import { Button, ChakraProvider } from '@chakra-ui/react';
import theme from './theme/theme';
import { Home } from './components/pages/Home';
import { Program } from './components/pages/Program';
import { Comments } from './components/pages/Comments';

function App() {
return (
      <ChakraProvider theme={theme}>
            <Comments></Comments>
      </ChakraProvider>
    )
}

export default App;
