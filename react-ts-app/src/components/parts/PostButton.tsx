import { Button, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure, VStack, FormControl, FormLabel, Input, Text, Textarea, Flex } from "@chakra-ui/react"
import { useState } from "react";
import { useMessage } from "./PostMessage";
import { SetTime } from "./SetTime";

export const PostButton = () => {
      // Modalに使用する関数
      const { isOpen, onOpen, onClose } = useDisclosure()

      // 文字数カウントをする
      const [ count, setCount ] = useState(0);
      // テキストエリアの文字が変わるたびにカウント
      const onChangeTextArea = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
            setCount(event.target.value.length)
      }
      const showPostModal = () => {
            setCount(0)
            onOpen()
      }
      // 投稿ボタン押した後のメッセージ
      const { showMessage } = useMessage();

      const onClickPostButton = () => {
            console.log(count);
            if (count <= 140) {
                  showMessage( {title: '成功', status: 'success'})
            } else {
                  showMessage( {title: '文字数オーバー', status: 'error'})
            }
            
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
                                                <Input />
                                          </FormControl> 
                                         <FormControl>
                                                <FormLabel>放送日</FormLabel>
                                                <SetTime />
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
}