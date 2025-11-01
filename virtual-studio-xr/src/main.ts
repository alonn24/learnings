// src/main.ts
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { createObjects } from './scene/createObjects';
import { startRenderLoop } from './loop/renderLoop';
import { setupResizeHandler } from './utils/resizeHandler';

function main() {
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);

  // add VR button
  const vrButton = VRButton.createButton(renderer);
  document.body.appendChild(vrButton);

  // scene + camera
  const { scene, camera } = setupScene();
  createObjects(scene);

  // resize
  setupResizeHandler(camera, renderer);

  // start XR render loop
  startRenderLoop(renderer, scene, camera);
}

main();

