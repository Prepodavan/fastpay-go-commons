package cc_errors

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

const (

	/** Код ошибки при успешном выполнении операции */
	ErrorSuccess = 0

	/** Аккаунт существует */
	ErrorAccountExist = 60100

	/** Идентификатор аккаунта существует */
	ErrorIdentifierExist = 60101

	/** Изменение статуса идентифицированного аккаунта доступно банку, который определил его юридический тип */
	ErrorEditIdentifiedAccountStatusForbidden = 60102

	/** Банк с таким БИК уже существует */
	ErrorBankWithSuchBikExist = 60103

	/** Валюта существует */
	ErrorCurrencyExist = 60104

	/** Валюта не может быть созданна так как не существует чейнкодов для работы с ней */
	ErrorCurrencyCreateConflict = 60105

	/** Требуется передать идентификатор получателя */
	ErrorReceiverIdentificatorRequired = 60106

	/** Списание средств недоступно для клиентского банка  */
	ErrorClientBankWithdrawForbidden = 60107

	/** Операция отмены списания средств недоступна для клиентского банка */
	ErrorClientBankWithdrawRejectForbidden = 60108

	/** Операция подтверждения списания средств недоступна для клиентского банка */
	ErrorClientBankWithdrawConfirmForbidden = 60109

	/** Аккаунт не существует */
	ErrorAccountNotExist = 60110

	/** Идентификатор аккаунта не существует */
	ErrorIdentifierNotExist = 60111

	/** Банк не существует */
	ErrorBankNotExist = 60112

	/** Валюта не существует */
	ErrorCurrencyNotExist = 60114

	/** Клиентский банк заблокирован */
	ErrorClientBankIsBlocked = 60117

	/** Аккаунт не является эскроу-счётом */
	ErrorAccountNotEscrow = 60118

	/** Аккаунт не идентифицирован */
	ErrorAccountNotIdentified = 60119 //+

	/** Аккаунт недоступен */
	ErrorAccountNotAvailable = 60120

	/** Аккаунт заблокирован */
	ErrorAccountIsBlocked = 60121

	/** Банк недоступен */
	ErrorBankNotAvailable = 60122

	/** Банк заблокирован */
	ErrorBankIsBlocked = 60123

	/** Банк с указанными параметрами уже существует в системе */
	ErrorBankExist = 60124

	/** Контракт не найден */
	ErrorContractNotExists = 60125

	/** Контракт с указанными параметрами уже существует в системе */
	ErrorContractExists = 60126

	/** Информация о клиенте банка не найдена */
	ErrorCustomerNotFound = 60128

	/** Клиентский банк не доступен */
	ErrorClientBankNotAvailable = 60129

	/** Редактирование контракта недоступно */
	ErrorEditExchangeContractForbidden = 60130

	/** Требуется передать идентификатор клиента */
	ErrorClientIdentifierRequired = 60131

	/** Аккаунт отправителя не доступен */
	ErrorAccountSenderNotAvailable = 60132

	/** Аккаунт получателя не доступен */
	ErrorAccountReceiverNotAvailable = 60133

	/** Изменение статуса неидентифицированного аккаунта доступно банку владельцу */
	ErrorEditUnidentifiedAccountStatusForbidden = 60134

	/** Некорректный код валюты отправителя */
	ErrorIncorrectCurrencyCodeSenderAccount = 60135

	/** Некорректный код валюты получателя */
	ErrorIncorrectCurrencyCodeReceiverAccount = 60136

	/** Переданный тип контракта недоступен указанному клиентскому банку */
	ErrorClientBankContractTypeForbidden = 60137

	/** Переданный тип контракта недоступен указанному аккаунту */
	ErrorAccountContractTypeForbidden = 60138

	/** Адрес клиентского банка не соответствует заявленному в аккаунте */
	ErrorClientBankAddressNotMatchAccount = 60139

	/** Изменение доступных типов контрактов клиентского банка доступно только его банку владельцу */
	ErrorUpdateClientBankContractTypesForbidden = 60140

	/** Ошибка проверки сигнатуры */
	ErrorSignVerify = 60200

	/** Недостаточно средств */
	ErrorFundsNotEnough = 60240

	/** Ошибка валидации */
	ErrorValidateDefault = 60300

	/** Неверный формат адреса аккаунта */
	ErrorAddressNotFollowingRegex = 60301

	/** Требуется передать адрес аккаунта */
	ErrorAddressNotPassed = 60302

	/** Неверный формат идентификатора аккаунта */
	ErrorIdentifierNotFolowingRegex = 60303

	/** Требуется передать идентификатор */
	ErrorIdentifierNotPassed = 60304

	/** Неверный формат типа идентификации аккаунта */
	ErrorIdentityTypeIncorrect = 60305

	/** Неверный формат статуса аккаунта */
	ErrorStateIncorrect = 60306

	/** Неверный формат юридического типа аккаунта */
	ErrorJuridicalTypeIncorrect = 60307

	/** Неверный формат роли аккаунта */
	ErrorRoleIncorrect = 60308

	/** Требуется передать имя канала чейнкода */
	ErrorChannelNotPassed = 60309

	/** Требуется передать имя чейнкода */
	ErrorNameNotPassed = 60310

	/** Требуется передать версию чейнкода */
	ErrorVersionNotPassed = 60311

	/** Значение суммы не может быть отрицательным */
	ErrorAmountNegative = 60312

	/** Требуется передать сумму */
	ErrorAmountNotPassed = 60313

	/** Требуется передать тип идентификации аккаунта */
	ErrorIdentityTypeNotPassed = 60314

	/** Значение не может быть отрицательным */
	ErrorValueNegative = 60315

	/** Требуется передать значение */
	ErrorValueNotPassed = 60316

	/** Значение кода валюты должно входить в диапазон от 0 до 999 */
	ErrorCurrencyCodeRange = 60317

	/** Требуется передать код валюты */
	ErrorCurrencyCodeNotPassed = 60318

	/** Значение баланса не может быть отрицательным */
	ErrorBalanceNegative = 60319

	/** Требуется передать баланс */
	ErrorBalanceNotPassed = 60320

	/** Неверный формат сообщения */
	ErrorMsgHashNotFolowingRegex = 60321

	/** Требуется передать сообщение */
	ErrorMsgHashNotPassed = 60322

	/** Неверный формат R сигнатуры */
	ErrorSigRNotFolowingRegex = 60323

	/** Требуется передать R сигнатуры */
	ErrorSigRHashNotPassed = 60324

	/** Неверный формат S сигнатуры */
	ErrorSigSNotFolowingRegex = 60325

	/** Требуется передать S сигнатуры */
	ErrorSigSNotPassed = 60326

	/** Неверное значение V сигнатуры */
	ErrorSigVIncorrect = 60327

	/** Неверный формат V сигнатуры */
	ErrorSigVNotFollowingRegex = 60328

	/** Требуется передать идентификатор транзакции */
	ErrorTxIdSNotPassed = 60331

	/** Требуется передать время */
	ErrorTimestampNotPassed = 60332

	/** Значение времени не может быть отрицательным */
	ErrorTimestampNegative = 60333

	/** Требуется передать тип транзакции */
	ErrorTxTypeNotPassed = 60334

	/** Неверный формат типа транзакции */
	ErrorTxTypeIncorrect = 60335

	/** Значение размера страницы меньше минимально допустимого */
	ErrorPageSizeMinValue = 60336

	/** Требуется передать значение символа валюты */
	ErrorSymbolNotPassed = 60337

	/** Требуется передать значение количества знаков после запятой */
	ErrorDecimalsNotPassed = 60338

	/** Значение количества знаков должно входить в диапазон от 0 до 8 */
	ErrorDecimalsRange = 60339

	/** Значение количества знаков после запятой больше максимально допустимого */
	ErrorDecimalsMaxValue = 60340

	/** Требуется передать значение признака активности */
	ErrorEnabledNotPassed = 60341

	/** Требуется передать идентификатор аккаунта */
	ErrorIdentifiersNotPassed = 60342

	/** Количество идентификаторв больше максимально допустимого */
	ErrorIdentifiersMaxValue = 60343

	/** Требуется передать статус аккаунта */
	ErrorStateNotPassed = 60344

	/** Требуется передать юридический тип аккаунта */
	ErrorJuridicalTypeNotPassed = 60345

	/** Требуется передать юридический тип банка-владельца аккаунта */
	ErrorJuridicalTypeBankSetterNotPassed = 60346

	/** Требуется передать тип аккаунта */
	ErrorAccountTypeNotPassed = 60347

	/** Неверное значение типа аккаунта */
	ErrorAccountTypeIncorrect = 60348

	/** Требуется передать роль */
	ErrorRoleNotPassed = 60349

	/** Требуется передать ID */
	ErrorIdNotPassed = 60350

	/** Неверный формат ID */
	ErrorIdNotFolowingRegex = 60351

	/** Требуется передать номер счета */
	ErrorNumberNotPassed = 60352

	/** Неверный формат номера счета */
	ErrorNumberInvoiceNotFolowingRegex = 60353

	/** Неверное значение состояния счета */
	ErrorInvoiceStateIncorrect = 60354

	/** Требуется передать состояние счета */
	ErrorInvoiceStateNotPassed = 60355

	/** Значение лимита эмисси не может быть отрицательным */
	ErrorIssueLimitNegative = 60356

	/** Неверное значение статуса транзакции */
	ErrorStatusTransactionIncorrect = 60357

	/** Требуется передать статус транзакции */
	ErrorStatusTransactionPassed = 60358

	/** MSP ID не совпадают */
	ErrorMspIdNotMatch = 60359

	/** Адреса не совпадают */
	ErrorAddressNotMatch = 60360

	/** Неверный формат публичного ключа аккаунта */
	ErrorPublicKeyNotFolowingRegex = 60361

	/** Требуется передать публичный ключ аккаунта */
	ErrorPublicKeyNotPassed = 60362

	/** Требуется передать дополнительную конфигурацию банка */
	ErrorConfNotPassed = 60363

	/** Требуется передать адрес банка */
	ErrorBankAddressNotPassed = 60364

	/** Требуется передать код страны */
	ErrorCountryCodeNotPassed = 60365

	/** Требуется передать идентификатор транзакции */
	ErrorTransactionIdNotPassed = 60366

	/** Неверный формат идентификатора или адреса аккаунта */
	ErrorAddressOrIdentifierNotFolowingRegex = 60367

	/** Требуется передать идентификатор или адрес аккаунта */
	ErrorAddressOrIdentifierNotPassed = 60368

	/** Требуется передать отображаемое имя банка */
	ErrorBankDisplayNameNotPassed = 60369

	/** Ошибка при получении MSP_ID */
	ErrorGetMspId = 60370

	/** Время подписи истекло */
	ErrorExpirationSign = 60371

	/** Подпись ранее была использована */
	ErrorSignUsed = 60372

	/** Некорректный код валюты аккаунта */
	ErrorIncorrectCurrencyCodeAccount = 60387

	/* Количество переданного актива не соответствует переданному курсу */
	ErrorAmountNotCorrespondPrice = 60403

	/* Безопасная сделка не найдена */
	ErrorSafeDealNotFound = 60404

	/* Участник безопасной сделки проголосовал ранее */
	ErrorCounterpartVoted = 60405

	/* Некорректная сумма эскроу-счета */
	ErrorIncorrectEscrowAmount = 60406

	/* Адрес аккаунта уже принимает участие в сделке */
	ErrorAccountAlreadyInDeal = 60407

	/** Информация о доступных платформах не найдена */
	ErrorAvailablePlatformsNotFound = 60411

	/** Превышен операционный лимит перевода средств */
	ErrorLimitOperation = 60420

	/** Превышен дневной лимит перевода средств */
	ErrorLimitDaily = 60430

	/** Превышен месячный лимит перевода средств */
	ErrorLimitMonthly = 60440

	/** Превышен лимит баланса */
	ErrorLimitBalance = 60450

	/** Адреса отправителя и получателя совпадают */
	ErrorAddressFromToEqual = 60461

	/** Невозможно заблокировать банк, который является единственным владельцем системы */
	ErrorSingleBankOwnerCannotBeBlocked = 60462

	/** История по трансграничному переводу не найдена */
	ErrorCrossTransactionHistoryNotFound = 60463

	/** Невозможно выполнить операцию с текущим статусом транзакции */
	ErrorStatusConflict = 60464

	/** Неверный адрес аккаунта для выполнения вывода средств в рамках трансграничного перевода */
	ErrorCrossTransactionWithdrawIncorrectAddress = 60465

	/** Неверная сумма для выполнения вывода средств в рамках трансграничномого перевода */
	ErrorCrossTransactionWithdrawIncorrectAmount = 60466

	/** Другая ошибка */
	ErrorDefault = 60500

	/** Счет не найден */
	ErrorInvoiceNotExist = 60501

	/** Cуммы cчета не совпадают */
	ErrorInvoiceTransactionAmount = 60504

	/** Попытка оплатить ранее оплаченный счет */
	ErrorInvoiceTransactionPaid = 60505

	/** Попытка оплатить счета с истекшим сроком жизни */
	ErrorInvoiceTransactionExpired = 60506

	/** Счёт с указанными номером уже существует */
	ErrorInvoiceExist = 60507

	/** Ошибка обновления статуса счета */
	ErrorInvoiceUpdate = 60508

	/** Указанный адрес платильщика не совпадает с тем, который указан в счете */
	ErrorPayerNotEqual = 60509

	/** Ошибка сохрания данных в БД */
	ErrorPutState = 60510

	/** Ошибка при получении данных в БД */
	ErrorGetState = 60511

	/** Ошибка при удалении данных в БД */
	ErrorDeleteState = 60512

	/** Ошибка создании композитного ключа */
	ErrorCreateCompositeKey = 60513

	/** Ошибка при серилизации JSON */
	ErrorJsonMarshal = 60514

	/** Ошибка при десерилизации JSON */
	ErrorJsonUnmarshal = 60515

	/** Ошибка при получении времени создания транзакции */
	ErrorGetTxTime = 60516

	/** Не достаточно прав для вызова метода смарт-контракта */
	ErrorForbidden = 60601

	/** Эскроу-счёт не найден */
	ErrorDepositNotFound = 60602

	/** Метод чейнкода недоступен для использования */
	ErrorForbiddenForExternal = 60603

	/** Ошибка валидации контракта */
	ErrorContractValidate = 60606

	/** Адрес банка в сертификате не совпадает */
	ErrorCertificateBankAddressNotMatch = 60609

	/** Клиринг. Расхождение требований и обязательств */
	ErrorClearingBadClaimsLiabilities = 60701

	/** Клиринг. Информация о клиринге не найдена */
	ErrorClearingInfoNotFound = 60702

	/** Сертификат. Отсутвует атрибут address в сертификате */
	ErrorCertificateBankAddressNotFound = 60801

	/** Клиентский банк с указанными параметрами уже существует в системе */
	ErrorClientBankExists = 60802

	/** Информация о клиентском банке не найдена */
	ErrorClientBankNotFound = 60803

	/** Опорный банк(отправитель) не является владельцем клиентского банка */
	ErrorClientBankOwnerNotEqualSender = 60804

	/** Некорректное значение типа лимита */
	ErrorLimitTypeIncorrect = 60805

	/** Адрес банка отправителя не совпадает с адресом технического аккаунта */
	ErrorAccountTechnicalNotEqlSender = 60806

	/** Переданный номер счета не совпадает с номером счета, который предназначен для пополнения эскроу-счета */
	ErrorPaymentInvoiceNumberNotEqlForEscrowAccount = 60807

	/** Аккаунт для создания арбитра недоступен**/
	ErrorAccountNotNotAvailableForCreateArbitrator = 61001

	/* Арбитр ранее уже был добавлен */
	ErrorArbitratorExist = 61002

	/* Арбитра не существует */
	ErrorArbitratorNotExist = 61003

	/* Арбитр не может частично исполнить сделку */
	ErrorArbitratorCannotPartialPerform = 61004

	/* В сделке не задействован физический актив */
	ErrorPhysicalAssetNotPresent = 61005

	/* Внести данные о частичном исполнении сделки может только первый заявитель об исполнении сделки */
	ErrorPartialPerformOrder = 61006

	/* Требуется указать адрес аккаунта требуемого актива */
	ErrorOfferAccountSellOnly = 61007

	/* Переданный контракт не является предложением */
	ErrorContractNotOffer = 61008

	/* В сделке не указан адрес инициатора предложения */
	ErrorOfferOwnerMissing = 61009

	/* Коды валют не соответствует заявленным в предложении */
	ErrorCurrencyNotMatchOffer = 61010

	/* Курс не соответствует заявленному в предложении */
	ErrorPriceNotMatchOffer = 61011

	/* Количество актива не соответствуют интервалу минимального и максимального количества */
	ErrorAssetAmountNotCorrespondMaxMinInterval = 61012

	/* Количество актива инициатора не соответствует минимальному и максимальному значениям */
	ErrorAmountInitiatorNotCorrespondMaxMinInterval = 61013

	/* Количество актива акцептанта не соответствует минимальному и максимальному значениям */
	ErrorAmountAcceptorNotCorrespondMaxMinInterval = 61014

	/* Физический актив не может начисляться на эскроу-счет */
	ErrorPhysicalAssetEscrow = 61015

	/** Действие запрещено на текущем этапе сделки */
	ErrorChangeStateSafeDealForbidden = 61016

	/** Действие запрещено, до тех пор, когда все приглашения не будут приняты */
	ErrorNotAcceptedAllInvitations = 61017

	/** Невалидная конфигурация машины состояний для безопасной сделки */
	ErrorSettingsFSMSafeDeal = 61018

	/** Сделка с переданным идентификатором уже существует */
	ErrorSafeDealExists = 61019

	/** Неверный адрес акцептанта безопасной сделки */
	ErrorNotValidAddressAcceptor = 61020

	/** Адрес аккаунта не является контрагентом сделки */
	ErrorAddressNotCounterpartDeal = 61021

	/** Адрес аккаунта не является участником сделки */
	ErrorAddressNotParticipantOrCounterpartDeal = 61022

	/** Запрещено приглашать арбитра аккаунтом, который не является участником сделки */
	ErrorInviteArbitratorByAddressNotCounterpartDeal = 61023

	/** Запрещено принимать заявку аккаунтом, который не является участником сделки  */
	ErrorConfirmInvitationByAddressNotCounterpartDeal = 61024

	/** Контрагент ранее проголосовал за расторжение(заключение) контракта */
	ErrorCounterpartMadeChoice = 61025

	/** Контрагент ранее проголосовал за исполнение контракта */
	ErrorCounterpartMadePerformChoice = 61026

	/** Контрагент ранее проголосовал за расторжение контракта */
	ErrorCounterpartMadeCancelChoice = 61027

	/** Операция запрещена для аккаунта, который не является акцептантом или арбитром */
	ErrorNotValidAddressAcceptorOrArbitrator = 61028

	/** Операция запрещена для аккаунта, который не является инициатором */
	ErrorNotValidAddressInitiator = 61029

	/** Переданный контрагент ранее был приглашён */
	ErrorCounterpartWasInvited = 61030

	/** Участник не был приглашен */
	ErrorParticipantNotInvited = 61031

	/** Операция запрещена для аккаунта, который не является инициатором сделки */
	ErrorNotValidAddressOwner = 61032

	/** Приглашено достаточно участников с такой ролью */
	ErrorEnoughParticipants = 61033

	/** Не может быть приглашено больше одного арбитра */
	ErrorEnoughArbitrators = 61034

	/** Не может быть приглашено больше одного акцептанта */
	ErrorEnoughAcceptors = 61035

	/** Некорректная сумма сделки */
	ErrorIncorrectSafeDealAmount = 61036

	/** Эскроу-счет безопасной сделки не найден или был пополнен ранее */
	ErrorSafeDealEscrowNotFound = 61037

	/** Адрес инициатора не может совпадать с адресом акцептанта */
	ErrorCreateSafeDealSameParticipantAddresses = 61038

	/** Сумма обязательной комиссии не может превышать сумму комиссии, которая оплачивается за участие арбитра */
	ErrorObligatoryCommissionExceedsFullCommission = 61039

	/** Эскроу-счет не задействован в текущей безопасной сделке */
	ErrorEscrowAccountNotBelongSafeDeal = 61040

	/** Эскроу-аккаунт имеет некорректный код валюты для оплаты эскроу-счёта */
	ErrorEscrowAccountIncorrectCurrencyCodeForPaymentDeposit = 61041

	/** Аккаунт инициатора сделки недоступен */
	ErrorOwnerSafeDealAccountNotAvailable = 61042

	/** Аккаунт инициатора сделки недоступен */
	ErrorInitiatorSafeDealAccountNotAvailable = 61043

	/** Аккаунт акцептанта сделки недоступен */
	ErrorAcceptorSafeDealAccountNotAvailable = 61044

	/** Аккаунт арбитра сделки недоступен */
	ErrorArbitratorSafeDealAccountNotAvailable = 61045

	/** Аккаунт участника сделки недоступен */
	ErrorParticipantSafeDealAccountNotAvailable = 61046

	/** Аккаунт приглашенного участником сделки недоступен */
	ErrorInvitedSafeDealAccountNotAvailable = 61047

	/** Некорректный код валюты аккаунта инициатора */
	ErrorIncorrectCurrencyCodeAccountInitiator = 61048

	/** Некорректный код валюты аккаунта акцептора */
	ErrorIncorrectCurrencyCodeAccountAcceptor = 61049
)

