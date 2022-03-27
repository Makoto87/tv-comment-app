import { SearchIcon } from '@chakra-ui/icons'
import { Box, Button, Center, Flex, Heading, HStack, IconButton, Input, InputGroup, InputRightAddon, Text, VStack, Wrap, WrapItem } from '@chakra-ui/react'

export const Home = () => {
      return (
            <>
                  <Heading as='h1' size='lg' mt={5} pl={8} pb={5}>TV Comment App</Heading>
                  <InputGroup px={8} paddingBottom={5} size='lg' >
                        <Input bg="white" placeholder='検索' />
                        <InputRightAddon p={0} children={<IconButton aria-label='Search database' icon={<SearchIcon />} />} />
                  </InputGroup>

                  <Flex paddingX={8} w='100%'>
                        {/* 新規投稿ボタンとカテゴリーリスト */}
                        <VStack w='15%' spacing={5} paddingTop={10}>
                              <Button colorScheme='blue' size='lg' w='100%'>
                                    新規投稿
                              </Button>
                              <Box w='100%'>
                                    <Heading as='h3' size='md'>カテゴリ</Heading>
                                    <VStack spacing={3}>
                                          <Button marginTop={3} colorScheme='blackAlpha' size='lg' w='100%'>
                                                バラエティ
                                          </Button>
                                          <Button marginTop={3} colorScheme='blackAlpha' size='lg' w='100%'>
                                                ドラマ
                                          </Button>
                                    </VStack>
                              </Box>
                        </VStack>

                        <VStack w='85%'>
                              <Text w='100%' paddingX={7} fontSize='lg'>ホーム</Text>
                              <Wrap w='95%' py={{ base: 8}} px={{ base: 10}} justify='space-around' >
                                    <WrapItem mx='auto' p={3}>
                                          <Button colorScheme='green' w='300px' h='150px' fontSize='2xl'>
                                                バラエティ
                                          </Button>
                                    </WrapItem>
                              </Wrap>
                        </VStack>
                  </Flex>
            </>
      )
}