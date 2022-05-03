import { Wrap, WrapItem, Button } from "@chakra-ui/react"
import { memo, ReactNode, VFC } from "react";
import { useNavigate } from "react-router-dom";

import { useQuery, gql } from "@apollo/client";

const FETCH_PROGRAMS = gql`
      query getPrograms($search: String!) {
            programs(search: $search) {
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

type Props = {
      title: string
}

export const ProgramList: VFC<Props> = memo((props) => {
      const { title } = props;
      const navigate = useNavigate();
      console.log(title)

      const { loading, error, data } = useQuery<ProgramsData>(FETCH_PROGRAMS, 
            {variables: {
                  search: title
            }}
      );
      console.log(data, loading, error);

      if (loading) return (
            <h1>Loading Now</h1>
      );

      if (error)   return (
            <h1>Server Error</h1>
      );

      return (
            <Wrap w='95%' py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  { data!.programs.map((program) => (
                        <WrapItem mx='auto' p={3}>
                              <Button key={program.id} onClick={() => navigate("episodes", { state: { id: program.id }})} whiteSpace='unset' colorScheme='green' w={{ base: '200px', md: '300px' }} h={{ base: '100px', md: '150px'}} fontSize={{ md: '2xl'}}>
                                    {program.name}
                              </Button>
                        </WrapItem>
                  ))}
            </Wrap>
      )
});