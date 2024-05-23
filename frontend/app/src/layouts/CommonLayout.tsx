import { ThemeProvider } from "@mui/material";
import CssBaseline from "@mui/material/CssBaseline";
import { Analytics } from "@vercel/analytics/react";
import { AnimatePresence, motion } from "framer-motion";
import Head from "next/head";
import { useRouter } from "next/router";
import Script from "next/script";
import * as React from "react";

import { GA_TRACKING_ID } from "@/libs/gtag";
import theme from "@/theme";

type LayoutProps = Required<{
  readonly children: React.ReactNode;
}>;
const CommonLayout = ({ children }: LayoutProps) => {
  const pageTransition = {
    hidden: {
      opacity: 0,
    },
    show: {
      opacity: 1,
      transition: {
        duration: 0.5,
      },
    },
    exit: {
      opacity: 0,
      transition: {
        duration: 0.2,
      },
    },
  };
  const router = useRouter();

  return (
    <>
      <Head>
        <title>YMovieHelper</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <Analytics />
      {GA_TRACKING_ID && (
        <>
          <Script
            id="load-ga"
            src={`https://www.googletagmanager.com/gtag/js?id=${GA_TRACKING_ID}`}
            strategy="afterInteractive"
          />
          <Script
            id="load-ga-script"
            strategy="afterInteractive"
            dangerouslySetInnerHTML={{
              __html: `
                window.dataLayer = window.dataLayer || [];
                function gtag(){dataLayer.push(arguments);}
                gtag('js', new Date());

                gtag('config', '${GA_TRACKING_ID}');
                `,
            }}
          ></Script>
        </>
      )}
      <AnimatePresence mode="wait">
        <motion.div
          key={router.route}
          variants={pageTransition}
          initial="hidden"
          animate="show"
          exit="exit"
        >
          <ThemeProvider theme={theme}>
            <CssBaseline />
            {children}
          </ThemeProvider>
        </motion.div>
      </AnimatePresence>
    </>
  );
};

export default CommonLayout;
