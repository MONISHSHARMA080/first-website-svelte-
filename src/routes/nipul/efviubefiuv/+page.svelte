<script>
  let showIntro = true;
  let showAbout = false;
  let showContact = false;

  function introEnd() {
    showIntro = false;
    showAbout = true;
  }

  function aboutEnd() {
    showAbout = false;
    showContact = true;
  }
</script>

<style>
  :global(body) {
    margin: 0;
    font-family: 'Montserrat', sans-serif;
    background: #F8F9FA; /* light gray */
  }

  .container {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }

  .curve {
    width: 100%;
    height: 100vh;
    background: #F8F9FA;
    border-radius: 50%;
    transform: rotate(45deg);
    position: absolute;
    top: 0;
    left: 0;
    z-index: -1;
  }

  .jewelry-logo {
    font-size: 36px;
    color: #F7DC6F; /* gold */
    font-weight: bold;
    margin-bottom: 20px;
  }

  .gold-button {
    background: #F7DC6F;
    color: #333;
    border: none;
    padding: 10px 20px;
    font-size: 18px;
    cursor: pointer;
    border-radius: 20px;
  }

  .silver-button {
    background: #B1B1B1; /* silver */
    color: #333;
    border: none;
    padding: 10px 20px;
    font-size: 18px;
    cursor: pointer;
    border-radius: 20px;
  }

  .animation {
    animation: fadeIn 2s;
  }

  @keyframes fadeIn {
    0% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }
</style>

{#if showIntro}
  <div class="curve"></div>
  <div class="container">
    <h1 class="jewelry-logo">Luxury Jewels</h1>
    <p class="animation">Experience the art of luxury jewelry</p>
    <button class="gold-button" on:click={introEnd}>Discover</button>
  </div>
{/if}

{#if showAbout}
  <div class="container">
    <h1 class="jewelry-logo">Our Story</h1>
    <p class="animation">We create exclusive, handcrafted jewelry for the discerning individual</p>
    <button class="silver-button" on:click={aboutEnd}>Explore Our Collection</button>
  </div>
{/if}

{#if showContact}
  <div class="container">
    <h1 class="jewelry-logo">Get in Touch</h1>
    <p class="animation">Join our exclusive club for personalized service and bespoke creations</p>
    <button class="gold-button">Contact Us</button>
  </div>
{/if}


This code generates a website with a Material Design-inspired layout, featuring curved elements and a luxurious color scheme with gold and silver accents. The website has an intro screen with a CTA button, an about page, and a contact page. The animations are triggered on page load and between page transitions. The design is optimized for desktop and mobile devices