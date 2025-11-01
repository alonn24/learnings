# Virtual Studio XR - Project Progress

## Week 1 – Project Setup & Basic WebXR Scene ✅

**Goal**: Set up a modern TS project that can render a 3D scene and enter VR with a single button.

### Tasks:
- [x] Create Vite + TypeScript project structure
- [x] Install Three.js and @types/three
- [x] Create folder structure (scene/, loop/, utils/)
- [x] Create `src/main.ts` - app entrypoint with VRButton
- [x] Create `src/scene/setupScene.ts` - scene, camera, lights
- [x] Create `src/scene/createObjects.ts` - spinning cube
- [x] Create `src/loop/renderLoop.ts` - XR-aware animation loop
- [x] Create `src/utils/resizeHandler.ts` - window resize handler
- [x] Enable WebXR (renderer.xr.enabled = true)
- [x] Add VRButton to DOM
- [x] Run dev server and verify: grey background, spinning cube, VR button

**Status**: ✅ Complete - Verified working with screenshot

---

## Week 2 – Interaction & Controllers ⏳

**Goal**: Add input to the scene - VR controller → raycast → select object → change it.

### Tasks:
- [ ] Create `src/managers/` folder
- [ ] Create `src/managers/InteractionManager.ts`
  - [ ] Set up XR controller via `renderer.xr.getController(0)`
  - [ ] Implement raycasting with THREE.Raycaster
  - [ ] Add selectstart event listener
  - [ ] Implement color change on object selection
- [ ] Create `src/scene/createMultipleObjects.ts`
  - [ ] Add cube mesh
  - [ ] Add sphere mesh
  - [ ] Add cone mesh
  - [ ] Return array of interactable objects
- [ ] Update `src/main.ts` to use InteractionManager
- [ ] Register created objects as interactable
- [ ] Test: Point at objects with controller and change color on select

**Expected Result**: Scene has 3 objects (cube, sphere, cone) that change color when selected with controller

---

## Week 3 – AR, Lighting, and Physics ⏳

**Goal**: Add AR mode and basic physics so objects can rest on a surface and look more real.

### Tasks:
- [ ] Install cannon-es physics engine (`npm install cannon-es`)
- [ ] Create `src/physics/` folder
- [ ] Create `src/physics/PhysicsManager.ts`
  - [ ] Set up CANNON.World with gravity
  - [ ] Create ground plane
  - [ ] Implement addMeshWithPhysics method
  - [ ] Implement update method to sync physics → Three.js
  - [ ] Create shape from mesh helper method
- [ ] Create `src/ar/` folder
- [ ] Create `src/ar/setupAR.ts`
  - [ ] Enable AR session using ARButton
  - [ ] Configure required features (hit-test)
- [ ] Update `src/main.ts` (Week 3 version)
  - [ ] Add AR button
  - [ ] Create PhysicsManager instance
  - [ ] Register all objects with physics
  - [ ] Update render loop to step physics world
- [ ] Test: Objects should fall and rest on ground plane
- [ ] Test: AR button appears (requires HTTPS + supported device)

**Expected Result**: Objects have physics simulation, AR mode available

---

## Week 4 – XR UI, Scene Persistence, Deployment ⏳

**Goal**: Make it feel like a real app with 3D UI, object creation, and scene saving.

### Tasks:
- [ ] Create `src/state/` folder
- [ ] Create `src/state/SceneStateManager.ts`
  - [ ] Define SavedObject interface
  - [ ] Implement save method (to localStorage)
  - [ ] Implement load method (from localStorage)
  - [ ] Handle object type, position, and color
- [ ] Create `src/ui/` folder
- [ ] Create `src/ui/XRMenu.ts`
  - [ ] Create 3D panel background
  - [ ] Add "Add" button mesh
  - [ ] Add "Reset" button mesh
  - [ ] Add "Save" button mesh
  - [ ] Define XRMenuCallbacks type
  - [ ] Return menu group positioned in space
- [ ] Update `src/main.ts` (Week 4 version)
  - [ ] Create SceneStateManager instance
  - [ ] Create XR menu with callbacks
  - [ ] Implement onAdd callback (create new mesh)
  - [ ] Implement onReset callback (clear scene)
  - [ ] Implement onSave callback (persist to localStorage)
  - [ ] Add menu to scene
  - [ ] Optional: Load saved scene on startup
- [ ] Test: Add objects dynamically
- [ ] Test: Save and load scene state
- [ ] Test: Reset clears scene

