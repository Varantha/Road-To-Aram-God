import { useNavigate } from 'react-router-dom';

export default function SearchBox() {
  const navigateTo = useNavigate();

  function handleSubmit(event) {
    event.preventDefault();
    var summoner = document.getElementById("summoner").value
    var region = document.getElementById("region").value
    navigateTo(`/${region}/${summoner}`)
}

  return (
    <div className="search-body">
    <form className="search-container" onSubmit={event => handleSubmit(event)}>
        <label htmlFor="region">Region:</label>
        <select id="region" className="search-dropdown">
            <option value="NA">North America</option>
            <option selected="selected" value="EUW">Europe West</option>
            <option value="EUNE">Europe Nordic & East</option>
            <option value="KR">Korea</option>
            <option value="OCE">Oceania</option>
        </select>
        <label htmlFor="summoner">Summoner:</label>
        <div className="search-input-container">
            <input type="text" id="summoner" className="search-input" placeholder="Enter summoner name..." />
            <button type="submit" className="search-btn">Search</button>
        </div>
    </form>
    
    </div>
  );
}



