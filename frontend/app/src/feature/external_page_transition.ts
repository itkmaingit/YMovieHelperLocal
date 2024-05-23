import { NextRouter } from "next/router";

const handlePageTransition = (router: NextRouter, url: string) => async () => {
  await router.push(url);
};

export default handlePageTransition;
