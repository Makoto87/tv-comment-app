import { SearchIcon } from "@chakra-ui/icons"
import { Heading, IconButton, Input, InputGroup, InputRightAddon } from "@chakra-ui/react"
import { memo, ReactNode, VFC, Dispatch, SetStateAction, createContext, useState, ChangeEvent, KeyboardEvent } from "react"
import { Link, useNavigate } from "react-router-dom";

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
      const navigate = useNavigate();

      const onChangeSearchBar = (e: ChangeEvent<HTMLInputElement>) => {
            setsearchText(() => e.target.value)
      }
      const onPressEnter = (e: KeyboardEvent) => {
            if ( e.key === "Enter" ) {
                  setProgramSubstr("%" + searchText + "%")
                  navigate("/")
            }
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
                        <Input onChange={onChangeSearchBar} onKeyDown={onPressEnter} bg="white" placeholder='番組名を検索' />
                        <Link to="/" onClick={onClickSearchBar}>
                              <InputRightAddon p={0}>
                                    <IconButton aria-label='Search database' icon={<SearchIcon />} />
                              </InputRightAddon>
                        </Link>
                  </InputGroup>
                  {children}
            </ProgramContext.Provider>
      );
});