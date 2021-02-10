import React from 'react'
import { Translation } from 'react-i18next'
import { useNavigation } from '@react-navigation/native'

import { servicesAuthViaDefault, useAccountServices } from '@berty-tech/store/services'
import { useMsgrContext, useNotificationsInhibitor } from '@berty-tech/store/hooks'
import { PersistentOptionsKeys } from '@berty-tech/store/context'

import SwiperCard from './SwiperCard'
import OnboardingWrapper from './OnboardingWrapper'
import { Routes } from '@berty-tech/navigation'
import { RouteProp } from '@react-navigation/native'

const ServicesAuthBody: React.FC<{ next: () => void; isFromModal: boolean }> = ({
	next,
	isFromModal,
}) => {
	const ctx = useMsgrContext()
	const accountServices = useAccountServices() || []
	const { goBack, reset } = useNavigation()

	React.useEffect(() => {
		if (accountServices.length > 0) {
			next()
		}
	}, [next, accountServices.length])

	return (
		<Translation>
			{(t) => (
				<SwiperCard
					header={t('onboarding.services-auth.header')}
					label={t('onboarding.services-auth.recommended')}
					title={t('onboarding.services-auth.title')}
					description={t('onboarding.services-auth.desc')}
					button={
						accountServices.length > 0
							? undefined
							: {
									text: t('onboarding.services-auth.button'),
									onPress: async () => {
										await servicesAuthViaDefault(ctx)
										await ctx.setPersistentOption({
											type: PersistentOptionsKeys.Configurations,
											payload: {
												...ctx.persistentOptions.configurations,
												notification: {
													...ctx.persistentOptions.configurations.notification,
													state: 'added',
												},
											},
										})
										if (isFromModal) {
											reset({
												index: 0,
												routes: [
													{
														name: Routes.Main.Home,
														params: {
															isFromModal: true,
														},
													},
												],
											})
										} else {
											goBack()
										}
									},
							  }
					}
					skip={{
						text: t('onboarding.notifications.skip'),
						onPress: next,
					}}
				/>
			)}
		</Translation>
	)
}

export const ServicesAuth: React.FC<{ route: RouteProp<any, any> }> = ({ route }) => {
	const isFromModal = route.params?.isFromModal
	useNotificationsInhibitor(() => true)
	const { goBack, reset } = useNavigation()
	const { persistentOptions, setPersistentOption } = useMsgrContext()

	return (
		<OnboardingWrapper>
			<ServicesAuthBody
				isFromModal={isFromModal}
				next={async () => {
					await setPersistentOption({
						type: PersistentOptionsKeys.Configurations,
						payload: {
							...persistentOptions.configurations,
							notification: {
								...persistentOptions.configurations.notification,
								state: 'skipped',
							},
						},
					})
					if (isFromModal) {
						reset({
							index: 0,
							routes: [
								{
									name: Routes.Main.Home,
									params: {
										isFromModal: true,
									},
								},
							],
						})
					} else {
						goBack()
					}
				}}
			/>
		</OnboardingWrapper>
	)
}

export default ServicesAuth
