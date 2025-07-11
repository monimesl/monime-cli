"use client";

export default function Spinner() {
    return (
        <div className="flex items-center justify-center">
            <div className="relative w-6 h-6">
                {[...Array(12)].map((_, i) => (
                    <span
                        key={i}
                        className="absolute left-1/2 top-0 h-2 w-0.5 origin-bottom rounded-full bg-white opacity-0 animate-spinner"
                        style={{
                            transform: `rotate(${i * 30}deg) translateY(-100%)`,
                            animationDelay: `${i * 0.083}s`,
                        }}
                    />
                ))}
            </div>
        </div>
    );
}