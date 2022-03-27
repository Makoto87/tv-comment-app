import { Wrap, WrapItem, Button, VStack, Box, Text, Flex } from "@chakra-ui/react"

export const CommentList = () => {
      return (
            <VStack w='95%' spacing={5} py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  <Box bg='white' border='3px' borderWidth={5} borderColor='red.500' shadow='sm' rounded='md' w='100%'  fontSize='xl' p={5} pb={3}>
                        <Text>この番組面白すぎ！ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ</Text>
                        <Flex pt={2} fontSize='md' direction={{ base: 'column', md: 'row' }} justifyContent='space-between' alignItems={{md: 'center'}} textColor='red.300'>
                              <Text flexGrow='2' pb={{base: '1', md: '0'}}>ゲストユーザー</Text>
                              <Text flexGrow='4' pb={{base: '3', md: '0'}}>2020/12/31</Text>
                              <Flex alignItems='center' pb={{base: '2', md: '0'}}>
                                    <Button flexGrow='1' colorScheme='red.200' variant='outline' size='sm'>Good</Button>
                                    <Text flexGrow='10' pl='3'>100</Text>
                              </Flex>
                              
                        </Flex>
                  </Box>
            </VStack>
      )
}