import { ChevronRightIcon } from "@chakra-ui/icons"
import { Flex, VStack, Text, Button, HStack, Breadcrumb, BreadcrumbItem, BreadcrumbLink } from "@chakra-ui/react"
import { Link, useNavigate, useLocation } from "react-router-dom"

import { SubList } from "../parts/SubList"
import { CommentList } from "../parts/CommentList"
import { PostButton } from "../parts/PostButton"
import { HeaderLayout } from "../templates/HeaderLayout"
import { memo, useState } from "react"

export const Comments = memo(() => {
      const navigate = useNavigate();
      const location = useLocation()
      const [selectId, setSelectId] = useState<{ episodeID: number }>(location.state as { episodeID: number })
      console.log(selectId.episodeID)
      return (
            // ヘッダーを取得
            <HeaderLayout>
                  <Flex paddingX={8} w='100%' flexWrap={ {base: 'wrap', md: 'nowrap'}}>
                        {/* 新規投稿ボタンと放送回リスト */}
                        <VStack w={{ base: '100%', md: '15%' }} spacing={{ base: '4', md: '5'}} pt={{ base: '3', md: '10'}} pb={ {base: '6', md: '0'}}>
                              <PostButton />
                              <SubList />
                        </VStack>

                        <VStack w={{ base: '100%', md: '85%'}}>
                              {/* パンくずリスト */}
                              <Breadcrumb w='100%' paddingX={{ base: '1', md: '7'}} fontSize='lg' spacing='8px' separator={<ChevronRightIcon color='gray.500' />}>
                                    <BreadcrumbItem>
                                          <BreadcrumbLink as = { Link } to='/'>ホーム</BreadcrumbLink>
                                    </BreadcrumbItem>
                                    <BreadcrumbItem>
                                          <BreadcrumbLink>
                                                <span onClick={() => navigate(-1)}>放送回</span>
                                          </BreadcrumbLink>
                                    </BreadcrumbItem>
                                    <BreadcrumbItem>
                                          <BreadcrumbLink isCurrentPage>コメント</BreadcrumbLink>
                                    </BreadcrumbItem>
                              </Breadcrumb>
                              {/* コメント一覧を表示 */}
                              <CommentList episodeID={selectId.episodeID} />
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
});