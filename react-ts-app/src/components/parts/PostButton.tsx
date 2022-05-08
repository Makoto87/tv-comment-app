import { Button, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure, VStack, FormControl, FormLabel, Input, Text, Textarea, Flex } from "@chakra-ui/react"
import { memo, useState } from "react";
import { gql, isApolloError, useMutation } from "@apollo/client";

import { useMessage } from "../../hooks/useMessage";
import { SetTime } from "./SetTime";

const SAVE_COMMENT = gql`
      mutation createComment($input: NewComment!) {
            createComment(input: $input)
      }
`;

interface NewComment {
      comment: string
      programName: string
      episodeDate: number
      userID: number
}

export const PostButton = memo(() => {
      const { isOpen, onOpen, onClose } = useDisclosure()

      const [ count, setCount ] = useState(0);
      const [ text, setText ] = useState("");
      const [ programName, setProgramName ] = useState("");
      const [ year, setYear ] = useState("");
      const [ month, setMonth ] = useState("");
      const [ day, setDay ] = useState("");

      const onChangeTextArea = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
            setCount(event.target.value.length)
            setText(event.target.value)
      }

      const showPostModal = () => {
            setCount(0)
            onOpen()
      }

      const onChangeProgramName = (event: React.ChangeEvent<HTMLInputElement>) => {
            setProgramName(event.target.value)
      }

      const { showMessage } = useMessage();

      const [createComment] = useMutation<{createComment: string}>(SAVE_COMMENT)

      const onClickPostButton = () => {
            if (programName == "") {
                  showMessage( {title: '番組名を記入してください', status: 'error'});
                  return
            }
            if (count <= 0) {
                  showMessage( {title: 'コメントを記入してください', status: 'error'});
                  return
            }
            if (count > 140) {
                  showMessage( {title: '文字数オーバー', status: 'error'});
                  return;
            }

            let numberYear = Number(year);
            let numberMonth = Number(month);
            let numberDay = Number(day);

            let date = numberYear * (10**4) + numberMonth * (10 ** 2) + numberDay;

            createComment({
                  variables: {
                        input: {
                              comment: text,
                              programName: programName,
                              episodeDate: date,
                              userID: 1,
                        }
                  }
            }).then(() => {
                  showMessage( {title: 'コメント投稿成功', status: 'success'});
                  onClose();
            }).catch((err) => {
                  if (isApolloError(err)) {
                        showMessage( {title: '失敗：　' + err.graphQLErrors[0]["message"], status: 'error'});
                  } else {
                        showMessage( {title: '失敗：　サーバーエラー', status: 'error'});
                  }
            });
      }

      return (
            <>
                  <Button onClick={showPostModal} colorScheme='blue' size='lg' w='100%'>
                        新規投稿
                  </Button>

                  <Modal isOpen={isOpen} onClose={onClose} autoFocus={false} >
                        <ModalOverlay />
                        <ModalContent>
                              <ModalHeader>新規投稿</ModalHeader>
                              <ModalCloseButton />
                              <ModalBody>
                                    <VStack spacing={4}>
                                         <FormControl>
                                                <FormLabel>番組名</FormLabel>
                                                <Input onChange={(event) => onChangeProgramName(event)} />
                                          </FormControl> 
                                         <FormControl>
                                                <FormLabel>放送日</FormLabel>
                                                <SetTime setYear={setYear} setMonth={setMonth} setDay={setDay} />
                                          </FormControl> 
                                         <FormControl>
                                                <FormLabel>コメント(140字まで)</FormLabel>
                                                <Textarea onChange={(event) => onChangeTextArea(event)} />
                                                <Flex alignContent='flex-end'>
                                                      <Text>現在: {count}文字</Text>
                                                </Flex>
                                          </FormControl> 
                                    </VStack>
                              </ModalBody>

                              <ModalFooter py={5}>
                                    <Button variant='ghost' mr={3} onClick={onClose}>キャンセル</Button>
                                    <Button onClick={onClickPostButton} colorScheme='blue' >
                                          投稿
                                    </Button>
                              </ModalFooter>
                        </ModalContent>
                  </Modal>
            </>
      )
});