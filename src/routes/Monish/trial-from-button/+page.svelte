<script>
  let page = 'home';
  let animate = false;

  function navigate(to) {
    page = to;
    animate = true;
    setTimeout(() => { animate = false; }, 500);
  }
</script>

<style>
  :global(body) {
    margin: 0;
    font-family: 'Montserrat', sans-serif;
    background-color: #f0f0f0;
  }

  .container {
    max-width: 1200px;
    margin: 40px auto;
    padding: 20px;
    border-radius: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    background-color: #fff;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
    border-bottom: 1px solid #ddd;
  }

  .nav a {
    text-decoration: none;
    color: #337ab7;
    transition: color 0.2s ease;
  }

  .nav a:hover {
    color: #23527c;
  }

  .nav button {
    background-color: #337ab7;
    color: #fff;
    border: none;
    padding: 10px 20px;
    border-radius: 20px;
    cursor: pointer;
  }

  .nav button:hover {
    background-color: #23527c;
  }

  .hero {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .hero h1 {
    font-size: 48px;
    font-weight: bold;
    color: #337ab7;
  }

  .hero p {
    font-size: 24px;
    color: #666;
  }

  .revolution {
    background-image: linear-gradient(to bottom, #f7f7f7, #fff);
    padding: 40px;
    border-radius: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }

  .revolution h2 {
    font-size: 36px;
    font-weight: bold;
    color: #337ab7;
    margin-bottom: 20px;
  }

  .revolution p {
    font-size: 18px;
    color: #666;
  }

  .features {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }

  .feature {
    margin: 20px;
    padding: 20px;
    border-radius: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    width: calc(33.33% - 20px);
  }

  .feature img {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    margin-bottom: 20px;
  }

  .feature h3 {
    font-size: 24px;
    font-weight: bold;
    color: #337ab7;
    margin-bottom: 10px;
  }

  .feature p {
    font-size: 18px;
    color: #666;
  }

  .call-to-action {
    background-color: #337ab7;
    padding: 20px;
    border-radius: 20px;
    text-align: center;
    color: #fff;
  }

  .call-to-action a {
    color: #fff;
    text-decoration: none;
  }

  .call-to-action a:hover {
    color: #ccc;
  }

  /* Animations */

  .animate {
    transition: opacity 0.5s ease;
  }

  .animate.in {
    opacity: 1;
  }

  .animate.out {
    opacity: 0;
  }

  /* Page Transitions */

  .page {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    opacity: 0;
  }

  .page.in {
    opacity: 1;
  }

  .page.out {
    opacity: 0;
  }

  .page.home {
    background-image: linear-gradient(to bottom, #337ab7, #fff);
  }

  .page.about {
    background-color: #f7f7f7;
  }

  .page.services {
    background-color: #337ab7;
  }

  .page.contact {
    background-color: #337ab7;
  }
</style>

<div class="container">
  <nav class="nav">
    <a href="#" on:click={() => navigate('home')}>Home</a>
    <a href="#" on:click={() => navigate('about')}>About</a>
    <a href="#" on:click={() => navigate('services')}>Services</a>
    <a href="#" on:click={() => navigate('contact')}>Contact</a>
    <button on:click={() => console.log('Join the Revolution!')}>Join the Revolution!</button>
  </nav>

  {#if page === 'home'}
  <div class="hero">
    <h1>Welcome to our GYM!</h1>
    <p>Where fitness meets revolution!</p>
  </div>
  {/if}

  {#if page === 'about'}
  <div class="revolution">
    <h2>About Us</h2>
    <p>We're a group of fitness enthusiasts who believe in creating a revolution in the fitness industry!</p>
  </div>
  {/if}

  {#if page === 'services'}
  <div class="features">
    {#each ['Personal Training', 'Group Fitness', 'Nutrition'] as feature}
    <div class="feature">
      <img src="https://via.placeholder.com/50" alt={feature} />
      <h3>{feature}</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
    </div>
    {/each}
  </div>
  {/if}

  {#if page === 'contact'}
  <div class="call-to-action">
    <h2>Get in Touch!</h2>
    <a href="mailto:info@gym.com">info@gym.com</a>
    <p>Or visit us at 123 Main St, Anytown, USA</p>
  </div>
  {/if}
</div>

{#if animate}
  <div class="page home {page === 'home' ? 'in' : 'out'}">
    <!-- Home page content -->
  </div>

  <div class="page about {page === 'about' ? 'in' : 'out'}">
    <!-- About page content -->
  </div>

  <div class="page services {page === 'services' ? 'in' : 'out'}">
    <!-- Services page content -->
  </div>

  <div class="page contact {page === 'contact' ? 'in' : 'out'}">
    <!-- Contact page content -->
  </div>
{/if}
