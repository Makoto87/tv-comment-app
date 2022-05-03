import { Box, Heading, VStack, Button, Stack, HStack, Wrap, WrapItem } from "@chakra-ui/react"
import { memo, VFC, Dispatch, SetStateAction } from "react"

type Props = {
      setText: Dispatch<SetStateAction<String>>;
}

export const SubList: VFC<Props> = memo((props) => {
      const {setText} = props
      const hiraganaArray: String[] = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん".split('');
      const changeText = (char: String) => {
            setText(char);
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