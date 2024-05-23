import { Box, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import useSWR from "swr";

import DynamicItemTable from "@/components/ShowItems/DynamicItemTable";
import MultipleItemTable from "@/components/ShowItems/MultipleItemTable";
import SingleItemTable from "@/components/ShowItems/SingleItemTable";
import FullScreenLoading from "@/components/utils/FullScreenLoading";
import GlobalAddButton from "@/components/utils/GlobalAddButton";
import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import SectionTitle from "@/components/utils/SectionTitle";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useFeedback from "@/hooks/use_feedback";
import { clientAxios } from "@/libs/axios";

export default function SimpleTable() {
  const router = useRouter();
  const { software_id, project_id } = router.query;
  const handleNavigateToItemCreate = async () => {
    router.push(`/mypage/${software_id}/${project_id}/items/create_item`);
  };

  const fetcher = async (url: string) => {
    const token = parseCookies().token;
    const getConfig = {
      method: "get",
      url: url,
      headers: { Authorization: `Bearer ${token}` },
    };
    const response = await clientAxios(getConfig);
    return response.data;
  };

  const { data, mutate } = useSWR(
    `/softwares/${software_id}/projects/${project_id}/items`,
    fetcher
  );

  const [setCode, setState, snackbarProps] = useFeedback();

  return (
    <>
      <LoadingAndSnackbar {...snackbarProps} />
      <GlobalAddButton clickFunction={handleNavigateToItemCreate} />
      <FullScreenLoading open={data === undefined}></FullScreenLoading>
      <SectionTitle title="Items"></SectionTitle>
      <FunctionHandlersContext.Provider
        value={{ mutate: mutate, setCode: setCode, setState: setState }}
      >
        <Box sx={{ my: 4 }}>
          <Typography
            variant="h4"
            gutterBottom
            component="div"
            sx={{ marginBottom: 4 }}
          >
            Dynamic Items
          </Typography>
          <DynamicItemTable items={data?.dynamicItems} />
        </Box>

        <Box sx={{ my: 4 }}>
          <Typography
            variant="h4"
            gutterBottom
            component="div"
            sx={{ marginBottom: 4 }}
          >
            Single Items
          </Typography>
          <SingleItemTable items={data?.singleItems} />
        </Box>

        <Box sx={{ my: 4 }}>
          <Typography
            variant="h4"
            gutterBottom
            component="div"
            sx={{ marginBottom: 4 }}
          >
            Multiple Items
          </Typography>
          <MultipleItemTable items={data?.multipleItems} />
        </Box>
      </FunctionHandlersContext.Provider>
    </>
  );
}
