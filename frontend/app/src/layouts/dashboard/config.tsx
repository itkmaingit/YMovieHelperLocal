import DescriptionIcon from "@mui/icons-material/Description";
import GavelIcon from "@mui/icons-material/Gavel";
import HelpIcon from "@mui/icons-material/Help";
import HomeIcon from "@mui/icons-material/Home";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import SlideshowIcon from "@mui/icons-material/Slideshow";
import SpeakerNotesIcon from "@mui/icons-material/SpeakerNotes";
import UploadFileIcon from "@mui/icons-material/UploadFile";
import { SvgIcon } from "@mui/material";

export type Belong = "None" | "Software" | "Project";
type Item = {
  title: string;
  path: string;
  icon: React.ReactNode;
  belong: Belong;
  isBlank: boolean;
  isExternal: boolean;
};

export const items: Item[] = [
  {
    //item.path==pathnameなら、サイドバーのアイコンがハイライトされる
    title: "Home",
    path: "/mypage/dashboard",
    icon: (
      <SvgIcon fontSize="large">
        <HomeIcon />
      </SvgIcon>
    ),
    belong: "None",
    isBlank: false,
    isExternal: false,
  },
  {
    title: "Characters",
    path: "characters",
    icon: (
      <SvgIcon fontSize="large">
        <PeopleAltIcon />
      </SvgIcon>
    ),
    belong: "Software",
    isBlank: false,
    isExternal: false,
  },
  {
    title: "Items",
    path: "items",
    icon: (
      <SvgIcon fontSize="large">
        <UploadFileIcon />
      </SvgIcon>
    ),
    belong: "Project",
    isBlank: false,
    isExternal: false,
  },
  {
    title: "Rules",
    path: "rules",
    icon: (
      <SvgIcon fontSize="large">
        <GavelIcon />
      </SvgIcon>
    ),
    belong: "Project",
    isBlank: false,
    isExternal: false,
  },
  {
    title: "Create",
    path: "make_ymmp",
    icon: (
      <SvgIcon fontSize="large">
        <SlideshowIcon />
      </SvgIcon>
    ),
    belong: "Project",
    isBlank: false,
    isExternal: false,
  },
  {
    title: "How to",
    path: "/how_to",
    icon: (
      <SvgIcon fontSize="large">
        <HelpIcon />
      </SvgIcon>
    ),
    belong: "None",
    isBlank: true,
    isExternal: false,
  },
  {
    title: "Terms of use",
    path: "/terms_of_use",
    icon: (
      <SvgIcon fontSize="large">
        <DescriptionIcon />
      </SvgIcon>
    ),
    belong: "None",
    isBlank: true,
    isExternal: false,
  },
  {
    title: "Inquiry",
    path: "https://forms.gle/Gh3ZpW9DS64eb1qS7",
    icon: (
      <SvgIcon fontSize="large">
        <SpeakerNotesIcon />
      </SvgIcon>
    ),
    belong: "None",
    isBlank: true,
    isExternal: true,
  },
];
