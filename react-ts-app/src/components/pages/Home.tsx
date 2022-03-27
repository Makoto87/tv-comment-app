import { SearchIcon } from '@chakra-ui/icons'
import { Box, Button, Container, Flex, Heading, HStack, IconButton, Input, InputGroup, InputLeftAddon, InputRightAddon, VStack } from '@chakra-ui/react'

export const Home = () => {
      return (
            <>
                  <InputGroup size='lg' p={8}>
                        <Input bg="white" placeholder='検索' />
                        <InputRightAddon children={<SearchIcon />} />
                  </InputGroup>

                  <Flex paddingX={8}>
                        {/* 新規投稿ボタンとカテゴリーリスト */}
                        <VStack w='250px' spacing={5} paddingTop={3}>
                              <Button colorScheme='teal' size='lg' w='100%'>
                                    新規投稿
                              </Button>
                              <Box w='100%'>
                                    <Heading as='h3' size='md'>カテゴリ</Heading>
                                    <VStack spacing={3}>
                                          <Button marginTop={3} colorScheme='green' size='lg' w='100%'>
                                                バラエティ
                                          </Button>
                                          <Button marginTop={3} colorScheme='green' size='lg' w='100%'>
                                                ドラマ
                                          </Button>
                                    </VStack>
                              </Box>
                        </VStack>

                        <div>
                              <h3>ホーム</h3>
                              <div>
                                    <Button colorScheme='teal' size='md'>
                                          バラエティ
                                    </Button>
                              </div>
                        </div>
                  </Flex>
            </>
      )
}