var ErrorCodeMessagesMap = map[int]string{
	ErrorAccountExist:                                        "Аккаунт существует",
	ErrorIdentifierExist:                                     "Идентификатор аккаунта существует",
	ErrorEditIdentifiedAccountStatusForbidden:                "Изменение статуса идентифицированного аккаунта доступно банку, который определил его юридический тип",
	ErrorBankWithSuchBikExist:                                "Банк с таким БИК уже существует",
	ErrorCurrencyExist:                                       "Валюта существует",
	ErrorCurrencyCreateConflict:                              "Валюта не может быть созданна так как не существует чейнкодов для работы с ней",
	ErrorReceiverIdentificatorRequired:                       "Требуется передать идентификатор получателя",
	ErrorClientBankWithdrawForbidden:                         "Списание средств недоступно для клиентского банка",
	ErrorClientBankWithdrawRejectForbidden:                   "Операция отмены списания средств недоступна для клиентского банка",
	ErrorClientBankWithdrawConfirmForbidden:                  "Операция подтверждения списания средств недоступна для клиентского банка",
	ErrorAccountNotExist:                                     "Аккаунт не существует",
	ErrorIdentifierNotExist:                                  "Идентификатор аккаунта не существует",
	ErrorBankNotExist:                                        "Банк не существует",
	ErrorCurrencyNotExist:                                    "Валюта не существует",
	ErrorClientBankIsBlocked:                                 "Клиентский банк заблокирован",
	ErrorAccountNotIdentified:                                "Аккаунт не идентифицирован",
	ErrorAccountNotAvailable:                                 "Аккаунт недоступен",
	ErrorAccountIsBlocked:                                    "Аккаунт заблокирован",
	ErrorBankNotAvailable:                                    "Банк недоступен",
	ErrorBankIsBlocked:                                       "Банк заблокирован",
	ErrorBankExist:                                           "Банк с указанными параметрами уже существует в системе",
	ErrorContractNotExists:                                   "Контракт не найден",
	ErrorContractExists:                                      "Контракт с указанными параметрами уже существует в системе",
	ErrorCustomerNotFound:                                    "Информация о клиенте банка не найдена",
	ErrorClientBankNotAvailable:                              "Клиентский банк недоступен",
	ErrorEditExchangeContractForbidden:                       "Редактирование контракта недоступно",
	ErrorClientIdentifierRequired:                            "Требуется передать идентификатор клиента",
	ErrorAccountSenderNotAvailable:                           "Аккаунт отправителя недоступен",
	ErrorAccountReceiverNotAvailable:                         "Аккаунт получателя недоступен",
	ErrorEditUnidentifiedAccountStatusForbidden:              "Изменение статуса неидентифицированного аккаунта доступно банку владельцу",
	ErrorIncorrectCurrencyCodeSenderAccount:                  "Некорректный код валюты отправителя",
	ErrorIncorrectCurrencyCodeReceiverAccount:                "Некорректный код валюты получателя",
	ErrorClientBankContractTypeForbidden:                     "Переданный тип контракта недоступен указанному клиентскому банку",
	ErrorAccountContractTypeForbidden:                        "Переданный тип контракта недоступен указанному аккаунту",
	ErrorClientBankAddressNotMatchAccount:                    "Адрес клиентского банка не соответствует заявленному в аккаунте",
	ErrorUpdateClientBankContractTypesForbidden:              "Изменение доступных типов контрактов клиентского банка доступно только его банку владельцу",
	ErrorSignVerify:                                          "Ошибка проверки сигнатуры",
	ErrorFundsNotEnough:                                      "Недостаточно средств",
	ErrorValidateDefault:                                     "Ошибка валидации",
	ErrorAddressNotFollowingRegex:                            "Неверный формат адреса аккаунта",
	ErrorAddressNotPassed:                                    "Требуется передать адрес аккаунта",
	ErrorIdentifierNotFolowingRegex:                          "Неверный формат идентификатора аккаунта",
	ErrorIdentifierNotPassed:                                 "Требуется передать идентификатор",
	ErrorIdentityTypeIncorrect:                               "Неверный формат типа идентификации аккаунта",
	ErrorStateIncorrect:                                      "Неверный формат статуса аккаунта",
	ErrorJuridicalTypeIncorrect:                              "Неверный формат юридического типа аккаунта",
	ErrorRoleIncorrect:                                       "Неверный формат роли аккаунта",
	ErrorChannelNotPassed:                                    "Требуется передать имя канала чейнкода",
	ErrorNameNotPassed:                                       "Требуется передать имя чейнкода",
	ErrorVersionNotPassed:                                    "Требуется передать версию чейнкода",
	ErrorAmountNegative:                                      "Значение суммы не может быть отрицательным",
	ErrorAmountNotPassed:                                     "Требуется передать сумму",
	ErrorIdentityTypeNotPassed:                               "Требуется передать тип идентификации аккаунта",
	ErrorValueNegative:                                       "Значение не может быть отрицательным",
	ErrorValueNotPassed:                                      "Требуется передать значение",
	ErrorCurrencyCodeRange:                                   "Значение кода валюты должно входить в диапазон от 0 до 999",
	ErrorCurrencyCodeNotPassed:                               "Требуется передать код валюты",
	ErrorBalanceNegative:                                     "Значение баланса не может быть отрицательным",
	ErrorBalanceNotPassed:                                    "Требуется передать баланс",
	ErrorMsgHashNotFolowingRegex:                             "Неверный формат сообщения",
	ErrorMsgHashNotPassed:                                    "Требуется передать сообщение",
	ErrorSigRNotFolowingRegex:                                "Неверный формат R сигнатуры",
	ErrorSigRHashNotPassed:                                   "Требуется передать R сигнатуры",
	ErrorSigSNotFolowingRegex:                                "Неверный формат S сигнатуры",
	ErrorSigSNotPassed:                                       "Требуется передать S сигнатуры",
	ErrorSigVIncorrect:                                       "Неверное значение V сигнатуры",
	ErrorSigVNotFollowingRegex:                               "Неверный формат V сигнатуры",
	ErrorTxIdSNotPassed:                                      "Требуется передать идентификатор транзакции",
	ErrorTimestampNotPassed:                                  "Требуется передать время",
	ErrorTimestampNegative:                                   "Значение времени не может быть отрицательным",
	ErrorTxTypeNotPassed:                                     "Требуется передать тип транзакции",
	ErrorTxTypeIncorrect:                                     "Неверный формат типа транзакции",
	ErrorPageSizeMinValue:                                    "Значение размера страницы меньше минимально допустимого",
	ErrorSymbolNotPassed:                                     "Требуется передать значение символа валюты",
	ErrorDecimalsNotPassed:                                   "Требуется передать значение количества знаков после запятой",
	ErrorDecimalsRange:                                       "Значение количества знаков должно входить в диапазон от 0 до 8",
	ErrorDecimalsMaxValue:                                    "Значение количества знаков после запятой больше максимально допустимого",
	ErrorEnabledNotPassed:                                    "Требуется передать значение признака активности",
	ErrorIdentifiersNotPassed:                                "Требуется передать идентификатор аккаунта",
	ErrorIdentifiersMaxValue:                                 "Количество идентификаторв больше максимально допустимого",
	ErrorStateNotPassed:                                      "Требуется передать статус аккаунта",
	ErrorJuridicalTypeNotPassed:                              "Требуется передать юридический тип аккаунта",
	ErrorJuridicalTypeBankSetterNotPassed:                    "Требуется передать юридический тип банка-владельца аккаунта",
	ErrorAccountTypeNotPassed:                                "Требуется передать тип аккаунта",
	ErrorAccountTypeIncorrect:                                "Неверное значение типа аккаунта",
	ErrorRoleNotPassed:                                       "Требуется передать роль",
	ErrorIdNotPassed:                                         "Требуется передать ID",
	ErrorIdNotFolowingRegex:                                  "Неверный формат ID",
	ErrorNumberNotPassed:                                     "Требуется передать номер счета",
	ErrorNumberInvoiceNotFolowingRegex:                       "Неверный формат номера счета",
	ErrorInvoiceStateIncorrect:                               "Неверное значение состояния счета",
	ErrorInvoiceStateNotPassed:                               "Требуется передать состояние счета",
	ErrorIssueLimitNegative:                                  "Значение лимита эмисси не может быть отрицательным",
	ErrorStatusTransactionIncorrect:                          "Неверное значение статуса транзакции",
	ErrorStatusTransactionPassed:                             "Требуется передать статус транзакции",
	ErrorMspIdNotMatch:                                       "MSP ID не совпадают",
	ErrorAddressNotMatch:                                     "Адреса не совпадают",
	ErrorPublicKeyNotFolowingRegex:                           "Неверный формат публичного ключа аккаунта",
	ErrorPublicKeyNotPassed:                                  "Требуется передать публичный ключ аккаунта",
	ErrorConfNotPassed:                                       "Требуется передать дополнительную конфигурацию банка",
	ErrorBankAddressNotPassed:                                "Требуется передать адрес банка",
	ErrorCountryCodeNotPassed:                                "Требуется передать код страны",
	ErrorTransactionIdNotPassed:                              "Требуется передать идентификатор транзакции",
	ErrorAddressOrIdentifierNotFolowingRegex:                 "Неверный формат идентификатора или адреса аккаунта",
	ErrorAddressOrIdentifierNotPassed:                        "Требуется передать идентификатор или адрес аккаунта",
	ErrorBankDisplayNameNotPassed:                            "Требуется передать отображаемое имя банка",
	ErrorGetMspId:                                            "Ошибка при получении MSP_ID",
	ErrorExpirationSign:                                      "Время подписи истекло",
	ErrorSignUsed:                                            "Подпись ранее была использована",
	ErrorChangeStateSafeDealForbidden:                        "Действие запрещено на текущем этапе сделки",
	ErrorSettingsFSMSafeDeal:                                 "Невалидная конфигурация машины состояний для безопасной сделки",
	ErrorSafeDealExists:                                      "Сделка с переданным идентификатором уже существет",
	ErrorNotValidAddressAcceptor:                             "Неверный адрес акцептанта безопасной сделки",
	ErrorAddressNotCounterpartDeal:                           "Адрес аккаунта не является контрагентом сделки",
	ErrorCounterpartMadeChoice:                               "Контрагент ранее проголосовал за расторжение(заключение) контракта",
	ErrorNotValidAddressAcceptorOrArbitrator:                 "Операция запрещена для аккаунта, который не является акцептантом или арбитром",
	ErrorNotValidAddressInitiator:                            "Операция запрещена для аккаунта, который не является инициатором",
	ErrorCounterpartWasInvited:                               "Переданный контрагент ранее был приглашён",
	ErrorParticipantNotInvited:                               "Участник не был приглашен",
	ErrorNotValidAddressOwner:                                "Операция запрещена для аккаунта, который не является инициатором сделки",
	ErrorEnoughParticipants:                                  "Приглашено достаточно участников с такой ролью",
	ErrorIncorrectSafeDealAmount:                             "Некорректная сумма сделки",
	ErrorSafeDealEscrowNotFound:                              "Эскроу-счет безопасной сделки не найден или был пополнен ранее",
	ErrorIncorrectCurrencyCodeAccount:                        "Некорректный код валюты аккаунта",
	ErrorArbitratorExist:                                     "Арбитр ранее уже был добавлен",
	ErrorArbitratorNotExist:                                  "Арбитра не существует",
	ErrorArbitratorCannotPartialPerform:                      "Арбитр не может частично исполнить сделку",
	ErrorPhysicalAssetNotPresent:                             "В сделке не задействован физический актив",
	ErrorPartialPerformOrder:                                 "Внести данные о частичном исполнении сделки может только первый заявитель об исполнении сделки",
	ErrorOfferAccountSellOnly:                                "Требуется указать адрес аккаунта требуемого актива",
	ErrorContractNotOffer:                                    "Переданный контракт не является предложением",
	ErrorOfferOwnerMissing:                                   "В сделке не указан адрес инициатора предложения",
	ErrorCurrencyNotMatchOffer:                               "Коды валют не соответствует заявленным в предложении",
	ErrorPriceNotMatchOffer:                                  "Курс не соответствует заявленному в предложении",
	ErrorAssetAmountNotCorrespondMaxMinInterval:              "Количество актива не соответствуют интервалу минимального и максимального количества",
	ErrorPhysicalAssetEscrow:                                 "Физический актив не может начисляться на эскроу-счет",
	ErrorAmountNotCorrespondPrice:                            "Количество переданного актива не соответствует переданному курсу",
	ErrorSafeDealNotFound:                                    "Безопасная сделка не найдена",
	ErrorCounterpartVoted:                                    "Участник безопасной сделки проголосовал ранее",
	ErrorIncorrectEscrowAmount:                               "Некорректная сумма эскроу-счета",
	ErrorAccountAlreadyInDeal:                                "Адрес аккаунта уже принимает участие в сделке",
	ErrorAvailablePlatformsNotFound:                          "Информация о доступных платформах не найдена",
	ErrorLimitOperation:                                      "Превышен операционный лимит перевода средств",
	ErrorLimitDaily:                                          "Превышен дневной лимит перевода средств",
	ErrorLimitMonthly:                                        "Превышен месячный лимит перевода средств",
	ErrorLimitBalance:                                        "Превышен лимит баланса",
	ErrorAddressFromToEqual:                                  "Адреса отправителя и получателя совпадают",
	ErrorSingleBankOwnerCannotBeBlocked:                      "Невозможно заблокировать банк, который является единственным владельцем системы",
	ErrorCrossTransactionHistoryNotFound:                     "История по трансграничному переводу не найдена",
	ErrorStatusConflict:                                      "Невозможно выполнить операцию с текущим статусом транзакции",
	ErrorCrossTransactionWithdrawIncorrectAddress:            "Неверный адрес аккаунта для выполнения вывода средств в рамках трансграничного перевода",
	ErrorCrossTransactionWithdrawIncorrectAmount:             "Неверная сумма для выполнения вывода средств в рамках трансграничномого перевода",
	ErrorDefault:                                             "Другая ошибка",
	ErrorInvoiceNotExist:                                     "Счет не найден",
	ErrorInvoiceTransactionAmount:                            "Cуммы cчета не совпадают",
	ErrorInvoiceTransactionPaid:                              "Попытка оплатить ранее оплаченный счет",
	ErrorInvoiceTransactionExpired:                           "Попытка оплатить счета с истекшим сроком жизни",
	ErrorInvoiceExist:                                        "Счёт с указанными номером уже существует",
	ErrorInvoiceUpdate:                                       "Ошибка обновления статуса счета",
	ErrorPayerNotEqual:                                       "Указанный адрес платильщика не совпадает с тем, который указан в счете",
	ErrorPutState:                                            "Ошибка сохрания данных в БД",
	ErrorGetState:                                            "Ошибка при получении данных в БД",
	ErrorDeleteState:                                         "Ошибка при удалении данных в БД",
	ErrorCreateCompositeKey:                                  "Ошибка создании композитного ключа",
	ErrorJsonMarshal:                                         "Ошибка при серилизации JSON",
	ErrorJsonUnmarshal:                                       "Ошибка при десерилизации JSON",
	ErrorGetTxTime:                                           "Ошибка при получении времени создания транзакции",
	ErrorForbidden:                                           "Не достаточно прав для вызова метода смарт-контракта",
	ErrorDepositNotFound:                                     "Эскроу-счёт не найден",
	ErrorContractValidate:                                    "Ошибка валидации контракта",
	ErrorCertificateBankAddressNotMatch:                      "Адрес банка в сертификате не совпадает",
	ErrorClearingBadClaimsLiabilities:                        "Клиринг. Расхождение требований и обязательств",
	ErrorClearingInfoNotFound:                                "Клиринг. Информация о клиринге не найдена",
	ErrorCertificateBankAddressNotFound:                      "Сертификат. Отсутвует атрибут address в сертификате",
	ErrorClientBankExists:                                    "Клиентский банк с указанными параметрами уже существует в системе",
	ErrorClientBankNotFound:                                  "Информация о клиентском банке не найдена",
	ErrorClientBankOwnerNotEqualSender:                       "Опорный банк(отправитель) не является владельцем клиентского банка",
	ErrorLimitTypeIncorrect:                                  "Некорректное значение типа лимита",
	ErrorAccountTechnicalNotEqlSender:                        "Адрес банка отправителя не совпадает с адресом технического аккаунта",
	ErrorAccountNotEscrow:                                    "Аккаунт не является эскроу-счетом",
	ErrorPaymentInvoiceNumberNotEqlForEscrowAccount:          "Переданный номер счета не совпадает с номером счета, который предназначен для пополнения эскроу-счета",
	ErrorAccountNotNotAvailableForCreateArbitrator:           "Аккаунт для создания арбитра недоступен",
	ErrorInviteArbitratorByAddressNotCounterpartDeal:         "Запрещено приглашать арбитра аккаунтом, который не является участником сделки",
	ErrorConfirmInvitationByAddressNotCounterpartDeal:        "Запрещено принимать заявку аккаунтом, который не является участником сделки",
	ErrorCreateSafeDealSameParticipantAddresses:              "Адрес инициатора не может совпадать с адресом акцептанта",
	ErrorEnoughArbitrators:                                   "Не может быть приглашено больше одного арбитра",
	ErrorEnoughAcceptors:                                     "Не может быть приглашено больше одного акцептанта",
	ErrorObligatoryCommissionExceedsFullCommission:           "Сумма обязательной комиссии не может превышать сумму комиссии, которая оплачивается за участие арбитра",
	ErrorCounterpartMadePerformChoice:                        "Контрагент ранее проголосовал за исполнение контракта",
	ErrorCounterpartMadeCancelChoice:                         "Контрагент ранее проголосовал за расторжение контракта",
	ErrorForbiddenForExternal:                                "Метод чейнкода недоступен для использования",
	ErrorEscrowAccountNotBelongSafeDeal:                      "Эскроу-счет не задействован в текущей безопасной сделке",
	ErrorAddressNotParticipantOrCounterpartDeal:              "Адрес аккаунта не является участником сделки",
	ErrorEscrowAccountIncorrectCurrencyCodeForPaymentDeposit: "Эскроу-аккаунт имеет некорректный код валюты для оплаты эскроу-счёта",
	ErrorNotAcceptedAllInvitations:                           "Действие запрещено, до тех пор, когда все приглашения не будут приняты",
	ErrorOwnerSafeDealAccountNotAvailable:                    "Аккаунт инициатора сделки недоступен",
	ErrorInitiatorSafeDealAccountNotAvailable:                "Аккаунт инициатора сделки недоступен",
	ErrorAcceptorSafeDealAccountNotAvailable:                 "Аккаунт акцептанта сделки недоступен",
	ErrorArbitratorSafeDealAccountNotAvailable:               "Аккаунт арбитра сделки недоступен",
	ErrorAmountInitiatorNotCorrespondMaxMinInterval:          "Количество актива инициатора не соответствует минимальному и максимальному значениям",
	ErrorAmountAcceptorNotCorrespondMaxMinInterval:           "Количество актива акцептанта не соответствует минимальному и максимальному значениям",
	ErrorIncorrectCurrencyCodeAccountInitiator:               "Некорректный код валюты аккаунта инициатора",
	ErrorIncorrectCurrencyCodeAccountAcceptor:                "Некорректный код валюты аккаунта акцептора",
	ErrorParticipantSafeDealAccountNotAvailable:              "Аккаунт участника сделки недоступен",
	ErrorInvitedSafeDealAccountNotAvailable:                  "Аккаунт приглашенного участника сделки недоступен",
}

