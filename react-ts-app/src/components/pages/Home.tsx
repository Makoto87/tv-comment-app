import { Flex, Text, VStack } from '@chakra-ui/react'
import { memo } from 'react'

import { SubList } from '../parts/SubList'
import { PostButton } from '../parts/PostButton'
import { ProgramList } from '../parts/ProgramList'
import { HeaderLayout } from '../templates/HeaderLayout'

export const Home = memo(() => {
      return (
            <HeaderLayout>
                  <Flex paddingX={8} w='100%' flexWrap={ {base: 'wrap', md: 'nowrap'}}>
                        <VStack w={{ base: '100%', md: '15%' }} spacing={{ base: '4', md: '5'}} pt={{ base: '3', md: '10'}} pb={ {base: '6', md: '0'}}>
                              <PostButton />
                              <SubList />
                        </VStack>

                        <VStack w={{ base: '100%', md: '85%'}}>
                              <Text w='100%' paddingX={{ base: '1', md: '7'}} fontSize='lg'>ホーム</Text>
                              <ProgramList />
                        </VStack>
                  </Flex>
            </HeaderLayout>
      )
});