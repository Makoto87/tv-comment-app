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
      const [searchText, setsearchText] = useState("");

      const onChangeSearchBar = (e: React.ChangeEvent<HTMLInputElement>) => {
            setsearchText(() => e.target.value)
      }
      const onClickSearchBar = () => {
            setProgramSubstr("%" + searchText + "%")
      }
      return (
            <ProgramContext.Provider value={{programSubstr, setProgramSubstr}}>
                  <Link to="/" onClick={() => {setProgramSubstr("")}}>
                        <Heading as='h1' size='lg' mt={5} pl={8} pb={5}>TV Comment App</Heading>
                  </Link>
                  <InputGroup px={8} paddingBottom={5} size='lg' >
                        <Input onChange={onChangeSearchBar} bg="white" placeholder='番組名を検索' />
                        <InputRightAddon onClick={onClickSearchBar} p={0}>
                              <Link to="/">
                                    <IconButton aria-label='Search database' icon={<SearchIcon />} />
                              </Link>
                        </InputRightAddon>
                  </InputGroup>
                  {children}
            </ProgramContext.Provider>
      );
});