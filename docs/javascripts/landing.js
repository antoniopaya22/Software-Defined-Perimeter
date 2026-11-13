/* Landing page interactions (GSAP 3 + ScrollTrigger).
 * - Animated access flow with step navigation, pause and restart.
 * - Scroll reveals and growing result bars (IntersectionObserver + CSS).
 * - Copy-to-clipboard buttons.
 * Works with Material's instant navigation (document$) and respects
 * prefers-reduced-motion: everything is shown in its final state. */
(function () {
  "use strict";

  var cleanups = [];

  function cleanup() {
    while (cleanups.length) {
      try { cleanups.pop()(); } catch (e) { /* ignore */ }
    }
  }

  function reducedMotion() {
    return window.matchMedia && window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  }

  /* ---------------- copy buttons ---------------- */
  function initCopy(root) {
    root.querySelectorAll("[data-copy], [data-copy-target]").forEach(function (btn) {
      btn.addEventListener("click", function () {
        var text = btn.getAttribute("data-copy");
        if (!text) {
          var target = document.getElementById(btn.getAttribute("data-copy-target"));
          text = target ? target.textContent : "";
        }
        var done = function () {
          var label = btn.textContent;
          btn.classList.add("is-copied");
          btn.textContent = "Copied";
          setTimeout(function () { btn.classList.remove("is-copied"); btn.textContent = label; }, 1600);
        };
        if (navigator.clipboard && navigator.clipboard.writeText) {
          navigator.clipboard.writeText(text.trim()).then(done, function () {});
        }
      });
    });
  }

  /* ---------------- access flow ----------------
   * Generic driver: step buttons ([data-step], with data-status text),
   * pause/restart controls and visibility handling. The animation itself
   * is selected by <svg class="lp-diagram" data-flow="...">. */

  var FLOWS = {};

  // Software-Defined Perimeter lab: controller, client, gateway, service.
  FLOWS.sdp = {
    reset: function (q) {
      ["lp-l-cl-gw", "lp-l-gw-svc"].forEach(function (id) { q(id).classList.remove("is-open"); });
      ["lp-l-ctrl-gw", "lp-l-cl-ctrl"].forEach(function (id) { q(id).classList.remove("is-active"); });
    },
    final: function (q) {
      ["lp-l-cl-gw", "lp-l-gw-svc"].forEach(function (id) { q(id).classList.add("is-open"); });
    },
    build: function (tl, q, step, focus) {
      // 1. gateway enrolls
      step(0);
      tl.call(function () { q("lp-l-ctrl-gw").classList.add("is-active"); focus("lp-n-ctrl", true); });
      tl.fromTo("#lp-p-policy", { attr: { cy: 96 }, opacity: 1 }, { attr: { cy: 168 }, duration: 0.9, ease: "power1.inOut" });
      tl.to("#lp-p-policy", { opacity: 0, duration: 0.2 });
      tl.call(function () { q("lp-l-ctrl-gw").classList.remove("is-active"); focus("lp-n-ctrl", false); });
      tl.to({}, { duration: 0.8 });
      // 2. hidden by default
      step(1);
      tl.fromTo("#lp-p-attempt", { attr: { cx: 172 }, opacity: 1 }, { attr: { cx: 246 }, duration: 0.7, ease: "power1.in" });
      tl.to("#lp-p-attempt", { opacity: 0, duration: 0.15 });
      tl.fromTo("#lp-p-drop", { opacity: 1, scale: 0.6, transformOrigin: "50% 50%" }, { scale: 1, duration: 0.25 }, "<");
      tl.to("#lp-p-drop", { opacity: 0, duration: 0.4, delay: 0.7 });
      tl.to({}, { duration: 0.4 });
      // 3. single packet
      step(2);
      tl.call(function () { focus("lp-n-client", true); });
      tl.fromTo("#lp-p-spa", { attr: { cx: 172 }, opacity: 1 }, { attr: { cx: 300 }, duration: 0.9, ease: "power2.inOut" });
      tl.to("#lp-p-spa", { opacity: 0, duration: 0.2 });
      tl.call(function () { focus("lp-n-client", false); });
      tl.to({}, { duration: 0.5 });
      // 4. verified and opened
      step(3);
      openGate(tl, q, "lp-n-gw", ["lp-l-cl-gw", "lp-l-gw-svc"], focus);
      // 5. connected
      step(4);
      tl.call(function () { q("lp-n-svc").classList.add("is-on"); q("lp-l-cl-ctrl").classList.add("is-active"); });
      tl.fromTo("#lp-p-data", { attr: { cx: 172 }, opacity: 1 }, { attr: { cx: 470 }, duration: 1.2, ease: "none", repeat: 1 });
      tl.fromTo("#lp-p-cred", { attr: { cx: 110, cy: 170 }, opacity: 1 }, { attr: { cx: 255, cy: 80 }, duration: 1.1, ease: "power1.inOut" }, "<0.3");
      tl.to(["#lp-p-data", "#lp-p-cred"], { opacity: 0, duration: 0.2 });
      tl.call(function () { q("lp-l-cl-ctrl").classList.remove("is-active"); });
      tl.to({}, { duration: 1.6 });
    }
  };

  // SecureSDP Bridge model: client and bridge in domain-b, root, gateway
  // and service in domain-a.
  FLOWS.bridge = {
    reset: function (q) {
      ["lp-l-cl-gw", "lp-l-gw-svc"].forEach(function (id) { q(id).classList.remove("is-open"); });
      ["lp-l-cl-br", "lp-l-br-root", "lp-l-root-gw"].forEach(function (id) { q(id).classList.remove("is-active"); });
    },
    final: function (q) {
      ["lp-l-cl-gw", "lp-l-gw-svc"].forEach(function (id) { q(id).classList.add("is-open"); });
    },
    build: function (tl, q, step, focus) {
      // 1. authenticate locally (client -> bridge)
      step(0);
      tl.call(function () { q("lp-l-cl-br").classList.add("is-active"); focus("lp-n-client", true); });
      tl.fromTo("#lp-p-req", { attr: { cx: 120, cy: 218 }, opacity: 1 }, { attr: { cy: 122 }, duration: 0.9, ease: "power1.inOut" });
      tl.call(function () { focus("lp-n-client", false); focus("lp-n-bridge", true); });
      tl.to({}, { duration: 0.4 });
      // 2. relay to the root (bridge -> root)
      step(1);
      tl.call(function () { q("lp-l-br-root").classList.add("is-active"); });
      tl.fromTo("#lp-p-req", { attr: { cx: 197, cy: 90 } }, { attr: { cx: 278 }, duration: 0.9, ease: "power1.inOut" });
      tl.to("#lp-p-req", { opacity: 0, duration: 0.2 });
      tl.call(function () { focus("lp-n-bridge", false); focus("lp-n-root", true); });
      tl.to({}, { duration: 0.4 });
      // 3. decide and issue (root -> gateway: stanza; root -> client: grant)
      step(2);
      tl.call(function () { q("lp-l-root-gw").classList.add("is-active"); });
      tl.fromTo("#lp-p-grant", { attr: { cx: 355, cy: 122 }, opacity: 1 }, { attr: { cy: 218 }, duration: 0.9, ease: "power1.inOut" });
      tl.fromTo("#lp-p-reply", { attr: { cx: 278, cy: 90 }, opacity: 1 }, { attr: { cx: 197 }, duration: 0.9, ease: "power1.inOut" }, "<");
      tl.to("#lp-p-grant", { opacity: 0, duration: 0.2 });
      tl.to("#lp-p-reply", { attr: { cx: 120, cy: 122 }, duration: 0.3 }, "<");
      tl.to("#lp-p-reply", { attr: { cy: 218 }, duration: 0.7, ease: "power1.inOut" });
      tl.to("#lp-p-reply", { opacity: 0, duration: 0.2 });
      tl.call(function () {
        focus("lp-n-root", false);
        ["lp-l-cl-br", "lp-l-br-root", "lp-l-root-gw"].forEach(function (id) { q(id).classList.remove("is-active"); });
      });
      tl.to({}, { duration: 0.3 });
      // 4. one packet (SPA client -> gateway) and the gate opens
      step(3);
      tl.call(function () { focus("lp-n-client", true); });
      tl.fromTo("#lp-p-spa", { attr: { cx: 197, cy: 250 }, opacity: 1 }, { attr: { cx: 300 }, duration: 0.8, ease: "power2.inOut" });
      tl.to("#lp-p-spa", { opacity: 0, duration: 0.2 });
      tl.call(function () { focus("lp-n-client", false); });
      openGate(tl, q, "lp-n-gw", ["lp-l-cl-gw", "lp-l-gw-svc"], focus);
      // 5. connected
      step(4);
      tl.call(function () { q("lp-n-svc").classList.add("is-on"); });
      tl.fromTo("#lp-p-data", { attr: { cx: 197, cy: 250 }, opacity: 1 }, { attr: { cx: 462 }, duration: 1.2, ease: "none", repeat: 1 });
      tl.to("#lp-p-data", { opacity: 0, duration: 0.2 });
      tl.to({}, { duration: 1.6 });
    }
  };

  // shared: the gateway badge turns from locked to open
  function openGate(tl, q, gwId, links, focus) {
    tl.call(function () { focus(gwId, true); q("lp-badge").classList.add("is-open"); });
    tl.to("#lp-badge", { scale: 1.25, transformOrigin: "50% 50%", duration: 0.2, yoyo: true, repeat: 1 });
    tl.to("#lp-badge-closed", { opacity: 0, duration: 0.2 }, "<");
    tl.to("#lp-badge-open", { opacity: 1, duration: 0.2 }, "<");
    tl.call(function () { links.forEach(function (id) { q(id).classList.add("is-open"); }); });
    tl.to({}, { duration: 1.1 });
    tl.call(function () { focus(gwId, false); });
  }

  function initFlow(root) {
    var svg = root.querySelector(".lp-diagram");
    if (!svg) return;
    var flow = FLOWS[svg.getAttribute("data-flow") || "sdp"];
    if (!flow) return;
    var q = function (id) { return svg.querySelector("#" + id); };
    var stepButtons = Array.prototype.slice.call(root.querySelectorAll("[data-step]"));
    var toggle = root.querySelector('[data-flow="toggle"]');
    var restart = root.querySelector('[data-flow="restart"]');
    var status = q("lp-status");
    var labels = stepButtons.map(function (_, i) { return "s" + i; });

    function setStep(i) {
      stepButtons.forEach(function (b, j) {
        if (j === i) b.setAttribute("aria-current", "step");
        else b.removeAttribute("aria-current");
      });
      if (status && stepButtons[i]) status.textContent = stepButtons[i].getAttribute("data-status") || "";
    }

    function focus(id, on) { var n = q(id); if (n) n.classList.toggle("is-focus", on); }

    function finalState() {
      flow.final(q);
      var svc = q("lp-n-svc"); if (svc) svc.classList.add("is-on");
      q("lp-badge").classList.add("is-open");
      q("lp-badge-closed").setAttribute("opacity", "0");
      q("lp-badge-open").setAttribute("opacity", "1");
      setStep(stepButtons.length - 1);
    }

    if (!window.gsap || reducedMotion()) {
      finalState();
      stepButtons.forEach(function (b, i) { b.addEventListener("click", function () { setStep(i); }); });
      if (toggle) toggle.hidden = true;
      if (restart) restart.hidden = true;
      return;
    }

    var gsap = window.gsap;
    if (window.ScrollTrigger) gsap.registerPlugin(window.ScrollTrigger);
    var tl = gsap.timeline({ repeat: -1, repeatDelay: 1.2, paused: true });

    // reset at the start of every loop
    tl.call(function () {
      flow.reset(q);
      svg.querySelectorAll(".lp-d-node").forEach(function (n) { n.classList.remove("is-focus", "is-on"); });
      q("lp-badge").classList.remove("is-open");
    });
    tl.set(svg.querySelectorAll(".lp-d-pkt, #lp-p-drop"), { opacity: 0 });
    tl.set("#lp-badge-closed", { opacity: 1 });
    tl.set("#lp-badge-open", { opacity: 0 });

    function step(i) { tl.addLabel(labels[i]); tl.call(setStep, [i]); }
    flow.build(tl, q, step, focus);

    function setPlaying(playing) {
      if (!toggle) return;
      toggle.setAttribute("aria-label", playing ? "Pause animation" : "Play animation");
      toggle.classList.toggle("is-paused", !playing);
    }

    stepButtons.forEach(function (b, i) {
      b.addEventListener("click", function () { tl.play(labels[i]); setPlaying(true); });
    });
    if (toggle) toggle.addEventListener("click", function () {
      if (tl.paused()) { tl.play(); setPlaying(true); } else { tl.pause(); setPlaying(false); }
    });
    if (restart) restart.addEventListener("click", function () { tl.restart(); setPlaying(true); });

    // play only while visible
    if (window.ScrollTrigger) {
      var playIfAllowed = function () { if (!toggle || !toggle.classList.contains("is-paused")) tl.play(); };
      var st = window.ScrollTrigger.create({
        trigger: svg,
        start: "top 85%",
        end: "bottom 10%",
        onEnter: playIfAllowed,
        onEnterBack: playIfAllowed,
        onLeave: function () { tl.pause(); },
        onLeaveBack: function () { tl.pause(); }
      });
      cleanups.push(function () { st.kill(); });
    } else {
      tl.play();
    }
    cleanups.push(function () { tl.kill(); });
  }

  /* ---------------- reveals and bars ----------------
   * Content is always visible (no fade-ins that leave gaps while scrolling).
   * Only the result bars grow when they enter the viewport. */
  function initReveals(root) {
    var items = Array.prototype.slice.call(root.querySelectorAll(".lp-bars[data-reveal]"));
    if (reducedMotion() || !("IntersectionObserver" in window)) return;

    root.classList.add("lp-js");
    items.forEach(function (el) { el.classList.add("is-pending"); });

    var io = new IntersectionObserver(function (entries) {
      var visible = entries.filter(function (e) { return e.isIntersecting; });
      visible.forEach(function (entry, i) {
        entry.target.style.transitionDelay = Math.min(i, 5) * 70 + "ms";
        entry.target.classList.remove("is-pending");
        io.unobserve(entry.target);
      });
    }, { rootMargin: "0px 0px -8% 0px", threshold: 0.05 });

    items.forEach(function (el) { io.observe(el); });
    cleanups.push(function () {
      io.disconnect();
      items.forEach(function (el) { el.classList.remove("is-pending"); el.style.transitionDelay = ""; });
    });
  }

  function init() {
    cleanup();
    var root = document.querySelector("[data-landing]");
    if (!root) return;
    initCopy(root);
    initFlow(root);
    initReveals(root);
  }

  if (window.document$ && typeof window.document$.subscribe === "function") {
    window.document$.subscribe(init);
  } else if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();
