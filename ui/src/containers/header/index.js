import React from 'react'
import styled from 'styled-components'
import { grid, space, color, shadow, border } from 'styled-system'

import Brand from '_containers/side-nav/brand'
import Flex from '_components/flex'
import Link from '_components/link'

export const Container = styled.div`
  display: flex;

  align-items: center;
  justify-content: space-between;

  border-bottom: 1px solid ${(props) => props.theme.colors.neutralLighter};
  position: fixed;

  top: 0;
  right: 0;
  left: 0;

  z-index: 199;

  ${border}
  ${shadow}
  ${color}
  ${space}
  ${grid}
`

Container.defaultProps = {
  backgroundColor: 'white',
  boxShadow: 'light',
  border: 'dark',
}

const StyledLink = styled(Link)`
  margin: auto;
  padding-right: 18px;
  :hover {
    color: ${({ theme }) => theme.colors.neutralDarker};
  }
`

const Header = (props) => (
  <Container {...props}>
    <Brand />
    <Flex>
      <StyledLink to="settings/tokens" color="neutralDark">
        ACL Tokens
      </StyledLink>
    </Flex>
  </Container>
)

Header.defaultProps = {
  padding: 3,
}

export default Header
