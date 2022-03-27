import { Flex, VStack, Text } from "@chakra-ui/react"

import { CategoryList } from "../parts/CategoryList"
import { CommentList } from "../parts/CommentList"
import { PostButton } from "../parts/PostButton"
import { HeaderLayout } from "../templates/HeaderLayout"

export const Comments = () => {
      return (
            // ヘッダーを取得
            <HeaderLayout>
                  <Flex paddingX={8} w='100%' flexWrap={ {base: 'wrap', md: 'nowrap'}}>
                        {/* 新規投稿ボタンと放送回リスト */}
                        <VStack w={{ base: '100%', md: '15%' }} spacing={{ base: '4', md: '5'}} pt={{ base: '3', md: '10'}} pb={ {base: '6', md: '0'}}>
                              <PostButton />
                              <CategoryList />
                        </VStack>

                        {/* 番組一覧を表示 */}
                        <VStack w={{ base: '100%', md: '85%'}}>
                              <Text w='100%' paddingX={{ base: '1', md: '7'}} fontSize='lg'>ホーム</Text>
                              <CommentList />
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
}