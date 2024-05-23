export const GA_TRACKING_ID = process.env.NEXT_PUBLIC_GA_ID || ""; // あなたのGA4の測定IDに置き換えてください

// https://developers.google.com/analytics/devguides/collection/ga4
export const pageview = (url: URL): void => {
  if (typeof window !== "undefined" && typeof window.gtag === "function") {
    window.gtag("config", GA_TRACKING_ID, {
      page_path: url,
    });
  }
};

type GTagEvent = {
  action: string;
  category: string;
  label: string;
  value: number;
};

export const event = ({ action, category, label, value }: GTagEvent): void => {
  if (typeof window !== "undefined" && typeof window.gtag === "function") {
    window.gtag("event", action, {
      event_category: category,
      event_label: label,
      value: value,
    });
  }
};
