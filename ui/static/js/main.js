var navLinks = document.querySelectorAll("nav a");

for(var i = 0; i < navLinks.length; i++){
  var link = navLinks[i]
  if(link.getAttribute("href") == window.location.pathname) {
      link.classList.add("live");
      break;
  }
}

var colors = ['#4AA5FA', '#73CA86', '#FF9330', '#FF6565']; 

var links = document.querySelectorAll( 'a' );


links.forEach(link => {
  link.onmouseover = function() {
    var randomcolor = colors[Math.floor(Math.random() * colors.length)];
    this.style.color = randomcolor;
  };
  link.onmouseout = function() {
    this.style.color = '#F8EBD8';
  };
});




