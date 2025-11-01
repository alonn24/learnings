// src/scene/setupScene.ts
import * as THREE from 'three';

export function setupScene() {
  const scene = new THREE.Scene();
  scene.background = new THREE.Color(0x808080);

  const camera = new THREE.PerspectiveCamera(
    70,
    window.innerWidth / window.innerHeight,
    0.1,
    100
  );
  camera.position.set(0, 1.6, 3);

  // basic light
  const hemiLight = new THREE.HemisphereLight(0xffffff, 0x444444, 1);
  hemiLight.position.set(0, 1, 0);
  scene.add(hemiLight);

  return { scene, camera };
}

