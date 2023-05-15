import PropTypes from 'prop-types';
import Challenge from './Challenge';


const thresholdsOrder = ['IRON', 'BRONZE', 'SILVER', 'GOLD', 'PLATINUM', 'DIAMOND']

function getNextThreshold(currentThreshold) {
    const thresholds = thresholdsOrder
  
    const index = thresholds.indexOf(currentThreshold.toUpperCase());
  
    if (index === -1 || index === thresholds.length - 1) {
      // If the current threshold is not found or already at the highest level, return currentThreshold.
      return currentThreshold;
    }
  
    return thresholds[index + 1];
  }

  function getHighestThreshold(challenge) {
    const reversedThresholdsOrder = thresholdsOrder.slice().reverse();

    for (const threshold of reversedThresholdsOrder) {
        if (Object.prototype.hasOwnProperty.call(challenge.thresholds, threshold)) {
        return threshold;
        }
    }
    return null; // Return null if no threshold is found.
  }

  
  
const ChallengeCategory = ({ category }) => {
  return (
    <div className="category">
      <h2 className="category-title">{category.categoryName}</h2>
      <div className="challenges">
        {category.challenges.map((challenge) => (
            <Challenge key={challenge.challengeId} challenge={challenge} />
        ))}
      </div>
    </div>
  );
};

ChallengeCategory.propTypes = {
    category: PropTypes.shape({
      categoryName: PropTypes.string,
      challenges: PropTypes.arrayOf(
        PropTypes.shape({
          challengeId: PropTypes.number,
          challengeName: PropTypes.string,
          value: PropTypes.number,
          level: PropTypes.string,
          thresholds: PropTypes.object,
        })
      ),
      categoryChallenge: PropTypes.shape({
        challengeId: PropTypes.number,
        challengeName: PropTypes.string,
      }),
    }).isRequired,
  };

export default ChallengeCategory;
