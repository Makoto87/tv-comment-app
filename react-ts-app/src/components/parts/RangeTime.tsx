import { Box, Heading, Stack, Text, Button, NumberDecrementStepper, NumberIncrementStepper, NumberInput, NumberInputField, NumberInputStepper, VStack } from "@chakra-ui/react"
import { memo, useState, VFC } from "react";
import { SetTime } from "./SetTime";

type Props = {
      setFromDate: React.Dispatch<React.SetStateAction<number>>
      setToDate: React.Dispatch<React.SetStateAction<number>>
}

export const RangeTime: VFC<Props> = memo((props) => {
      const { setFromDate, setToDate } = props
      
      const [ fromYear, setFromYear ] = useState("");
      const [ fromMonth, setFromMonth ] = useState("");
      const [ fromDay, setFromDay ] = useState("");

      const [ toYear, setToYear ] = useState("");
      const [ toMonth, setToMonth ] = useState("");
      const [ toDay, setToDay ] = useState("");

      const onClickSetTime = () => {
            let numberFromYear = Number(fromYear);
            let numberFromMonth = Number(fromMonth);
            let numberFromDay = Number(fromDay);
            let numberToYear = Number(toYear);
            let numberToMonth = Number(toMonth);
            let numberToDay = Number(toDay);

            let fromDate = numberFromYear * (10**4) + numberFromMonth * (10 ** 2) + numberFromDay;
            let toDate = numberToYear * (10**4) + numberToMonth * (10 ** 2) + numberToDay;

            if (toDate == 0) { toDate = 99991231 }

            setFromDate(fromDate)
            setToDate(toDate)
      }
      return (
            <Box w='100%' mt='5' >
                  <Heading as='h3' size='md'>時間</Heading>
                  <VStack pt='3' spacing={5} >
                        <SetTime setYear={setFromYear} setMonth={setFromMonth} setDay={setFromDay} />
                        <Text w='100%' pl='1' whiteSpace='nowrap' fontSize='lg'>から</Text>
                        <SetTime setYear={setToYear} setMonth={setToMonth} setDay={setToDay} />
                        <Text w='100%' pl='1' whiteSpace='nowrap' fontSize='lg'>まで</Text>

                        <Button onClick={onClickSetTime} colorScheme='blackAlpha' size='lg' w={{ md: '100%' }}>
                              時間指定
                        </Button>
                  </VStack>
            </Box>
      )
});