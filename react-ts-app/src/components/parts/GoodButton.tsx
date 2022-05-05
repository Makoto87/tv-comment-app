import { Button } from "@chakra-ui/react"
import { memo, VFC } from "react";
import { gql, useMutation } from "@apollo/client";

type Props = {
      commentID: number
      setLikes: React.Dispatch<React.SetStateAction<number>>
}

const PUSH_LIKE = gql`
      mutation pushLike($commentID: Int!) {
            pushLike(commentID: $commentID)
      }
`

export const GoodButton: VFC<Props> = memo((props) => {
      const {commentID, setLikes} = props

      const [pushLike] = useMutation<{pushLike: number}>(PUSH_LIKE, {
            variables: {
                  commentID: commentID,
            }
      })

      const onClickPushLike = () =>  {
            pushLike().then((res) => {
                  setLikes(res.data!.pushLike)
            })
      }

      return (
            <Button onClick={onClickPushLike} flexGrow={1} colorScheme='red.200' variant='outline' size='sm'>Good</Button>
      )
})