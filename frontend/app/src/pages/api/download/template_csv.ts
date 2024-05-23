import axios from "axios";
import { NextApiRequest, NextApiResponse } from "next";

export default async (req: NextApiRequest, res: NextApiResponse) => {
  const { uid, projectID } = req.query; // get the fileId from query parameters

  try {
    const response = await axios.get(
      `${process.env.CLOUDFRONT_ENDPOINT}/temp/${uid}/${projectID}/template.csv`,
      {
        responseType: "arraybuffer",
      }
    );
    res.setHeader("Content-Disposition", "attachment; filename=template.csv");
    res.send(response.data);
  } catch (error) {
    res.status(500).send("Error downloading file");
  }
};
