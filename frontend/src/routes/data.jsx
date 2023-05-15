import  { useState, useEffect } from 'react';
import ChallengeCategory from '../../components/ChallengeCategory';
export default function Data() {

  const [categories, setCategories] = useState([]);

  useEffect(() => {
    // Replace this with your actual API call
    async function fetchData() {
      const response = await fetch('http://127.0.0.1:8080/euw1/getCombinedChallengeInfo/Varantha');
      const data = await response.json();
      setCategories(data);
    }
    fetchData();
  }, []);

  return (
      <div className="categories">
        {categories.map((category) => (
          <ChallengeCategory key={category.categoryChallenge.challengeId} category={category} />
        ))}
      </div>
  );
}
