<style global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  .loading {
    position: relative;
    display: flex;
  }

  .loading::before {
    content: "";
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: inherit;
    border-radius: inherit;
  }

  .loading::after {
    content: "";
    z-index: 10;
    position: absolute;
    top: calc(50% - (1.5em / 2));
    left: calc(50% - (1.5em / 2));
    width: 1.5em;
    height: 1.5em;
    border-radius: 50%;
    border: 0.25em solid #fff;
    border-color: #f7fafc #a0aec0 #f7fafc #a0aec0;
    animation: loading 0.625s linear infinite;
  }

  @keyframes loading {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
