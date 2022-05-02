import { SearchIcon } from '@chakra-ui/icons'
import { Box, Button, Center, Flex, Heading, HStack, IconButton, Input, InputGroup, InputRightAddon, Text, VStack, Wrap, WrapItem } from '@chakra-ui/react'

import { SubList } from '../parts/SubList'
import { PostButton } from '../parts/PostButton'
import { ProgramList } from '../parts/ProgramList'
import { HeaderLayout } from '../templates/HeaderLayout'
import { memo } from 'react'

export const Home = memo(() => {
      // テストデータ
      return (
            // ヘッダーを取得
            <HeaderLayout>
                  <Flex paddingX={8} w='100%' flexWrap={ {base: 'wrap', md: 'nowrap'}}>
                        {/* 新規投稿ボタンとカテゴリーリスト */}
                        <VStack w={{ base: '100%', md: '15%' }} spacing={{ base: '4', md: '5'}} pt={{ base: '3', md: '10'}} pb={ {base: '6', md: '0'}}>
                              <PostButton />
                              <SubList />
                        </VStack>

                        {/* 番組一覧を表示 */}
                        <VStack w={{ base: '100%', md: '85%'}}>
                              <Text w='100%' paddingX={{ base: '1', md: '7'}} fontSize='lg'>ホーム</Text>
                              <ProgramList title={""}/>
                              {/* <ProgramList buttonType='programs' titles={programTitles} /> */}
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
});