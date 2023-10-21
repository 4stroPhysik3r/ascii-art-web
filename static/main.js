(function setGlowEffectRx() {
  const glowEffects = document.querySelectorAll('.glow-effect')

  glowEffects.forEach(glowEffect => {
    const glowLines = glowEffect.querySelectorAll('rect')
    const rx = getComputedStyle(glowEffect).borderRadius

    glowLines.forEach((line) => {
      line.setAttribute('rx', rx)
    })
  })
})()
