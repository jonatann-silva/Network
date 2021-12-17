import React from 'react'
import styled from 'styled-components'

import Box from '_components/box'
import Text from '_components/text'
import { illustrations } from '_assets/'
import Flex from '_components/flex'

const StyledBox = styled(Box).attrs({
  border: '',
  height: 'auto',
})`
  svg {
    height: 300px;
    width: auto;
  }
  padding: 20px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
`

const Container = styled(Flex)`
  flex-direction: column;
`
const NotFound = () => (
  <Container>
    <StyledBox>
      <illustrations.NotFound />
      <Text textStyle="subtitle" mt={4}>
        página não encontrada
      </Text>
      <Text textStyle="description" my={3}>
        Oops! Parece que a página que você está tentando acessar não existe ou foi movida.
      </Text>
    </StyledBox>
  </Container>
)

export default NotFound
