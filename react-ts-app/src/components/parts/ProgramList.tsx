import { Wrap, WrapItem, Button } from "@chakra-ui/react"
import { ReactNode, VFC } from "react";
import { useNavigate } from "react-router-dom";

type Props = {
      buttonType: "programs" | "comments";
      titles: string[]
}

export const ProgramList: VFC<Props> = (props) => {
      const { buttonType, titles } = props;
      const navigate = useNavigate();
      return (
            <Wrap w='95%' py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  {titles.map((title) => (
                        <WrapItem mx='auto' p={3}>
                              <Button onClick={() => navigate(buttonType)} whiteSpace='unset' colorScheme='green' w={{ base: '200px', md: '300px' }} h={{ base: '100px', md: '150px'}} fontSize={{ md: '2xl'}}>
                                    {title}
                              </Button>
                        </WrapItem>
                  ))}
            </Wrap>
      )
}