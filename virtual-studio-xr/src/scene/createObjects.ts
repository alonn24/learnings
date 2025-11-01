// src/scene/createObjects.ts
import * as THREE from 'three';

export function createObjects(scene: THREE.Scene) {
  const geometry = new THREE.BoxGeometry(1, 1, 1);
  const material = new THREE.MeshStandardMaterial({ color: 0x44aa88 });
  const cube = new THREE.Mesh(geometry, material);
  cube.position.y = 1;
  cube.name = 'mainCube';

  scene.add(cube);
}

