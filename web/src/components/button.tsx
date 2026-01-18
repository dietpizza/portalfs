import React from "react";
import { Ripple } from "./ripple";
import { cn } from "../utils";

type TButton = {
  variant?: "filled" | "outlined" | "text";
  corner: "sm" | "md" | "xl" | "full";
  label?: string;
  icon?: React.ReactNode;
};

export function Button({
  variant = "text",
  corner = "md",
  label,
  icon,
}: TButton) {
  const filledStyles = "bg-primary-container text-on-primary-container";
  const outlinedStyles = "border-primary-container border-2";
  const textStyles = "text-on-surface";

  return (
    <Ripple
      className={cn(
        "inline-block p-2",
        `rounded-${corner}`,
        variant == "filled" && filledStyles,
        variant == "outlined" && outlinedStyles,
        variant == "text" && textStyles,
      )}
    >
      {icon}
      {label && <span className="text-md">{label}</span>}
    </Ripple>
  );
}
