import { ChevronRightIcon } from "@chakra-ui/icons"
import { Flex, VStack, Text, Breadcrumb, BreadcrumbItem, BreadcrumbLink } from "@chakra-ui/react"
import { memo, useState } from "react"
import { Link, useLocation } from "react-router-dom"

import { PostButton } from "../parts/PostButton"
import { EpisodeList } from "../parts/EpisodeList"
import { RangeTime } from "../parts/RangeTime"
import { HeaderLayout } from "../templates/HeaderLayout"

export const Episodes = memo(() => {
      const location = useLocation()
      const [selectId, setSelectId] = useState<{ id: number }>(location.state as { id: number })
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
                              <EpisodeList titleID={selectId.id}/>
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
});