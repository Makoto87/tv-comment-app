import { Wrap, WrapItem, Button, VStack, Box, Text, Flex } from "@chakra-ui/react"
import { memo, VFC } from "react";
import { useQuery, gql } from "@apollo/client";

const FETCH_COMMENTS = gql`
      query getCommentss($episodeID: Int!) {
            comments(episodeID: $episodeID) {
                  id,
                  comment,
                  likes,
                  user {
                        id,
                        name
                  },
                  postDate
            }
      }
`;

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

interface CommentData {
      comments: Comment[];
}

type Props = {
      episodeID: number
}

export const CommentList: VFC<Props> = memo((props) => {
      const { episodeID } = props

      const { loading, error, data } = useQuery<CommentData>(FETCH_COMMENTS,
            {variables: {
                  episodeID: episodeID
            }}
      );
      console.log(data, loading, error?.networkError )

      if (loading) return (
            <h1>Loading Now</h1>
      );

      if (error)   return (
            <h1>Server Error</h1>
      );

      return (
            <VStack w='95%' spacing={5} py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  {data?.comments.map((comment) => (
                        <Box bg='white' border='3px' borderWidth={5} borderColor='red.500' shadow='sm' rounded='md' w='100%'  fontSize='xl' p={5} pb={3}>
                              <Text>{comment.comment}</Text>
                              <Flex pt={2} fontSize='md' direction={{ base: 'column', md: 'row' }} justifyContent='space-between' alignItems={{md: 'center'}} textColor='red.300'>
                                    <Text flexGrow={2} pb={{base: '1', md: '0'}}>{comment.user.name}</Text>
                                    <Text flexGrow={4} pb={{base: '3', md: '0'}}>{comment.postDate}</Text>
                                    <Flex alignItems='center' pb={{base: '2', md: '0'}}>
                                          <Button flexGrow={1} colorScheme='red.200' variant='outline' size='sm'>Good</Button>
                                          <Text flexGrow={10} pl='3'>{comment.likes}</Text>
                                    </Flex>
                              </Flex>
                        </Box>
                  ))}
            </VStack>
      )
});