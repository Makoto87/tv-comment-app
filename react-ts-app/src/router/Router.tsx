import { Route, Routes } from "react-router-dom"
import { Comments } from "../components/pages/Comments"
import { Home } from "../components/pages/Home"
import { NotFound } from "../components/pages/NotFound"
import { Program } from "../components/pages/Program"

export const Router = () => {
      return (
            <Routes>
                  <Route index element={<Home />} />
                  <Route path="/programs">
                        <Route index element={<Program />} />
                        <Route path="comments" element={<Comments />} />
                  </Route>
                  <Route path="*" element={<NotFound />} />
            </Routes>
      )
}