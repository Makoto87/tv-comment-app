import { Flex, VStack, Text } from "@chakra-ui/react"

import { PostButton } from "../parts/PostButton"
import { ProgramList } from "../parts/ProgramList"
import { RangeTime } from "../parts/RangeTime"
import { HeaderLayout } from "../templates/HeaderLayout"

export const Program = () => {
      return (
            // ヘッダーを取得
            <HeaderLayout>
                  <Flex paddingX={8} w='100%' flexWrap={ {base: 'wrap', md: 'nowrap'}}>
                        {/* 新規投稿ボタンと時間指定 */}
                        <VStack w={{ base: '100%', md: '15%' }} spacing={{ base: '4', md: '5'}} pt={{ base: '3', md: '10'}} pb={ {base: '6', md: '0'}}>
                              <PostButton />
                              <RangeTime />
                        </VStack>

                        {/* 放送回を表示 */}
                        <VStack w={{ base: '100%', md: '85%'}}>
                              <Text w='100%' paddingX={{ base: '1', md: '7'}} fontSize='lg'>ホーム  &gt;  放送回一覧</Text>
                              <ProgramList />
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
}