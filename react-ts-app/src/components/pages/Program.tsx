import { ChevronRightIcon } from "@chakra-ui/icons"
import { Flex, VStack, Text, Breadcrumb, BreadcrumbItem, BreadcrumbLink } from "@chakra-ui/react"
import { memo } from "react"
import { Link } from "react-router-dom"

import { PostButton } from "../parts/PostButton"
import { ProgramList } from "../parts/ProgramList"
import { RangeTime } from "../parts/RangeTime"
import { HeaderLayout } from "../templates/HeaderLayout"

export const Program = memo(() => {
      // テストデータ
      const programNumbers: string[] = [...Array(50).fill("123")]
      return (
            // ヘッダーを取得
            <HeaderLayout>
                  <Flex paddingX={8} w='100%' flexWrap={ {base: 'wrap', md: 'nowrap'}}>
                        {/* 新規投稿ボタンと時間指定 */}
                        <VStack w={{ base: '100%', md: '15%' }} spacing={{ base: '4', md: '5'}} pt={{ base: '3', md: '10'}} pb={ {base: '6', md: '0'}}>
                              <PostButton />
                              <RangeTime />
                        </VStack>
                        
                        <VStack w={{ base: '100%', md: '85%'}}>
                              {/* パンくずリスト */}
                              <Breadcrumb w='100%' paddingX={{ base: '1', md: '7'}} fontSize='lg' spacing='8px' separator={<ChevronRightIcon color='gray.500' />}>
                                    <BreadcrumbItem>
                                          <BreadcrumbLink as = { Link } to='/'>ホーム</BreadcrumbLink>
                                    </BreadcrumbItem>

                                    <BreadcrumbItem>
                                          <BreadcrumbLink isCurrentPage>放送回</BreadcrumbLink>
                                    </BreadcrumbItem>
                              </Breadcrumb>

                              {/* 放送回を表示 */}
                              <ProgramList buttonType="comments" titles={programNumbers}/>
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
});