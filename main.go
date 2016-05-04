package boxpacker

/*
	var t * testing.T

	 if ok := testPackBoxThreeItemsFitExactly(ctx, t); !ok{
		 log.Printf("fail")
	 }
	if ok := testPackBoxThreeItemsFitEasily(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackBoxThreeItemsFitExactly(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackBoxThreeItemsFitExactlyNoRotation(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackBoxThreeItemsFitSizeButOverweight(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackBoxThreeItemsFitWeightBut2Oversize(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackTwoItemsFitExactlySideBySide(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackThreeItemsBottom2FitSideBySideOneExactlyOnTop(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackThreeItemsBottom2FitSideBySideWithSpareSpaceOneOverhangSlightlyOnTop(ctx, t); !ok{
		log.Printf("fail")
	}
	if ok := testPackSingleItemFitsBetterRotated(ctx, t); !ok{
		log.Printf("fail")
	}
*/


import ("log")
func Uline() {
	log.Printf("init\n")
	///	log.Printf("In pack");

	NewPacker()

	AddBox(&Box{Reference:"3 x 3 x 3 #200 ULINE", InnerLength: 3, InnerWidth: 3, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"4 x 4 x 4 #200 ULINE", InnerLength: 4, InnerWidth: 4, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"5 x 5 x 5 #200 ULINE", InnerLength: 5, InnerWidth: 5, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"6 x 6 x 6 #200 ULINE", InnerLength: 6, InnerWidth: 6, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"7 x 7 x 7 #200 ULINE", InnerLength: 7, InnerWidth: 7, InnerDepth: 7, MaxWeight: 40})
	AddBox(&Box{Reference:"8 x 8 x 8 #200 ULINE", InnerLength: 8, InnerWidth: 8, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"9 x 9 x 9 #200 ULINE", InnerLength: 9, InnerWidth: 9, InnerDepth: 9, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 10 x 10 #200 ULINE", InnerLength: 10, InnerWidth: 10, InnerDepth: 10, MaxWeight: 40})
	AddBox(&Box{Reference:"11 x 11 x 11 #200 ULINE", InnerLength: 11, InnerWidth: 11, InnerDepth: 11, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 12 x 12 #200 ULINE", InnerLength: 12, InnerWidth: 12, InnerDepth: 12, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 13 x 13 #200 ULINE", InnerLength: 13, InnerWidth: 13, InnerDepth: 13, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 14 x 14 #200 ULINE", InnerLength: 14, InnerWidth: 14, InnerDepth: 14, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 15 x 15 #200 ULINE", InnerLength: 15, InnerWidth: 15, InnerDepth: 15, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 16 x 16 #200 ULINE", InnerLength: 16, InnerWidth: 16, InnerDepth: 16, MaxWeight: 40})
	AddBox(&Box{Reference:"17 x 17 x 17 #200 ULINE", InnerLength: 17, InnerWidth: 17, InnerDepth: 17, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 18 x 18 #200 ULINE", InnerLength: 18, InnerWidth: 18, InnerDepth: 18, MaxWeight: 40})
	AddBox(&Box{Reference:"19 x 19 x 19 #200 ULINE", InnerLength: 19, InnerWidth: 19, InnerDepth: 19, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 20 x 20 #200 ULINE", InnerLength: 20, InnerWidth: 20, InnerDepth: 20, MaxWeight: 40})
	AddBox(&Box{Reference:"21 x 21 x 21 #200 ULINE", InnerLength: 21, InnerWidth: 21, InnerDepth: 21, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 22 x 22 #200 ULINE", InnerLength: 22, InnerWidth: 22, InnerDepth: 22, MaxWeight: 40})
	AddBox(&Box{Reference:"23 x 23 x 23 #200 ULINE", InnerLength: 23, InnerWidth: 23, InnerDepth: 23, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 24 x 24 #200 ULINE", InnerLength: 24, InnerWidth: 24, InnerDepth: 24, MaxWeight: 40})
	AddBox(&Box{Reference:"25 x 25 x 25 #200 ULINE", InnerLength: 25, InnerWidth: 25, InnerDepth: 25, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 26 x 26 #200 ULINE", InnerLength: 26, InnerWidth: 26, InnerDepth: 26, MaxWeight: 40})
	AddBox(&Box{Reference:"27 x 27 x 27 #200 ULINE", InnerLength: 27, InnerWidth: 27, InnerDepth: 27, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 28 x 28 #200 ULINE", InnerLength: 28, InnerWidth: 28, InnerDepth: 28, MaxWeight: 40})
	AddBox(&Box{Reference:"30 x 30 x 30 #200 ULINE", InnerLength: 30, InnerWidth: 30, InnerDepth: 30, MaxWeight: 40})
	AddBox(&Box{Reference:"32 x 32 x 32 #200 ULINE", InnerLength: 32, InnerWidth: 32, InnerDepth: 32, MaxWeight: 40})
	AddBox(&Box{Reference:"34 x 34 x 34 #200 ULINE", InnerLength: 34, InnerWidth: 34, InnerDepth: 34, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 36 x 36 #200 ULINE", InnerLength: 36, InnerWidth: 36, InnerDepth: 36, MaxWeight: 40})
	AddBox(&Box{Reference:"40 x 40 x 40 #200 ULINE", InnerLength: 40, InnerWidth: 40, InnerDepth: 40, MaxWeight: 40})
	AddBox(&Box{Reference:"48 x 48 x 48 #200 ULINE", InnerLength: 48, InnerWidth: 48, InnerDepth: 48, MaxWeight: 40})
	AddBox(&Box{Reference:"6 x 6 x 2 #200 ULINE", InnerLength: 6, InnerWidth: 6, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"7 x 7 x 3 #200 ULINE", InnerLength: 7, InnerWidth: 7, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"8 x 6 x 2 #200 ULINE", InnerLength: 8, InnerWidth: 6, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"8 x 6 x 3 #200 ULINE", InnerLength: 8, InnerWidth: 6, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"8 x 8 x 2 #200 ULINE", InnerLength: 8, InnerWidth: 8, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"8 x 8 x 3 #200 ULINE", InnerLength: 8, InnerWidth: 8, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"8 x 8 x 4 #200 ULINE", InnerLength: 8, InnerWidth: 8, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"9 x 6 x 2 #200 ULINE", InnerLength: 9, InnerWidth: 6, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"9 x 6 x 3 #200 ULINE", InnerLength: 9, InnerWidth: 6, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"9 x 7 x 3 #200 ULINE", InnerLength: 9, InnerWidth: 7, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"9 x 9 x 3 #200 ULINE", InnerLength: 9, InnerWidth: 9, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"9 x 9 x 4 #200 ULINE", InnerLength: 9, InnerWidth: 9, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 6 x 3 #200 ULINE", InnerLength: 10, InnerWidth: 6, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 7 x 3 #200 ULINE", InnerLength: 10, InnerWidth: 7, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 8 x 2 #200 ULINE", InnerLength: 10, InnerWidth: 8, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 8 x 3 #200 ULINE", InnerLength: 10, InnerWidth: 8, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 8 x 4 #200 ULINE", InnerLength: 10, InnerWidth: 8, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 9 x 4 #200 ULINE", InnerLength: 10, InnerWidth: 9, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 10 x 2 #200 ULINE", InnerLength: 10, InnerWidth: 10, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 10 x 3 #200 ULINE", InnerLength: 10, InnerWidth: 10, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 10 x 4 #200 ULINE", InnerLength: 10, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"10 x 10 x 5 #200 ULINE", InnerLength: 10, InnerWidth: 10, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"11 x 8 x 4 #200 ULINE", InnerLength: 11, InnerWidth: 8, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"11 x 11 x 3 #200 ULINE", InnerLength: 11, InnerWidth: 11, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"11 x 11 x 4 #200 ULINE", InnerLength: 11, InnerWidth: 11, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"11 x 11 x 5 #200 ULINE", InnerLength: 11, InnerWidth: 11, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 8 x 3 #200 ULINE", InnerLength: 12, InnerWidth: 8, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 8 x 4 #200 ULINE", InnerLength: 12, InnerWidth: 8, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 9 x 2 #200 ULINE", InnerLength: 12, InnerWidth: 9, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 9 x 3 #200 ULINE", InnerLength: 12, InnerWidth: 9, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 9 x 4 #200 ULINE", InnerLength: 12, InnerWidth: 9, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 10 x 2 #200 ULINE", InnerLength: 12, InnerWidth: 10, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 10 x 3 #200 ULINE", InnerLength: 12, InnerWidth: 10, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 10 x 4 #200 ULINE", InnerLength: 12, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 10 x 5 #200 ULINE", InnerLength: 12, InnerWidth: 10, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 12 x 2 #200 ULINE", InnerLength: 12, InnerWidth: 12, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 12 x 3 #200 ULINE", InnerLength: 12, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 12 x 4 #200 ULINE", InnerLength: 12, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 12 x 5 #200 ULINE", InnerLength: 12, InnerWidth: 12, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"12 x 12 x 6 #200 ULINE", InnerLength: 12, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 9 x 4 #200 ULINE", InnerLength: 13, InnerWidth: 9, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 10 x 2 #200 ULINE", InnerLength: 13, InnerWidth: 10, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 10 x 4 #200 ULINE", InnerLength: 13, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 10 x 5 #200 ULINE", InnerLength: 13, InnerWidth: 10, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 11 x 5 #200 ULINE", InnerLength: 13, InnerWidth: 11, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 13 x 2 #200 ULINE", InnerLength: 13, InnerWidth: 13, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 13 x 3 #200 ULINE", InnerLength: 13, InnerWidth: 13, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 13 x 4 #200 ULINE", InnerLength: 13, InnerWidth: 13, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 13 x 5 #200 ULINE", InnerLength: 13, InnerWidth: 13, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"13 x 13 x 6 #200 ULINE", InnerLength: 13, InnerWidth: 13, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 6 x 2 #200 ULINE", InnerLength: 14, InnerWidth: 6, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 8 x 4 #200 ULINE", InnerLength: 14, InnerWidth: 8, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 10 x 2 #200 ULINE", InnerLength: 14, InnerWidth: 10, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 10 x 3 #200 ULINE", InnerLength: 14, InnerWidth: 10, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 10 x 4 #200 ULINE", InnerLength: 14, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 10 x 5 #200 ULINE", InnerLength: 14, InnerWidth: 10, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 11 x 3 #200 ULINE", InnerLength: 14, InnerWidth: 11, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 12 x 3 #200 ULINE", InnerLength: 14, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 12 x 4 #200 ULINE", InnerLength: 14, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 14 x 2 #200 ULINE", InnerLength: 14, InnerWidth: 14, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 14 x 3 #200 ULINE", InnerLength: 14, InnerWidth: 14, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 14 x 4 #200 ULINE", InnerLength: 14, InnerWidth: 14, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 14 x 5 #200 ULINE", InnerLength: 14, InnerWidth: 14, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"14 x 14 x 6 #200 ULINE", InnerLength: 14, InnerWidth: 14, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 10 x 4 #200 ULINE", InnerLength: 15, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 10 x 5 #200 ULINE", InnerLength: 15, InnerWidth: 10, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 12 x 3 #200 ULINE", InnerLength: 15, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 12 x 4 #200 ULINE", InnerLength: 15, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 12 x 5 #200 ULINE", InnerLength: 15, InnerWidth: 12, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 15 x 3 #200 ULINE", InnerLength: 15, InnerWidth: 15, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 15 x 4 #200 ULINE", InnerLength: 15, InnerWidth: 15, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 15 x 5 #200 ULINE", InnerLength: 15, InnerWidth: 15, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"15 x 15 x 6 #200 ULINE", InnerLength: 15, InnerWidth: 15, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 9 x 3 #200 ULINE", InnerLength: 16, InnerWidth: 9, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 10 x 4 #200 ULINE", InnerLength: 16, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 12 x 3 #200 ULINE", InnerLength: 16, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 12 x 4 #200 ULINE", InnerLength: 16, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 12 x 6 #200 ULINE", InnerLength: 16, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 14 x 4 #200 ULINE", InnerLength: 16, InnerWidth: 14, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 14 x 6 #200 ULINE", InnerLength: 16, InnerWidth: 14, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 16 x 3 #200 ULINE", InnerLength: 16, InnerWidth: 16, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 16 x 4 #200 ULINE", InnerLength: 16, InnerWidth: 16, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 16 x 5 #200 ULINE", InnerLength: 16, InnerWidth: 16, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"16 x 16 x 6 #200 ULINE", InnerLength: 16, InnerWidth: 16, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"17 x 13 x 5 #200 ULINE", InnerLength: 17, InnerWidth: 13, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"17 x 14 x 5 #200 ULINE", InnerLength: 17, InnerWidth: 14, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"17 x 17 x 4 #200 ULINE", InnerLength: 17, InnerWidth: 17, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"17 x 17 x 6 #200 ULINE", InnerLength: 17, InnerWidth: 17, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 10 x 4 #200 ULINE", InnerLength: 18, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 12 x 2 #200 ULINE", InnerLength: 18, InnerWidth: 12, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 12 x 3 #200 ULINE", InnerLength: 18, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 12 x 4 #200 ULINE", InnerLength: 18, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 12 x 5 #200 ULINE", InnerLength: 18, InnerWidth: 12, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 12 x 6 #200 ULINE", InnerLength: 18, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 13 x 5 #200 ULINE", InnerLength: 18, InnerWidth: 13, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 14 x 4 #200 ULINE", InnerLength: 18, InnerWidth: 14, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 14 x 6 #200 ULINE", InnerLength: 18, InnerWidth: 14, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 16 x 2 #200 ULINE", InnerLength: 18, InnerWidth: 16, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 16 x 4 #200 ULINE", InnerLength: 18, InnerWidth: 16, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 16 x 6 #200 ULINE", InnerLength: 18, InnerWidth: 16, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 18 x 2 #200 ULINE", InnerLength: 18, InnerWidth: 18, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 18 x 4 #200 ULINE", InnerLength: 18, InnerWidth: 18, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 18 x 5 #200 ULINE", InnerLength: 18, InnerWidth: 18, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"18 x 18 x 6 #200 ULINE", InnerLength: 18, InnerWidth: 18, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"19 x 12 x 3 #200 ULINE", InnerLength: 19, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"19 x 12 x 4 #200 ULINE", InnerLength: 19, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"19 x 13 x 6 #200 ULINE", InnerLength: 19, InnerWidth: 13, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 10 x 4 #200 ULINE", InnerLength: 20, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 10 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 10, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 12 x 3 #200 ULINE", InnerLength: 20, InnerWidth: 12, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 12 x 4 #200 ULINE", InnerLength: 20, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 12 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 14 x 3 #200 ULINE", InnerLength: 20, InnerWidth: 14, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 14 x 4 #200 ULINE", InnerLength: 20, InnerWidth: 14, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 14 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 14, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 15 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 15, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 16 x 4 #200 ULINE", InnerLength: 20, InnerWidth: 16, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 16 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 16, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 18 x 4 #200 ULINE", InnerLength: 20, InnerWidth: 18, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 18 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 18, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 20 x 2 #200 ULINE", InnerLength: 20, InnerWidth: 20, InnerDepth: 2, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 20 x 4 #200 ULINE", InnerLength: 20, InnerWidth: 20, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 20 x 6 #200 ULINE", InnerLength: 20, InnerWidth: 20, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 20 x 7 #200 ULINE", InnerLength: 20, InnerWidth: 20, InnerDepth: 7, MaxWeight: 40})
	AddBox(&Box{Reference:"20 x 20 x 8 #200 ULINE", InnerLength: 20, InnerWidth: 20, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 10 x 4 #200 ULINE", InnerLength: 22, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 10 x 6 #200 ULINE", InnerLength: 22, InnerWidth: 10, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 12 x 6 #200 ULINE", InnerLength: 22, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 14 x 4 #200 ULINE", InnerLength: 22, InnerWidth: 14, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 14 x 6 #200 ULINE", InnerLength: 22, InnerWidth: 14, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 16 x 4 #200 ULINE", InnerLength: 22, InnerWidth: 16, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 16 x 6 #200 ULINE", InnerLength: 22, InnerWidth: 16, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 18 x 4 #200 ULINE", InnerLength: 22, InnerWidth: 18, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 18 x 6 #200 ULINE", InnerLength: 22, InnerWidth: 18, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 22 x 4 #200 ULINE", InnerLength: 22, InnerWidth: 22, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 22 x 6 #200 ULINE", InnerLength: 22, InnerWidth: 22, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"22 x 22 x 8 #200 ULINE", InnerLength: 22, InnerWidth: 22, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 10 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 10 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 10, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 12 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 12 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 14 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 14, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 14 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 14, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 16 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 16, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 16 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 16, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 17 x 3 #200 ULINE", InnerLength: 24, InnerWidth: 17, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 18 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 18, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 18 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 18, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 20 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 20, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 20 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 20, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 20 x 8 #200 ULINE", InnerLength: 24, InnerWidth: 20, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 24 x 4 #200 ULINE", InnerLength: 24, InnerWidth: 24, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 24 x 6 #200 ULINE", InnerLength: 24, InnerWidth: 24, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 24 x 7 #200 ULINE", InnerLength: 24, InnerWidth: 24, InnerDepth: 7, MaxWeight: 40})
	AddBox(&Box{Reference:"24 x 24 x 8 #200 ULINE", InnerLength: 24, InnerWidth: 24, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 10 x 4 #200 ULINE", InnerLength: 26, InnerWidth: 10, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 15 x 5 #200 ULINE", InnerLength: 26, InnerWidth: 15, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 15 x 7 #200 ULINE", InnerLength: 26, InnerWidth: 15, InnerDepth: 7, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 17 x 5 #200 ULINE", InnerLength: 26, InnerWidth: 17, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 20 x 4 #200 ULINE", InnerLength: 26, InnerWidth: 20, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 20 x 6 #200 ULINE", InnerLength: 26, InnerWidth: 20, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 26 x 4 #200 ULINE", InnerLength: 26, InnerWidth: 26, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 26 x 6 #200 ULINE", InnerLength: 26, InnerWidth: 26, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"26 x 26 x 8 #200 ULINE", InnerLength: 26, InnerWidth: 26, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 12 x 6 #200 ULINE", InnerLength: 28, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 16 x 5 #200 ULINE", InnerLength: 28, InnerWidth: 16, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 16 x 7 #200 ULINE", InnerLength: 28, InnerWidth: 16, InnerDepth: 7, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 17 x 5 #200 ULINE", InnerLength: 28, InnerWidth: 17, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 18 x 6 #200 ULINE", InnerLength: 28, InnerWidth: 18, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 18 x 8 #200 ULINE", InnerLength: 28, InnerWidth: 18, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 24 x 6 #200 ULINE", InnerLength: 28, InnerWidth: 24, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 28 x 6 #200 ULINE", InnerLength: 28, InnerWidth: 28, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"28 x 28 x 8 #200 ULINE", InnerLength: 28, InnerWidth: 28, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"29 x 17 x 3 #200 ULINE", InnerLength: 29, InnerWidth: 17, InnerDepth: 3, MaxWeight: 40})
	AddBox(&Box{Reference:"30 x 20 x 6 #200 ULINE", InnerLength: 30, InnerWidth: 20, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"30 x 20 x 8 #200 ULINE", InnerLength: 30, InnerWidth: 20, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"30 x 30 x 6 #200 ULINE", InnerLength: 30, InnerWidth: 30, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"30 x 30 x 8 #200 ULINE", InnerLength: 30, InnerWidth: 30, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"32 x 32 x 8 #200 ULINE", InnerLength: 32, InnerWidth: 32, InnerDepth: 8, MaxWeight: 40})
	AddBox(&Box{Reference:"34 x 21 x 6 #200 ULINE", InnerLength: 34, InnerWidth: 21, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 12 x 4 #200 ULINE", InnerLength: 36, InnerWidth: 12, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 12 x 6 #200 ULINE", InnerLength: 36, InnerWidth: 12, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 16 x 5 #200 ULINE", InnerLength: 36, InnerWidth: 16, InnerDepth: 5, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 18 x 6 #200 ULINE", InnerLength: 36, InnerWidth: 18, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 24 x 4 #200 ULINE", InnerLength: 36, InnerWidth: 24, InnerDepth: 4, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 24 x 6 #200 ULINE", InnerLength: 36, InnerWidth: 24, InnerDepth: 6, MaxWeight: 40})
	AddBox(&Box{Reference:"36 x 24 x 8 #200 ULINE", InnerLength: 36, InnerWidth: 24, InnerDepth: 8, MaxWeight: 40})

	// Some items and their priorities.
	// Length      int Depth       int Volume      int Weight      int
	//var i = &Item{Description:"Narwhal", Width: 20, Length: 30, Depth: 10, Weight: 6}
	//var k = &Item{Description:"Power Tower Add-on", Width: 10, Length: 3, Depth: 6, Weight: 15}
	var i = &Item{Description:"Baseball", Width: 3, Length: 3, Depth: 3, Weight: 1}
	//AddItem(ctx, i, 1)
	//AddItem(ctx, k, 1)
	AddItem(i, 2)

	packedBoxes := Pack()

	log.Printf("These items fitted into %d box(es).\n", packedBoxes.Len())
	for _, pb := range *packedBoxes {
		packedBox := pb.(PackedBox)
		boxType := packedBox.Box
		log.Printf("This box is a %s, it is %d in wide, %d in long and %d in high.\n", boxType.Reference, boxType.OuterWidth, boxType.OuterLength, boxType.OuterDepth)
		log.Printf("The combined weight of this box and the items inside it is %d lbs.", packedBox.BoxWeight());
		log.Printf("The items in this box are:\n");
		itemsInTheBox := packedBox.Items

		for _, i := range itemsInTheBox {
			item := i.(Item)
			log.Printf("%s", item.Description)
		}
	}
}




























































































































































































































