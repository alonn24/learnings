// src/loop/renderLoop.ts
import * as THREE from 'three';

export function startRenderLoop(
  renderer: THREE.WebGLRenderer,
  scene: THREE.Scene,
  camera: THREE.PerspectiveCamera
) {
  const clock = new THREE.Clock();

  renderer.setAnimationLoop(() => {
    const delta = clock.getDelta();

    // rotate the cube if exists
    const cube = scene.getObjectByName('mainCube');
    if (cube) {
      cube.rotation.y += delta;
    }

    renderer.render(scene, camera);
  });
}

