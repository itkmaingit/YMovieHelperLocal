import { Head, Html, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <>
      <Html lang="ja">
        <Head>
          <meta
            name="description"
            content="ゆっくり動画の作成支援ツールです。誰でも無料で使えます。"
          />
          <meta property="og:title" content="YMovieHelper" />
          <meta
            property="og:description"
            content="ゆっくり動画の作成支援ツールです。誰でも無料で使えます。"
          />
          <meta charSet="utf-8" />
        </Head>
        <body style={{ margin: 0, padding: 0 }}>
          <Main />
          <NextScript />
        </body>
      </Html>
    </>
  );
}
