import type React from "react";

type TPageScaffold = {
  children: any;
};

export const PageScaffold: React.FC<TPageScaffold> = ({ children }) => {
  return (
    <div className="h-screen w-screen bg-background text-on-background">
      {children}
    </div>
  );
};
