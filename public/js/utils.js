// JS Utilities

function g(elementId) {
  const el = document.getElementById(elementId);
  if (!el) {
    console.error('g: unable to find element by id:', elementId);
    return;
  }

  return el;
}
