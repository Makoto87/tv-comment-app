import { Container } from "@chakra-ui/react";
import { memo } from "react"
import { Link } from "react-router-dom";

export const NotFound = memo(() => {
      return (
            <>
                  <Container>
                        <p>404ページです</p>
                        <Link style={{color: 'red', textDecoration: 'underline'}} to="/">Homeへ</Link>
                  </Container>
            </>
      )
});