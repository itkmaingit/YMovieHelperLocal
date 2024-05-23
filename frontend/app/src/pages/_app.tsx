import "@/styles/global.css";

import { usePageView } from "@/hooks/use_page_view";
import DefaultLayout from "@/layouts/dashboard/DefaultLayout";

import type { NextPage } from "next";
import type { AppProps } from "next/app";
import type { ReactElement, ReactNode } from "react";

// On each page, there are various patterns to consider. These patterns can be broken down into two main categories:

// 1. Whether or not to apply non-UI logic components such as authentication guards.
// 2. Whether or not to apply a default layout with common headers and footers.
// With these combinations, there are numerous layout patterns to think about. However, in this project, we will specify the following layout patterns:

// By default, unless otherwise specified, we will implement an authentication guard and apply the default layout. If a different layout is required, it should be individually specified on each page.

// +-----------------------------+
// |                             |
// |     Authentication Guard    |
// |                             |
// +-----------------------------+
//                  |
// +-----------------------------+
// |                             |
// |         Default Layout      |
// |   (Header & Footer etc...)  |
// +-----------------------------+
//                  |
//                  v
//          Individual Page
//                  |
//                  |
//           Layout Specification
//         (if different from default)

export type NextPageWithLayout<P = object, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

export default function MyApp({ Component, pageProps }: AppPropsWithLayout) {
  usePageView();

  const getLayout =
    Component.getLayout ??
    ((page) => (
      <>
        <DefaultLayout>{page}</DefaultLayout>
      </>
    ));

  return getLayout(<Component {...pageProps} />);
}
