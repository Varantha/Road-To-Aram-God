import "./Challenge.css";
import PropTypes from "prop-types";

const Challenge = ({ challenge }) => {
  const thresholdsOrder = ["BRONZE", "SILVER", "GOLD", "PLATINUM", "DIAMOND"];

  Challenge.propTypes = {
    challenge: PropTypes.shape({
      challengeName: PropTypes.string.isRequired,
      value: PropTypes.number.isRequired,
      thresholds: PropTypes.object.isRequired,
      level: PropTypes.string.isRequired,
      challengeId: PropTypes.number.isRequired,
    }).isRequired,
  };

  function combineThresholds(orderArray, thresholds) {
    const newArray = [];
    const thresholdNames = Object.keys(thresholds);
    for (let i = 0; i <= orderArray.length - 1; i++) {
      if (thresholdNames.includes(orderArray[i])) {
        newArray.push(orderArray[i]);
      }
    }
    return newArray;
  }

  function getHighestThreshold(challenge) {
    const reversedThresholdsOrder = thresholdsOrder.slice().reverse();

    for (const threshold of reversedThresholdsOrder) {
      if (
        Object.prototype.hasOwnProperty.call(challenge.thresholds, threshold)
      ) {
        return threshold;
      }
    }
    return null; // Return null if no threshold is found.
  }
  //highestThreshold = getHighestThreshold(challenge);

  return (
    <div className="wrapper">
      <div className="challenge" key={challenge.challengeId}>
        <h3 className="challenge-name">{challenge.challengeName}</h3>
        <div className="margin-area">
          {combineThresholds(thresholdsOrder, challenge.thresholds).map(
            (threshold, index) => {
              if (!challenge.thresholds[threshold]) return null;

              const isCurrentThreshold = challenge.level === threshold;
              const isCompleted =
                thresholdsOrder.indexOf(challenge.level) >= index;
              const progressBarWidth =
                ((index + 1) / thresholdsOrder.length) * 100;

              return (
                <>
                  <div
                    className={`dot dot-${threshold}`}
                    style={{
                      left: `${progressBarWidth}%`,
                      background:
                        isCompleted || isCurrentThreshold ? "#0C84D9" : "#bbb",
                    }}
                  >
                    {index + 1}
                  </div>
                  {
                    <div
                      className={`progress-bar progress-bar-${threshold}`}
                      style={{
                        left: `${(index / thresholdsOrder.length) * 100}%`,
                        width: `${(1 / thresholdsOrder.length) * 100}%`,
                        background: isCompleted ? "#0C84D9" : "#bbb",
                      }}
                    />
                  }
                  <div
                    className={`message message-${threshold}`}
                    style={{
                      left: `${((index - 1) / thresholdsOrder.length) * 100}%`,
                      top: "40px",
                    }}
                  >
                    {threshold}
                  </div>
                </>
              );
            }
          )}
        </div>
      </div>
    </div>
  );
};

export default Challenge;
