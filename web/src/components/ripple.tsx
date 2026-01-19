import React, { useState, useRef } from "react";
import type { MouseEvent } from "react";
import { motion, AnimatePresence } from "motion/react";

interface RippleEffect {
  id: number;
  x: number;
  y: number;
  size: number;
  duration: number;
}

interface RippleProps {
  children: React.ReactNode;
  className?: string;
  onClick?: (e: MouseEvent<HTMLDivElement>) => void;
  disabled?: boolean;
  rippleColor?: string;
  hoverColor?: string;
}

export const Ripple: React.FC<RippleProps> = ({
  children,
  className = "",
  onClick,
  disabled = false,
  rippleColor = "currentColor",
  hoverColor = "currentColor",
}) => {
  const [ripples, setRipples] = useState<RippleEffect[]>([]);
  const [isHovered, setIsHovered] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);
  const [extraOp, setExtraOp] = useState<number | null>(null);
  const nextId = useRef(0);

  const getRippleExtaOpacity = (size: number) => {
    if (size >= 128) return 0;
    if (size >= 72) return 0.1;
    if (size >= 48) return 0.4;
    if (size >= 24) return 0.7;
    return 0.9;
  };

  const createRipple = (clientX: number, clientY: number) => {
    if (disabled || !containerRef.current) return;

    const rect = containerRef.current.getBoundingClientRect();
    const x = clientX - rect.left;
    const y = clientY - rect.top;

    // MD3 ripple is more compact than MD2
    const size = Math.max(rect.width, rect.height) * 2;
    if (extraOp === null) {
      console.info("Extra Opacity:", getRippleExtaOpacity(size / 2), size / 2);
      setExtraOp(getRippleExtaOpacity(size / 2));
    }

    const duration = 480;

    const newRipple: RippleEffect = {
      id: nextId.current++,
      x,
      y,
      size,
      duration,
    };

    setRipples((prev) => [...prev, newRipple]);

    // Remove ripple after animation completes
    setTimeout(() => {
      setRipples((prev) => prev.filter((r) => r.id !== newRipple.id));
    }, duration);
  };

  const handlePointerDown = (e: React.PointerEvent<HTMLDivElement>) => {
    if (disabled) return;
    if (
      e.pointerType === "mouse" ||
      e.pointerType === "touch" ||
      e.pointerType === "pen"
    ) {
      // Stop propagation to prevent parent ripples from triggering
      e.stopPropagation();

      createRipple(e.clientX, e.clientY);

      // On touch devices, clear hover state when touching
      if (e.pointerType === "touch") {
        setIsHovered(false);
      }
    }
  };

  const handlePointerUp = (e: React.PointerEvent<HTMLDivElement>) => {
    // On touch devices, ensure hover doesn't stick after tap
    if (e.pointerType === "touch") {
      setIsHovered(false);
    }
  };

  const handlePointerLeave = () => {
    setIsHovered(false);
  };

  const handlePointerEnter = (e: React.PointerEvent<HTMLDivElement>) => {
    // Only show hover state for mouse, not for touch
    if (!disabled && e.pointerType === "mouse") {
      setIsHovered(true);
    }
  };

  const handleClick = (e: MouseEvent<HTMLDivElement>) => {
    if (!disabled && onClick) {
      onClick(e);
    }
  };

  return (
    <div
      ref={containerRef}
      className={`relative overflow-hidden ${disabled ? "cursor-not-allowed opacity-38" : "cursor-pointer"} ${className}`}
      onPointerDown={handlePointerDown}
      onPointerUp={handlePointerUp}
      onPointerLeave={handlePointerLeave}
      onPointerEnter={handlePointerEnter}
      onClick={handleClick}
    >
      {children}

      {/* MD3 State Layers */}
      <span className="absolute inset-0 pointer-events-none overflow-hidden">
        {/* Hover state layer - MD3 uses persistent overlay */}
        <motion.span
          className="absolute inset-0"
          style={{
            backgroundColor: hoverColor,
          }}
          initial={{ opacity: 0 }}
          animate={{ opacity: isHovered ? 0.08 : 0 }}
          transition={{ duration: 0.3 }}
        />

        {/* Ripple effects with AnimatePresence */}
        <AnimatePresence>
          {ripples.map((ripple) => (
            <motion.span
              key={ripple.id}
              className="absolute rounded-full"
              style={{
                left: ripple.x,
                top: ripple.y,
                width: ripple.size,
                height: ripple.size,
                backgroundColor: rippleColor,
                filter: "blur(60px)",
              }}
              initial={{
                transform: "translate(-50%, -50%) scale(0)",
                opacity: 0,
              }}
              animate={{
                transform: "translate(-50%, -50%) scale(0.9)",
                opacity: 0.22 + (extraOp || 0),
              }}
              exit={{
                opacity: 0,
              }}
              transition={{
                duration: ripple.duration / 1000,
                ease: "easeOut",
              }}
            />
          ))}
        </AnimatePresence>
      </span>
    </div>
  );
};
