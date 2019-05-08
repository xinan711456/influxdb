// Libraries
import React, {useMemo, FunctionComponent} from 'react'
import {Table} from '@influxdata/vis'

// Components
import EmptyGraphMessage from 'src/shared/components/EmptyGraphMessage'

// Utils
import {latestValues as getLatestValues} from 'src/shared/utils/latestValues'

interface Props {
  table: Table
  children: (latestValue: number) => JSX.Element

  // If `quiet` is set and a single latest value can't be found, this component
  // will display nothing instead of an empty graph error message
  quiet?: boolean
}

const LatestValueTransform: FunctionComponent<Props> = ({
  table,
  quiet = false,
  children,
}) => {
  const latestValues = useMemo(() => getLatestValues(table), [table])

  if (latestValues.length !== 1 && quiet) {
    return null
  }

  if (latestValues.length === 0) {
    return <EmptyGraphMessage message="No latest value found" />
  }

  if (latestValues.length > 1) {
    return <EmptyGraphMessage message="Multiple latest values found" />
  }

  return children(latestValues[0])
}

export default LatestValueTransform
