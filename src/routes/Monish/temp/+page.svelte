<script>
  let homepageVisible = true;
  let collectionVisible = false;
  let contactVisible = false;

  function navigateTo(section) {
    homepageVisible = section === 'home';
    collectionVisible = section === 'collection';
    contactVisible = section === 'contact';
  }
</script>

<style>
  :global(body) {
    margin: 0;
    font-family: 'Montserrat', sans-serif;
    background-color: #F8F4FF; /* Soft, luxury-inspired background */
  }

  .main-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    padding: 2rem;
  }

  .header {
    background-color: #333;
    padding: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .logo {
    font-size: 1.5rem;
    font-weight: bold;
    color: #FFD700; /* Gold accents */
  }

  .nav-link {
    color: #fff;
    text-decoration: none;
    transition: color 0.2s ease;
  }

  .nav-link:hover {
    color: #FFD700;
  }

  .hero-section {
    background-image: linear-gradient(to bottom, #F8F4FF, #E5E5EA);
    padding: 4rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }

  .hero-text {
    font-size: 3rem;
    font-weight: bold;
    color: #333;
    text-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  }

  .hero-button {
    background-color: #FFD700;
    border: none;
    padding: 1rem 2rem;
    font-size: 1.5rem;
    font-weight: bold;
    color: #fff;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .hero-button:hover {
    background-color: #FFC400;
  }

  .collection-section {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }

  .collection-item {
    width: calc(25% - 1rem);
    margin: 0.5rem;
    background-color: #fff;
    padding: 1rem;
    border: 1px solid #ddd;
    border-radius: 10px;
    cursor: pointer;
  }

  .collection-item:hover {
    border-color: #FFD700;
  }

  .contact-section {
    padding: 2rem;
  }

  .contact-form {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .contact-input {
    padding: 1rem;
    border: none;
    border-bottom: 1px solid #ccc;
    font-size: 1.5rem;
  }

  .contact-button {
    background-color: #FFD700;
    border: none;
    padding: 1rem 2rem;
    font-size: 1.5rem;
    font-weight: bold;
    color: #fff;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .contact-button:hover {
    background-color: #FFC400;
  }

  /* Animations */
  .fade-in {
    animation: fade-in 0.5s ease;
  }

  @keyframes fade-in {
    0% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }

  .slide-up {
    animation: slide-up 0.5s ease;
  }

  @keyframes slide-up {
    0% {
      transform: translateY(50px);
      opacity: 0;
    }
    100% {
      transform: translateY(0);
      opacity: 1;
    }
  }
</style>

<div class="main-container">
  <header class="header">
    <div class="logo">Jewelry Luxe</div>
    <nav>
      <a class="nav-link" href="#" on:click={() => navigateTo('home')}>Home</a>
      <a class="nav-link" href="#" on:click={() => navigateTo('collection')}>Collection</a>
      <a class="nav-link" href="#" on:click={() => navigateTo('contact')}>Contact</a>
    </nav>
  </header>
  {#if homepageVisible}
  <section class="hero-section" data-animate="fade-in">
    <h1 class="hero-text">Experience Luxury Jewelry</h1>
    <button class="hero-button" on:click={() => navigateTo('collection')}>Explore Our Collection</button>
  </section>
  {/if}
  {#if collectionVisible}
  <section class="collection-section" data-animate="slide-up">
    <div class="collection-item">
      <img src="ring.jpg" alt="Ring">
      <h2>Diamond Ring</h2>
      <p>$10,000</p>
    </div>
    <div class="collection-item">
      <img src="necklace.jpg" alt="Necklace">
      <h2>Pearl Necklace</h2>
      <p>$5,000</p>
    </div>
    <!-- Add more collection items -->
  </section>
  {/if}
  {#if contactVisible}
  <section class="contact-section" data-animate="fade-in">
    <h2>Get in Touch</h2>
    <form class="contact-form">
      <input type="text" class="contact-input" placeholder="Name">
      <input type="email" class="contact-input" placeholder="Email">
      <textarea class="contact-input" placeholder="Message"></textarea>
      <button class="contact-button">Send Message</button>
    </form>
  </section>
  {/if}
</div>

This code uses Svelte's syntax and features like conditional statements ({#if ...}) and event handlers (on:click={() => ...}) to create a dynamic and interactive website. The design is inspired by Material Design, with a luxurious and posh theme featuring gold and silver accents, curved lines, and bold typography. The website has three main sections: a hero section with a call-to-action, a collection section showcasing jewelry items, and a contact section with a form. The code includes animations for a smooth user experience