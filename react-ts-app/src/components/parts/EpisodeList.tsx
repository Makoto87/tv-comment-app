import { Wrap, WrapItem, Button } from "@chakra-ui/react"
import { memo, ReactNode, VFC } from "react";
import { useNavigate } from "react-router-dom";

import { useQuery, gql } from "@apollo/client";

const FETCH_EPISODES = gql`
      query getEpisodes($input: QueryEpisodesInput!) {
            episodes(input: $input) {
                  id,
                  date
            }
      }
`;

interface Episode {
      id: number;
      date: number;
}

interface EpisodesData {
      episodes: Episode[];
}

interface QueryEpisodesInput {
      programID: number;
      fromDate?: number;
      toDate?: number;
  }

type Props = {
      titles: string[]
}

export const EpisodeList: VFC<Props> = memo((props) => {
      const { titles } = props;
      const navigate = useNavigate();

      // const args: QueryEpisodesInput = {programID: 2}

      const { data } = useQuery<EpisodesData>(FETCH_EPISODES,
            {variables: {
                  input: {
                        programID: 432,
                  }
            }}
      );
      console.log(data)
      console.log("read now")
      data && console.log("read data")
      data && data.episodes.map(episode => console.log(episode.date))

      return (
            <Wrap w='95%' py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  {titles.map((title) => (
                        <WrapItem mx='auto' p={3}>
                              <Button onClick={() => navigate("comments")} whiteSpace='unset' colorScheme='green' w={{ base: '200px', md: '300px' }} h={{ base: '100px', md: '150px'}} fontSize={{ md: '2xl'}}>
                                    {title}
                              </Button>
                        </WrapItem>
                  ))}
            </Wrap>
      )
});