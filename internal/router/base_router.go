package router

func (r router) BaseRouter() {
	nicepayHandler := r.handler.NicepayHandler()

	nicepayPrefix := r.route.Group("/nicepay")
	{
		directPrefix := nicepayPrefix.Group("/direct")
		{

			v2Prefix := directPrefix.Group("/v2")
			{
				v2Prefix.POST("/registration", nicepayHandler.NicepayRegister)
				v2Prefix.POST("/inquiry", nicepayHandler.NicepayStatus)
				v2Prefix.POST("/payment", nicepayHandler.NicepayPayment)
			}
		}
	}
}
