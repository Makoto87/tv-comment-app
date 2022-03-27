import { Wrap, WrapItem, Button } from "@chakra-ui/react"

export const ProgramList = () => {
      return (
            <Wrap w='95%' py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  <WrapItem mx='auto' p={3}>
                        <Button whiteSpace='unset' colorScheme='green' w={{ base: '200px', md: '300px' }} h={{ base: '100px', md: '150px'}} fontSize={{ md: '2xl'}}>
                              バラエティが面白いので何度も見てしまう病気にかかってしまった。
                        </Button>
                  </WrapItem>
            </Wrap>
      )
}