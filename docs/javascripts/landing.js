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

  /* ---------------- access flow ---------------- */
  function initFlow(root) {
    var svg = root.querySelector(".lp-diagram");
    if (!svg) return;
    var $ = function (id) { return svg.querySelector("#" + id); };
    var stepButtons = Array.prototype.slice.call(root.querySelectorAll("[data-step]"));
    var toggle = root.querySelector('[data-flow="toggle"]');
    var restart = root.querySelector('[data-flow="restart"]');
    var status = $("lp-status");
    var labels = ["s0", "s1", "s2", "s3", "s4"];
    var statusText = [
      "controller → gateway: access list (SDP ID 3 → service 1)",
      "iptables: DROP 10.231.0.40 → :8080 (no rule)",
      "SPA: udp/62201 · AES + HMAC · SDP ID 3, service 1",
      "fwknopd: DNAT :8080 → 10.231.1.30:80 for 10.231.0.40 · 30 s",
      "HTTP 200 · keys rotated with the controller"
    ];

    function setStep(i) {
      stepButtons.forEach(function (b, j) {
        if (j === i) b.setAttribute("aria-current", "step");
        else b.removeAttribute("aria-current");
      });
      if (status) status.textContent = statusText[i];
    }

    var nodes = {
      ctrl: $("lp-n-ctrl"), client: $("lp-n-client"), gw: $("lp-n-gw"), svc: $("lp-n-svc")
    };

    function finalState() {
      [$("lp-l-cl-gw"), $("lp-l-gw-svc")].forEach(function (l) { l.classList.add("is-open"); });
      nodes.svc.classList.add("is-on");
      $("lp-badge").classList.add("is-open");
      $("lp-badge-closed").setAttribute("opacity", "0");
      $("lp-badge-open").setAttribute("opacity", "1");
      setStep(4);
    }

    if (!window.gsap || reducedMotion()) {
      finalState();
      stepButtons.forEach(function (b, i) {
        b.addEventListener("click", function () { setStep(i); });
      });
      if (toggle) toggle.hidden = true;
      if (restart) restart.hidden = true;
      return;
    }

    var gsap = window.gsap;
    if (window.ScrollTrigger) gsap.registerPlugin(window.ScrollTrigger);
    var openLinks = [$("lp-l-cl-gw"), $("lp-l-gw-svc")];
    var tl = gsap.timeline({ repeat: -1, repeatDelay: 1.2, paused: true });

    // reset to the initial state at the start of every loop
    tl.call(function () {
      openLinks.forEach(function (l) { l.classList.remove("is-open"); });
      $("lp-l-ctrl-gw").classList.remove("is-active");
      $("lp-l-cl-ctrl").classList.remove("is-active");
      nodes.svc.classList.remove("is-on");
      Object.keys(nodes).forEach(function (k) { nodes[k].classList.remove("is-focus"); });
      $("lp-badge").classList.remove("is-open");
    });
    tl.set(["#lp-p-attempt", "#lp-p-policy", "#lp-p-cred", "#lp-p-spa", "#lp-p-data", "#lp-p-drop"], { opacity: 0 });
    tl.set("#lp-badge-closed", { opacity: 1 });
    tl.set("#lp-badge-open", { opacity: 0 });

    // 1. gateway enrolls: controller pushes the access list
    tl.addLabel(labels[0]);
    tl.call(setStep, [0]);
    tl.call(function () { $("lp-l-ctrl-gw").classList.add("is-active"); nodes.ctrl.classList.add("is-focus"); });
    tl.fromTo("#lp-p-policy", { attr: { cy: 96 }, opacity: 1 }, { attr: { cy: 168 }, duration: 0.9, ease: "power1.inOut" });
    tl.to("#lp-p-policy", { opacity: 0, duration: 0.2 });
    tl.call(function () { $("lp-l-ctrl-gw").classList.remove("is-active"); nodes.ctrl.classList.remove("is-focus"); });
    tl.to({}, { duration: 0.8 });

    // 2. hidden by default: a plain connection attempt is dropped
    tl.addLabel(labels[1]);
    tl.call(setStep, [1]);
    tl.fromTo("#lp-p-attempt", { attr: { cx: 172 }, opacity: 1 }, { attr: { cx: 246 }, duration: 0.7, ease: "power1.in" });
    tl.to("#lp-p-attempt", { opacity: 0, duration: 0.15 });
    tl.fromTo("#lp-p-drop", { opacity: 1, scale: 0.6, transformOrigin: "50% 50%" }, { scale: 1, duration: 0.25 }, "<");
    tl.to("#lp-p-drop", { opacity: 0, duration: 0.4, delay: 0.7 });
    tl.to({}, { duration: 0.4 });

    // 3. single packet authorization
    tl.addLabel(labels[2]);
    tl.call(setStep, [2]);
    tl.call(function () { nodes.client.classList.add("is-focus"); });
    tl.fromTo("#lp-p-spa", { attr: { cx: 172 }, opacity: 1 }, { attr: { cx: 300 }, duration: 0.9, ease: "power2.inOut" });
    tl.to("#lp-p-spa", { opacity: 0, duration: 0.2 });
    tl.call(function () { nodes.client.classList.remove("is-focus"); });
    tl.to({}, { duration: 0.5 });

    // 4. verified and opened
    tl.addLabel(labels[3]);
    tl.call(setStep, [3]);
    tl.call(function () { nodes.gw.classList.add("is-focus"); $("lp-badge").classList.add("is-open"); });
    tl.to("#lp-badge", { scale: 1.25, transformOrigin: "50% 50%", duration: 0.2, yoyo: true, repeat: 1 });
    tl.to("#lp-badge-closed", { opacity: 0, duration: 0.2 }, "<");
    tl.to("#lp-badge-open", { opacity: 1, duration: 0.2 }, "<");
    tl.call(function () { openLinks.forEach(function (l) { l.classList.add("is-open"); }); });
    tl.to({}, { duration: 1.1 });
    tl.call(function () { nodes.gw.classList.remove("is-focus"); });

    // 5. connected: traffic flows, credentials rotated with the controller
    tl.addLabel(labels[4]);
    tl.call(setStep, [4]);
    tl.call(function () { nodes.svc.classList.add("is-on"); $("lp-l-cl-ctrl").classList.add("is-active"); });
    tl.fromTo("#lp-p-data", { attr: { cx: 172 }, opacity: 1 }, { attr: { cx: 470 }, duration: 1.2, ease: "none", repeat: 1 });
    tl.fromTo("#lp-p-cred", { attr: { cx: 110, cy: 170 }, opacity: 1 }, { attr: { cx: 255, cy: 80 }, duration: 1.1, ease: "power1.inOut" }, "<0.3");
    tl.to(["#lp-p-data", "#lp-p-cred"], { opacity: 0, duration: 0.2 });
    tl.call(function () { $("lp-l-cl-ctrl").classList.remove("is-active"); });
    tl.to({}, { duration: 1.6 });

    function setPlaying(playing) {
      if (!toggle) return;
      toggle.setAttribute("aria-label", playing ? "Pause animation" : "Play animation");
      toggle.classList.toggle("is-paused", !playing);
    }

    stepButtons.forEach(function (b, i) {
      b.addEventListener("click", function () {
        tl.play(labels[i]);
        setPlaying(true);
      });
    });
    if (toggle) toggle.addEventListener("click", function () {
      if (tl.paused()) { tl.play(); setPlaying(true); } else { tl.pause(); setPlaying(false); }
    });
    if (restart) restart.addEventListener("click", function () { tl.restart(); setPlaying(true); });

    // play only while visible
    if (window.ScrollTrigger) {
      var st = window.ScrollTrigger.create({
        trigger: svg,
        start: "top 85%",
        end: "bottom 10%",
        onEnter: function () { if (!toggle || !toggle.classList.contains("is-paused")) tl.play(); },
        onEnterBack: function () { if (!toggle || !toggle.classList.contains("is-paused")) tl.play(); },
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
   * Progressive enhancement: without JavaScript everything is visible.
   * With it, sections fade in as they enter the viewport (IntersectionObserver
   * + CSS transitions), and the result bars grow when they become visible. */
  function initReveals(root) {
    var items = Array.prototype.slice.call(root.querySelectorAll("[data-reveal]"));
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
