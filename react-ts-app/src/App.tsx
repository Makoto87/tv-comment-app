import './App.css';
import { ChakraProvider } from '@chakra-ui/react';
import theme from './theme/theme';
import { Program } from './components/pages/Program';
import { BrowserRouter } from 'react-router-dom';
import { Router } from './router/Router';

import { useQuery, gql } from "@apollo/client";

const FETCH_BOOKS = gql`
      query {
            programs(search: "test%") {
                  id
                  name
            }
      }
`;

interface Program {
      id: number;
      name: string;
}

interface ProgramsData {
      programs: Program[];
}

function App() {
      const { data } = useQuery<ProgramsData>(FETCH_BOOKS);
      console.log(data)
      data && data.programs.map(program => console.log(program.name))
      
      return (
            <ChakraProvider theme={theme}>
                  <BrowserRouter>
                        <Router />
                  </BrowserRouter>
            </ChakraProvider>
      )
}

export default App;
