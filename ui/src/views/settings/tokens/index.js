import { useQuery } from '@apollo/client'
import { useLocation } from '@reach/router'
import { useFormik } from 'formik'
import React, { useEffect, useState } from 'react'
import styled from 'styled-components'
import * as Yup from 'yup'
import Box from '_components/box'
import Button from '_components/button'
import Flex from '_components/flex'
import TextInput from '_components/inputs/text-input'
import Text from '_components/text'
import { GET_SELF_TOKEN } from '_graphql/queries'
import useLocalStorage from '_utils/use-localstorage'

const Container = styled(Flex)`
  flex-direction: column;
`

const StyledButton = styled(Button).attrs({
  height: '40px',
  borderRadius: 3,
  mr: 3,
})`
  display: block;
  white-space: nowrap;
`

const TokensView = () => {
  const location = useLocation()

  const [token, setToken] = useState(undefined)
  const [tokenSecret, setTokenSecret] = useLocalStorage('dsn.settings.acl.token', undefined)
  const [isAuthenticated, setAuthenticated] = useState(undefined)

  const formik = useFormik({
    initialValues: {
      secret: tokenSecret || '',
    },
    validationSchema: Yup.object().shape({
      secret: Yup.string(),
    }),
  })

  const handleGetSelfQueryCompleted = (data) => {
    if (data.result !== null) {
      setToken(data.result)
    }
  }

  const handleGetSelfQueryError = (errors) => {
    if (errors.networkError.statusCode === 404) {
      setToken(undefined)
      setTokenSecret(undefined)
      formik.resetForm({ secret: '' })
    }
  }

  const getSelfQuery = useQuery(GET_SELF_TOKEN, {
    onCompleted: handleGetSelfQueryCompleted,
    onError: handleGetSelfQueryError,
  })

  useEffect(() => {}, [location])

  const handleSetTokenButtonClick = () => {
    setTokenSecret(formik.values.secret)
    getSelfQuery
      .refetch()
      .then((res) => {
        handleGetSelfQueryCompleted(res.data)
        setAuthenticated(true)
      })
      .catch((errors) => {
        handleGetSelfQueryError(errors)
        setAuthenticated(false)
      })
  }

  const handleClearTokenButtonClick = () => {
    setToken(null)
    setTokenSecret(undefined)
    setAuthenticated(undefined)
    formik.resetForm({ secret: '' })
  }

  return (
    <Container>
      <Box mb={3} mr="auto">
        <Text textStyle="title">Tokens de controle de acesso</Text>
      </Box>
      <Text textStyle="bodyText">
        Ao fornecer um token secreto, cada pedido futuro será autenticado, potencialmente
        autorizando o acesso de leitura a informações adicionais.
      </Text>
      <Box my={3} p={3} border="discrete">
        <div>
          <Text textStyle="subtitle" mb={2} mr="auto" fontSize="18px">
            Armazenamento de token
          </Text>
          <Text textStyle="bodyText">
            Os tokens são armazenados do lado do cliente no armazenamento local.Isso persistirá ao seu token
            sessões.Você pode limpar manualmente o seu token aqui.
          </Text>
        </div>
        <StyledButton variant="primaryInverted" mx={4} onClick={handleClearTokenButtonClick}>
          Limpar Token.
        </StyledButton>
      </Box>
      <>
        <Text my={3}>Chave</Text>
        <TextInput
          id="secret"
          name="secret"
          {...formik.getFieldProps('secret')}
          placeholder="XXXXXXXX-XXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
        />

        <Text textStyle="detail" mt={2}>
          Enviado a cada pedido para determinar a autorização
        </Text>
        <StyledButton my={3} mb={4} px={3} onClick={handleSetTokenButtonClick}>
          Configurar token
        </StyledButton>
      </>

      {isAuthenticated === true && (
        <Box mt={2} mb={3} p={3} bg="green" color="white">
          <div>
            <Text textStyle="subtitle" mb={2} mr="auto" fontSize="18px" color="white">
              Token autenticado!
            </Text>
            <Text textStyle="bodyText" color="white">
              Seu token é válido e autorizado para as seguintes políticas.
            </Text>
          </div>
        </Box>
      )}

      {isAuthenticated === false && (
        <Box mt={2} mb={3} p={3} bg="danger">
          <div>
            <Text textStyle="subtitle" mb={2} mr="auto" fontSize="18px" color="white">
              Token não autenticado!
            </Text>
            <Text textStyle="bodyText" color="white">
              O chave que você forneceu é inválida ou não corresponde a um token existente.
            </Text>
          </div>
        </Box>
      )}

      {token && (
        <>
          <Text textStyle="subtitle" fontSize="18px" mb={2} mt={2}>
            Token: {token.Name}
          </Text>
          <Text textStyle="bodyText" mb={2}>
            ID: {token.ID}
          </Text>
          {token.Secret && <Text textStyle="bodyText" mb={2}>
            Chave: {token.Secret}
          </Text>}
          <Text textStyle="subtitle" fontSize="18px" mb={2} mt={3}>
            Políticas
          </Text>
          <Box flexDirection="column" border="discrete" mb={3} p={3}>
            {token.Type === 'management'
              ? `The management token has all the permissions`
              : token.Policies.map((el) => (
                  <Text textStyle="body" key={el} mb={2}>
                    {el}
                  </Text>
                ))}
          </Box>
        </>
      )}
    </Container>
  )
}

export default TokensView
