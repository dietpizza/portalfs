import React, { useState, useRef } from "react";
import type { MouseEvent } from "react";

// interface RippleEffect {
//   id: number;
//   x: number;
//   y: number;
//   size: number;
// }

// interface RippleProps {
//   children: React.ReactNode;
//   className?: string;
//   onClick?: (e: MouseEvent<HTMLDivElement>) => void;
//   disabled?: boolean;
//   rippleColor?: string;
// }

// export const Ripple: React.FC<RippleProps> = ({
//   children,
//   className = "",
//   onClick,
//   disabled = false,
//   rippleColor = "rgba(255, 255, 255, 0.3)",
// }) => {
//   const [ripples, setRipples] = useState<RippleEffect[]>([]);
//   const containerRef = useRef<HTMLDivElement>(null);
//   const nextId = useRef(0);

//   const createRipple = (clientX: number, clientY: number) => {
//     if (disabled || !containerRef.current) return;

//     const rect = containerRef.current.getBoundingClientRect();
//     const x = clientX - rect.left;
//     const y = clientY - rect.top;

//     // Calculate ripple size - should be large enough to cover the entire element
//     const size = Math.max(rect.width, rect.height) * 2;

//     const newRipple: RippleEffect = {
//       id: nextId.current++,
//       x,
//       y,
//       size,
//     };

//     setRipples((prev) => [...prev, newRipple]);

//     // Remove ripple after animation completes
//     setTimeout(() => {
//       setRipples((prev) => prev.filter((r) => r.id !== newRipple.id));
//     }, 400);
//   };

//   const handlePointerDown = (e: React.PointerEvent<HTMLDivElement>) => {
//     if (
//       e.pointerType === "mouse" ||
//       e.pointerType === "touch" ||
//       e.pointerType === "pen"
//     ) {
//       createRipple(e.clientX, e.clientY);
//     }
//   };

//   const handleClick = (e: MouseEvent<HTMLDivElement>) => {
//     if (!disabled && onClick) {
//       onClick(e);
//     }
//   };

//   return (
//     <div
//       ref={containerRef}
//       className={`relative overflow-hidden ${disabled ? "cursor-not-allowed opacity-50" : "cursor-pointer"} ${className}`}
//       onPointerDown={handlePointerDown}
//       onClick={handleClick}
//     >
//       {children}

//       {/* Ripple container */}
//       <span className="absolute inset-0 pointer-events-none overflow-hidden">
//         {ripples.map((ripple) => (
//           <span
//             key={ripple.id}
//             className="absolute rounded-full animate-ripple"
//             style={{
//               left: ripple.x,
//               top: ripple.y,
//               width: ripple.size,
//               height: ripple.size,
//               transform: "translate(-50%, -50%) scale(0)",
//               backgroundColor: rippleColor,
//               animation: "ripple 400ms ease-out",
//             }}
//           />
//         ))}
//       </span>

//       <style jsx>{`
//         @keyframes ripple {
//           0% {
//             transform: translate(-50%, -50%) scale(0);
//             opacity: 1;
//           }
//           100% {
//             transform: translate(-50%, -50%) scale(1);
//             opacity: 0;
//           }
//         }
//       `}</style>
//     </div>
//   );
// };

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
  const [isPressed, setIsPressed] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);
  const nextId = useRef(0);

  const createRipple = (clientX: number, clientY: number) => {
    if (disabled || !containerRef.current) return;

    const rect = containerRef.current.getBoundingClientRect();
    const x = clientX - rect.left;
    const y = clientY - rect.top;

    // MD3 ripple is more compact than MD2
    const size = Math.max(rect.width, rect.height) * 2;

    // Calculate duration based on size for more natural feel
    // Using cubic easing for duration with refined timing
    const maxDimension = Math.max(rect.width, rect.height);
    const normalizedSize = Math.min(maxDimension / 100, 6); // Normalize to 0-6 range
    const duration = 320 + Math.pow(normalizedSize, 1.8) * 65; // Cubic curve, 320-660ms range

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

      setIsPressed(true);
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

    // Let ripple animation complete before removing press state
    // Find the longest duration ripple currently active
    const longestDuration =
      ripples.length > 0 ? Math.max(...ripples.map((r) => r.duration)) : 400;

    setTimeout(() => {
      setIsPressed(false);
    }, longestDuration);
  };

  const handlePointerLeave = () => {
    // Find the longest duration ripple currently active
    const longestDuration =
      ripples.length > 0 ? Math.max(...ripples.map((r) => r.duration)) : 400;

    setTimeout(() => {
      setIsPressed(false);
    }, longestDuration);
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
        <span
          className="absolute inset-0 transition-opacity duration-300"
          style={{
            backgroundColor: hoverColor,
            opacity: isHovered ? 0.08 : 0,
          }}
        />

        {/* Press state layer - MD3 shows this during press */}
        <span
          className="absolute inset-0 transition-opacity duration-300"
          style={{
            backgroundColor: rippleColor,
            opacity: isPressed ? 0.12 : 0,
          }}
        />

        {/* Ripple effects */}
        {ripples.map((ripple) => (
          <span
            key={ripple.id}
            className="absolute rounded-full"
            style={{
              left: ripple.x,
              top: ripple.y,
              width: ripple.size,
              height: ripple.size,
              transform: "translate(-50%, -50%) scale(0)",
              backgroundColor: rippleColor,
              filter: "blur(20px)",
              animation: `md3-ripple ${ripple.duration}ms cubic-bezier(0.2, 0, 0, 1)`,
            }}
          />
        ))}
      </span>

      <style>{`
        @keyframes md3-ripple {
          0% {
            transform: translate(-50%, -50%) scale(0);
            opacity: 0.12;
          }
          50% {
            opacity: 0.1;
          }
          75% {
            opacity: 0.06;
          }
          100% {
            transform: translate(-50%, -50%) scale(1);
            opacity: 0;
          }
        }
      `}</style>
    </div>
  );
};