var ErrorStringCodeMap = map[string]int{
	"ErrorAccountExist":                                        ErrorAccountExist,
	"ErrorIdentifierExist":                                     ErrorIdentifierExist,
	"ErrorEditIdentifiedAccountStatusForbidden":                ErrorEditIdentifiedAccountStatusForbidden,
	"ErrorBankWithSuchBikExist":                                ErrorBankWithSuchBikExist,
	"ErrorCurrencyExist":                                       ErrorCurrencyExist,
	"ErrorCurrencyCreateConflict":                              ErrorCurrencyCreateConflict,
	"ErrorReceiverIdentificatorRequired":                       ErrorReceiverIdentificatorRequired,
	"ErrorClientBankWithdrawForbidden":                         ErrorClientBankWithdrawForbidden,
	"ErrorClientBankWithdrawRejectForbidden":                   ErrorClientBankWithdrawRejectForbidden,
	"ErrorClientBankWithdrawConfirmForbidden":                  ErrorClientBankWithdrawConfirmForbidden,
	"ErrorAccountNotExist":                                     ErrorAccountNotExist,
	"ErrorIdentifierNotExist":                                  ErrorIdentifierNotExist,
	"ErrorBankNotExist":                                        ErrorBankNotExist,
	"ErrorCurrencyNotExist":                                    ErrorCurrencyNotExist,
	"ErrorClientBankIsBlocked":                                 ErrorClientBankIsBlocked,
	"ErrorAccountNotIdentified":                                ErrorAccountNotIdentified,
	"ErrorAccountNotAvailable":                                 ErrorAccountNotAvailable,
	"ErrorAccountIsBlocked":                                    ErrorAccountIsBlocked,
	"ErrorBankNotAvailable":                                    ErrorBankNotAvailable,
	"ErrorBankIsBlocked":                                       ErrorBankIsBlocked,
	"ErrorBankExist":                                           ErrorBankExist,
	"ErrorContractNotExists":                                   ErrorContractNotExists,
	"ErrorContractExists":                                      ErrorContractExists,
	"ErrorCustomerNotFound":                                    ErrorCustomerNotFound,
	"ErrorClientBankNotAvailable":                              ErrorClientBankNotAvailable,
	"ErrorEditExchangeContractForbidden":                       ErrorEditExchangeContractForbidden,
	"ErrorClientIdentifierRequired":                            ErrorClientIdentifierRequired,
	"ErrorAccountSenderNotAvailable":                           ErrorAccountSenderNotAvailable,
	"ErrorAccountReceiverNotAvailable":                         ErrorAccountReceiverNotAvailable,
	"ErrorEditUnidentifiedAccountStatusForbidden":              ErrorEditUnidentifiedAccountStatusForbidden,
	"ErrorIncorrectCurrencyCodeSenderAccount":                  ErrorIncorrectCurrencyCodeSenderAccount,
	"ErrorIncorrectCurrencyCodeReceiverAccount":                ErrorIncorrectCurrencyCodeReceiverAccount,
	"ErrorClientBankContractTypeForbidden":                     ErrorClientBankContractTypeForbidden,
	"ErrorAccountContractTypeForbidden":                        ErrorAccountContractTypeForbidden,
	"ErrorClientBankAddressNotMatchAccount":                    ErrorClientBankAddressNotMatchAccount,
	"ErrorUpdateClientBankContractTypesForbidden":              ErrorUpdateClientBankContractTypesForbidden,
	"ErrorSignVerify":                                          ErrorSignVerify,
	"ErrorFundsNotEnough":                                      ErrorFundsNotEnough,
	"ErrorValidateDefault":                                     ErrorValidateDefault,
	"ErrorAddressNotFollowingRegex":                            ErrorAddressNotFollowingRegex,
	"ErrorAddressNotPassed":                                    ErrorAddressNotPassed,
	"ErrorIdentifierNotFolowingRegex":                          ErrorIdentifierNotFolowingRegex,
	"ErrorIdentifierNotPassed":                                 ErrorIdentifierNotPassed,
	"ErrorIdentityTypeIncorrect":                               ErrorIdentityTypeIncorrect,
	"ErrorStateIncorrect":                                      ErrorStateIncorrect,
	"ErrorJuridicalTypeIncorrect":                              ErrorJuridicalTypeIncorrect,
	"ErrorRoleIncorrect":                                       ErrorRoleIncorrect,
	"ErrorChannelNotPassed":                                    ErrorChannelNotPassed,
	"ErrorNameNotPassed":                                       ErrorNameNotPassed,
	"ErrorVersionNotPassed":                                    ErrorVersionNotPassed,
	"ErrorAmountNegative":                                      ErrorAmountNegative,
	"ErrorAmountNotPassed":                                     ErrorAmountNotPassed,
	"ErrorIdentityTypeNotPassed":                               ErrorIdentityTypeNotPassed,
	"ErrorValueNegative":                                       ErrorValueNegative,
	"ErrorValueNotPassed":                                      ErrorValueNotPassed,
	"ErrorCurrencyCodeRange":                                   ErrorCurrencyCodeRange,
	"ErrorCurrencyCodeNotPassed":                               ErrorCurrencyCodeNotPassed,
	"ErrorBalanceNegative":                                     ErrorBalanceNegative,
	"ErrorBalanceNotPassed":                                    ErrorBalanceNotPassed,
	"ErrorMsgHashNotFolowingRegex":                             ErrorMsgHashNotFolowingRegex,
	"ErrorMsgHashNotPassed":                                    ErrorMsgHashNotPassed,
	"ErrorSigRNotFolowingRegex":                                ErrorSigRNotFolowingRegex,
	"ErrorSigRHashNotPassed":                                   ErrorSigRHashNotPassed,
	"ErrorSigSNotFolowingRegex":                                ErrorSigSNotFolowingRegex,
	"ErrorSigSNotPassed":                                       ErrorSigSNotPassed,
	"ErrorSigVIncorrect":                                       ErrorSigVIncorrect,
	"ErrorSigVNotFollowingRegex":                               ErrorSigVNotFollowingRegex,
	"ErrorTxIdSNotPassed":                                      ErrorTxIdSNotPassed,
	"ErrorTimestampNotPassed":                                  ErrorTimestampNotPassed,
	"ErrorTimestampNegative":                                   ErrorTimestampNegative,
	"ErrorTxTypeNotPassed":                                     ErrorTxTypeNotPassed,
	"ErrorTxTypeIncorrect":                                     ErrorTxTypeIncorrect,
	"ErrorPageSizeMinValue":                                    ErrorPageSizeMinValue,
	"ErrorSymbolNotPassed":                                     ErrorSymbolNotPassed,
	"ErrorDecimalsNotPassed":                                   ErrorDecimalsNotPassed,
	"ErrorDecimalsRange":                                       ErrorDecimalsRange,
	"ErrorDecimalsMaxValue":                                    ErrorDecimalsMaxValue,
	"ErrorEnabledNotPassed":                                    ErrorEnabledNotPassed,
	"ErrorIdentifiersNotPassed":                                ErrorIdentifiersNotPassed,
	"ErrorIdentifiersMaxValue":                                 ErrorIdentifiersMaxValue,
	"ErrorStateNotPassed":                                      ErrorStateNotPassed,
	"ErrorJuridicalTypeNotPassed":                              ErrorJuridicalTypeNotPassed,
	"ErrorJuridicalTypeBankSetterNotPassed":                    ErrorJuridicalTypeBankSetterNotPassed,
	"ErrorAccountTypeNotPassed":                                ErrorAccountTypeNotPassed,
	"ErrorAccountTypeIncorrect":                                ErrorAccountTypeIncorrect,
	"ErrorRoleNotPassed":                                       ErrorRoleNotPassed,
	"ErrorIdNotPassed":                                         ErrorIdNotPassed,
	"ErrorIdNotFolowingRegex":                                  ErrorIdNotFolowingRegex,
	"ErrorNumberNotPassed":                                     ErrorNumberNotPassed,
	"ErrorNumberInvoiceNotFolowingRegex":                       ErrorNumberInvoiceNotFolowingRegex,
	"ErrorInvoiceStateIncorrect":                               ErrorInvoiceStateIncorrect,
	"ErrorInvoiceStateNotPassed":                               ErrorInvoiceStateNotPassed,
	"ErrorIssueLimitNegative":                                  ErrorIssueLimitNegative,
	"ErrorStatusTransactionIncorrect":                          ErrorStatusTransactionIncorrect,
	"ErrorStatusTransactionPassed":                             ErrorStatusTransactionPassed,
	"ErrorMspIdNotMatch":                                       ErrorMspIdNotMatch,
	"ErrorAddressNotMatch":                                     ErrorAddressNotMatch,
	"ErrorPublicKeyNotFolowingRegex":                           ErrorPublicKeyNotFolowingRegex,
	"ErrorPublicKeyNotPassed":                                  ErrorPublicKeyNotPassed,
	"ErrorConfNotPassed":                                       ErrorConfNotPassed,
	"ErrorBankAddressNotPassed":                                ErrorBankAddressNotPassed,
	"ErrorCountryCodeNotPassed":                                ErrorCountryCodeNotPassed,
	"ErrorTransactionIdNotPassed":                              ErrorTransactionIdNotPassed,
	"ErrorAddressOrIdentifierNotFolowingRegex":                 ErrorAddressOrIdentifierNotFolowingRegex,
	"ErrorAddressOrIdentifierNotPassed":                        ErrorAddressOrIdentifierNotPassed,
	"ErrorBankDisplayNameNotPassed":                            ErrorBankDisplayNameNotPassed,
	"ErrorGetMspId":                                            ErrorGetMspId,
	"ErrorExpirationSign":                                      ErrorExpirationSign,
	"ErrorSignUsed":                                            ErrorSignUsed,
	"ErrorChangeStateSafeDealForbidden":                        ErrorChangeStateSafeDealForbidden,
	"ErrorSettingsFSMSafeDeal":                                 ErrorSettingsFSMSafeDeal,
	"ErrorSafeDealExists":                                      ErrorSafeDealExists,
	"ErrorNotValidAddressAcceptor":                             ErrorNotValidAddressAcceptor,
	"ErrorAddressNotCounterpartDeal":                           ErrorAddressNotCounterpartDeal,
	"ErrorCounterpartMadeChoice":                               ErrorCounterpartMadeChoice,
	"ErrorNotValidAddressAcceptorOrArbitrator":                 ErrorNotValidAddressAcceptorOrArbitrator,
	"ErrorNotValidAddressInitiator":                            ErrorNotValidAddressInitiator,
	"ErrorCounterpartWasInvited":                               ErrorCounterpartWasInvited,
	"ErrorParticipantNotInvited":                               ErrorParticipantNotInvited,
	"ErrorNotValidAddressOwner":                                ErrorNotValidAddressOwner,
	"ErrorEnoughParticipants":                                  ErrorEnoughParticipants,
	"ErrorIncorrectSafeDealAmount":                             ErrorIncorrectSafeDealAmount,
	"ErrorSafeDealEscrowNotFound":                              ErrorSafeDealEscrowNotFound,
	"ErrorIncorrectCurrencyCodeAccount":                        ErrorIncorrectCurrencyCodeAccount,
	"ErrorArbitratorExist":                                     ErrorArbitratorExist,
	"ErrorArbitratorNotExist":                                  ErrorArbitratorNotExist,
	"ErrorArbitratorCannotPartialPerform":                      ErrorArbitratorCannotPartialPerform,
	"ErrorPhysicalAssetNotPresent":                             ErrorPhysicalAssetNotPresent,
	"ErrorPartialPerformOrder":                                 ErrorPartialPerformOrder,
	"ErrorOfferAccountSellOnly":                                ErrorOfferAccountSellOnly,
	"ErrorContractNotOffer":                                    ErrorContractNotOffer,
	"ErrorOfferOwnerMissing":                                   ErrorOfferOwnerMissing,
	"ErrorCurrencyNotMatchOffer":                               ErrorCurrencyNotMatchOffer,
	"ErrorPriceNotMatchOffer":                                  ErrorPriceNotMatchOffer,
	"ErrorAssetAmountNotCorrespondMaxMinInterval":              ErrorAssetAmountNotCorrespondMaxMinInterval,
	"ErrorPhysicalAssetEscrow":                                 ErrorPhysicalAssetEscrow,
	"ErrorAmountNotCorrespondPrice":                            ErrorAmountNotCorrespondPrice,
	"ErrorSafeDealNotFound":                                    ErrorSafeDealNotFound,
	"ErrorCounterpartVoted":                                    ErrorCounterpartVoted,
	"ErrorIncorrectEscrowAmount":                               ErrorIncorrectEscrowAmount,
	"ErrorAccountAlreadyInDeal":                                ErrorAccountAlreadyInDeal,
	"ErrorAvailablePlatformsNotFound":                          ErrorAvailablePlatformsNotFound,
	"ErrorLimitOperation":                                      ErrorLimitOperation,
	"ErrorLimitDaily":                                          ErrorLimitDaily,
	"ErrorLimitMonthly":                                        ErrorLimitMonthly,
	"ErrorLimitBalance":                                        ErrorLimitBalance,
	"ErrorAddressFromToEqual":                                  ErrorAddressFromToEqual,
	"ErrorSingleBankOwnerCannotBeBlocked":                      ErrorSingleBankOwnerCannotBeBlocked,
	"ErrorCrossTransactionHistoryNotFound":                     ErrorCrossTransactionHistoryNotFound,
	"ErrorStatusConflict":                                      ErrorStatusConflict,
	"ErrorCrossTransactionWithdrawIncorrectAddress":            ErrorCrossTransactionWithdrawIncorrectAddress,
	"ErrorCrossTransactionWithdrawIncorrectAmount":             ErrorCrossTransactionWithdrawIncorrectAmount,
	"ErrorDefault":                                             ErrorDefault,
	"ErrorInvoiceNotExist":                                     ErrorInvoiceNotExist,
	"ErrorInvoiceTransactionAmount":                            ErrorInvoiceTransactionAmount,
	"ErrorInvoiceTransactionPaid":                              ErrorInvoiceTransactionPaid,
	"ErrorInvoiceTransactionExpired":                           ErrorInvoiceTransactionExpired,
	"ErrorInvoiceExist":                                        ErrorInvoiceExist,
	"ErrorInvoiceUpdate":                                       ErrorInvoiceUpdate,
	"ErrorPayerNotEqual":                                       ErrorPayerNotEqual,
	"ErrorPutState":                                            ErrorPutState,
	"ErrorGetState":                                            ErrorGetState,
	"ErrorDeleteState":                                         ErrorDeleteState,
	"ErrorCreateCompositeKey":                                  ErrorCreateCompositeKey,
	"ErrorJsonMarshal":                                         ErrorJsonMarshal,
	"ErrorJsonUnmarshal":                                       ErrorJsonUnmarshal,
	"ErrorGetTxTime":                                           ErrorGetTxTime,
	"ErrorForbidden":                                           ErrorForbidden,
	"ErrorDepositNotFound":                                     ErrorDepositNotFound,
	"ErrorContractValidate":                                    ErrorContractValidate,
	"ErrorCertificateBankAddressNotMatch":                      ErrorCertificateBankAddressNotMatch,
	"ErrorClearingBadClaimsLiabilities":                        ErrorClearingBadClaimsLiabilities,
	"ErrorClearingInfoNotFound":                                ErrorClearingInfoNotFound,
	"ErrorCertificateBankAddressNotFound":                      ErrorCertificateBankAddressNotFound,
	"ErrorClientBankExists":                                    ErrorClientBankExists,
	"ErrorClientBankNotFound":                                  ErrorClientBankNotFound,
	"ErrorClientBankOwnerNotEqualSender":                       ErrorClientBankOwnerNotEqualSender,
	"ErrorLimitTypeIncorrect":                                  ErrorLimitTypeIncorrect,
	"ErrorAccountTechnicalNotEqlSender":                        ErrorAccountTechnicalNotEqlSender,
	"ErrorAccountNotEscrow":                                    ErrorAccountNotEscrow,
	"ErrorPaymentInvoiceNumberNotEqlForEscrowAccount":          ErrorPaymentInvoiceNumberNotEqlForEscrowAccount,
	"ErrorAccountNotNotAvailableForCreateArbitrator":           ErrorAccountNotNotAvailableForCreateArbitrator,
	"ErrorInviteArbitratorByAddressNotCounterpartDeal":         ErrorInviteArbitratorByAddressNotCounterpartDeal,
	"ErrorConfirmInvitationByAddressNotCounterpartDeal":        ErrorConfirmInvitationByAddressNotCounterpartDeal,
	"ErrorCreateSafeDealSameParticipantAddresses":              ErrorCreateSafeDealSameParticipantAddresses,
	"ErrorEnoughArbitrators":                                   ErrorEnoughArbitrators,
	"ErrorEnoughAcceptors":                                     ErrorEnoughAcceptors,
	"ErrorObligatoryCommissionExceedsFullCommission":           ErrorObligatoryCommissionExceedsFullCommission,
	"ErrorCounterpartMadePerformChoice":                        ErrorCounterpartMadePerformChoice,
	"ErrorCounterpartMadeCancelChoice":                         ErrorCounterpartMadeCancelChoice,
	"ErrorForbiddenForExternal":                                ErrorForbiddenForExternal,
	"ErrorEscrowAccountNotBelongSafeDeal":                      ErrorEscrowAccountNotBelongSafeDeal,
	"ErrorAddressNotParticipantOrCounterpartDeal":              ErrorAddressNotParticipantOrCounterpartDeal,
	"ErrorEscrowAccountIncorrectCurrencyCodeForPaymentDeposit": ErrorEscrowAccountIncorrectCurrencyCodeForPaymentDeposit,
	"ErrorNotAcceptedAllInvitations":                           ErrorNotAcceptedAllInvitations,
	"ErrorOwnerSafeDealAccountNotAvailable":                    ErrorOwnerSafeDealAccountNotAvailable,
	"ErrorInitiatorSafeDealAccountNotAvailable":                ErrorInitiatorSafeDealAccountNotAvailable,
	"ErrorAcceptorSafeDealAccountNotAvailable":                 ErrorAcceptorSafeDealAccountNotAvailable,
	"ErrorArbitratorSafeDealAccountNotAvailable":               ErrorArbitratorSafeDealAccountNotAvailable,
	"ErrorAmountInitiatorNotCorrespondMaxMinInterval":          ErrorAmountInitiatorNotCorrespondMaxMinInterval,
	"ErrorAmountAcceptorNotCorrespondMaxMinInterval":           ErrorAmountAcceptorNotCorrespondMaxMinInterval,
	"ErrorIncorrectCurrencyCodeAccountInitiator":               ErrorIncorrectCurrencyCodeAccountInitiator,
	"ErrorIncorrectCurrencyCodeAccountAcceptor":                ErrorIncorrectCurrencyCodeAccountAcceptor,
	"ErrorParticipantSafeDealAccountNotAvailable":              ErrorParticipantSafeDealAccountNotAvailable,
	"ErrorInvitedSafeDealAccountNotAvailable":                  ErrorInvitedSafeDealAccountNotAvailable,
}
