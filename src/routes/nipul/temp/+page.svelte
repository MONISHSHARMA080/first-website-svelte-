
	
<main>
<div class="container">
  <h1>
	<span class="text-animation text">Start by describing us your website</span>
  </h1>
</div>
</main>

<style>
main {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh; /* This will make the main element take up the entire page height */
  background-color: #121212;
  color: #f5f5f5;
  font-family: sans-serif;
}

.container {
  text-align: center;
}

h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  position: relative;
}
.text{
 color: rgb(255, 2, 52);   
}

.text-animation {
  background: linear-gradient(
	90deg,
	rgb(255, 2, 120) 50%,
	rgb(1, 247, 165) 0%,
	rgb(1, 99, 247) 100%
  );
  background-size: 200% 200%;
  background-position: 100% 0;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: text-animation 3s ease 1 forwards;
}

@keyframes text-animation {
  0% {
	background-position: 100% 0;
  }
  50% {
	background-position: 0 0;
  }
  100% {
	background-position: -100% 0;
  }
}

</style>

	