**Expected Result**: Working 3D UI menu, ability to add/reset/save scene

---

## Week 5 – Spatial Audio & Voice Commands ⏳

**Goal**: Add immersive audio to the scene and enable voice-based interaction for hands-free control.

### Tasks:
- [ ] Create `src/audio/` folder
- [ ] Create `src/audio/AudioManager.ts`
  - [ ] Set up THREE.AudioListener on camera
  - [ ] Implement loadSound method for loading audio files
  - [ ] Implement playSound/stopSound methods
  - [ ] Create procedural beep feedback using Web Audio API
- [ ] Create `src/audio/SpatialAudioSource.ts`
  - [ ] Set up THREE.PositionalAudio
  - [ ] Configure distance model and rolloff
  - [ ] Attach audio to mesh objects
- [ ] Create `src/voice/` folder
- [ ] Create `src/voice/VoiceCommandManager.ts`
  - [ ] Implement Web Speech API (SpeechRecognition)
  - [ ] Register voice commands
  - [ ] Process transcript and trigger actions
  - [ ] Handle browser compatibility
- [ ] Update `src/main.ts` (Week 5 version)
  - [ ] Initialize AudioManager
  - [ ] Initialize VoiceCommandManager
  - [ ] Register commands: "add cube", "change color red", "save scene"
  - [ ] Add button to start voice recognition (requires user gesture)
  - [ ] Add audio feedback beeps to menu actions
- [ ] Test: Voice commands trigger correct actions
- [ ] Test: Audio feedback plays on interactions

**Expected Result**: Voice-controlled scene manipulation with audio feedback

---

## Week 6 – Networking & Multi-user XR ⏳

**Goal**: Enable multiple users to share the same XR space and interact with shared objects in real-time.

### Tasks:
- [ ] Install socket.io-client (`npm install socket.io-client @types/socket.io-client`)
- [ ] Install server dependencies (`npm install socket.io express`)
- [ ] Install concurrently for running both servers (`npm install -D concurrently`)
- [ ] Create `server/` folder
- [ ] Create `server/server.js`
  - [ ] Set up Express + Socket.io server
  - [ ] Handle user join/leave events
  - [ ] Broadcast position/rotation updates
  - [ ] Handle shared object movement
- [ ] Create `src/networking/` folder
- [ ] Create `src/networking/Avatar.ts`
  - [ ] Create simple avatar mesh (sphere head + cylinder body)
  - [ ] Add name label sprite above avatar
  - [ ] Implement position/rotation update methods
- [ ] Create `src/networking/NetworkManager.ts`
  - [ ] Connect to Socket.io server
  - [ ] Handle existing-users, user-joined, user-left events
  - [ ] Create/remove avatars for remote users
  - [ ] Broadcast local camera position every 100ms
  - [ ] Handle shared object synchronization
- [ ] Update `src/main.ts` (Week 6 version)
  - [ ] Initialize NetworkManager
  - [ ] Send position updates in render loop
- [ ] Add package.json scripts for running server
- [ ] Test: Multiple browser windows show each other as avatars
- [ ] Test: Object movement syncs across clients

**Expected Result**: Multi-user XR space with visible avatars and shared interactions

---

## Week 7 – Asset Pipeline & Optimization ⏳

**Goal**: Load realistic 3D models (GLTF/GLB) and optimize scene for smooth performance on XR devices.

### Tasks:
- [ ] Create `src/assets/` folder
- [ ] Create `src/assets/GLTFModelLoader.ts`
  - [ ] Set up GLTFLoader from Three.js examples
  - [ ] Configure DRACOLoader for compression
  - [ ] Implement async load with progress callback
- [ ] Create `src/assets/AssetManager.ts`
  - [ ] Implement model caching system
  - [ ] Add texture loading support
  - [ ] Implement preloadAssets for bulk loading
  - [ ] Add getCachedModel method for instancing
- [ ] Create `src/optimization/` folder
- [ ] Create `src/optimization/LODManager.ts`
  - [ ] Implement createLOD with multiple detail levels
  - [ ] Add autoGenerateLOD helper
  - [ ] Document distance thresholds (0-10, 10-30, 30+ units)
- [ ] Create `src/optimization/DrawCallOptimizer.ts`
  - [ ] Implement geometry merging
  - [ ] Implement instanced mesh creation
  - [ ] Add performance stats logging
  - [ ] Add getDrawCallCount helper
