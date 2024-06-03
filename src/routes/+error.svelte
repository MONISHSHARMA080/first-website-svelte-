<script>
    // You don't need to import anything in Svelte for this example
  </script>
  
  <style>
    .flex {
      display: flex;
    }
  
    .h-screen {
      height: 100vh;
    }
  
    .w-full {
      width: 100%;
    }
  
    .flex-col {
      flex-direction: column;
    }
  
    .items-center {
      align-items: center;
    }
  
    .justify-center {
      justify-content: center;
    }
  
    .bg-gray-950 {
      background-color: #242424;
    }
  
    .text-white {
      color: #ffffff;
    }
  
    .animate-pulse {
      /* Define your own pulse animation */
    }
  
    .space-y-4 > * + * {
      margin-top: 1rem;
    }
  
    .text-6xl {
      font-size: 4rem;
      line-height: 1;
    }
  
    .font-bold {
      font-weight: bold;
    }
  
    .text-red-500 {
      color: #dc2626;
    }
  
    .text-2xl {
      font-size: 1.5rem;
    }
  
    .platform {
      background-color: #ec7777;
      color: #242424;
      padding: 5rem;
      border-radius: 2.6rem;
      box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
    }
  </style>
  
  <div class="flex h-screen w-full flex-col items-center justify-center bg-gray-950 text-white">
    <div class="platform">
      <div class="animate-pulse space-y-4">
        <h1 class="text-6xl font-bold text-red-900">Oops!</h1>
      </div>
      <h3 class="text-2xl">An error occurred. Please regenerate your the website that will fix it.</h3>
    </div>
  </div>