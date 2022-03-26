import logo from './logo.svg';
import './App.css';
import { Button, ChakraProvider } from '@chakra-ui/react';

function App() {
return (
      <ChakraProvider>
        <Button>テスト</Button>
      </ChakraProvider>
    )
}

export default App;
