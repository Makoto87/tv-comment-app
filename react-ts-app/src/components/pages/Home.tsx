import { SearchIcon } from '@chakra-ui/icons'
import { Button, Container, HStack, IconButton, Input, InputGroup, InputLeftAddon, InputRightAddon, VStack } from '@chakra-ui/react'

export const Home = () => {
      return (
            <>
                  <InputGroup size='lg' p={8}>
                        <Input bg="white" placeholder='検索' />
                        <InputRightAddon children={<SearchIcon />} />
                  </InputGroup>

                  <div>
                        <Button colorScheme='teal' size='md'>
                              新規投稿
                        </Button>
                        <div>
                              <h3>カテゴリ</h3>
                              <Button colorScheme='teal' size='md'>
                                    バラエティ
                              </Button>
                        </div>
                  </div>

                  <div>
                        <h3>ホーム</h3>
                        <div>
                              <Button colorScheme='teal' size='md'>
                                    バラエティ
                              </Button>
                        </div>
                  </div>
            </>
      )
}