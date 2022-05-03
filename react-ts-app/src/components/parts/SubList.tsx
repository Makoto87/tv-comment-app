import { Box, Heading, VStack, Button, Stack, HStack, Wrap, WrapItem } from "@chakra-ui/react"
import { memo } from "react"

export const SubList = memo(() => {
      const hiraganaArray: String[] = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん".split('');
      return (
            <Box w='100%' >
                  <Heading as='h3' size='md' pb={5}>50音検索</Heading>
                  <Wrap w='100%' justify='space-around'>
                        {hiraganaArray.map((char) => (
                              <WrapItem w='25%' m={0}>
                                    <Button colorScheme='blackAlpha' size='lg'>
                                          {char}
                                    </Button>
                              </WrapItem>
                        ))}
                  </Wrap>
            </Box>
      )
});