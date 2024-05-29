<script>
  let currentPage = 'home';

  function navigate(page) {
    currentPage = page;
  }
</script>

<style>
  :global(*){
    font-family: 'Montserrat', sans-serif;
    box-sizing: border-box;
    margin: 0;
    padding: 0;
  }

  :global(body){
    overflow-x: hidden;
    height: 100vh;
    width: 100vw;
    display: flex;
    justify-content: center;
    align-items: center;
    background: radial-gradient(farthest-corner at 40% 50%, #FF69B4, #FFC0CB);
  }

  .container{
    max-width: 1200px;
    margin: 0 auto;
    padding: 2rem;
    border-radius: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    background-color: #FFFFFF;
  }

  .nav-bar{
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 60px;
    background-color: #FF69B4;
    padding: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #FFC0CB;
  }

  .nav-link{
    text-decoration: none;
    color: #FFFFFF;
    font-weight: bold;
    border-radius: 20px;
    padding: 0.5rem 1rem;
    transition: background-color 0.2s ease-in-out;
  }

  .nav-link:hover{
    background-color: #FFC0CB;
  }

  .home{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .hero{
    font-size: 3rem;
    font-weight: bold;
    color: #FF69B4;
    margin-bottom: 2rem;
  }

  .button{
    background-color: #FFC0CB;
    color: #FFFFFF;
    border: none;
    border-radius: 20px;
    padding: 1rem 2rem;
    font-size: 1.5rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.2s ease-in-out;
  }

  .button:hover{
    background-color: #FF69B4;
  }

  .about{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-image: url('https://images.pexels.com/photos/3421624/pexels-photo-3421624.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260');
    background-size: cover;
    background-position: center;
  }

  .about-text{
    color: #FFFFFF;
    font-size: 2rem;
    font-weight: bold;
    text-align: center;
    margin-bottom: 2rem;
  }

  .classes{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-image: url('https://images.pexels.com/photos/3421625/pexels-photo-3421625.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260');
    background-size: cover;
    background-position: center;
  }

  .class-link{
    text-decoration: none;
    color: #FFFFFF;
    font-weight: bold;
    border-radius: 20px;
    padding: 0.5rem 1rem;
    transition: background-color 0.2s ease-in-out;
  }

  .class-link:hover{
    background-color: #FFC0CB;
  }

  .contact{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-image: url('https://images.pexels.com/photos/3421633/pexels-photo-3421633.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260');
    background-size: cover;
    background-position: center;
  }

  .contact-form{
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 2rem;
    border-radius: 20px;
    background-color: #FFFFFF;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }

  .input-field{
    padding: 1rem;
    font-size: 1.5rem;
    border: none;
    border-radius: 20px;
    width: 100%;
    margin-bottom: 1rem;
  }

  .submit-btn{
    background-color: #FFC0CB;
    color: #FFFFFF;
    border: none;
    border-radius: 20px;
    padding: 1rem 2rem;
    font-size: 1.5rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.2s ease-in-out;
  }

  .submit-btn:hover{
    background-color: #FF69B4;
  }
</style>

{#if currentPage === 'home'}
  <div class="container">
    <nav class="nav-bar">
      <a href="#" class="nav-link" on:click={() => navigate('about')}>About</a>
      <a href="#" class="nav-link" on:click={() => navigate('classes')}>Classes</a>
      <a href="#" class="nav-link" on:click={() => navigate('contact')}>Contact</a>
    </nav>
    <main class="home">
      <h1 class="hero">Revolution Fitness</h1>
      <button class="button" on:click={() => navigate('about')}>Learn More</button>
    </main>
  </div>
{/if}

{#if currentPage === 'about'}
  <div class="container">
    <nav class="nav-bar">
      <a href="#" class="nav-link" on:click={() => navigate('home')}>Home</a>
      <a href="#" class="nav-link" on:click={() => navigate('classes')}>Classes</a>
      <a href="#" class="nav-link" on:click={() => navigate('contact')}>Contact</a>
    </nav>
    <main class="about">
      <h2 class="about-text">Welcome to Revolution Fitness</h2>
      <p>We're a team of fitness enthusiasts dedicated to helping you reach your fitness goals.</p>
    </main>
  </div>
{/if}

{#if currentPage === 'classes'}
  <div class="container">
    <nav class="nav-bar">
      <a href="#" class="nav-link" on:click={() => navigate('home')}>Home</a>
      <a href="#" class="nav-link" on:click={() => navigate('about')}>About</a>
      <a href="#" class="nav-link" on:click={() => navigate('contact')}>Contact</a>
    </nav>
    <main class="classes">
      <h2>Our Classes</h2>
      <ul>
        <li><a href="#" class="class-link">Yoga</a></li>
        <li><a href="#" class="class-link">Pilates</a></li>
        <li><a href="#" class="class-link">CrossFit</a></li>
      </ul>
    </main>
  </div>
{/if}

{#if currentPage === 'contact'}
  <div class="container">
    <nav class="nav-bar">
      <a href="#" class="nav-link" on:click={() => navigate('home')}>Home</a>
      <a href="#" class="nav-link" on:click={() => navigate('about')}>About</a>
      <a href="#" class="nav-link" on:click={() => navigate('classes')}>Classes</a>
    </nav>
    <main class="contact">
      <h2>Get in Touch</h2>
      <form class="contact-form">
        <input type="text" class="input-field" placeholder="Name" />
        <input type="email" class="input-field" placeholder="Email" />
        <textarea class="input-field" placeholder="Message" />
        <button class="submit-btn">Send</button>
      </form>
    </main>
  </div>
{/if}
