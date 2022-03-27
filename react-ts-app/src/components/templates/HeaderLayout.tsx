import { SearchIcon } from "@chakra-ui/icons"
import { Heading, IconButton, Input, InputGroup, InputRightAddon } from "@chakra-ui/react"
import { memo, ReactNode, VFC } from "react"

type Props = {
      children: ReactNode;
};

export const HeaderLayout: VFC<Props> = memo((props) => {
      const { children } = props;
      return (
            <>
                  <Heading as='h1' size='lg' mt={5} pl={8} pb={5}>TV Comment App</Heading>
                  <InputGroup px={8} paddingBottom={5} size='lg' >
                        <Input bg="white" placeholder='検索' />
                        <InputRightAddon p={0} children={<IconButton aria-label='Search database' icon={<SearchIcon />} />} />
                  </InputGroup>
                  {children}
            </>
      );
});