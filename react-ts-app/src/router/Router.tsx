import { Route, Routes } from "react-router-dom"
import { Comments } from "../components/pages/Comments"
import { Home } from "../components/pages/Home"
import { NotFound } from "../components/pages/NotFound"
import { Episodes } from "../components/pages/Episodes"

export const Router = () => {
      return (
            <Routes>
                  <Route index element={<Home />} />
                  <Route path="/episodes">
                        <Route index element={<Episodes />} />
                        <Route path="comments" element={<Comments />} />
                  </Route>
                  <Route path="*" element={<NotFound />} />
            </Routes>
      )
}