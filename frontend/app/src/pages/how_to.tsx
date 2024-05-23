import fs from "fs";
import path from "path";

import { Box } from "@mui/material";
import { ReactElement } from "react";

import { NextPageWithLayout } from "./_app";

import DefaultLayout from "@/layouts/dashboard/DefaultLayout";
import MarkdownToHTML from "@/libs/convertMarkdown";

type HowToPageProps = {
  htmlText: string;
};

const HowToPage: NextPageWithLayout<HowToPageProps> = ({ htmlText }) => {
  return (
    <Box>
      <Box
        dangerouslySetInnerHTML={{ __html: htmlText }}
        className="markdown"
        sx={{
          "@media (max-width:600px)": {
            // position: "absolute",

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
  const filePath = path.join(process.cwd(), "public", `how_to.md`);
  const fileContents = fs.readFileSync(filePath, "utf8");
  const htmlText = await MarkdownToHTML(fileContents);
  return {
    props: {
      htmlText,
    },
  };
}

HowToPage.getLayout = function getLayout(page: ReactElement) {
  return <DefaultLayout>{page}</DefaultLayout>;
};

export default HowToPage;
