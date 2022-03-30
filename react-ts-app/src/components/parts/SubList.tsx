import { Box, Heading, VStack, Button, Stack } from "@chakra-ui/react"
import { memo } from "react"

export const SubList = memo(() => {
      return (
            <Box w='100%' >
                  <Heading as='h3' size='md'>カテゴリ</Heading>
                  <Stack direction={{ base: 'row', md: 'column' }} pt='3' spacing={3} >
                        <Button colorScheme='blackAlpha' size='lg' w={{ md: '100%' }}>
                              バラエティ
                        </Button>
                        <Button  colorScheme='blackAlpha' size='lg' w={{ md: '100%' }}>
                              バラエティ
                        </Button>
                  </Stack>
            </Box>
      )
});