import PropTypes from 'prop-types'
import React from 'react'
import styled, { keyframes } from 'styled-components'
import { layout } from 'styled-system'

const rotate = keyframes`
  0% {
    transform: rotate(0);
  }
  100% {
    transform: rotate(-360deg);
  }
`

const Overlay = styled.div`
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  display: grid;
  align-items: center;
  justify-items: center;
`

const Container = styled.div`
  ${layout}

  #dsnn {
    will-change: transform;
    transform-origin: center center;
    animation: ${rotate} 0.8s infinite linear;
  }
`

const DSNn = ({ size, color, ...props }) => (
  <Overlay {...props}>
    <Container size={size}>
      <svg viewBox="0 0 100 100" id="dsnn" width="100%">
        <g
          fill={color}
          fillOpacity="0.9"
          transform="matrix(3.033652,0,0,3.033652,-216.55808,-52.042203)"
        >
          <path d="m 84.774931,48.082027 c -5.382014,-1.559736 -6.314469,-2.617115 -8.35259,-4.861912 -3.64551,-4.035142 -4.38983,-9.379617 -2.12741,-15.275404 0.27727,-0.722563 -0.46418,-0.347766 -0.97748,0.494104 -0.28836,0.472943 -0.52473,0.751294 -0.52528,0.618558 -0.002,-0.597529 1.39907,-2.536251 2.46475,-3.409495 l 1.17177,-0.96018 -1.04849,2.052192 c -3.96309,7.756878 0.27045,16.365359 9.19095,18.688851 8.14862,2.122448 16.168649,-5.115455 14.71032,-13.275753 -0.181,-1.012829 -0.61089,-2.449734 -0.95531,-3.193123 -1.22873,-2.652088 -4.08081,-5.079845 -7.07708,-6.024173 -1.58155,-0.498454 -2.29182,-0.347444 -1.40743,0.299235 0.5967,0.436315 0.22012,0.49037 -0.77649,0.111459 -0.4317,-0.164134 -0.6343,-0.167519 -0.53631,-0.009 -1.845412,0.419865 -2.699594,-0.134509 -3.77523,-0.06044 -1.3605,0.09492 -2.09774,7.94e-4 -3.15149,-0.402353 -1.47991,-0.566198 -1.66146,-0.531939 -1.67777,0.316608 -0.009,0.479049 -0.0633,0.453956 -0.41333,-0.191672 -0.69543,-1.282778 -0.43994,-1.576161 1.46272,-1.67963 1.13052,-0.06148 2.28306,-0.347421 3.43958,-0.853345 0.95663,-0.418485 2.40464,-0.953943 3.2178,-1.189906 1.64019,-0.475957 5.4031,-0.596432 5.13612,-0.164441 -0.0899,0.145521 -0.50666,0.264584 -0.92605,0.264584 -1.08722,0 -1.22351,0.452509 -0.20074,0.666525 0.47421,0.09923 1.57658,0.509917 2.4497,0.912639 0.87313,0.402721 2.27211,0.932619 3.10886,1.177549 1.74434,0.510599 1.97378,0.722725 0.52916,0.489237 l -0.99218,-0.160365 1.42185,1.557114 c 0.78201,0.856412 1.6905,1.977858 2.018859,2.492103 0.32836,0.514244 0.91985,1.232646 1.31442,1.596448 0.39458,0.363802 0.8417,0.899583 0.99361,1.190625 0.23281,0.446043 0.16665,0.428408 -0.42117,-0.112266 -0.81662,-0.751127 -0.81445,-0.759926 -0.30049,1.213279 0.21828,0.838038 0.39687,2.138995 0.39687,2.891013 0,0.752019 0.1786,1.794742 0.39688,2.317165 0.21828,0.52242 0.3915,1.221571 0.38494,1.553665 -0.0107,0.539676 -0.0495,0.519501 -0.36518,-0.189939 -0.38266,-0.859851 -0.39295,-0.846828 -0.94111,1.190625 -0.17618,0.654844 -0.65505,1.845469 -1.06416,2.645833 -0.4091,0.800365 -0.812589,2.024698 -0.896649,2.720737 -0.0841,0.696042 -0.31769,1.580187 -0.5192,1.96477 -0.34362,0.65583 -0.35802,0.610113 -0.23186,-0.736362 l 0.1345,-1.435603 -1.44918,1.363747 c -3.34591,3.148658 -8.13854,4.417079 -12.83398,3.396652 z m 2.85031,-26.862384 c 0.28148,-0.150641 0.51178,-0.388766 0.51178,-0.529167 0,-0.423605 -0.69593,-0.287931 -1.35975,0.26509 -0.54245,0.451903 -0.56142,0.521522 -0.14421,0.529167 0.26423,0.0048 0.71071,-0.114449 0.99218,-0.26509 z" />
        </g>
      </svg>
    </Container>
  </Overlay>
)

DSNn.propTypes = {
  size: PropTypes.number,
  color: PropTypes.string,
}

DSNn.defaultProps = {
  size: 90,
  color: '#eee',
}

export default DSNn
