import fs from "fs";
import path from "path";

import { Box } from "@mui/material";
import { ReactElement } from "react";

import { NextPageWithLayout } from "./_app";

import DefaultLayout from "@/layouts/dashboard/DefaultLayout";
import MarkdownToHTML from "@/libs/convertMarkdown";

type TermsOfUsePageProps = {
  htmlText: string;
};

const TermsOfUsePage = ({
  htmlText,
}: TermsOfUsePageProps & NextPageWithLayout) => {
  return (
    <Box>
      <Box
        dangerouslySetInnerHTML={{ __html: htmlText }}
        className="markdown"
        sx={{
          "@media (max-width:600px)": {
            width: "80vw",
            ml: "-100px",
            pr: "50px",
          },
        }}
      ></Box>
    </Box>
  );
};

export async function getStaticProps() {
  const filePath = path.join(process.cwd(), "public", `terms_of_use.md`);
  const fileContents = fs.readFileSync(filePath, "utf8");
  const htmlText = await MarkdownToHTML(fileContents);
  return {
    props: {
      htmlText,
    },
  };
}

TermsOfUsePage.getLayout = function getLayout(page: ReactElement) {
  return <DefaultLayout>{page}</DefaultLayout>;
};

export default TermsOfUsePage;
