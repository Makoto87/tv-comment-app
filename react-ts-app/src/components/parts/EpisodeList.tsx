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
      titleID: number
}

export const EpisodeList: VFC<Props> = memo((props) => {
      const { titleID } = props;
      const navigate = useNavigate();

      const { loading, error, data } = useQuery<EpisodesData>(FETCH_EPISODES,
            {variables: {
                  input: {
                        programID: titleID,
                  }
            }}
      );
      console.log(data, loading, error)

      if (loading) return (
            <h1>Loading Now</h1>
      );

      if (error)   return (
            <h1>Server Error</h1>
      );

      return (
            <Wrap w='95%' py={{ base: 8}} px={{ md: 10 }} justify='space-around' >
                  {data?.episodes.map((episode) => (
                        <WrapItem mx='auto' p={3}>
                              <Button onClick={() => navigate("comments")} whiteSpace='unset' colorScheme='green' w={{ base: '200px', md: '300px' }} h={{ base: '100px', md: '150px'}} fontSize={{ md: '2xl'}}>
                                    {episode.date}
                              </Button>
                        </WrapItem>
                  ))}
            </Wrap>
      )
});