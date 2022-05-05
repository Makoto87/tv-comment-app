import { Wrap, WrapItem, Button, VStack, Box, Text, Flex } from "@chakra-ui/react"
import { memo, VFC } from "react";
import { useQuery, gql } from "@apollo/client";
import { CommentBox } from "./CommentBox";

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
                        <CommentBox comment={comment}></CommentBox>
                  ))}
            </VStack>
      )
});