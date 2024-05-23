import { useRouter } from "next/router";
import { useEffect } from "react";

import * as gtag from "@/libs/gtag";

export const usePageView = () => {
  const router = useRouter();
  useEffect(() => {
    if (typeof window !== undefined) {
      // Client-side-only code
      const handleRouteChange = (url: URL) => {
        gtag.pageview(url);
      };

      router.events.on("routeChangeComplete", handleRouteChange);

      return () => {
        router.events.off("routeChangeComplete", handleRouteChange);
      };
    }
  }, [router.events]);
};
