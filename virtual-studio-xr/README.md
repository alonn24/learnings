````markdown
# WebXR Course (TypeScript + Vite + Three.js + Cursor)
A step-by-step, incremental WebXR project built over 4 weeks.

You’ll build **one** continuous app called **Virtual Studio XR** — an XR scene that runs in the browser (VR + AR), lets you place/manipulate objects, add UI in 3D space, and persist the scene.

This README is written so you can:
1. paste parts of it into **Cursor** as prompts,
2. copy the structures as-is,
3. see **TypeScript-first** examples.

---

## Table of Contents
1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Recommended Tools & Extensions](#recommended-tools--extensions)
4. [Week 1 – Project Setup & Basic WebXR Scene](#week-1--project-setup--basic-webxr-scene)
5. [Week 2 – Interaction & Controllers](#week-2--interaction--controllers)
6. [Week 3 – AR, Lighting, and Physics](#week-3--ar-lighting-and-physics)
7. [Week 4 – XR UI, Scene Persistence, Deployment](#week-4--xr-ui-scene-persistence-deployment)
8. [Next Steps](#next-steps)

---

## Overview

- **Language**: TypeScript
- **Build tool**: Vite
- **3D engine**: Three.js
- **XR layer**: WebXR (via Three.js helpers)
- **IDE workflow**: Cursor, with prompt-driven file generation
- **Target devices**: Browser with WebXR support (e.g. Chrome), and XR headsets (e.g. Meta Quest 3)

Each week:
- starts with: **what you’ll learn**
- defines: **desired outcome**
- defines: **new files / folders**
- adds: **new TS concepts / XR concepts**
- ends with: **Cursor prompt** to generate the week’s base code

---

## Prerequisites

- Node.js ≥ 22.20.0
- npm ≥ 10
- Basic familiarity with:
  - modules (`import` / `export`)
  - async / event-based code
  - 3D coordinate systems (optional but helps)
- A device that can open `https://...` over HTTPS (for WebXR in real headset)

---

## Recommended Tools & Extensions

### Browser Extensions

#### Chrome/Edge Extensions for WebXR Development

1. **WebXR API Emulator**
   - **Install**: [Chrome Web Store](https://chromewebstore.google.com/detail/webxr-api-emulator/mjddjgeghkdijejnciaefnkjmkafnnje)
   - **Purpose**: Emulate WebXR devices (headsets, controllers) in desktop browser
   - **Features**:
     - Simulate VR/AR sessions without a physical headset
     - Test controller inputs with keyboard/mouse
     - Switch between different headset types (Quest, Vive, etc.)
     - Adjust FOV, IPD, and other headset properties
   - **Usage**: Essential for development when you don't have constant access to a physical headset

2. **Three.js Developer Tools**
   - **Install**: [Chrome Web Store](https://chromewebstore.google.com/detail/threejs-developer-tools/ebpnegggocnnhleeicgljbedjkganaek)
   - **Purpose**: Debug Three.js scenes directly in DevTools
   - **Features**:
     - Inspect scene graph hierarchy
     - View mesh, material, and geometry properties
     - Monitor performance metrics
     - Toggle object visibility
     - Real-time property editing

3. **Immersive Web Emulator**
   - **Install**: [Chrome Web Store](https://chromewebstore.google.com/detail/immersive-web-emulator/cgffilbpcibhmcfbgggfhfolhkfbhmik)
   - **Purpose**: Alternative WebXR emulator with different features
   - **Features**:
     - Device emulation for testing
     - Room-scale tracking simulation
     - Hand tracking emulation

---

### Development Tools

#### IDE Extensions (for Cursor/VS Code)

1. **Three.js Snippets**
   - Autocomplete for Three.js classes and methods
   - Code snippets for common patterns

2. **GLSL Lint**
   - Syntax highlighting for shader code
   - Error checking for GLSL

3. **Live Server** (if not using Vite's dev server)
   - Serves your app over HTTPS (required for WebXR)
   - Auto-reload on file changes

#### Command-line Tools

1. **nvm** (Node Version Manager)
   - **Install**: [nvm installation guide](https://github.com/nvm-sh/nvm)
   - **Usage**: `nvm use` in project directory (reads `.nvmrc`)
   - **Purpose**: Ensures consistent Node.js version across team

2. **serve** (for testing production builds)
   ```bash
   npm install -g serve
   serve -s dist -l 5000 --ssl-cert cert.pem --ssl-key key.pem
   ```

3. **mkcert** (for local HTTPS)
   ```bash
   # Install mkcert
   brew install mkcert  # macOS
   choco install mkcert # Windows
   
   # Create local CA
   mkcert -install
   
   # Generate certificates
   mkcert localhost 127.0.0.1 ::1
   ```

---

### Testing & Debugging Tools

#### Browser DevTools Tips for WebXR

1. **Remote Debugging** (for testing on Quest/other headsets)
   - **Quest**: Connect via USB, enable developer mode, use `chrome://inspect`
   - **Purpose**: Debug your WebXR app running on actual device
   - **Setup**:
     ```bash
     adb devices  # Verify device connection
     # Open chrome://inspect in desktop Chrome
     # Your device's browser tabs will appear
     ```

2. **Performance Monitor**
   - Press `Ctrl+Shift+P` (Chrome DevTools)
   - Type "Show Performance Monitor"
   - Monitor FPS, CPU, memory in real-time

3. **WebXR Device Logging**
   - Check `about:gpu` for WebGL/WebGPU capabilities
   - Check `chrome://device-log` for XR device events

#### Testing Frameworks

1. **Playwright** (for automated testing)
   ```bash
   npm install -D @playwright/test
   ```
   - Can test WebXR flows programmatically
   - Supports headless testing

2. **WebXR Test API** (experimental)
   - Part of WebXR spec for automated testing
   - Not yet widely supported

---

### Recommended NPM Packages

#### Essential for this course:
- ✅ `three` - 3D engine
- ✅ `@types/three` - TypeScript definitions
- ✅ `vite` - Build tool
- ✅ `cannon-es` - Physics engine (Week 3)

#### Optional but helpful:

1. **three-mesh-ui**
   ```bash
   npm install three-mesh-ui
   ```
   - Better 3D UI components than basic meshes
   - Proper text rendering, buttons, panels

2. **@theatre/core + @theatre/studio**
   ```bash
   npm install @theatre/core @theatre/studio
   ```
   - Animation and sequencing tool
   - Visual timeline editor

3. **leva**
   ```bash
   npm install leva
   ```
   - GUI controls for tweaking parameters during development
   - Works great for material/light adjustments

4. **stats.js** (included in Three.js examples)
   ```typescript
   import Stats from 'three/examples/jsm/libs/stats.module.js';
   ```
   - FPS/memory monitor overlay
   - Essential for performance optimization

5. **drei** (React Three Fiber ecosystem - if using React)
   - Pre-built helpers and abstractions
   - Not covered in this course (pure Three.js)

---

### Hardware Recommendations

#### For Testing

1. **Meta Quest 3 / Quest Pro**
   - Best developer experience
   - Built-in browser with excellent WebXR support
   - USB-C connection for remote debugging

2. **Meta Quest 2**
   - More affordable option
   - Good WebXR support
   - Large user base

3. **Phone with ARCore/ARKit** (for AR testing)
   - Modern Android (ARCore) or iPhone (ARKit)
   - Chrome or Safari with WebXR support

#### Development Hardware

1. **Good GPU** (NVIDIA RTX 3060+ or equivalent)
   - Smooth development experience
   - Faster build times
   - Better preview rendering

2. **USB-C Cable** (for Quest debugging)
   - Link Cable or quality USB-C 3.0 cable
   - For `adb` connection and remote debugging

---

### Learning & Reference Resources

1. **WebXR Device API Spec**
   - https://immersiveweb.dev/
   - Official specs and examples

2. **Three.js Documentation**
   - https://threejs.org/docs/
   - Examples: https://threejs.org/examples/

3. **WebXR Samples**
   - https://immersive-web.github.io/webxr-samples/
   - Real working examples of WebXR features

4. **Three.js Journey** (paid course)
   - https://threejs-journey.com/
   - Excellent for deep Three.js knowledge

5. **Mozilla WebXR Guide**
   - https://developer.mozilla.org/en-US/docs/Web/API/WebXR_Device_API
   - Comprehensive WebXR documentation

---

### Quick Setup Checklist

Before starting Week 1:
- [ ] Install Node.js 22.20.0 (use `nvm`)
- [ ] Install **WebXR API Emulator** Chrome extension
- [ ] Install **Three.js Developer Tools** Chrome extension
- [ ] Set up HTTPS for local development (optional for Week 1, required for device testing)
- [ ] If you have a Quest: Enable Developer Mode and set up `adb`
- [ ] Clone/create project and run `npm install`

---

## Week 1 – Project Setup & Basic WebXR Scene

### Goal
Set up a modern TS project that can **render a 3D scene** and **enter VR** with a single button.

### What you will learn
- How to bootstrap a TS + Vite project for WebXR
- Three.js core objects: `Scene`, `PerspectiveCamera`, `WebGLRenderer`
- How to enable WebXR in Three.js (`renderer.xr.enabled = true`)
- How to organize the code so later weeks can extend it

### Expected result
Opening `http://localhost:5173/` will show:
- a grey background
- a spinning cube
- a **“Enter VR”** button (from Three.js)  
In a Quest browser → press → it enters VR.

---

### 1. Create the project

```bash
# create project
npm create vite@latest virtual-studio-xr -- --template vanilla-ts

cd virtual-studio-xr

# install 3D deps
npm install three @types/three

# run
npm run dev
````

---

### 2. Suggested folder structure

```text
virtual-studio-xr/
│
├── index.html
├── tsconfig.json
├── vite.config.ts
│
└── src/
    ├── main.ts               # app entrypoint
    ├── scene/
    │   ├── setupScene.ts     # create scene, camera, lights
    │   └── createObjects.ts  # add initial meshes (cube)
    ├── loop/
    │   └── renderLoop.ts     # XR-aware animation loop
    └── utils/
        └── resizeHandler.ts  # handle window resize
```

This structure is important because **Week 2** will add `managers/`, **Week 3** will add `physics/`, and **Week 4** will add `ui/` and `state/`.

---

### 3. `src/main.ts`

```ts
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
```

---

### 4. `src/scene/setupScene.ts`

```ts
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
```

---

### 5. `src/scene/createObjects.ts`

```ts
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
```

---

### 6. `src/loop/renderLoop.ts`

```ts
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
```

---

### 7. `src/utils/resizeHandler.ts`

```ts
// src/utils/resizeHandler.ts
import * as THREE from 'three';

export function setupResizeHandler(
  camera: THREE.PerspectiveCamera,
  renderer: THREE.WebGLRenderer
) {
  window.addEventListener('resize', () => {
    const { innerWidth, innerHeight } = window;
    camera.aspect = innerWidth / innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(innerWidth, innerHeight);
  });
}
```

---

### 8. Cursor prompt (Week 1)

```text
Create a WebXR-ready TypeScript project structure for Vite and Three.js.

Requirements:
- main.ts as entrypoint
- scene/setupScene.ts to create a scene, camera, and basic light
- scene/createObjects.ts to add a single spinning cube
- loop/renderLoop.ts to handle renderer.setAnimationLoop and rotate the cube
- utils/resizeHandler.ts to handle window resize
- enable WebXR with renderer.xr.enabled = true and add VRButton
- code must be modular and strongly typed
```

---

## Week 2 – Interaction & Controllers

### Goal

Add **input** to the scene: VR controller → raycast → select object → change it.

### What you will learn

* WebXR input device flow (`renderer.xr.getController(...)`)
* Raycasting in Three.js
* How to structure an **InteractionManager**
* How to extend an existing scene without rewriting it

### Expected result

* Scene has 3 objects (cube, sphere, cone)
* You can point at them with the controller and **on select** → the target changes color

---

### 1. New folders / files

Add:

```text
src/
  managers/
    InteractionManager.ts
  scene/
    createMultipleObjects.ts
```

We keep Week 1 code — **we don’t delete it**.

---

### 2. Install (optional) helpers

No extra npm deps are strictly required for Week 2 — we use Three.js built-ins.

---

### 3. `src/scene/createMultipleObjects.ts`

```ts
// src/scene/createMultipleObjects.ts
import * as THREE from 'three';

export function createMultipleObjects(scene: THREE.Scene) {
  const objects: THREE.Mesh[] = [];

  const cube = new THREE.Mesh(
    new THREE.BoxGeometry(0.6, 0.6, 0.6),
    new THREE.MeshStandardMaterial({ color: 0x00aaff })
  );
  cube.position.set(-1, 1, -1);
  objects.push(cube);

  const sphere = new THREE.Mesh(
    new THREE.SphereGeometry(0.4, 32, 32),
    new THREE.MeshStandardMaterial({ color: 0xff5500 })
  );
  sphere.position.set(0, 1.2, -1.5);
  objects.push(sphere);

  const cone = new THREE.Mesh(
    new THREE.ConeGeometry(0.4, 0.8, 32),
    new THREE.MeshStandardMaterial({ color: 0x55ff55 })
  );
  cone.position.set(1, 1, -1);
  objects.push(cone);

  objects.forEach((obj) => {
    obj.castShadow = false;
    scene.add(obj);
  });

  return objects;
}
```

---

### 4. `src/managers/InteractionManager.ts`

```ts
// src/managers/InteractionManager.ts
import * as THREE from 'three';

export class InteractionManager {
  private scene: THREE.Scene;
  private camera: THREE.PerspectiveCamera;
  private renderer: THREE.WebGLRenderer;
  private raycaster: THREE.Raycaster;
  private tempMatrix: THREE.Matrix4;
  private controllers: THREE.Group[] = [];
  private interactable: THREE.Object3D[] = [];

  constructor(
    scene: THREE.Scene,
    camera: THREE.PerspectiveCamera,
    renderer: THREE.WebGLRenderer
  ) {
    this.scene = scene;
    this.camera = camera;
    this.renderer = renderer;
    this.raycaster = new THREE.Raycaster();
    this.tempMatrix = new THREE.Matrix4();

    this.setupControllers();
  }

  public setInteractable(objects: THREE.Object3D[]) {
    this.interactable = objects;
  }

  private setupControllers() {
    const controller = this.renderer.xr.getController(0);
    controller.addEventListener('selectstart', (event) =>
      this.onSelectStart(event.target as THREE.Group)
    );
    this.scene.add(controller);
    this.controllers.push(controller);
  }

  private onSelectStart(controller: THREE.Group) {
    const intersection = this.getIntersection(controller);
    if (intersection && intersection.object) {
      const mesh = intersection.object as THREE.Mesh;
      const material = mesh.material as THREE.MeshStandardMaterial;
      material.color.setHex(Math.random() * 0xffffff);
    }
  }

  private getIntersection(controller: THREE.Group) {
    this.tempMatrix.identity().extractRotation(controller.matrixWorld);
    const origin = controller.position.clone();
    const direction = new THREE.Vector3(0, 0, -1).applyMatrix4(this.tempMatrix);

    this.raycaster.set(origin, direction);

    return this.raycaster.intersectObjects(this.interactable, false)[0];
  }
}
```

---

### 5. Update `src/main.ts` to use the manager

```ts
// src/main.ts (Week 2 version)
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { startRenderLoop } from './loop/renderLoop';
import { setupResizeHandler } from './utils/resizeHandler';
import { createMultipleObjects } from './scene/createMultipleObjects';
import { InteractionManager } from './managers/InteractionManager';

function main() {
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);
  document.body.appendChild(VRButton.createButton(renderer));

  const { scene, camera } = setupScene();
  const objects = createMultipleObjects(scene);

  const interactionManager = new InteractionManager(scene, camera, renderer);
  interactionManager.setInteractable(objects);

  setupResizeHandler(camera, renderer);
  startRenderLoop(renderer, scene, camera);
}

main();
```

---

### 6. Cursor prompt (Week 2)

```text
Extend the existing WebXR TypeScript project to support VR controller interaction.

Requirements:
- create src/managers/InteractionManager.ts
- in that class, set up XR controller via renderer.xr.getController(0)
- use a Raycaster to detect which scene object is pointed at
- on selectstart, change the object's material color
- add a new factory function createMultipleObjects(scene) to create 3 different meshes
- update main.ts to register the InteractionManager and set the created objects as interactables
- keep the project modular (scene/, managers/, loop/, utils/)
```

---

## Week 3 – AR, Lighting, and Physics

### Goal

Add **AR mode** and **basic physics** so objects can rest on a surface and look more real.

### What you will learn

* How to start an `immersive-ar` session
* How to plug a physics engine (Cannon-es) into a render loop
* How to sync physics bodies → Three.js meshes
* How to upgrade lighting for AR

> **Note**: actual device AR requires HTTPS + supported hardware. The structure below is still valid.

---

### 1. Install physics

```bash
npm install cannon-es
```

---

### 2. New folders / files

```text
src/
  physics/
    PhysicsManager.ts
  ar/
    setupAR.ts
```

---

### 3. `src/physics/PhysicsManager.ts`

```ts
// src/physics/PhysicsManager.ts
import * as CANNON from 'cannon-es';
import * as THREE from 'three';

export class PhysicsManager {
  public world: CANNON.World;
  private meshBodyMap = new Map<THREE.Object3D, CANNON.Body>();

  constructor() {
    this.world = new CANNON.World();
    this.world.gravity.set(0, -9.82, 0);

    // ground
    const groundBody = new CANNON.Body({
      mass: 0,
      shape: new CANNON.Plane(),
    });
    groundBody.quaternion.setFromEuler(-Math.PI / 2, 0, 0);
    this.world.addBody(groundBody);
  }

  addMeshWithPhysics(mesh: THREE.Mesh, mass = 1) {
    const shape = this.createShapeFromMesh(mesh);
    const body = new CANNON.Body({
      mass,
      shape,
      position: new CANNON.Vec3(mesh.position.x, mesh.position.y, mesh.position.z),
    });
    this.world.addBody(body);
    this.meshBodyMap.set(mesh, body);
  }

  update(delta: number) {
    this.world.step(1 / 60, delta, 3);
    this.meshBodyMap.forEach((body, mesh) => {
      mesh.position.set(body.position.x, body.position.y, body.position.z);
      mesh.quaternion.set(
        body.quaternion.x,
        body.quaternion.y,
        body.quaternion.z,
        body.quaternion.w
      );
    });
  }

  private createShapeFromMesh(mesh: THREE.Mesh): CANNON.Shape {
    const box = new CANNON.Box(
      new CANNON.Vec3(
        mesh.scale.x * 0.5,
        mesh.scale.y * 0.5,
        mesh.scale.z * 0.5
      )
    );
    return box;
  }
}
```

---

### 4. `src/ar/setupAR.ts`

```ts
// src/ar/setupAR.ts
import { ARButton } from 'three/examples/jsm/webxr/ARButton.js';
import * as THREE from 'three';

export function enableAR(renderer: THREE.WebGLRenderer) {
  const arButton = ARButton.createButton(renderer, {
    requiredFeatures: ['hit-test'],
  });
  document.body.appendChild(arButton);
}
```

---

### 5. Update `main.ts` (Week 3 version)

```ts
// src/main.ts (Week 3)
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { createMultipleObjects } from './scene/createMultipleObjects';
import { startRenderLoop } from './loop/renderLoop';
import { setupResizeHandler } from './utils/resizeHandler';
import { InteractionManager } from './managers/InteractionManager';
import { PhysicsManager } from './physics/PhysicsManager';
import { enableAR } from './ar/setupAR';

function main() {
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);

  // VR + AR buttons
  document.body.appendChild(VRButton.createButton(renderer));
  enableAR(renderer);

  const { scene, camera } = setupScene();
  const objects = createMultipleObjects(scene);

  const interactionManager = new InteractionManager(scene, camera, renderer);
  interactionManager.setInteractable(objects);

  const physicsManager = new PhysicsManager();
  // add physics for each object
  objects.forEach((obj) => {
    if (obj instanceof THREE.Mesh) {
      physicsManager.addMeshWithPhysics(obj);
    }
  });

  setupResizeHandler(camera, renderer);

  // custom loop with physics update
  const clock = new THREE.Clock();
  renderer.setAnimationLoop(() => {
    const delta = clock.getDelta();
    physicsManager.update(delta);
    renderer.render(scene, camera);
  });
}

main();
```

---

### 6. Cursor prompt (Week 3)

```text
Extend the existing WebXR TS project with AR and physics.

Requirements:
- add src/physics/PhysicsManager.ts using cannon-es
- PhysicsManager should be able to register a Three.js mesh and keep it in sync with a physics body
- update main.ts to create a PhysicsManager and register all scene objects with it
- add src/ar/setupAR.ts that enables an immersive-ar session using ARButton
- in the render loop, update the physics world before rendering
- keep the project structure consistent (scene/, managers/, physics/, ar/, utils/)
```

---

## Week 4 – XR UI, Scene Persistence, Deployment

### Goal

Make it feel like a **real app**: add 3D UI, let the user add objects, reset, and **save** the scene.

### What you will learn

* Spatial / 3D UI patterns
* Managing scene state
* Persisting scene to LocalStorage
* Getting the project ready for deployment (HTTPS)

---

### 1. New folders / files

```text
src/
  ui/
    XRMenu.ts
  state/
    SceneStateManager.ts
```

---

### 2. `src/state/SceneStateManager.ts`

```ts
// src/state/SceneStateManager.ts
import * as THREE from 'three';

export interface SavedObject {
  type: 'cube' | 'sphere' | 'cone';
  position: { x: number; y: number; z: number };
  color: number;
}

const STORAGE_KEY = 'virtual-studio-xr-state';

export class SceneStateManager {
  save(objects: THREE.Mesh[]) {
    const data: SavedObject[] = objects.map((obj) => {
      const material = obj.material as THREE.MeshStandardMaterial;
      return {
        type: obj.userData.type ?? 'cube',
        position: {
          x: obj.position.x,
          y: obj.position.y,
          z: obj.position.z,
        },
        color: material.color.getHex(),
      };
    });
    localStorage.setItem(STORAGE_KEY, JSON.stringify(data));
  }

  load(): SavedObject[] {
    const raw = localStorage.getItem(STORAGE_KEY);
    if (!raw) return [];
    return JSON.parse(raw) as SavedObject[];
  }
}
```

---

### 3. `src/ui/XRMenu.ts` (simple 3D panel)

```ts
// src/ui/XRMenu.ts
import * as THREE from 'three';

export type XRMenuCallbacks = {
  onAdd?: () => void;
  onReset?: () => void;
  onSave?: () => void;
};

export function createXRMenu(callbacks: XRMenuCallbacks) {
  const group = new THREE.Group();
  group.position.set(0, 1.5, -1);

  const panelGeo = new THREE.PlaneGeometry(0.6, 0.3);
  const panelMat = new THREE.MeshBasicMaterial({ color: 0x222222, transparent: true, opacity: 0.7 });
  const panel = new THREE.Mesh(panelGeo, panelMat);
  group.add(panel);

  // very naive “buttons” – just smaller planes
  const buttonLabels = ['Add', 'Reset', 'Save'] as const;

  buttonLabels.forEach((label, index) => {
    const btnGeo = new THREE.PlaneGeometry(0.15, 0.08);
    const btnMat = new THREE.MeshBasicMaterial({ color: 0x555555 });
    const btn = new THREE.Mesh(btnGeo, btnMat);
    btn.position.set(-0.18 + index * 0.18, 0, 0.01);
    btn.userData.label = label;
    group.add(btn);
  });

  // you can later reuse the InteractionManager to detect clicks on these

  return group;
}
```

---

### 4. Update `main.ts` (Week 4 version)

```ts
// src/main.ts (Week 4)
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { createMultipleObjects } from './scene/createMultipleObjects';
import { setupResizeHandler } from './utils/resizeHandler';
import { InteractionManager } from './managers/InteractionManager';
import { PhysicsManager } from './physics/PhysicsManager';
import { enableAR } from './ar/setupAR';
import { SceneStateManager } from './state/SceneStateManager';
import { createXRMenu } from './ui/XRMenu';

function main() {
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);
  document.body.appendChild(VRButton.createButton(renderer));
  enableAR(renderer);

  const { scene, camera } = setupScene();
  const objects = createMultipleObjects(scene);

  const interactionManager = new InteractionManager(scene, camera, renderer);
  interactionManager.setInteractable(objects);

  const physicsManager = new PhysicsManager();
  objects.forEach((obj) => {
    if (obj instanceof THREE.Mesh) {
      physicsManager.addMeshWithPhysics(obj);
    }
  });

  const stateManager = new SceneStateManager();

  // XR menu with callbacks (placeholder)
  const menu = createXRMenu({
    onAdd: () => {
      // you would create a new mesh, add to scene, add to physics, add to interaction
    },
    onReset: () => {
      // remove objects from scene
    },
    onSave: () => {
      const meshes = scene.children.filter((c) => c instanceof THREE.Mesh) as THREE.Mesh[];
      stateManager.save(meshes);
    },
  });
  scene.add(menu);

  setupResizeHandler(camera, renderer);

  const clock = new THREE.Clock();
  renderer.setAnimationLoop(() => {
    const delta = clock.getDelta();
    physicsManager.update(delta);
    renderer.render(scene, camera);
  });
}

main();
```

---

### 5. Cursor prompt (Week 4)

```text
Extend the WebXR TS project to include an in-world XR menu and scene persistence.

Requirements:
- add src/state/SceneStateManager.ts that can save and load mesh data to localStorage
- add src/ui/XRMenu.ts that creates a 3D floating panel with 3 buttons: Add, Reset, Save
- integrate the menu into main.ts and hook the callbacks to scene/state logic
- keep the project modular (scene/, managers/, physics/, ar/, ui/, state/)
- code must remain in TypeScript
```

---

## Deployment Notes

* Vite already builds to static files: `npm run build`
* Deploy to **Vercel** / **Netlify** / **GitHub Pages**
* For XR on device → make sure it’s served over **HTTPS**

---

## Next Steps

* add **hand tracking** (if device supports)
* add **networking** (multi-user) via WebRTC / Colyseus
* add **better UI** via `three-mesh-ui`
* add **asset loading** (GLTF) instead of primitive geometries

---

That’s it — this is the full **TypeScript-first** README for your WebXR learning path, with explicit setup, folder structure, and incremental weekly additions.

```
```

