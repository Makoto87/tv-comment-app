import { SearchIcon } from "@chakra-ui/icons"
import { Heading, IconButton, Input, InputGroup, InputRightAddon } from "@chakra-ui/react"
import { memo, ReactNode, VFC, Dispatch, SetStateAction, createContext, useState } from "react"
import { Link } from "react-router-dom";

type Props = {
      children: ReactNode;
      setText?: Dispatch<SetStateAction<String>>;
};

export const ProgramContext = createContext({} as {
      programSubstr: string
      setProgramSubstr: Dispatch<SetStateAction<string>>
})

export const HeaderLayout: VFC<Props> = memo((props) => {
      const { children } = props;
      const [programSubstr, setProgramSubstr] = useState("");
      return (
            <ProgramContext.Provider value={{programSubstr, setProgramSubstr}}>
                  <Link to="/" onClick={() => {setProgramSubstr("")}}>
                        <Heading as='h1' size='lg' mt={5} pl={8} pb={5}>TV Comment App</Heading>
                  </Link>
                  <InputGroup px={8} paddingBottom={5} size='lg' >
                        <Input bg="white" placeholder='検索' />
                        <InputRightAddon p={0} children={<IconButton aria-label='Search database' icon={<SearchIcon />} />} />
                  </InputGroup>
                  {children}
            </ProgramContext.Provider>
      );
});