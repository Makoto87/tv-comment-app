import { Stack, NumberInput, NumberInputField, NumberInputStepper, NumberIncrementStepper, NumberDecrementStepper } from "@chakra-ui/react";
import { memo, VFC } from "react";

type Props = {
      setYear: React.Dispatch<React.SetStateAction<string>>
      setMonth: React.Dispatch<React.SetStateAction<string>>
      setDay: React.Dispatch<React.SetStateAction<string>>
}

export const SetTime: VFC<Props> = memo((props) => {
      const dateData: string[] = [ "年", "月", "日" ];
      const { setYear, setMonth, setDay } = props
      const dateHooks: React.Dispatch<React.SetStateAction<string>>[] = [setYear, setMonth, setDay]
      return (
            <Stack direction={{ base: 'row', md: 'column' }}>
                  {dateData.map((data, index) => {
                        return (
                        <NumberInput key={index}>
                              <NumberInputField onChange={(event) => dateHooks[index](event.target.value)} bg="white" placeholder={data} />
                              <NumberInputStepper>
                                    <NumberIncrementStepper />
                                    <NumberDecrementStepper />
                              </NumberInputStepper>
                        </NumberInput>
                  )})}
                  
            </Stack>
      )
});