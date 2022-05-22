import { Button } from "@chakra-ui/react"
import { Dispatch, memo, VFC } from "react";
import { gql, useMutation } from "@apollo/client";

type Props = {
      commentID: number
      setLikes: Dispatch<React.SetStateAction<number>>
      setPushedButton: Dispatch<React.SetStateAction<boolean>>
}

const PUSH_LIKE = gql`
      mutation pushLike($commentID: Int!) {
            pushLike(commentID: $commentID)
      }
`

export const GoodButton: VFC<Props> = memo((props) => {
      const {commentID, setLikes, setPushedButton} = props

      const [pushLike] = useMutation<{pushLike: number}>(PUSH_LIKE, {
            variables: {
                  commentID: commentID,
            }
      })

      const onClickPushLike = () =>  {
            pushLike().then((res) => {
                  setPushedButton(true)
                  setLikes(res.data!.pushLike)
            })
      }

      return (
            <Button onClick={onClickPushLike} flexGrow={1} colorScheme='red.200' variant='outline' size='sm'>Good</Button>
      )
})