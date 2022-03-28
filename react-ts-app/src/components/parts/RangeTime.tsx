import { Box, Heading, Stack, Text, Button, NumberDecrementStepper, NumberIncrementStepper, NumberInput, NumberInputField, NumberInputStepper, VStack } from "@chakra-ui/react"
import { SetTime } from "./SetTime";

export const RangeTime = () => {
      const setDateTexts: string[] = ["から", "まで"]
      return (
            <Box w='100%' mt='5' >
                  <Heading as='h3' size='md'>時間</Heading>
                  <VStack pt='3' spacing={5} >

                        {setDateTexts.map((text, index) => (
                              <>
                              <SetTime />
                              <Text key={index} w='100%' pl='1' whiteSpace='nowrap' fontSize='lg'>{text}</Text>
                              </>
                        ))}
                        
                        <Button  colorScheme='blackAlpha' size='lg' w={{ md: '100%' }}>
                              時間指定
                        </Button>
                  </VStack>
            </Box>
      )
}