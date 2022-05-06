import { Box, Heading, Button, Wrap, WrapItem } from "@chakra-ui/react"
import { memo, useContext } from "react"
import { ProgramContext } from '../templates/HeaderLayout'

export const SubList = memo(() => {
      const { setProgramSubstr } = useContext(ProgramContext);
      const hiraganaArray: String[] = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん".split('');
      const changeText = (char: string) => {
            setProgramSubstr(char);
      }
      return (
            <Box w='100%' >
                  <Heading as='h3' size='md' pb={5}>50音検索</Heading>
                  <Wrap w='100%' justify='space-around'>
                        {hiraganaArray.map((char) => (
                              <WrapItem w='25%' m={0}>
                                    <Button onClick={() => changeText(char + "%")} colorScheme='blackAlpha' size='lg'>
                                          {char}
                                    </Button>
                              </WrapItem>
                        ))}
                  </Wrap>
            </Box>
      )
});