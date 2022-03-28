import { Stack, NumberInput, NumberInputField, NumberInputStepper, NumberIncrementStepper, NumberDecrementStepper } from "@chakra-ui/react";

export const SetTime = () => {
      const dateData: string[] = [ "年", "月", "日" ];
      return (
            <Stack direction={{ base: 'row', md: 'column' }}>
                  {dateData.map((data, index) => {
                        return (
                        <NumberInput key={index}>
                              <NumberInputField  bg="white" placeholder={data} />
                              <NumberInputStepper>
                                    <NumberIncrementStepper />
                                    <NumberDecrementStepper />
                              </NumberInputStepper>
                        </NumberInput>
                  )})}
                  
            </Stack>
      )
}