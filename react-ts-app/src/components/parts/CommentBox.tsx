import { Box, Flex, Text } from "@chakra-ui/react";
import { memo, useState, VFC, Dispatch } from "react";

import { GoodButton } from "./GoodButton";

type Props = {
      comment: Comment
      setPushedButton: Dispatch<React.SetStateAction<boolean>>
}

interface Comment {
      id: number;
      comment: String;
      likes: number;
      user: User;
      postDate: String;
}

interface User {
      id: number;
      name: String;
}

export const CommentBox: VFC<Props> = memo((props) => {
      const { comment, setPushedButton } = props

      const [ likes, setLikes ] = useState(comment.likes)

      return (
            <Box bg='white' border='3px' borderWidth={5} borderColor='red.500' shadow='sm' rounded='md' w='100%'  fontSize='xl' p={5} pb={3}>
                  <Text>{comment.comment}</Text>
                  <Flex pt={2} fontSize='md' direction={{ base: 'column', md: 'row' }} justifyContent='space-between' alignItems={{md: 'center'}} textColor='red.300'>
                        <Text flexGrow={2} pb={{base: '1', md: '0'}}>{comment.user.name}</Text>
                        <Text flexGrow={4} pb={{base: '3', md: '0'}}>{comment.postDate}</Text>
                        <Flex alignItems='center' pb={{base: '2', md: '0'}}>
                              <GoodButton commentID={comment.id} setLikes={setLikes} setPushedButton={setPushedButton}></GoodButton>
                              <Text flexGrow={10} pl='3'>{likes}</Text>
                        </Flex>
                  </Flex>
            </Box>
      )
})