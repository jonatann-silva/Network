import { useMutation, useQuery } from '@apollo/client'
import { useLocation, useNavigate, useParams } from '@reach/router'
import React, { useEffect, useState } from 'react'
import styled from 'styled-components'
import { icons } from '_assets/'
import BackLink from '_components/back-link'
import Box from '_components/box'
import Button from '_components/button'
import { useConfirmationDialog } from '_components/confirmation-dialog'
import EmptyState from '_components/empty-state'
import Icon from '_components/icon'
import List from '_components/list'
import { Dragon as Spinner } from '_components/spinner'
import Text from '_components/text'
import { CREATE_INTERFACE, DELETE_INTERFACE } from '_graphql/mutations'
import { GET_NETWORK_WITH_INTERFACES } from '_graphql/queries'
import AdmitNodeModal from '_modals/admit-node'
import PeerCard from './peer-card'

const Container = styled(Box)`
  flex-direction: column;
`

const StyledIcon = styled(Icon)`
  border-radius: 4px;
  background: ${(props) => props.theme.colors.neutralLighter};
  align-items: center;
  justify-content: center;
`

const NetworkDetails = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const { networkId } = useParams()

  const { confirm } = useConfirmationDialog()

  const [isAdmitNodeModalOpen, setAdmitNodeModalOpen] = useState(false)

  const getNetworkQuery = useQuery(GET_NETWORK_WITH_INTERFACES, {
    variables: { id: networkId },
  })

  const [createInterface, createInterfaceMutation] = useMutation(CREATE_INTERFACE, {
    variables: { networkId },
  })

  const [deleteInterface, deleteInterfaceMutation] = useMutation(DELETE_INTERFACE)

  useEffect(() => {
    window.scrollTo(0, 0)
    getNetworkQuery.refetch()
  }, [location])

  const handleAdmitNodeButtonClick = () => {
    setAdmitNodeModalOpen(true)
  }

  const handleAdmitNode = (id) => {
    createInterface({
      variables: {
        networkId,
        nodeId: id,
      },
    })
      .then(() => {
        getNetworkQuery.refetch()
      })
      .catch(() => {})
  }

  const handlePeerCardClick = (id) => {
    navigate(`/ui/clients/${id}`)
  }

  const handlePeerDelete = (id) => {
    confirm({
      title: 'Tem certeza?',
      details: `Isso removerá o nó da rede, destruindo sua interface e conexões.`,
      isDestructive: true,
      onConfirm: () => {
        deleteInterface({
          variables: { id },
        })
          .then(() => {
            getNetworkQuery.refetch()
          })
          .catch(() => {})
      },
    })
  }

  const isLoading =
    getNetworkQuery.loading || createInterfaceMutation.loading || deleteInterfaceMutation.loading

  if (isLoading) {
    return <Spinner />
  }

  const network = getNetworkQuery.data ? getNetworkQuery.data.result : { Interfaces: [] }

  if (isLoading) {
    return <Spinner />
  }

  return (
    <Container>
      <AdmitNodeModal
        isOpen={isAdmitNodeModalOpen}
        onAdmit={handleAdmitNode}
        onClose={() => setAdmitNodeModalOpen(false)}
      />

      <BackLink to="/ui/networks" text="Redes" mb={3} />
      <Box alignItems="center" mb={4}>
        <StyledIcon mr="12px" p={2} icon={<icons.Network />} size="48px" color="neutralDarker" />
        <div>
          <Text textStyle="title" mb={1}>
            {network.Name}
          </Text>
          <Text textStyle="detail">{network.ID}</Text>
        </div>
      </Box>

      <Box px={3} py={2} border="discrete" alignItems="center">
        <Box mr={4}>
          <Text textStyle="strong" fontSize="12px" mr={2}>
           Range IP
          </Text>
          <Text textStyle="detail">{network.AddressRange}</Text>
        </Box>
      </Box>

      <Box alignItems="center" mt={4} mb={3}>
        <Text textStyle="subtitle">Hosts</Text>
        {network.Interfaces.length > 0 && (
          <Button variant="primary" ml="auto" onClick={handleAdmitNodeButtonClick}>
            Permitir host.
          </Button>
        )}
      </Box>
      <List>
        {network.Interfaces.map((el) => (
          <PeerCard
            key={el.ID}
            id={el.NodeID}
            name={el.Name}
            address={el.Address}
            nodeName={el.Node.Name}
            nodeStatus={el.Node.Status}
            hasPublicKey={el.HasPublicKey}
            nodeAdvertiseAddress={el.Node.AdvertiseAddress}
            onClick={() => handlePeerCardClick(el.Node.ID)}
            onDelete={() => handlePeerDelete(el.ID)}
          />
        ))}
      </List>
      {network.Interfaces.length === 0 && (
        <EmptyState
          title="Sem hosts."
          description="Hosts podem ser adicionados a uma rede a qualquer momento."
          image={<Icon my={4} icon={<icons.Host />} size="48px" color="neutral" />}
          extra={
            <Box alignItems="center" mt="24px">
              <Button variant="primary" onClick={handleAdmitNodeButtonClick}>
                Permitir Host
              </Button>
            </Box>
          }
        />
      )}
    </Container>
  )
}

NetworkDetails.propTypes = {}

export default NetworkDetails
