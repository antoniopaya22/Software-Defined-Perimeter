/* Landing page of the SDP research sites (GSAP 3 + ScrollTrigger).
 * Shared by Software-Defined-Perimeter and Secure-SDP; each page enables
 * the parts it contains. Without JavaScript all content is visible; with
 * prefers-reduced-motion every figure shows its final state. */
(function () {
  "use strict";
  var root = document.getElementById("lp");
  if (!root) return;
  var gsap = window.gsap;
  if (gsap && window.ScrollTrigger) gsap.registerPlugin(window.ScrollTrigger);
  var reduce = window.matchMedia && window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  var animate = !!gsap && !reduce;

  function whenVisible(el, fn, margin) {
    if (!("IntersectionObserver" in window)) { fn(); return; }
    var io = new IntersectionObserver(function (entries) {
      entries.forEach(function (e) { if (e.isIntersecting) { io.disconnect(); fn(); } });
    }, { rootMargin: margin || "0px 0px -10% 0px" });
    io.observe(el);
  }

  /* ---------------- navigation: solid header and section dots ---------------- */
  (function nav() {
    var bar = document.getElementById("lp-nav");
    var onScroll = function () { if (bar) bar.classList.toggle("is-solid", window.scrollY > 24); };
    window.addEventListener("scroll", onScroll, { passive: true });
    onScroll();
    var dots = Array.prototype.slice.call(document.querySelectorAll(".lp-dots a"));
    if (!dots.length || !("IntersectionObserver" in window)) return;
    var io = new IntersectionObserver(function (entries) {
      entries.forEach(function (e) {
        if (!e.isIntersecting) return;
        dots.forEach(function (d) { d.classList.toggle("is-on", d.getAttribute("data-dot") === e.target.id); });
      });
    }, { rootMargin: "-45% 0px -50% 0px" });
    dots.forEach(function (d) {
      var sec = document.getElementById(d.getAttribute("data-dot"));
      if (sec) io.observe(sec);
    });
  })();

  /* ---------------- copy buttons ---------------- */
  root.querySelectorAll("[data-copy], [data-copy-target]").forEach(function (btn) {
    btn.addEventListener("click", function () {
      var text = btn.getAttribute("data-copy");
      if (!text) {
        var t = document.getElementById(btn.getAttribute("data-copy-target"));
        text = t ? t.textContent : "";
      }
      if (!navigator.clipboard) return;
      navigator.clipboard.writeText(text.trim()).then(function () {
        var label = btn.textContent;
        btn.textContent = "Copied";
        setTimeout(function () { btn.textContent = label; }, 1500);
      });
    });
  });

  /* ---------------- terminal replay ----------------
   * <pre class="lp-term"> holds one <span class="ln"> per line (visible
   * without JS). When it scrolls into view the lines are replayed. */
  root.querySelectorAll("pre.lp-term").forEach(function (term) {
    if (reduce) return;
    var lines = Array.prototype.slice.call(term.querySelectorAll(".ln"));
    if (!lines.length) return;
    lines.forEach(function (l) { l.style.display = "none"; });
    var cur = document.createElement("span");
    cur.className = "cur";
    cur.textContent = "▋";
    term.appendChild(cur);
    whenVisible(term, function () {
      var i = 0;
      (function next() {
        if (i >= lines.length) { cur.remove(); return; }
        var l = lines[i++];
        l.style.display = "";
        term.insertBefore(cur, l.nextSibling);
        setTimeout(next, l.classList.contains("cmd") ? 650 : 170);
      })();
    });
  });

  /* ---------------- duel: exposed port vs SDP gateway ---------------- */
  (function duel() {
    var open = document.getElementById("lp-duel-open");
    var sdp = document.getElementById("lp-duel-sdp");
    if (!open || !sdp) return;
    var sOpen = document.querySelector('[data-duel="open"]');
    var sSdp = document.querySelector('[data-duel="sdp"]');
    function finalState() {
      open.querySelector(".lp-dl-svc").classList.add("is-exposed");
      if (sOpen) { sOpen.textContent = "8 of 8 probes answered"; sOpen.className = "lp-duel__status is-bad"; }
      if (sSdp) { sSdp.textContent = "0 of 8 probes answered"; sSdp.className = "lp-duel__status is-good"; }
    }
    if (!animate) { finalState(); return; }
    var NS = "http://www.w3.org/2000/svg";
    function run(svg, blocked, status, label) {
      var g = svg.querySelector(".lp-dl-probes");
      var svc = svg.querySelector(".lp-dl-svc");
      var x = svg.querySelector(".lp-dl-x");
      var answered = 0, sent = 0;
      var tl = gsap.timeline({ repeat: -1, repeatDelay: 1.2, onRepeat: function () { answered = 0; sent = 0; svc.classList.remove("is-exposed"); } });
      tl.call(function () { while (g.firstChild) g.removeChild(g.firstChild); status.textContent = "scanning…"; status.className = "lp-duel__status"; });
      for (var i = 0; i < 8; i++) {
        (function (k) {
          var y = 40 + k * 17;
          tl.call(function () {
            var c = document.createElementNS(NS, "circle");
            c.setAttribute("class", "lp-dl-probe"); c.setAttribute("r", "4"); c.setAttribute("cx", "40"); c.setAttribute("cy", String(y));
            g.appendChild(c);
            gsap.to(c, { attr: { cx: blocked ? 158 : 240, cy: blocked ? y : 100 }, duration: 0.55, ease: "power1.in", onComplete: function () {
              sent++;
              if (blocked) {
                gsap.to(c, { opacity: 0, duration: 0.25 });
                gsap.fromTo(x, { opacity: 1 }, { opacity: 0, duration: 0.5 });
              } else {
                answered++;
                svc.classList.add("is-exposed");
                gsap.to(c, { attr: { cx: 40, cy: y }, opacity: 0.25, duration: 0.6, ease: "power1.out" });
              }
              status.textContent = answered + " of " + sent + " probes answered";
              status.className = "lp-duel__status " + (blocked ? "is-good" : "is-bad");
            } });
          }, null, k * 0.22);
        })(i);
      }
      tl.to({}, { duration: 1.8 });
      return tl;
    }
    whenVisible(open, function () {
      run(open, false, sOpen);
      run(sdp, true, sSdp);
    });
  })();

  /* ---------------- access flow ---------------- */
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


  (function flow() {
    var svg = root.querySelector(".lp-diagram");
    if (!svg) return;
    var def = FLOWS[svg.getAttribute("data-flow") || "sdp"];
    if (!def) return;
    var q = function (id) { return svg.querySelector("#" + id); };
    var buttons = Array.prototype.slice.call(root.querySelectorAll("[data-step]"));
    var caption = document.getElementById("lp-flow-caption");
    var toggle = root.querySelector('[data-flow="toggle"]');
    var restart = root.querySelector('[data-flow="restart"]');
    var status = q("lp-status");
    var labels = buttons.map(function (_, i) { return "s" + i; });

    function setStep(i) {
      buttons.forEach(function (b, j) {
        if (j === i) b.setAttribute("aria-current", "step"); else b.removeAttribute("aria-current");
      });
      var b = buttons[i];
      if (!b) return;
      if (status) status.textContent = b.getAttribute("data-status") || "";
      if (caption) caption.textContent = b.getAttribute("data-caption") || "";
    }
    function focus(id, on) { var n = q(id); if (n) n.classList.toggle("is-focus", on); }
    function finalState() {
      def.final(q);
      var svc = q("lp-n-svc"); if (svc) svc.classList.add("is-on");
      q("lp-badge").classList.add("is-open");
      q("lp-badge-closed").setAttribute("opacity", "0");
      q("lp-badge-open").setAttribute("opacity", "1");
      setStep(buttons.length - 1);
    }

    if (!animate) {
      finalState();
      buttons.forEach(function (b, i) { b.addEventListener("click", function () { setStep(i); }); });
      if (toggle) toggle.hidden = true;
      if (restart) restart.hidden = true;
      return;
    }

    var tl = gsap.timeline({ repeat: -1, repeatDelay: 1.2, paused: true });
    tl.call(function () {
      def.reset(q);
      svg.querySelectorAll(".lp-d-node").forEach(function (n) { n.classList.remove("is-focus", "is-on"); });
      q("lp-badge").classList.remove("is-open");
    });
    tl.set(svg.querySelectorAll(".lp-d-pkt, #lp-p-drop"), { opacity: 0 });
    tl.set("#lp-badge-closed", { opacity: 1 });
    tl.set("#lp-badge-open", { opacity: 0 });
    function step(i) { tl.addLabel(labels[i]); tl.call(setStep, [i]); }
    def.build(tl, q, step, focus);

    function setPlaying(p) {
      if (!toggle) return;
      toggle.setAttribute("aria-label", p ? "Pause animation" : "Play animation");
      toggle.textContent = p ? "❚❚" : "▶";
      toggle.dataset.paused = p ? "" : "1";
    }
    buttons.forEach(function (b, i) { b.addEventListener("click", function () { tl.play(labels[i]); setPlaying(true); }); });
    if (toggle) toggle.addEventListener("click", function () {
      if (tl.paused()) { tl.play(); setPlaying(true); } else { tl.pause(); setPlaying(false); }
    });
    if (restart) restart.addEventListener("click", function () { tl.restart(); setPlaying(true); });

    var allowed = function () { return !toggle || !toggle.dataset.paused; };
    if (window.ScrollTrigger) {
      window.ScrollTrigger.create({
        trigger: svg, start: "top 90%", end: "bottom 10%",
        onEnter: function () { if (allowed()) tl.play(); },
        onEnterBack: function () { if (allowed()) tl.play(); },
        onLeave: function () { tl.pause(); },
        onLeaveBack: function () { tl.pause(); }
      });
    } else {
      tl.play();
    }
  })();
})();
