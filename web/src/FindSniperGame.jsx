import React, { useState } from "react";
import { Stage, Layer, Image, Circle } from "react-konva";
import useImage from "use-image";

// If the image is in the 'src/assets' folder, import it
import imageSrc from "./assets/stock-photo-basics-word-made-with-building-blocks-584253658.jpg";

// If the image is in the 'public/assets' folder, use a relative path
// const imageSrc = "./assets/stock-photo-basics-word-made-with-building-blocks-584253658.jpg";

const FindSniperGame = ({ prompt }) => {
    const [clickPosition, setClickPosition] = useState(null);
    const [image] = useImage(imageSrc); // Load image using the hook

    const handleStageClick = (e) => {
        const { x, y } = e.target.getStage().getPointerPosition();
        setClickPosition({ x, y });
    };

    return (
        <div style={{ textAlign: "center" }}>
            <h1>{prompt}</h1>
            <Stage
                width={800}
                height={image ? (image.height * 800) / image.width : 600}
                onClick={handleStageClick}
                style={{ border: "1px solid black" }}
            >
                <Layer>
                    {image && (
                        <Image
                            image={image}
                            width={800}
                            height={(image.height * 800) / image.width}
                            listening={false}
                        />
                    )}
                </Layer>

                <Layer>
                    {clickPosition && (
                        <Circle
                            x={clickPosition.x}
                            y={clickPosition.y}
                            radius={50}
                            stroke="red"
                            strokeWidth={5}
                        />
                    )}
                </Layer>
            </Stage>
        </div>
    );
};

export default FindSniperGame;
