package query

var (
	Register = `
		INSERT INTO 
		nicepay(
			iMid,
			payMethod,
			currency,
			amt,
			referenceNo,
			goodsNm,
			billingNm,
			billingPhone,
			billingEmail,
			billingAddr,
			billingCity,
			billingState,
			billingPostCd,
			billingCountry,
			deliveryNm,
			deliveryPhone,
			deliveryAddr,
			deliveryCity,
			deliveryState,
			deliveryPostCd,
			deliveryCountry,
			dbProcessUrl,
			merchantToken,
			reqDt,
			reqTm,
			reqDomain,
			reqServerIP,
			reqClientVer,
			userIP,
			userSessionID,
			userAgent,
			cartData,
			instmntType,
			instmntMon,
			recurrOpt
		)VALUES(
			:iMid,
			:payMethod,
			:currency,
			:amt,
			:referenceNo,
			:goodsNm,
			:billingNm,
			:billingPhone,
			:billingEmail,
			:billingAddr,
			:billingCity,
			:billingState,
			:billingPostCd,
			:billingCountry,
			:deliveryNm,
			:deliveryPhone,
			:deliveryAddr,
			:deliveryCity,
			:deliveryState,
			:deliveryPostCd,
			:deliveryCountry,
			:dbProcessUrl,
			:merchantToken,
			:reqDt,
			:reqTm,
			:reqDomain,
			:reqServerIP,
			:reqClientVer,
			:userIP,
			:userSessionID,
			:userAgent,
			:cartData,
			:instmntType,
			:instmntMon,
			:recurrOpt
		);
	`
	FindStatus = `
    SELECT
    %v
    FROM nicepay
    WHERE %v LIMIT 1
  `
)
