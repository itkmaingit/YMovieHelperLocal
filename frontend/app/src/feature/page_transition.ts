import { NextRouter } from "next/router";

const handlePageTransition =
  (router?: NextRouter, url = "", isBlank = false, isExternal = false) =>
  async () => {
    if (isBlank) {
      window.open(url, "_blank");
    } else {
      if (isExternal) {
        window.open(url);
      } else {
        await router?.push(url);
      }
    }
  };

export default handlePageTransition;
