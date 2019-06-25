import * as React from 'react'
import {Box2, Text} from '../../common-adapters'
import {Background} from '../../common-adapters/text'
import * as Styles from '../../styles'

export type Props = {
  background: Background
  offerAdvancedSendFormbool?: boolean
  onAction?: (() => void) | null
  reviewProofs?: boolean
  sendFailed?: boolean
  text: string
}

const Banner = (props: Props) => (
  <Box2
    direction="vertical"
    fullWidth={true}
    centerChildren={true}
    style={Styles.collapseStyles([
      styles.container,
      {backgroundColor: Styles.backgroundModeToColor[props.background]},
    ])}
  >
    <Text
      center={true}
      type="BodySmallSemibold"
      negative={true}
      style={{
        color: Styles.backgroundModeToTextColor(props.background),
      }}
    >
      {props.text}
      {props.reviewProofs && props.onAction && (
        <Text
          type="BodySmallSemiboldPrimaryLink"
          center={true}
          style={styles.secondText}
          negative={true}
          onClick={props.onAction}
        >
          Please review.
        </Text>
      )}
      {props.offerAdvancedSendFormbool && props.onAction && (
        <Box2 direction={'horizontal'}>
          This user is accepting other assets than XLM Lumens.
          <Text
            type="BodySmallSemiboldPrimaryLink"
            center={true}
            style={styles.secondText}
            negative={true}
            onClick={props.onAction}
          >
            Send other assets?
          </Text>
        </Box2>
      )}
    </Text>
    {props.sendFailed && props.onAction && (
      <Text
        type="BodySmallSemiboldPrimaryLink"
        center={true}
        style={styles.secondText}
        negative={true}
        onClick={props.onAction}
      >
        Review payments
      </Text>
    )}
  </Box2>
)

const styles = Styles.styleSheetCreate({
  container: {
    minHeight: 40,
    padding: Styles.globalMargins.small,
    paddingBottom: Styles.globalMargins.tiny,
    paddingTop: Styles.globalMargins.tiny,
  },
  secondText: {paddingLeft: Styles.globalMargins.xtiny},
})

export default Banner
