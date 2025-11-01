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
8. [Week 5 – Spatial Audio & Voice Commands](#week-5--spatial-audio--voice-commands)
9. [Week 6 – Networking & Multi-user XR](#week-6--networking--multi-user-xr)
10. [Week 7 – Asset Pipeline & Optimization](#week-7--asset-pipeline--optimization)
11. [Week 8 – UX & Accessibility in XR](#week-8--ux--accessibility-in-xr)
12. [Deployment Notes](#deployment-notes)
13. [Next Steps](#next-steps)

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

## Week 5 – Spatial Audio & Voice Commands

### Goal

Add **immersive audio** to the scene and enable **voice-based interaction** for hands-free control.

### What you will learn

* WebAudio API and positional audio in Three.js
* Creating audio sources that respond to distance and orientation
* Web Speech API for voice recognition
* Command parsing and scene control via voice
* Audio feedback for user actions

### Expected result

* Background ambient audio
* 3D positional sound sources attached to objects
* Voice commands like "add cube", "change color red", "save scene"
* Audio feedback when actions complete

---

### 1. New folders / files

```text
src/
  audio/
    AudioManager.ts
    SpatialAudioSource.ts
  voice/
    VoiceCommandManager.ts
```

---

### 2. Install audio assets (optional)

For this week, you can use:
- Web-generated tones
- Public domain audio from [Freesound.org](https://freesound.org)
- Procedural audio via Web Audio API

---

### 3. `src/audio/AudioManager.ts`

```ts
// src/audio/AudioManager.ts
import * as THREE from 'three';

export class AudioManager {
  private listener: THREE.AudioListener;
  private audioLoader: THREE.AudioLoader;
  private sounds: Map<string, THREE.Audio> = new Map();

  constructor(camera: THREE.PerspectiveCamera) {
    this.listener = new THREE.AudioListener();
    camera.add(this.listener);
    this.audioLoader = new THREE.AudioLoader();
  }

  async loadSound(name: string, url: string): Promise<THREE.Audio> {
    return new Promise((resolve, reject) => {
      const sound = new THREE.Audio(this.listener);
      this.audioLoader.load(
        url,
        (buffer) => {
          sound.setBuffer(buffer);
          this.sounds.set(name, sound);
          resolve(sound);
        },
        undefined,
        reject
      );
    });
  }

  playSound(name: string, loop = false, volume = 1) {
    const sound = this.sounds.get(name);
    if (sound && !sound.isPlaying) {
      sound.setLoop(loop);
      sound.setVolume(volume);
      sound.play();
    }
  }

  stopSound(name: string) {
    const sound = this.sounds.get(name);
    if (sound && sound.isPlaying) {
      sound.stop();
    }
  }

  // Create simple tone feedback
  createBeep(frequency = 440, duration = 0.1) {
    const context = this.listener.context;
    const oscillator = context.createOscillator();
    const gainNode = context.createGain();

    oscillator.connect(gainNode);
    gainNode.connect(context.destination);

    oscillator.frequency.value = frequency;
    oscillator.type = 'sine';

    gainNode.gain.setValueAtTime(0.3, context.currentTime);
    gainNode.gain.exponentialRampToValueAtTime(0.01, context.currentTime + duration);

    oscillator.start(context.currentTime);
    oscillator.stop(context.currentTime + duration);
  }
}
```

---

### 4. `src/audio/SpatialAudioSource.ts`

```ts
// src/audio/SpatialAudioSource.ts
import * as THREE from 'three';

export class SpatialAudioSource {
  private positionalAudio: THREE.PositionalAudio;

  constructor(
    listener: THREE.AudioListener,
    mesh: THREE.Object3D,
    audioBuffer: AudioBuffer
  ) {
    this.positionalAudio = new THREE.PositionalAudio(listener);
    this.positionalAudio.setBuffer(audioBuffer);
    this.positionalAudio.setRefDistance(1);
    this.positionalAudio.setRolloffFactor(1);
    this.positionalAudio.setDistanceModel('inverse');
    this.positionalAudio.setLoop(true);

    mesh.add(this.positionalAudio);
  }

  play() {
    if (!this.positionalAudio.isPlaying) {
      this.positionalAudio.play();
    }
  }

  stop() {
    if (this.positionalAudio.isPlaying) {
      this.positionalAudio.stop();
    }
  }

  setVolume(volume: number) {
    this.positionalAudio.setVolume(volume);
  }
}
```

---

### 5. `src/voice/VoiceCommandManager.ts`

```ts
// src/voice/VoiceCommandManager.ts
export type VoiceCommand = {
  phrase: string;
  action: () => void;
};

export class VoiceCommandManager {
  private recognition: SpeechRecognition | null = null;
  private commands: VoiceCommand[] = [];
  private isListening = false;

  constructor() {
    // Check for browser support
    const SpeechRecognition = 
      (window as any).SpeechRecognition || (window as any).webkitSpeechRecognition;

    if (SpeechRecognition) {
      this.recognition = new SpeechRecognition();
      this.recognition.continuous = true;
      this.recognition.interimResults = false;
      this.recognition.lang = 'en-US';

      this.recognition.onresult = (event) => {
        const last = event.results.length - 1;
        const transcript = event.results[last][0].transcript.toLowerCase().trim();
        console.log('Heard:', transcript);
        this.processCommand(transcript);
      };

      this.recognition.onerror = (event) => {
        console.error('Speech recognition error:', event.error);
      };
    } else {
      console.warn('Speech Recognition not supported in this browser');
    }
  }

  registerCommand(phrase: string, action: () => void) {
    this.commands.push({ phrase: phrase.toLowerCase(), action });
  }

  private processCommand(transcript: string) {
    for (const command of this.commands) {
      if (transcript.includes(command.phrase)) {
        console.log('Executing command:', command.phrase);
        command.action();
        return;
      }
    }
  }

  startListening() {
    if (this.recognition && !this.isListening) {
      this.recognition.start();
      this.isListening = true;
      console.log('Voice recognition started');
    }
  }

  stopListening() {
    if (this.recognition && this.isListening) {
      this.recognition.stop();
      this.isListening = false;
      console.log('Voice recognition stopped');
    }
  }
}
```

---

### 6. Update `src/main.ts` (Week 5 version)

```ts
// src/main.ts (Week 5)
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
import { AudioManager } from './audio/AudioManager';
import { VoiceCommandManager } from './voice/VoiceCommandManager';

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
  const audioManager = new AudioManager(camera);

  // Set up voice commands
  const voiceManager = new VoiceCommandManager();
  
  voiceManager.registerCommand('add cube', () => {
    console.log('Adding cube via voice');
    audioManager.createBeep(440, 0.1);
    // Add cube logic here
  });

  voiceManager.registerCommand('change color red', () => {
    console.log('Changing color to red');
    audioManager.createBeep(550, 0.1);
    objects[0].material.color.setHex(0xff0000);
  });

  voiceManager.registerCommand('save scene', () => {
    console.log('Saving scene via voice');
    audioManager.createBeep(660, 0.1);
    const meshes = scene.children.filter((c) => c instanceof THREE.Mesh) as THREE.Mesh[];
    stateManager.save(meshes);
  });

  // Button to start voice recognition (requires user gesture)
  const voiceButton = document.createElement('button');
  voiceButton.textContent = 'Start Voice Commands';
  voiceButton.style.position = 'absolute';
  voiceButton.style.top = '10px';
  voiceButton.style.left = '10px';
  voiceButton.style.padding = '10px';
  voiceButton.style.zIndex = '1000';
  voiceButton.onclick = () => {
    voiceManager.startListening();
    voiceButton.textContent = 'Voice Listening...';
    voiceButton.disabled = true;
  };
  document.body.appendChild(voiceButton);

  const menu = createXRMenu({
    onAdd: () => {
      audioManager.createBeep(440, 0.1);
    },
    onReset: () => {
      audioManager.createBeep(220, 0.2);
    },
    onSave: () => {
      audioManager.createBeep(880, 0.1);
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

### 7. Cursor prompt (Week 5)

```text
Extend the WebXR TS project to include spatial audio and voice commands.

Requirements:
- add src/audio/AudioManager.ts for managing audio sources and creating feedback sounds
- add src/audio/SpatialAudioSource.ts for 3D positional audio
- add src/voice/VoiceCommandManager.ts using Web Speech API for voice recognition
- register voice commands: "add cube", "change color red", "save scene"
- add audio feedback beeps when commands execute
- add a button to start voice recognition (requires user gesture)
- integrate AudioManager and VoiceCommandManager into main.ts
- keep the project modular (audio/, voice/, scene/, managers/, physics/, ar/, ui/, state/)
```

---

## Week 6 – Networking & Multi-user XR

### Goal

Enable **multiple users** to share the same XR space, see each other's avatars, and interact with shared objects in real-time.

### What you will learn

* WebSocket communication for real-time updates
* Broadcasting position/rotation of users and objects
* Creating simple avatar representations
* Handling network latency and state synchronization
* Optional: WebRTC for peer-to-peer communication

### Expected result

* Multiple users can join the same session
* Each user sees others as simple avatar meshes
* Shared object manipulation (pick up object → everyone sees it move)
* User ID/name display above avatars

---

### 1. Install networking dependencies

```bash
npm install socket.io-client
npm install -D @types/socket.io-client
```

For a simple local server, you can use:
```bash
npm install socket.io express
```

---

### 2. New folders / files

```text
src/
  networking/
    NetworkManager.ts
    Avatar.ts
server/
  server.js  (simple Node.js WebSocket server)
```

---

### 3. Simple WebSocket server (`server/server.js`)

```js
// server/server.js
const express = require('express');
const { createServer } = require('http');
const { Server } = require('socket.io');

const app = express();
const httpServer = createServer(app);
const io = new Server(httpServer, {
  cors: {
    origin: '*',
  },
});

const users = new Map();

io.on('connection', (socket) => {
  console.log('User connected:', socket.id);

  socket.on('join', (userData) => {
    users.set(socket.id, {
      id: socket.id,
      name: userData.name || `User_${socket.id.slice(0, 4)}`,
      position: { x: 0, y: 1.6, z: 0 },
      rotation: { x: 0, y: 0, z: 0 },
    });

    // Send existing users to new user
    socket.emit('existing-users', Array.from(users.values()));

    // Broadcast new user to others
    socket.broadcast.emit('user-joined', users.get(socket.id));
  });

  socket.on('update-position', (data) => {
    const user = users.get(socket.id);
    if (user) {
      user.position = data.position;
      user.rotation = data.rotation;
      socket.broadcast.emit('user-moved', {
        id: socket.id,
        position: data.position,
        rotation: data.rotation,
      });
    }
  });

  socket.on('object-moved', (data) => {
    socket.broadcast.emit('object-moved', data);
  });

  socket.on('disconnect', () => {
    console.log('User disconnected:', socket.id);
    users.delete(socket.id);
    socket.broadcast.emit('user-left', socket.id);
  });
});

httpServer.listen(3000, () => {
  console.log('Socket.io server running on http://localhost:3000');
});
```

---

### 4. `src/networking/Avatar.ts`

```ts
// src/networking/Avatar.ts
import * as THREE from 'three';

export class Avatar {
  public mesh: THREE.Group;
  private nameLabel: THREE.Sprite | null = null;

  constructor(public userId: string, public userName: string) {
    this.mesh = new THREE.Group();

    // Simple avatar: head + body
    const headGeo = new THREE.SphereGeometry(0.15, 16, 16);
    const bodyGeo = new THREE.CylinderGeometry(0.15, 0.2, 0.6, 16);
    const material = new THREE.MeshStandardMaterial({ color: 0x3388ff });

    const head = new THREE.Mesh(headGeo, material);
    head.position.y = 0.75;

    const body = new THREE.Mesh(bodyGeo, material);
    body.position.y = 0.3;

    this.mesh.add(head);
    this.mesh.add(body);

    // Optional: add name label above head
    this.createNameLabel(userName);
  }

  private createNameLabel(name: string) {
    const canvas = document.createElement('canvas');
    const context = canvas.getContext('2d')!;
    canvas.width = 256;
    canvas.height = 64;

    context.fillStyle = 'rgba(0, 0, 0, 0.6)';
    context.fillRect(0, 0, canvas.width, canvas.height);

    context.font = '32px Arial';
    context.fillStyle = 'white';
    context.textAlign = 'center';
    context.fillText(name, 128, 42);

    const texture = new THREE.CanvasTexture(canvas);
    const spriteMaterial = new THREE.SpriteMaterial({ map: texture });
    this.nameLabel = new THREE.Sprite(spriteMaterial);
    this.nameLabel.scale.set(1, 0.25, 1);
    this.nameLabel.position.y = 1.2;

    this.mesh.add(this.nameLabel);
  }

  updatePosition(position: { x: number; y: number; z: number }) {
    this.mesh.position.set(position.x, position.y, position.z);
  }

  updateRotation(rotation: { x: number; y: number; z: number }) {
    this.mesh.rotation.set(rotation.x, rotation.y, rotation.z);
  }
}
```

---

### 5. `src/networking/NetworkManager.ts`

```ts
// src/networking/NetworkManager.ts
import { io, Socket } from 'socket.io-client';
import * as THREE from 'three';
import { Avatar } from './Avatar';

export class NetworkManager {
  private socket: Socket;
  private avatars: Map<string, Avatar> = new Map();
  private scene: THREE.Scene;
  private localCamera: THREE.PerspectiveCamera;

  constructor(scene: THREE.Scene, camera: THREE.PerspectiveCamera, serverUrl = 'http://localhost:3000') {
    this.scene = scene;
    this.localCamera = camera;
    this.socket = io(serverUrl);

    this.setupSocketListeners();
  }

  private setupSocketListeners() {
    this.socket.on('connect', () => {
      console.log('Connected to server:', this.socket.id);
      this.socket.emit('join', { name: `User_${Date.now()}` });
    });

    this.socket.on('existing-users', (users: any[]) => {
      users.forEach((user) => {
        if (user.id !== this.socket.id) {
          this.createAvatar(user.id, user.name, user.position, user.rotation);
        }
      });
    });

    this.socket.on('user-joined', (user: any) => {
      console.log('User joined:', user.name);
      this.createAvatar(user.id, user.name, user.position, user.rotation);
    });

    this.socket.on('user-moved', (data: any) => {
      const avatar = this.avatars.get(data.id);
      if (avatar) {
        avatar.updatePosition(data.position);
        avatar.updateRotation(data.rotation);
      }
    });

    this.socket.on('user-left', (userId: string) => {
      console.log('User left:', userId);
      this.removeAvatar(userId);
    });

    this.socket.on('object-moved', (data: any) => {
      // Handle shared object movement
      const obj = this.scene.getObjectByName(data.objectName);
      if (obj) {
        obj.position.set(data.position.x, data.position.y, data.position.z);
      }
    });
  }

  private createAvatar(userId: string, userName: string, position: any, rotation: any) {
    const avatar = new Avatar(userId, userName);
    avatar.updatePosition(position);
    avatar.updateRotation(rotation);
    this.avatars.set(userId, avatar);
    this.scene.add(avatar.mesh);
  }

  private removeAvatar(userId: string) {
    const avatar = this.avatars.get(userId);
    if (avatar) {
      this.scene.remove(avatar.mesh);
      this.avatars.delete(userId);
    }
  }

  updateLocalPosition() {
    const pos = this.localCamera.position;
    const rot = this.localCamera.rotation;
    this.socket.emit('update-position', {
      position: { x: pos.x, y: pos.y, z: pos.z },
      rotation: { x: rot.x, y: rot.y, z: rot.z },
    });
  }

  broadcastObjectMove(objectName: string, position: THREE.Vector3) {
    this.socket.emit('object-moved', {
      objectName,
      position: { x: position.x, y: position.y, z: position.z },
    });
  }
}
```

---

### 6. Update `src/main.ts` (Week 6 version)

```ts
// src/main.ts (Week 6)
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { createMultipleObjects } from './scene/createMultipleObjects';
import { setupResizeHandler } from './utils/resizeHandler';
import { NetworkManager } from './networking/NetworkManager';

function main() {
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);
  document.body.appendChild(VRButton.createButton(renderer));

  const { scene, camera } = setupScene();
  const objects = createMultipleObjects(scene);

  // Initialize networking
  const networkManager = new NetworkManager(scene, camera);

  setupResizeHandler(camera, renderer);

  const clock = new THREE.Clock();
  let lastNetworkUpdate = 0;

  renderer.setAnimationLoop(() => {
    const time = clock.getElapsedTime();

    // Send position updates every 100ms
    if (time - lastNetworkUpdate > 0.1) {
      networkManager.updateLocalPosition();
      lastNetworkUpdate = time;
    }

    renderer.render(scene, camera);
  });
}

main();
```

---

### 7. Running the server

Add to `package.json`:
```json
"scripts": {
  "dev": "vite",
  "server": "node server/server.js",
  "dev:all": "concurrently \"npm run dev\" \"npm run server\""
}
```

Install concurrently:
```bash
npm install -D concurrently
```

Then run: `npm run dev:all`

---

### 8. Cursor prompt (Week 6)

```text
Extend the WebXR TS project to support multi-user networking.

Requirements:
- create server/server.js using Socket.io for WebSocket communication
- add src/networking/NetworkManager.ts to handle socket events
- add src/networking/Avatar.ts to represent remote users
- implement position/rotation broadcasting every 100ms
- handle user join/leave events
- create simple avatar meshes (sphere head + cylinder body) with name labels
- update main.ts to initialize NetworkManager
- broadcast local camera position to other users
- keep the project modular
```

---

## Week 7 – Asset Pipeline & Optimization

### Goal

Load **realistic 3D models** (GLTF/GLB) and optimize the scene for **smooth performance** on XR devices.

### What you will learn

* Loading and parsing GLTF/GLB models
* Asset management and caching
* Level of Detail (LOD) techniques
* Draw call optimization (instancing, merging)
* Texture compression and optimization
* Performance profiling for XR (90fps target)

### Expected result

* Load complex GLTF models instead of primitive shapes
* Implement LOD for distant objects
* Reduce draw calls from 100+ to <50
* Maintain 72-90 FPS on Quest devices
* Loading screen with progress indicator

---

### 1. Install GLTF loader (included in Three.js)

No extra dependencies needed! GLTFLoader is in Three.js examples.

---

### 2. New folders / files

```text
src/
  assets/
    AssetManager.ts
    GLTFModelLoader.ts
  optimization/
    LODManager.ts
    DrawCallOptimizer.ts
  ui/
    LoadingScreen.ts
public/
  models/
    (place your .gltf/.glb files here)
```

---

### 3. `src/assets/GLTFModelLoader.ts`

```ts
// src/assets/GLTFModelLoader.ts
import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';
import { DRACOLoader } from 'three/examples/jsm/loaders/DRACOLoader.js';

export class GLTFModelLoader {
  private loader: GLTFLoader;
  private dracoLoader: DRACOLoader;

  constructor() {
    this.loader = new GLTFLoader();
    
    // Optional: Enable Draco compression for smaller files
    this.dracoLoader = new DRACOLoader();
    this.dracoLoader.setDecoderPath('https://www.gstatic.com/draco/v1/decoders/');
    this.loader.setDRACOLoader(this.dracoLoader);
  }

  async load(url: string, onProgress?: (progress: number) => void): Promise<THREE.Group> {
    return new Promise((resolve, reject) => {
      this.loader.load(
        url,
        (gltf) => {
          resolve(gltf.scene);
        },
        (xhr) => {
          const progress = (xhr.loaded / xhr.total) * 100;
          if (onProgress) onProgress(progress);
        },
        reject
      );
    });
  }

  dispose() {
    this.dracoLoader.dispose();
  }
}
```

---

### 4. `src/assets/AssetManager.ts`

```ts
// src/assets/AssetManager.ts
import * as THREE from 'three';
import { GLTFModelLoader } from './GLTFModelLoader';

export class AssetManager {
  private modelLoader: GLTFModelLoader;
  private cache: Map<string, THREE.Group> = new Map();
  private textureLoader: THREE.TextureLoader;

  constructor() {
    this.modelLoader = new GLTFModelLoader();
    this.textureLoader = new THREE.TextureLoader();
  }

  async loadModel(
    name: string,
    url: string,
    onProgress?: (progress: number) => void
  ): Promise<THREE.Group> {
    // Check cache first
    if (this.cache.has(name)) {
      return this.cache.get(name)!.clone();
    }

    const model = await this.modelLoader.load(url, onProgress);
    this.cache.set(name, model);
    return model.clone();
  }

  async loadTexture(url: string): Promise<THREE.Texture> {
    return new Promise((resolve, reject) => {
      this.textureLoader.load(url, resolve, undefined, reject);
    });
  }

  preloadAssets(assets: { name: string; url: string }[]): Promise<void[]> {
    const promises = assets.map((asset) => this.loadModel(asset.name, asset.url));
    return Promise.all(promises);
  }

  getCachedModel(name: string): THREE.Group | undefined {
    const model = this.cache.get(name);
    return model ? model.clone() : undefined;
  }

  dispose() {
    this.cache.clear();
    this.modelLoader.dispose();
  }
}
```

---

### 5. `src/optimization/LODManager.ts`

```ts
// src/optimization/LODManager.ts
import * as THREE from 'three';

export class LODManager {
  createLOD(
    highDetail: THREE.Object3D,
    mediumDetail: THREE.Object3D,
    lowDetail: THREE.Object3D
  ): THREE.LOD {
    const lod = new THREE.LOD();
    
    lod.addLevel(highDetail, 0);    // 0-10 units
    lod.addLevel(mediumDetail, 10); // 10-30 units
    lod.addLevel(lowDetail, 30);    // 30+ units

    return lod;
  }

  // Simplify mesh by reducing geometry
  simplifyMesh(mesh: THREE.Mesh, targetRatio = 0.5): THREE.Mesh {
    if (!mesh.geometry) return mesh;

    // Simple approach: create lower poly version
    // For production, use libraries like three-simplify or SimplifyModifier
    const geometry = mesh.geometry.clone();
    
    // This is a placeholder - in production use proper decimation
    const simplified = new THREE.Mesh(geometry, mesh.material);
    simplified.scale.copy(mesh.scale);
    
    return simplified;
  }

  autoGenerateLOD(object: THREE.Object3D): THREE.LOD {
    const lod = new THREE.LOD();
    
    const high = object.clone();
    const medium = object.clone();
    const low = object.clone();

    // Scale down for simple LOD effect
    medium.scale.multiplyScalar(0.8);
    low.scale.multiplyScalar(0.5);

    lod.addLevel(high, 0);
    lod.addLevel(medium, 15);
    lod.addLevel(low, 40);

    return lod;
  }
}
```

---

### 6. `src/optimization/DrawCallOptimizer.ts`

```ts
// src/optimization/DrawCallOptimizer.ts
import * as THREE from 'three';

export class DrawCallOptimizer {
  // Merge multiple meshes with same material into one
  mergeGeometries(meshes: THREE.Mesh[]): THREE.Mesh | null {
    if (meshes.length === 0) return null;

    const geometries: THREE.BufferGeometry[] = [];
    const material = meshes[0].material;

    meshes.forEach((mesh) => {
      const geo = mesh.geometry.clone();
      geo.applyMatrix4(mesh.matrixWorld);
      geometries.push(geo);
    });

    const mergedGeometry = THREE.BufferGeometryUtils.mergeGeometries(geometries);
    if (!mergedGeometry) return null;

    return new THREE.Mesh(mergedGeometry, material);
  }

  // Create instanced mesh for repeated objects
  createInstancedMesh(
    geometry: THREE.BufferGeometry,
    material: THREE.Material,
    positions: THREE.Vector3[]
  ): THREE.InstancedMesh {
    const count = positions.length;
    const instancedMesh = new THREE.InstancedMesh(geometry, material, count);

    const matrix = new THREE.Matrix4();
    positions.forEach((pos, i) => {
      matrix.setPosition(pos);
      instancedMesh.setMatrixAt(i, matrix);
    });

    instancedMesh.instanceMatrix.needsUpdate = true;
    return instancedMesh;
  }

  // Get draw call count (for debugging)
  getDrawCallCount(renderer: THREE.WebGLRenderer): number {
    return renderer.info.render.calls;
  }

  // Print performance stats
  logPerformanceStats(renderer: THREE.WebGLRenderer) {
    const info = renderer.info;
    console.log('=== Render Stats ===');
    console.log('Draw Calls:', info.render.calls);
    console.log('Triangles:', info.render.triangles);
    console.log('Points:', info.render.points);
    console.log('Lines:', info.render.lines);
    console.log('Textures:', info.memory.textures);
    console.log('Geometries:', info.memory.geometries);
  }
}
```

---

### 7. `src/ui/LoadingScreen.ts`

```ts
// src/ui/LoadingScreen.ts
export class LoadingScreen {
  private container: HTMLDivElement;
  private progressBar: HTMLDivElement;
  private progressText: HTMLDivElement;

  constructor() {
    this.container = document.createElement('div');
    this.container.style.cssText = `
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: #000;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      z-index: 9999;
    `;

    this.progressText = document.createElement('div');
    this.progressText.style.cssText = `
      color: white;
      font-size: 24px;
      margin-bottom: 20px;
    `;
    this.progressText.textContent = 'Loading... 0%';

    this.progressBar = document.createElement('div');
    this.progressBar.style.cssText = `
      width: 300px;
      height: 20px;
      background: #333;
      border-radius: 10px;
      overflow: hidden;
    `;

    const progressFill = document.createElement('div');
    progressFill.id = 'progress-fill';
    progressFill.style.cssText = `
      width: 0%;
      height: 100%;
      background: linear-gradient(90deg, #00aaff, #00ff88);
      transition: width 0.3s;
    `;

    this.progressBar.appendChild(progressFill);
    this.container.appendChild(this.progressText);
    this.container.appendChild(this.progressBar);
    document.body.appendChild(this.container);
  }

  updateProgress(progress: number) {
    const fill = document.getElementById('progress-fill');
    if (fill) {
      fill.style.width = `${progress}%`;
      this.progressText.textContent = `Loading... ${Math.round(progress)}%`;
    }
  }

  hide() {
    this.container.style.display = 'none';
  }

  show() {
    this.container.style.display = 'flex';
  }

  remove() {
    document.body.removeChild(this.container);
  }
}
```

---

### 8. Update `src/main.ts` (Week 7 version)

```ts
// src/main.ts (Week 7)
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { setupResizeHandler } from './utils/resizeHandler';
import { AssetManager } from './assets/AssetManager';
import { LODManager } from './optimization/LODManager';
import { DrawCallOptimizer } from './optimization/DrawCallOptimizer';
import { LoadingScreen } from './ui/LoadingScreen';

async function main() {
  const loadingScreen = new LoadingScreen();
  
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);
  document.body.appendChild(VRButton.createButton(renderer));

  const { scene, camera } = setupScene();
  const assetManager = new AssetManager();
  const lodManager = new LODManager();
  const optimizer = new DrawCallOptimizer();

  try {
    // Load assets with progress
    loadingScreen.updateProgress(10);
    
    // Example: Load a GLTF model
    // const model = await assetManager.loadModel(
    //   'room',
    //   '/models/room.glb',
    //   (progress) => loadingScreen.updateProgress(10 + progress * 0.8)
    // );
    // scene.add(model);

    // For demo: create some objects with LOD
    const geometry = new THREE.BoxGeometry(1, 1, 1);
    const material = new THREE.MeshStandardMaterial({ color: 0x00aaff });
    
    for (let i = 0; i < 5; i++) {
      const mesh = new THREE.Mesh(geometry, material);
      mesh.position.set(i * 3 - 6, 1, -5);
      const lod = lodManager.autoGenerateLOD(mesh);
      scene.add(lod);
    }

    loadingScreen.updateProgress(100);
    setTimeout(() => loadingScreen.remove(), 500);

  } catch (error) {
    console.error('Error loading assets:', error);
    loadingScreen.remove();
  }

  setupResizeHandler(camera, renderer);

  const clock = new THREE.Clock();
  let frameCount = 0;

  renderer.setAnimationLoop(() => {
    frameCount++;
    
    // Log performance stats every 300 frames (~5 seconds at 60fps)
    if (frameCount % 300 === 0) {
      optimizer.logPerformanceStats(renderer);
    }

    renderer.render(scene, camera);
  });
}

main();
```

---

### 9. Cursor prompt (Week 7)

```text
Extend the WebXR TS project to support GLTF asset loading and performance optimization.

Requirements:
- add src/assets/GLTFModelLoader.ts using GLTFLoader with Draco support
- add src/assets/AssetManager.ts for caching and managing loaded assets
- add src/optimization/LODManager.ts for Level of Detail management
- add src/optimization/DrawCallOptimizer.ts for merging geometries and instancing
- add src/ui/LoadingScreen.ts with progress bar
- update main.ts to load assets with loading screen
- implement LOD for multiple objects in scene
- log performance stats (draw calls, triangles) periodically
- keep the project modular
```

---

## Week 8 – UX & Accessibility in XR

### Goal

Implement **best practices** for comfort, usability, and accessibility in XR experiences.

### What you will learn

* Motion comfort techniques (vignetting, teleportation)
* Gaze-based UI for controller-free interaction
* Accessibility: colorblind modes, text scaling, audio cues
* Performance targets for different devices
* UX patterns: onboarding, tutorials, error handling

### Expected result

* Smooth locomotion options (teleport + smooth movement)
* Comfort vignette during movement
* Gaze cursor for UI interaction
* Configurable accessibility settings
* Tutorial overlay for first-time users

---

### 1. New folders / files

```text
src/
  locomotion/
    TeleportController.ts
    SmoothLocomotion.ts
  comfort/
    VignetteEffect.ts
    ComfortSettings.ts
  accessibility/
    AccessibilityManager.ts
    GazeController.ts
  tutorial/
    TutorialManager.ts
```

---

### 2. `src/locomotion/TeleportController.ts`

```ts
// src/locomotion/TeleportController.ts
import * as THREE from 'three';

export class TeleportController {
  private scene: THREE.Scene;
  private camera: THREE.PerspectiveCamera;
  private marker: THREE.Mesh;
  private raycaster: THREE.Raycaster;
  private isActive = false;

  constructor(scene: THREE.Scene, camera: THREE.PerspectiveCamera) {
    this.scene = scene;
    this.camera = camera;
    this.raycaster = new THREE.Raycaster();

    // Create teleport marker
    const geometry = new THREE.CircleGeometry(0.5, 32);
    const material = new THREE.MeshBasicMaterial({ 
      color: 0x00ff00, 
      transparent: true, 
      opacity: 0.5 
    });
    this.marker = new THREE.Mesh(geometry, material);
    this.marker.rotation.x = -Math.PI / 2;
    this.marker.visible = false;
    this.scene.add(this.marker);
  }

  showMarker(position: THREE.Vector3) {
    this.marker.position.copy(position);
    this.marker.visible = true;
  }

  hideMarker() {
    this.marker.visible = false;
  }

  teleport(position: THREE.Vector3) {
    // Keep camera height, only move X and Z
    this.camera.position.x = position.x;
    this.camera.position.z = position.z;
    this.hideMarker();
  }

  update(controller: THREE.Group, ground: THREE.Object3D) {
    const tempMatrix = new THREE.Matrix4();
    tempMatrix.identity().extractRotation(controller.matrixWorld);

    const origin = controller.position.clone();
    const direction = new THREE.Vector3(0, 0, -1).applyMatrix4(tempMatrix);

    this.raycaster.set(origin, direction);
    const intersects = this.raycaster.intersectObject(ground, true);

    if (intersects.length > 0) {
      this.showMarker(intersects[0].point);
      return intersects[0].point;
    } else {
      this.hideMarker();
      return null;
    }
  }
}
```

---

### 3. `src/comfort/VignetteEffect.ts`

```ts
// src/comfort/VignetteEffect.ts
import * as THREE from 'three';

export class VignetteEffect {
  private material: THREE.ShaderMaterial;
  private mesh: THREE.Mesh;
  private targetIntensity = 0;
  private currentIntensity = 0;

  constructor(camera: THREE.PerspectiveCamera) {
    const geometry = new THREE.PlaneGeometry(2, 2);
    
    this.material = new THREE.ShaderMaterial({
      transparent: true,
      depthTest: false,
      depthWrite: false,
      vertexShader: `
        varying vec2 vUv;
        void main() {
          vUv = uv;
          gl_Position = vec4(position, 1.0);
        }
      `,
      fragmentShader: `
        uniform float intensity;
        varying vec2 vUv;
        
        void main() {
          vec2 center = vec2(0.5);
          float dist = distance(vUv, center);
          float vignette = smoothstep(0.3, 1.0, dist);
          gl_FragColor = vec4(0.0, 0.0, 0.0, vignette * intensity);
        }
      `,
      uniforms: {
        intensity: { value: 0.0 }
      }
    });

    this.mesh = new THREE.Mesh(geometry, this.material);
    this.mesh.frustumCulled = false;
    this.mesh.renderOrder = 999;
    
    camera.add(this.mesh);
    this.mesh.position.z = -0.1;
  }

  setIntensity(value: number) {
    this.targetIntensity = THREE.MathUtils.clamp(value, 0, 1);
  }

  update(delta: number) {
    // Smooth transition
    this.currentIntensity = THREE.MathUtils.lerp(
      this.currentIntensity,
      this.targetIntensity,
      delta * 5
    );
    this.material.uniforms.intensity.value = this.currentIntensity;
  }
}
```

---

### 4. `src/accessibility/AccessibilityManager.ts`

```ts
// src/accessibility/AccessibilityManager.ts
import * as THREE from 'three';

export type AccessibilitySettings = {
  colorblindMode: 'none' | 'protanopia' | 'deuteranopia' | 'tritanopia';
  textScale: number;
  audioDescriptions: boolean;
  reducedMotion: boolean;
  highContrast: boolean;
};

export class AccessibilityManager {
  private settings: AccessibilitySettings = {
    colorblindMode: 'none',
    textScale: 1.0,
    audioDescriptions: false,
    reducedMotion: false,
    highContrast: false,
  };

  constructor() {
    this.loadSettings();
  }

  private loadSettings() {
    const saved = localStorage.getItem('accessibility-settings');
    if (saved) {
      this.settings = { ...this.settings, ...JSON.parse(saved) };
    }
  }

  saveSettings() {
    localStorage.setItem('accessibility-settings', JSON.stringify(this.settings));
  }

  setColorblindMode(mode: AccessibilitySettings['colorblindMode']) {
    this.settings.colorblindMode = mode;
    this.saveSettings();
  }

  setTextScale(scale: number) {
    this.settings.textScale = THREE.MathUtils.clamp(scale, 0.5, 2.0);
    this.saveSettings();
  }

  setReducedMotion(enabled: boolean) {
    this.settings.reducedMotion = enabled;
    this.saveSettings();
  }

  getSettings(): AccessibilitySettings {
    return { ...this.settings };
  }

  // Apply color filter for colorblind mode
  applyColorblindFilter(color: THREE.Color): THREE.Color {
    if (this.settings.colorblindMode === 'none') return color;

    // Simplified colorblind simulation
    const rgb = { r: color.r, g: color.g, b: color.b };

    switch (this.settings.colorblindMode) {
      case 'protanopia': // Red-blind
        return new THREE.Color(
          0.567 * rgb.r + 0.433 * rgb.g,
          0.558 * rgb.r + 0.442 * rgb.g,
          0.242 * rgb.g + 0.758 * rgb.b
        );
      case 'deuteranopia': // Green-blind
        return new THREE.Color(
          0.625 * rgb.r + 0.375 * rgb.g,
          0.7 * rgb.r + 0.3 * rgb.g,
          0.3 * rgb.g + 0.7 * rgb.b
        );
      case 'tritanopia': // Blue-blind
        return new THREE.Color(
          0.95 * rgb.r + 0.05 * rgb.g,
          0.433 * rgb.g + 0.567 * rgb.b,
          0.475 * rgb.g + 0.525 * rgb.b
        );
      default:
        return color;
    }
  }
}
```

---

### 5. `src/accessibility/GazeController.ts`

```ts
// src/accessibility/GazeController.ts
import * as THREE from 'three';

export class GazeController {
  private raycaster: THREE.Raycaster;
  private cursor: THREE.Mesh;
  private gazeDuration = 0;
  private gazeThreshold = 2.0; // seconds to activate
  private currentTarget: THREE.Object3D | null = null;
  private onGazeSelect?: (object: THREE.Object3D) => void;

  constructor(scene: THREE.Scene, camera: THREE.PerspectiveCamera) {
    this.raycaster = new THREE.Raycaster();

    // Create gaze cursor
    const geometry = new THREE.RingGeometry(0.015, 0.02, 32);
    const material = new THREE.MeshBasicMaterial({ 
      color: 0xffffff, 
      transparent: true,
      opacity: 0.8
    });
    this.cursor = new THREE.Mesh(geometry, material);
    camera.add(this.cursor);
    this.cursor.position.z = -1;
  }

  setOnGazeSelect(callback: (object: THREE.Object3D) => void) {
    this.onGazeSelect = callback;
  }

  update(camera: THREE.PerspectiveCamera, interactables: THREE.Object3D[], delta: number) {
    // Raycast from center of view
    this.raycaster.setFromCamera(new THREE.Vector2(0, 0), camera);
    const intersects = this.raycaster.intersectObjects(interactables, true);

    if (intersects.length > 0) {
      const target = intersects[0].object;

      if (target === this.currentTarget) {
        // Continue gazing at same object
        this.gazeDuration += delta;
        
        // Visual feedback: scale cursor based on gaze progress
        const progress = Math.min(this.gazeDuration / this.gazeThreshold, 1);
        this.cursor.scale.setScalar(1 + progress * 0.5);
        (this.cursor.material as THREE.MeshBasicMaterial).color.setHex(
          progress < 1 ? 0xffffff : 0x00ff00
        );

        // Trigger selection
        if (this.gazeDuration >= this.gazeThreshold && this.onGazeSelect) {
          this.onGazeSelect(target);
          this.gazeDuration = 0;
        }
      } else {
        // New target
        this.currentTarget = target;
        this.gazeDuration = 0;
        this.cursor.scale.setScalar(1);
      }
    } else {
      // No target
      this.currentTarget = null;
      this.gazeDuration = 0;
      this.cursor.scale.setScalar(1);
      (this.cursor.material as THREE.MeshBasicMaterial).color.setHex(0xffffff);
    }
  }
}
```

---

### 6. `src/tutorial/TutorialManager.ts`

```ts
// src/tutorial/TutorialManager.ts
export class TutorialManager {
  private steps: string[] = [];
  private currentStep = 0;
  private overlay: HTMLDivElement;
  private onComplete?: () => void;

  constructor(steps: string[]) {
    this.steps = steps;

    this.overlay = document.createElement('div');
    this.overlay.style.cssText = `
      position: fixed;
      bottom: 20px;
      left: 50%;
      transform: translateX(-50%);
      background: rgba(0, 0, 0, 0.8);
      color: white;
      padding: 20px 30px;
      border-radius: 10px;
      font-size: 18px;
      max-width: 600px;
      text-align: center;
      z-index: 1000;
    `;
  }

  start(onComplete?: () => void) {
    this.onComplete = onComplete;
    this.currentStep = 0;
    this.showStep();
    document.body.appendChild(this.overlay);
  }

  private showStep() {
    if (this.currentStep < this.steps.length) {
      this.overlay.innerHTML = `
        <div>${this.steps[this.currentStep]}</div>
        <div style="margin-top: 15px; font-size: 14px; color: #888;">
          Step ${this.currentStep + 1} of ${this.steps.length}
        </div>
      `;
    } else {
      this.complete();
    }
  }

  nextStep() {
    this.currentStep++;
    if (this.currentStep < this.steps.length) {
      this.showStep();
    } else {
      this.complete();
    }
  }

  private complete() {
    document.body.removeChild(this.overlay);
    if (this.onComplete) {
      this.onComplete();
    }
  }

  skip() {
    this.complete();
  }
}
```

---

### 7. Update `src/main.ts` (Week 8 version)

```ts
// src/main.ts (Week 8)
import * as THREE from 'three';
import { VRButton } from 'three/examples/jsm/webxr/VRButton.js';
import { setupScene } from './scene/setupScene';
import { createMultipleObjects } from './scene/createMultipleObjects';
import { setupResizeHandler } from './utils/resizeHandler';
import { TeleportController } from './locomotion/TeleportController';
import { VignetteEffect } from './comfort/VignetteEffect';
import { AccessibilityManager } from './accessibility/AccessibilityManager';
import { GazeController } from './accessibility/GazeController';
import { TutorialManager } from './tutorial/TutorialManager';

function main() {
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.xr.enabled = true;
  document.body.appendChild(renderer.domElement);
  document.body.appendChild(VRButton.createButton(renderer));

  const { scene, camera } = setupScene();
  const objects = createMultipleObjects(scene);

  // Add ground for teleportation
  const groundGeo = new THREE.PlaneGeometry(20, 20);
  const groundMat = new THREE.MeshStandardMaterial({ color: 0x808080 });
  const ground = new THREE.Mesh(groundGeo, groundMat);
  ground.rotation.x = -Math.PI / 2;
  ground.name = 'ground';
  scene.add(ground);

  // Initialize systems
  const teleportController = new TeleportController(scene, camera);
  const vignetteEffect = new VignetteEffect(camera);
  const accessibilityManager = new AccessibilityManager();
  const gazeController = new GazeController(scene, camera);

  // Tutorial
  const tutorial = new TutorialManager([
    '👀 Look around to explore the scene',
    '🎮 Use your controller to interact with objects',
    '✨ Gaze at objects for 2 seconds to select them',
    '🚀 Point at the ground and trigger to teleport',
    '⚙️ Press Menu for accessibility settings'
  ]);

  // Start tutorial on first visit
  if (!localStorage.getItem('tutorial-completed')) {
    setTimeout(() => {
      tutorial.start(() => {
        localStorage.setItem('tutorial-completed', 'true');
      });
    }, 2000);
  }

  // Gaze interaction
  gazeController.setOnGazeSelect((object) => {
    if (object instanceof THREE.Mesh) {
      const settings = accessibilityManager.getSettings();
      let newColor = new THREE.Color(Math.random() * 0xffffff);
      newColor = accessibilityManager.applyColorblindFilter(newColor);
      (object.material as THREE.MeshStandardMaterial).color.copy(newColor);
    }
  });

  // Controller for teleportation
  const controller = renderer.xr.getController(0);
  controller.addEventListener('selectstart', () => {
    const pos = teleportController.update(controller, ground);
    if (pos) {
      teleportController.teleport(pos);
      vignetteEffect.setIntensity(0.7); // Show vignette during teleport
      setTimeout(() => vignetteEffect.setIntensity(0), 300);
    }
  });
  scene.add(controller);

  setupResizeHandler(camera, renderer);

  const clock = new THREE.Clock();
  renderer.setAnimationLoop(() => {
    const delta = clock.getDelta();
    
    vignetteEffect.update(delta);
    gazeController.update(camera, objects, delta);
    teleportController.update(controller, ground);
    
    renderer.render(scene, camera);
  });
}

main();
```

---

### 8. Cursor prompt (Week 8)

```text
Extend the WebXR TS project with UX and accessibility features.

Requirements:
- add src/locomotion/TeleportController.ts for teleportation with ground marker
- add src/comfort/VignetteEffect.ts for motion comfort vignette shader
- add src/accessibility/AccessibilityManager.ts with colorblind modes and settings
- add src/accessibility/GazeController.ts for gaze-based interaction (2s dwell time)
- add src/tutorial/TutorialManager.ts for onboarding steps
- update main.ts to integrate all systems
- show tutorial overlay on first visit
- implement gaze selection with visual feedback on cursor
- apply vignette during teleportation for comfort
- keep the project modular
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