- [ ] Update `src/ui/LoadingScreen.ts`
  - [ ] Create progress bar overlay
  - [ ] Update progress percentage display
  - [ ] Implement show/hide/remove methods
- [ ] Create `public/models/` folder for GLTF assets
- [ ] Update `src/main.ts` (Week 7 version)
  - [ ] Initialize AssetManager and optimizers
  - [ ] Show loading screen during asset loading
  - [ ] Create LOD objects for distant objects
  - [ ] Log performance stats periodically
- [ ] Test: GLTF models load with progress indicator
- [ ] Test: LOD switches at different distances
- [ ] Test: Draw calls reduced significantly

**Expected Result**: Optimized scene with GLTF models, LOD, and <50 draw calls maintaining 72-90 FPS

---

## Week 8 – UX & Accessibility in XR ⏳

**Goal**: Implement best practices for comfort, usability, and accessibility in XR experiences.

### Tasks:
- [ ] Create `src/locomotion/` folder
- [ ] Create `src/locomotion/TeleportController.ts`
  - [ ] Create teleport marker mesh
  - [ ] Implement raycasting to ground
  - [ ] Show/hide marker based on valid targets
  - [ ] Implement teleport (move camera X/Z, keep Y)
- [ ] Create `src/comfort/` folder
- [ ] Create `src/comfort/VignetteEffect.ts`
  - [ ] Create shader material for vignette effect
  - [ ] Implement smooth intensity transitions
  - [ ] Attach vignette plane to camera
- [ ] Create `src/accessibility/` folder
- [ ] Create `src/accessibility/AccessibilityManager.ts`
  - [ ] Define AccessibilitySettings type
  - [ ] Implement colorblind mode filters (protanopia, deuteranopia, tritanopia)
  - [ ] Add text scale settings
  - [ ] Implement localStorage persistence
  - [ ] Add reduced motion option
- [ ] Create `src/accessibility/GazeController.ts`
  - [ ] Create gaze cursor (ring geometry)
  - [ ] Implement 2-second dwell time activation
  - [ ] Add visual feedback (scale + color change)
  - [ ] Raycast from camera center
- [ ] Create `src/tutorial/` folder
- [ ] Create `src/tutorial/TutorialManager.ts`
  - [ ] Create tutorial overlay UI
  - [ ] Implement multi-step tutorial flow
  - [ ] Add step counter display
  - [ ] Check localStorage to show only on first visit
- [ ] Update `src/main.ts` (Week 8 version)
  - [ ] Add ground plane for teleportation
  - [ ] Initialize all UX systems
  - [ ] Set up tutorial with 5 steps
  - [ ] Implement gaze selection with colorblind filter
  - [ ] Add controller teleportation with vignette
  - [ ] Update all systems in render loop
- [ ] Test: Teleportation works smoothly with vignette
- [ ] Test: Gaze selection activates after 2s
- [ ] Test: Tutorial shows on first visit only
- [ ] Test: Colorblind filters apply correctly

**Expected Result**: Comfortable, accessible XR experience with teleportation, gaze UI, and tutorial

---

## Deployment 📦

### Tasks:
- [ ] Run `npm run build` to create production build
- [ ] Test production build locally (`npm run preview`)
- [ ] Choose deployment platform:
  - [ ] Option A: Vercel
  - [ ] Option B: Netlify
  - [ ] Option C: GitHub Pages
- [ ] Configure HTTPS (required for WebXR on devices)
- [ ] Deploy and test on actual XR device
- [ ] Verify VR mode works
- [ ] Verify AR mode works (if device supports)

---

## Optional Enhancements 🚀

### Future Ideas:
- [ ] Add hand tracking support
- [ ] Add networking/multiplayer (WebRTC / Colyseus)
- [ ] Improve UI with `three-mesh-ui` library
- [ ] Add GLTF model loading instead of primitive shapes
- [ ] Add teleportation locomotion
- [ ] Add haptic feedback
- [ ] Add spatial audio
- [ ] Add custom shaders/materials
- [ ] Add shadows and advanced lighting
- [ ] Add post-processing effects

---

## Notes

- **Course Duration**: 8 weeks total
- **Current Status**: Week 1 Complete ✅ (7 weeks remaining)
- **Next Step**: Begin Week 2 - Interaction & Controllers
- **Dev Server**: Running on http://localhost:5173/
- **XR Testing**: Requires HTTPS and WebXR-capable device (e.g., Meta Quest 3)
- **Advanced Features**: Weeks 5-8 cover spatial audio, networking, optimization, and accessibility

