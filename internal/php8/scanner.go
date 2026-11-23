// line internal/php8/scanner.rl:1
package php8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rectorphp/php-parser-in-go/pkg/token"
)

// line internal/php8/scanner.go:13
const lexer_start int = 138
const lexer_first_final int = 138
const lexer_error int = 0

const lexer_en_main int = 138
const lexer_en_html int = 142
const lexer_en_php int = 151
const lexer_en_property int = 519
const lexer_en_nowdoc int = 526
const lexer_en_heredoc int = 530
const lexer_en_backqote int = 537
const lexer_en_template_string int = 543
const lexer_en_heredoc_end int = 549
const lexer_en_string_var int = 551
const lexer_en_string_var_index int = 557
const lexer_en_string_var_name int = 567
const lexer_en_halt_compiller_open_parenthesis int = 569
const lexer_en_halt_compiller_close_parenthesis int = 578
const lexer_en_halt_compiller_close_semicolon int = 587
const lexer_en_halt_compiller_end int = 596

// line internal/php8/scanner.rl:17

func initLexer(lex *Lexer) {

	// line internal/php8/scanner.go:39
	{
		lex.cs = lexer_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

	// line internal/php8/scanner.rl:21
}

func (lex *Lexer) Lex() *token.Token {
	eof := lex.pe
	var tok token.ID

	tkn := lex.tokenPool.Get()

	lblStart := 0
	lblEnd := 0

	_, _ = lblStart, lblEnd

	// line internal/php8/scanner.go:61
	{
		var _widec int16
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 138:
			goto st138
		case 139:
			goto st139
		case 1:
			goto st1
		case 140:
			goto st140
		case 141:
			goto st141
		case 142:
			goto st142
		case 143:
			goto st143
		case 144:
			goto st144
		case 145:
			goto st145
		case 146:
			goto st146
		case 147:
			goto st147
		case 148:
			goto st148
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
		case 149:
			goto st149
		case 150:
			goto st150
		case 151:
			goto st151
		case 152:
			goto st152
		case 153:
			goto st153
		case 154:
			goto st154
		case 155:
			goto st155
		case 156:
			goto st156
		case 157:
			goto st157
		case 158:
			goto st158
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 8:
			goto st8
		case 159:
			goto st159
		case 160:
			goto st160
		case 161:
			goto st161
		case 162:
			goto st162
		case 163:
			goto st163
		case 164:
			goto st164
		case 165:
			goto st165
		case 166:
			goto st166
		case 167:
			goto st167
		case 9:
			goto st9
		case 10:
			goto st10
		case 168:
			goto st168
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 41:
			goto st41
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 169:
			goto st169
		case 170:
			goto st170
		case 171:
			goto st171
		case 172:
			goto st172
		case 173:
			goto st173
		case 66:
			goto st66
		case 174:
			goto st174
		case 67:
			goto st67
		case 68:
			goto st68
		case 175:
			goto st175
		case 69:
			goto st69
		case 176:
			goto st176
		case 70:
			goto st70
		case 71:
			goto st71
		case 72:
			goto st72
		case 73:
			goto st73
		case 177:
			goto st177
		case 178:
			goto st178
		case 179:
			goto st179
		case 74:
			goto st74
		case 75:
			goto st75
		case 180:
			goto st180
		case 76:
			goto st76
		case 181:
			goto st181
		case 77:
			goto st77
		case 182:
			goto st182
		case 183:
			goto st183
		case 184:
			goto st184
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 185:
			goto st185
		case 186:
			goto st186
		case 187:
			goto st187
		case 188:
			goto st188
		case 189:
			goto st189
		case 82:
			goto st82
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 190:
			goto st190
		case 191:
			goto st191
		case 86:
			goto st86
		case 87:
			goto st87
		case 88:
			goto st88
		case 192:
			goto st192
		case 193:
			goto st193
		case 194:
			goto st194
		case 195:
			goto st195
		case 196:
			goto st196
		case 197:
			goto st197
		case 89:
			goto st89
		case 198:
			goto st198
		case 199:
			goto st199
		case 200:
			goto st200
		case 201:
			goto st201
		case 202:
			goto st202
		case 203:
			goto st203
		case 90:
			goto st90
		case 204:
			goto st204
		case 205:
			goto st205
		case 206:
			goto st206
		case 207:
			goto st207
		case 208:
			goto st208
		case 209:
			goto st209
		case 210:
			goto st210
		case 211:
			goto st211
		case 212:
			goto st212
		case 213:
			goto st213
		case 214:
			goto st214
		case 215:
			goto st215
		case 91:
			goto st91
		case 92:
			goto st92
		case 216:
			goto st216
		case 217:
			goto st217
		case 218:
			goto st218
		case 219:
			goto st219
		case 220:
			goto st220
		case 221:
			goto st221
		case 222:
			goto st222
		case 223:
			goto st223
		case 224:
			goto st224
		case 225:
			goto st225
		case 226:
			goto st226
		case 227:
			goto st227
		case 228:
			goto st228
		case 229:
			goto st229
		case 230:
			goto st230
		case 231:
			goto st231
		case 232:
			goto st232
		case 233:
			goto st233
		case 234:
			goto st234
		case 235:
			goto st235
		case 236:
			goto st236
		case 237:
			goto st237
		case 238:
			goto st238
		case 239:
			goto st239
		case 240:
			goto st240
		case 241:
			goto st241
		case 242:
			goto st242
		case 243:
			goto st243
		case 244:
			goto st244
		case 245:
			goto st245
		case 246:
			goto st246
		case 247:
			goto st247
		case 248:
			goto st248
		case 249:
			goto st249
		case 250:
			goto st250
		case 251:
			goto st251
		case 252:
			goto st252
		case 253:
			goto st253
		case 254:
			goto st254
		case 255:
			goto st255
		case 256:
			goto st256
		case 257:
			goto st257
		case 258:
			goto st258
		case 259:
			goto st259
		case 260:
			goto st260
		case 261:
			goto st261
		case 262:
			goto st262
		case 263:
			goto st263
		case 264:
			goto st264
		case 265:
			goto st265
		case 266:
			goto st266
		case 267:
			goto st267
		case 268:
			goto st268
		case 269:
			goto st269
		case 270:
			goto st270
		case 271:
			goto st271
		case 272:
			goto st272
		case 273:
			goto st273
		case 274:
			goto st274
		case 275:
			goto st275
		case 276:
			goto st276
		case 277:
			goto st277
		case 278:
			goto st278
		case 279:
			goto st279
		case 280:
			goto st280
		case 281:
			goto st281
		case 282:
			goto st282
		case 283:
			goto st283
		case 284:
			goto st284
		case 285:
			goto st285
		case 286:
			goto st286
		case 287:
			goto st287
		case 288:
			goto st288
		case 289:
			goto st289
		case 290:
			goto st290
		case 291:
			goto st291
		case 292:
			goto st292
		case 293:
			goto st293
		case 294:
			goto st294
		case 295:
			goto st295
		case 296:
			goto st296
		case 297:
			goto st297
		case 298:
			goto st298
		case 299:
			goto st299
		case 300:
			goto st300
		case 301:
			goto st301
		case 302:
			goto st302
		case 303:
			goto st303
		case 304:
			goto st304
		case 305:
			goto st305
		case 306:
			goto st306
		case 307:
			goto st307
		case 308:
			goto st308
		case 309:
			goto st309
		case 310:
			goto st310
		case 311:
			goto st311
		case 312:
			goto st312
		case 313:
			goto st313
		case 314:
			goto st314
		case 315:
			goto st315
		case 316:
			goto st316
		case 317:
			goto st317
		case 318:
			goto st318
		case 319:
			goto st319
		case 320:
			goto st320
		case 321:
			goto st321
		case 322:
			goto st322
		case 323:
			goto st323
		case 324:
			goto st324
		case 325:
			goto st325
		case 326:
			goto st326
		case 327:
			goto st327
		case 328:
			goto st328
		case 329:
			goto st329
		case 330:
			goto st330
		case 331:
			goto st331
		case 332:
			goto st332
		case 333:
			goto st333
		case 334:
			goto st334
		case 335:
			goto st335
		case 336:
			goto st336
		case 337:
			goto st337
		case 338:
			goto st338
		case 339:
			goto st339
		case 340:
			goto st340
		case 341:
			goto st341
		case 342:
			goto st342
		case 343:
			goto st343
		case 344:
			goto st344
		case 345:
			goto st345
		case 346:
			goto st346
		case 347:
			goto st347
		case 348:
			goto st348
		case 349:
			goto st349
		case 350:
			goto st350
		case 351:
			goto st351
		case 352:
			goto st352
		case 353:
			goto st353
		case 354:
			goto st354
		case 355:
			goto st355
		case 356:
			goto st356
		case 357:
			goto st357
		case 358:
			goto st358
		case 359:
			goto st359
		case 360:
			goto st360
		case 361:
			goto st361
		case 362:
			goto st362
		case 363:
			goto st363
		case 364:
			goto st364
		case 365:
			goto st365
		case 366:
			goto st366
		case 367:
			goto st367
		case 368:
			goto st368
		case 369:
			goto st369
		case 370:
			goto st370
		case 371:
			goto st371
		case 372:
			goto st372
		case 373:
			goto st373
		case 93:
			goto st93
		case 374:
			goto st374
		case 375:
			goto st375
		case 376:
			goto st376
		case 377:
			goto st377
		case 378:
			goto st378
		case 379:
			goto st379
		case 380:
			goto st380
		case 381:
			goto st381
		case 382:
			goto st382
		case 383:
			goto st383
		case 384:
			goto st384
		case 94:
			goto st94
		case 95:
			goto st95
		case 96:
			goto st96
		case 97:
			goto st97
		case 385:
			goto st385
		case 386:
			goto st386
		case 387:
			goto st387
		case 388:
			goto st388
		case 389:
			goto st389
		case 390:
			goto st390
		case 391:
			goto st391
		case 98:
			goto st98
		case 99:
			goto st99
		case 100:
			goto st100
		case 101:
			goto st101
		case 392:
			goto st392
		case 393:
			goto st393
		case 394:
			goto st394
		case 395:
			goto st395
		case 396:
			goto st396
		case 102:
			goto st102
		case 103:
			goto st103
		case 104:
			goto st104
		case 105:
			goto st105
		case 397:
			goto st397
		case 398:
			goto st398
		case 399:
			goto st399
		case 400:
			goto st400
		case 401:
			goto st401
		case 402:
			goto st402
		case 403:
			goto st403
		case 404:
			goto st404
		case 405:
			goto st405
		case 406:
			goto st406
		case 407:
			goto st407
		case 408:
			goto st408
		case 409:
			goto st409
		case 410:
			goto st410
		case 411:
			goto st411
		case 412:
			goto st412
		case 413:
			goto st413
		case 414:
			goto st414
		case 415:
			goto st415
		case 416:
			goto st416
		case 417:
			goto st417
		case 418:
			goto st418
		case 419:
			goto st419
		case 420:
			goto st420
		case 421:
			goto st421
		case 422:
			goto st422
		case 423:
			goto st423
		case 424:
			goto st424
		case 425:
			goto st425
		case 426:
			goto st426
		case 427:
			goto st427
		case 428:
			goto st428
		case 429:
			goto st429
		case 430:
			goto st430
		case 431:
			goto st431
		case 432:
			goto st432
		case 433:
			goto st433
		case 434:
			goto st434
		case 435:
			goto st435
		case 436:
			goto st436
		case 437:
			goto st437
		case 438:
			goto st438
		case 439:
			goto st439
		case 440:
			goto st440
		case 441:
			goto st441
		case 442:
			goto st442
		case 443:
			goto st443
		case 444:
			goto st444
		case 445:
			goto st445
		case 446:
			goto st446
		case 447:
			goto st447
		case 448:
			goto st448
		case 449:
			goto st449
		case 106:
			goto st106
		case 107:
			goto st107
		case 108:
			goto st108
		case 109:
			goto st109
		case 110:
			goto st110
		case 111:
			goto st111
		case 450:
			goto st450
		case 451:
			goto st451
		case 112:
			goto st112
		case 452:
			goto st452
		case 453:
			goto st453
		case 454:
			goto st454
		case 455:
			goto st455
		case 456:
			goto st456
		case 457:
			goto st457
		case 458:
			goto st458
		case 459:
			goto st459
		case 460:
			goto st460
		case 461:
			goto st461
		case 462:
			goto st462
		case 463:
			goto st463
		case 464:
			goto st464
		case 465:
			goto st465
		case 466:
			goto st466
		case 467:
			goto st467
		case 468:
			goto st468
		case 469:
			goto st469
		case 470:
			goto st470
		case 471:
			goto st471
		case 472:
			goto st472
		case 473:
			goto st473
		case 474:
			goto st474
		case 475:
			goto st475
		case 476:
			goto st476
		case 477:
			goto st477
		case 478:
			goto st478
		case 479:
			goto st479
		case 480:
			goto st480
		case 481:
			goto st481
		case 482:
			goto st482
		case 483:
			goto st483
		case 484:
			goto st484
		case 485:
			goto st485
		case 486:
			goto st486
		case 487:
			goto st487
		case 488:
			goto st488
		case 489:
			goto st489
		case 490:
			goto st490
		case 491:
			goto st491
		case 492:
			goto st492
		case 493:
			goto st493
		case 494:
			goto st494
		case 495:
			goto st495
		case 496:
			goto st496
		case 497:
			goto st497
		case 498:
			goto st498
		case 499:
			goto st499
		case 500:
			goto st500
		case 501:
			goto st501
		case 502:
			goto st502
		case 503:
			goto st503
		case 504:
			goto st504
		case 505:
			goto st505
		case 506:
			goto st506
		case 507:
			goto st507
		case 508:
			goto st508
		case 509:
			goto st509
		case 510:
			goto st510
		case 511:
			goto st511
		case 512:
			goto st512
		case 513:
			goto st513
		case 514:
			goto st514
		case 515:
			goto st515
		case 516:
			goto st516
		case 517:
			goto st517
		case 518:
			goto st518
		case 519:
			goto st519
		case 520:
			goto st520
		case 521:
			goto st521
		case 522:
			goto st522
		case 523:
			goto st523
		case 524:
			goto st524
		case 113:
			goto st113
		case 525:
			goto st525
		case 526:
			goto st526
		case 0:
			goto st0
		case 527:
			goto st527
		case 528:
			goto st528
		case 529:
			goto st529
		case 530:
			goto st530
		case 531:
			goto st531
		case 114:
			goto st114
		case 532:
			goto st532
		case 533:
			goto st533
		case 534:
			goto st534
		case 535:
			goto st535
		case 536:
			goto st536
		case 537:
			goto st537
		case 115:
			goto st115
		case 116:
			goto st116
		case 538:
			goto st538
		case 539:
			goto st539
		case 540:
			goto st540
		case 541:
			goto st541
		case 542:
			goto st542
		case 543:
			goto st543
		case 117:
			goto st117
		case 118:
			goto st118
		case 544:
			goto st544
		case 545:
			goto st545
		case 546:
			goto st546
		case 547:
			goto st547
		case 548:
			goto st548
		case 549:
			goto st549
		case 550:
			goto st550
		case 551:
			goto st551
		case 552:
			goto st552
		case 553:
			goto st553
		case 554:
			goto st554
		case 119:
			goto st119
		case 555:
			goto st555
		case 120:
			goto st120
		case 121:
			goto st121
		case 556:
			goto st556
		case 557:
			goto st557
		case 558:
			goto st558
		case 559:
			goto st559
		case 560:
			goto st560
		case 561:
			goto st561
		case 562:
			goto st562
		case 563:
			goto st563
		case 122:
			goto st122
		case 123:
			goto st123
		case 564:
			goto st564
		case 124:
			goto st124
		case 565:
			goto st565
		case 566:
			goto st566
		case 567:
			goto st567
		case 568:
			goto st568
		case 125:
			goto st125
		case 569:
			goto st569
		case 570:
			goto st570
		case 571:
			goto st571
		case 572:
			goto st572
		case 573:
			goto st573
		case 574:
			goto st574
		case 575:
			goto st575
		case 576:
			goto st576
		case 577:
			goto st577
		case 126:
			goto st126
		case 127:
			goto st127
		case 128:
			goto st128
		case 129:
			goto st129
		case 578:
			goto st578
		case 579:
			goto st579
		case 580:
			goto st580
		case 581:
			goto st581
		case 582:
			goto st582
		case 583:
			goto st583
		case 584:
			goto st584
		case 585:
			goto st585
		case 586:
			goto st586
		case 130:
			goto st130
		case 131:
			goto st131
		case 132:
			goto st132
		case 133:
			goto st133
		case 587:
			goto st587
		case 588:
			goto st588
		case 589:
			goto st589
		case 590:
			goto st590
		case 591:
			goto st591
		case 592:
			goto st592
		case 593:
			goto st593
		case 594:
			goto st594
		case 595:
			goto st595
		case 134:
			goto st134
		case 135:
			goto st135
		case 136:
			goto st136
		case 137:
			goto st137
		case 596:
			goto st596
		case 597:
			goto st597
		case 598:
			goto st598
		case 599:
			goto st599
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 1:
			goto st_case_1
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 168:
			goto st_case_168
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 66:
			goto st_case_66
		case 174:
			goto st_case_174
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 175:
			goto st_case_175
		case 69:
			goto st_case_69
		case 176:
			goto st_case_176
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 180:
			goto st_case_180
		case 76:
			goto st_case_76
		case 181:
			goto st_case_181
		case 77:
			goto st_case_77
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 89:
			goto st_case_89
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 203:
			goto st_case_203
		case 90:
			goto st_case_90
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 221:
			goto st_case_221
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
		case 225:
			goto st_case_225
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 237:
			goto st_case_237
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 249:
			goto st_case_249
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 258:
			goto st_case_258
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		case 280:
			goto st_case_280
		case 281:
			goto st_case_281
		case 282:
			goto st_case_282
		case 283:
			goto st_case_283
		case 284:
			goto st_case_284
		case 285:
			goto st_case_285
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 288:
			goto st_case_288
		case 289:
			goto st_case_289
		case 290:
			goto st_case_290
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
		case 293:
			goto st_case_293
		case 294:
			goto st_case_294
		case 295:
			goto st_case_295
		case 296:
			goto st_case_296
		case 297:
			goto st_case_297
		case 298:
			goto st_case_298
		case 299:
			goto st_case_299
		case 300:
			goto st_case_300
		case 301:
			goto st_case_301
		case 302:
			goto st_case_302
		case 303:
			goto st_case_303
		case 304:
			goto st_case_304
		case 305:
			goto st_case_305
		case 306:
			goto st_case_306
		case 307:
			goto st_case_307
		case 308:
			goto st_case_308
		case 309:
			goto st_case_309
		case 310:
			goto st_case_310
		case 311:
			goto st_case_311
		case 312:
			goto st_case_312
		case 313:
			goto st_case_313
		case 314:
			goto st_case_314
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 318:
			goto st_case_318
		case 319:
			goto st_case_319
		case 320:
			goto st_case_320
		case 321:
			goto st_case_321
		case 322:
			goto st_case_322
		case 323:
			goto st_case_323
		case 324:
			goto st_case_324
		case 325:
			goto st_case_325
		case 326:
			goto st_case_326
		case 327:
			goto st_case_327
		case 328:
			goto st_case_328
		case 329:
			goto st_case_329
		case 330:
			goto st_case_330
		case 331:
			goto st_case_331
		case 332:
			goto st_case_332
		case 333:
			goto st_case_333
		case 334:
			goto st_case_334
		case 335:
			goto st_case_335
		case 336:
			goto st_case_336
		case 337:
			goto st_case_337
		case 338:
			goto st_case_338
		case 339:
			goto st_case_339
		case 340:
			goto st_case_340
		case 341:
			goto st_case_341
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 346:
			goto st_case_346
		case 347:
			goto st_case_347
		case 348:
			goto st_case_348
		case 349:
			goto st_case_349
		case 350:
			goto st_case_350
		case 351:
			goto st_case_351
		case 352:
			goto st_case_352
		case 353:
			goto st_case_353
		case 354:
			goto st_case_354
		case 355:
			goto st_case_355
		case 356:
			goto st_case_356
		case 357:
			goto st_case_357
		case 358:
			goto st_case_358
		case 359:
			goto st_case_359
		case 360:
			goto st_case_360
		case 361:
			goto st_case_361
		case 362:
			goto st_case_362
		case 363:
			goto st_case_363
		case 364:
			goto st_case_364
		case 365:
			goto st_case_365
		case 366:
			goto st_case_366
		case 367:
			goto st_case_367
		case 368:
			goto st_case_368
		case 369:
			goto st_case_369
		case 370:
			goto st_case_370
		case 371:
			goto st_case_371
		case 372:
			goto st_case_372
		case 373:
			goto st_case_373
		case 93:
			goto st_case_93
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 379:
			goto st_case_379
		case 380:
			goto st_case_380
		case 381:
			goto st_case_381
		case 382:
			goto st_case_382
		case 383:
			goto st_case_383
		case 384:
			goto st_case_384
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 385:
			goto st_case_385
		case 386:
			goto st_case_386
		case 387:
			goto st_case_387
		case 388:
			goto st_case_388
		case 389:
			goto st_case_389
		case 390:
			goto st_case_390
		case 391:
			goto st_case_391
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 392:
			goto st_case_392
		case 393:
			goto st_case_393
		case 394:
			goto st_case_394
		case 395:
			goto st_case_395
		case 396:
			goto st_case_396
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 397:
			goto st_case_397
		case 398:
			goto st_case_398
		case 399:
			goto st_case_399
		case 400:
			goto st_case_400
		case 401:
			goto st_case_401
		case 402:
			goto st_case_402
		case 403:
			goto st_case_403
		case 404:
			goto st_case_404
		case 405:
			goto st_case_405
		case 406:
			goto st_case_406
		case 407:
			goto st_case_407
		case 408:
			goto st_case_408
		case 409:
			goto st_case_409
		case 410:
			goto st_case_410
		case 411:
			goto st_case_411
		case 412:
			goto st_case_412
		case 413:
			goto st_case_413
		case 414:
			goto st_case_414
		case 415:
			goto st_case_415
		case 416:
			goto st_case_416
		case 417:
			goto st_case_417
		case 418:
			goto st_case_418
		case 419:
			goto st_case_419
		case 420:
			goto st_case_420
		case 421:
			goto st_case_421
		case 422:
			goto st_case_422
		case 423:
			goto st_case_423
		case 424:
			goto st_case_424
		case 425:
			goto st_case_425
		case 426:
			goto st_case_426
		case 427:
			goto st_case_427
		case 428:
			goto st_case_428
		case 429:
			goto st_case_429
		case 430:
			goto st_case_430
		case 431:
			goto st_case_431
		case 432:
			goto st_case_432
		case 433:
			goto st_case_433
		case 434:
			goto st_case_434
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 443:
			goto st_case_443
		case 444:
			goto st_case_444
		case 445:
			goto st_case_445
		case 446:
			goto st_case_446
		case 447:
			goto st_case_447
		case 448:
			goto st_case_448
		case 449:
			goto st_case_449
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 450:
			goto st_case_450
		case 451:
			goto st_case_451
		case 112:
			goto st_case_112
		case 452:
			goto st_case_452
		case 453:
			goto st_case_453
		case 454:
			goto st_case_454
		case 455:
			goto st_case_455
		case 456:
			goto st_case_456
		case 457:
			goto st_case_457
		case 458:
			goto st_case_458
		case 459:
			goto st_case_459
		case 460:
			goto st_case_460
		case 461:
			goto st_case_461
		case 462:
			goto st_case_462
		case 463:
			goto st_case_463
		case 464:
			goto st_case_464
		case 465:
			goto st_case_465
		case 466:
			goto st_case_466
		case 467:
			goto st_case_467
		case 468:
			goto st_case_468
		case 469:
			goto st_case_469
		case 470:
			goto st_case_470
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 473:
			goto st_case_473
		case 474:
			goto st_case_474
		case 475:
			goto st_case_475
		case 476:
			goto st_case_476
		case 477:
			goto st_case_477
		case 478:
			goto st_case_478
		case 479:
			goto st_case_479
		case 480:
			goto st_case_480
		case 481:
			goto st_case_481
		case 482:
			goto st_case_482
		case 483:
			goto st_case_483
		case 484:
			goto st_case_484
		case 485:
			goto st_case_485
		case 486:
			goto st_case_486
		case 487:
			goto st_case_487
		case 488:
			goto st_case_488
		case 489:
			goto st_case_489
		case 490:
			goto st_case_490
		case 491:
			goto st_case_491
		case 492:
			goto st_case_492
		case 493:
			goto st_case_493
		case 494:
			goto st_case_494
		case 495:
			goto st_case_495
		case 496:
			goto st_case_496
		case 497:
			goto st_case_497
		case 498:
			goto st_case_498
		case 499:
			goto st_case_499
		case 500:
			goto st_case_500
		case 501:
			goto st_case_501
		case 502:
			goto st_case_502
		case 503:
			goto st_case_503
		case 504:
			goto st_case_504
		case 505:
			goto st_case_505
		case 506:
			goto st_case_506
		case 507:
			goto st_case_507
		case 508:
			goto st_case_508
		case 509:
			goto st_case_509
		case 510:
			goto st_case_510
		case 511:
			goto st_case_511
		case 512:
			goto st_case_512
		case 513:
			goto st_case_513
		case 514:
			goto st_case_514
		case 515:
			goto st_case_515
		case 516:
			goto st_case_516
		case 517:
			goto st_case_517
		case 518:
			goto st_case_518
		case 519:
			goto st_case_519
		case 520:
			goto st_case_520
		case 521:
			goto st_case_521
		case 522:
			goto st_case_522
		case 523:
			goto st_case_523
		case 524:
			goto st_case_524
		case 113:
			goto st_case_113
		case 525:
			goto st_case_525
		case 526:
			goto st_case_526
		case 0:
			goto st_case_0
		case 527:
			goto st_case_527
		case 528:
			goto st_case_528
		case 529:
			goto st_case_529
		case 530:
			goto st_case_530
		case 531:
			goto st_case_531
		case 114:
			goto st_case_114
		case 532:
			goto st_case_532
		case 533:
			goto st_case_533
		case 534:
			goto st_case_534
		case 535:
			goto st_case_535
		case 536:
			goto st_case_536
		case 537:
			goto st_case_537
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 538:
			goto st_case_538
		case 539:
			goto st_case_539
		case 540:
			goto st_case_540
		case 541:
			goto st_case_541
		case 542:
			goto st_case_542
		case 543:
			goto st_case_543
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 544:
			goto st_case_544
		case 545:
			goto st_case_545
		case 546:
			goto st_case_546
		case 547:
			goto st_case_547
		case 548:
			goto st_case_548
		case 549:
			goto st_case_549
		case 550:
			goto st_case_550
		case 551:
			goto st_case_551
		case 552:
			goto st_case_552
		case 553:
			goto st_case_553
		case 554:
			goto st_case_554
		case 119:
			goto st_case_119
		case 555:
			goto st_case_555
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 556:
			goto st_case_556
		case 557:
			goto st_case_557
		case 558:
			goto st_case_558
		case 559:
			goto st_case_559
		case 560:
			goto st_case_560
		case 561:
			goto st_case_561
		case 562:
			goto st_case_562
		case 563:
			goto st_case_563
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 564:
			goto st_case_564
		case 124:
			goto st_case_124
		case 565:
			goto st_case_565
		case 566:
			goto st_case_566
		case 567:
			goto st_case_567
		case 568:
			goto st_case_568
		case 125:
			goto st_case_125
		case 569:
			goto st_case_569
		case 570:
			goto st_case_570
		case 571:
			goto st_case_571
		case 572:
			goto st_case_572
		case 573:
			goto st_case_573
		case 574:
			goto st_case_574
		case 575:
			goto st_case_575
		case 576:
			goto st_case_576
		case 577:
			goto st_case_577
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 578:
			goto st_case_578
		case 579:
			goto st_case_579
		case 580:
			goto st_case_580
		case 581:
			goto st_case_581
		case 582:
			goto st_case_582
		case 583:
			goto st_case_583
		case 584:
			goto st_case_584
		case 585:
			goto st_case_585
		case 586:
			goto st_case_586
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 587:
			goto st_case_587
		case 588:
			goto st_case_588
		case 589:
			goto st_case_589
		case 590:
			goto st_case_590
		case 591:
			goto st_case_591
		case 592:
			goto st_case_592
		case 593:
			goto st_case_593
		case 594:
			goto st_case_594
		case 595:
			goto st_case_595
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 596:
			goto st_case_596
		case 597:
			goto st_case_597
		case 598:
			goto st_case_598
		case 599:
			goto st_case_599
		}
		goto st_out
	tr0:
		lex.cs = 138
		// line internal/php8/scanner.rl:152
		(lex.p) = (lex.te) - 1
		{
			lex.cs = 142
			lex.ungetCnt(1)
		}
		goto _again
	tr218:
		lex.cs = 138
		// line internal/php8/scanner.rl:152
		lex.te = (lex.p) + 1
		{
			lex.cs = 142
			lex.ungetCnt(1)
		}
		goto _again
	tr220:
		lex.cs = 138
		// line internal/php8/scanner.rl:152
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.cs = 142
			lex.ungetCnt(1)
		}
		goto _again
	tr221:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:149
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st138
	st138:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof138
		}
	st_case_138:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:2529
		if lex.data[(lex.p)] == 35 {
			goto tr219
		}
		goto tr218
	tr219:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st139
	st139:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		// line internal/php8/scanner.go:2544
		if lex.data[(lex.p)] == 33 {
			goto st1
		}
		goto tr220
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr3
		}
		goto st1
	tr2:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st140
	st140:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof140
		}
	st_case_140:
		// line internal/php8/scanner.go:2578
		goto tr221
	tr3:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st141
	st141:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof141
		}
	st_case_141:
		// line internal/php8/scanner.go:2597
		if lex.data[(lex.p)] == 10 {
			goto tr2
		}
		goto tr221
	tr4:
		lex.cs = 142
		// line internal/php8/scanner.rl:165
		(lex.p) = (lex.te) - 1
		{
			lex.addOpenTagToken(tkn)
			lex.cs = 151
		}
		goto _again
	tr7:
		lex.cs = 142
		// line internal/php8/scanner.rl:169
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
			lex.cs = 151
		}
		goto _again
	tr226:
		// line internal/php8/scanner.rl:159
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(tkn)
			tok = token.T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 142
				goto _out
			}
		}
		goto st142
	tr228:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:159
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(tkn)
			tok = token.T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 142
				goto _out
			}
		}
		goto st142
	tr235:
		lex.cs = 142
		// line internal/php8/scanner.rl:165
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addOpenTagToken(tkn)
			lex.cs = 151
		}
		goto _again
	tr236:
		lex.cs = 142
		// line internal/php8/scanner.rl:174
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ECHO
			lex.cs = 151
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr238:
		lex.cs = 142
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:169
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
			lex.cs = 151
		}
		goto _again
	st142:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof142
		}
	st_case_142:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:2690
		switch lex.data[(lex.p)] {
		case 10:
			goto tr223
		case 13:
			goto tr224
		case 60:
			goto st147
		}
		goto st143
	tr229:
		// line internal/php8/scanner.rl:54

		goto st143
	st143:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof143
		}
	st_case_143:
		// line internal/php8/scanner.go:2709
		switch lex.data[(lex.p)] {
		case 10:
			goto tr223
		case 13:
			goto tr224
		case 60:
			goto st146
		}
		goto st143
	tr223:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st144
	tr230:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st144
	tr233:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st144
	st144:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		// line internal/php8/scanner.go:2764
		switch lex.data[(lex.p)] {
		case 10:
			goto tr230
		case 13:
			goto tr231
		case 60:
			goto tr232
		}
		goto tr229
	tr224:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st145
	tr231:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st145
	st145:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof145
		}
	st_case_145:
		// line internal/php8/scanner.go:2805
		switch lex.data[(lex.p)] {
		case 10:
			goto tr233
		case 13:
			goto tr231
		case 60:
			goto tr232
		}
		goto tr229
	tr232:
		// line internal/php8/scanner.rl:54

		goto st146
	st146:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof146
		}
	st_case_146:
		// line internal/php8/scanner.go:2824
		switch lex.data[(lex.p)] {
		case 10:
			goto tr223
		case 13:
			goto tr224
		case 60:
			goto st146
		case 63:
			goto tr226
		}
		goto st143
	st147:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr223
		case 13:
			goto tr224
		case 60:
			goto st146
		case 63:
			goto tr234
		}
		goto st143
	tr234:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st148
	st148:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof148
		}
	st_case_148:
		// line internal/php8/scanner.go:2862
		switch lex.data[(lex.p)] {
		case 61:
			goto tr236
		case 80:
			goto st2
		case 112:
			goto st2
		}
		goto tr235
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 72:
			goto st3
		case 104:
			goto st3
		}
		goto tr4
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 80:
			goto st4
		case 112:
			goto st4
		}
		goto tr4
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr7
		case 10:
			goto tr8
		case 13:
			goto tr9
		case 32:
			goto tr7
		}
		goto tr4
	tr8:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st149
	st149:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof149
		}
	st_case_149:
		// line internal/php8/scanner.go:2929
		goto tr238
	tr9:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st150
	st150:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof150
		}
	st_case_150:
		// line internal/php8/scanner.go:2948
		if lex.data[(lex.p)] == 10 {
			goto tr8
		}
		goto tr238
	tr10:
		lex.cs = 151
		// line NONE:1
		switch lex.act {
		case 10:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 2, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 12:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 8, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 13:
			{
				(lex.p) = (lex.te) - 1

				base := 10
				if lex.data[lex.ts] == '0' {
					base = 8
				}

				s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, base, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 16, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAME_RELATIVE
				{
					(lex.p)++
					goto _out
				}
			}
		case 16:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAME_QUALIFIED
				{
					(lex.p)++
					goto _out
				}
			}
		case 19:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ABSTRACT
				{
					(lex.p)++
					goto _out
				}
			}
		case 20:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ARRAY
				{
					(lex.p)++
					goto _out
				}
			}
		case 21:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_AS
				{
					(lex.p)++
					goto _out
				}
			}
		case 22:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_BREAK
				{
					(lex.p)++
					goto _out
				}
			}
		case 23:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CALLABLE
				{
					(lex.p)++
					goto _out
				}
			}
		case 24:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CASE
				{
					(lex.p)++
					goto _out
				}
			}
		case 25:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 26:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CLASS
				{
					(lex.p)++
					goto _out
				}
			}
		case 27:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CLONE
				{
					(lex.p)++
					goto _out
				}
			}
		case 28:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CONST
				{
					(lex.p)++
					goto _out
				}
			}
		case 29:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CONTINUE
				{
					(lex.p)++
					goto _out
				}
			}
		case 30:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 31:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DEFAULT
				{
					(lex.p)++
					goto _out
				}
			}
		case 32:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DO
				{
					(lex.p)++
					goto _out
				}
			}
		case 33:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ECHO
				{
					(lex.p)++
					goto _out
				}
			}
		case 34:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ELSE
				{
					(lex.p)++
					goto _out
				}
			}
		case 35:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ELSEIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 36:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EMPTY
				{
					(lex.p)++
					goto _out
				}
			}
		case 37:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDDECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 38:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDFOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 39:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDFOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 40:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 41:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDSWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 42:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDWHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 43:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EVAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 44:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EXIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 45:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EXTENDS
				{
					(lex.p)++
					goto _out
				}
			}
		case 46:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FINAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 47:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FINALLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 48:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 49:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 50:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FUNCTION
				{
					(lex.p)++
					goto _out
				}
			}
		case 51:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FN
				{
					(lex.p)++
					goto _out
				}
			}
		case 52:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_GLOBAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 53:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_GOTO
				{
					(lex.p)++
					goto _out
				}
			}
		case 54:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_IF
				{
					(lex.p)++
					goto _out
				}
			}
		case 55:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ISSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 56:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_IMPLEMENTS
				{
					(lex.p)++
					goto _out
				}
			}
		case 57:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INSTANCEOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 58:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INSTEADOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 59:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INTERFACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 60:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LIST
				{
					(lex.p)++
					goto _out
				}
			}
		case 61:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAMESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 62:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PRIVATE
				{
					(lex.p)++
					goto _out
				}
			}
		case 64:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PUBLIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 66:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PRINT
				{
					(lex.p)++
					goto _out
				}
			}
		case 67:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PROTECTED
				{
					(lex.p)++
					goto _out
				}
			}
		case 69:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_READONLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 70:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_RETURN
				{
					(lex.p)++
					goto _out
				}
			}
		case 71:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_STATIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 72:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_SWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 73:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_MATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 74:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_THROW
				{
					(lex.p)++
					goto _out
				}
			}
		case 75:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_TRAIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 76:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_TRY
				{
					(lex.p)++
					goto _out
				}
			}
		case 77:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_UNSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 78:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_USE
				{
					(lex.p)++
					goto _out
				}
			}
		case 79:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_VAR
				{
					(lex.p)++
					goto _out
				}
			}
		case 80:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_WHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 82:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_YIELD
				{
					(lex.p)++
					goto _out
				}
			}
		case 83:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INCLUDE
				{
					(lex.p)++
					goto _out
				}
			}
		case 84:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INCLUDE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 85:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_REQUIRE
				{
					(lex.p)++
					goto _out
				}
			}
		case 86:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_REQUIRE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 87:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CLASS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 88:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DIR
				{
					(lex.p)++
					goto _out
				}
			}
		case 89:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 90:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FUNC_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 91:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LINE
				{
					(lex.p)++
					goto _out
				}
			}
		case 92:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 93:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_METHOD_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 94:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_TRAIT_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 95:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_HALT_COMPILER
				lex.cs = 569
				{
					(lex.p)++
					goto _out
				}
			}
		case 96:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NEW
				{
					(lex.p)++
					goto _out
				}
			}
		case 97:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LOGICAL_AND
				{
					(lex.p)++
					goto _out
				}
			}
		case 98:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LOGICAL_OR
				{
					(lex.p)++
					goto _out
				}
			}
		case 99:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LOGICAL_XOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 128:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_SL
				{
					(lex.p)++
					goto _out
				}
			}
		case 148:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_STRING
				{
					(lex.p)++
					goto _out
				}
			}
		case 154:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('"'))
				lex.cs = 543
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr13:
		// line internal/php8/scanner.rl:383
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr17:
		// line internal/php8/scanner.rl:406
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st151
	tr21:
		// line internal/php8/scanner.rl:369
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr36:
		// line internal/php8/scanner.rl:354
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ARRAY_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr43:
		// line internal/php8/scanner.rl:359
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr47:
		// line internal/php8/scanner.rl:355
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOL_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr55:
		// line internal/php8/scanner.rl:356
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOUBLE_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr62:
		// line internal/php8/scanner.rl:357
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INT_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr71:
		// line internal/php8/scanner.rl:358
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr80:
		// line internal/php8/scanner.rl:360
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_UNSET_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr84:
		// line internal/php8/scanner.rl:361
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VOID_CAST
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr85:
		// line internal/php8/scanner.rl:321
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ELLIPSIS
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr88:
		// line internal/php8/scanner.rl:187
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr99:
		// line internal/php8/scanner.rl:67
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st151
	tr100:
		// line internal/php8/scanner.rl:208
		(lex.p) = (lex.te) - 1
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 151
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr129:
		lex.cs = 151
		// line internal/php8/scanner.rl:381
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			lex.cs = 519
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr131:
		// line internal/php8/scanner.rl:378
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr134:
		// line internal/php8/scanner.rl:282
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PRIVATE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr138:
		// line internal/php8/scanner.rl:283
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PRIVATE_SET
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr139:
		// line internal/php8/scanner.rl:287
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PROTECTED
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr143:
		// line internal/php8/scanner.rl:288
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PROTECTED_SET
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr144:
		// line internal/php8/scanner.rl:284
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PUBLIC
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr148:
		// line internal/php8/scanner.rl:285
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PUBLIC_SET
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr149:
		// line internal/php8/scanner.rl:302
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr161:
		// line internal/php8/scanner.rl:301
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD_FROM
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr162:
		// line internal/php8/scanner.rl:236
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_FULLY_QUALIFIED
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr239:
		// line internal/php8/scanner.rl:406
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st151
	tr251:
		// line internal/php8/scanner.rl:369
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr289:
		lex.cs = 151
		// line internal/php8/scanner.rl:403
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('`'))
			lex.cs = 537
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr290:
		// line internal/php8/scanner.rl:375
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('{'))
			lex.call(151, 151)
			goto _out
		}
		goto st151
	tr292:
		// line internal/php8/scanner.rl:376
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('}'))
			lex.ret(1)
			goto _out
		}
		goto st151
	tr293:
		// line internal/php8/scanner.rl:183
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st151
	tr295:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:183
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st151
	tr300:
		// line internal/php8/scanner.rl:369
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr302:
		// line internal/php8/scanner.rl:339
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr303:
		// line internal/php8/scanner.rl:340
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_IDENTICAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr304:
		lex.cs = 151
		// line internal/php8/scanner.rl:404
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('"'))
			lex.cs = 543
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr305:
		// line internal/php8/scanner.rl:364
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st151
	tr307:
		// line internal/php8/scanner.rl:320
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ATTRIBUTE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr308:
		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st151
	tr311:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st151
	tr317:
		// line internal/php8/scanner.rl:377
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr318:
		// line internal/php8/scanner.rl:334
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MOD_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr319:
		// line internal/php8/scanner.rl:323
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOLEAN_AND
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr320:
		// line internal/php8/scanner.rl:325
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_AND_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr321:
		// line internal/php8/scanner.rl:406
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st151
	tr323:
		// line internal/php8/scanner.rl:328
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MUL_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr324:
		// line internal/php8/scanner.rl:347
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_POW
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr325:
		// line internal/php8/scanner.rl:329
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_POW_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr326:
		// line internal/php8/scanner.rl:336
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INC
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr327:
		// line internal/php8/scanner.rl:331
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PLUS_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr328:
		// line internal/php8/scanner.rl:335
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DEC
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr329:
		// line internal/php8/scanner.rl:332
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MINUS_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr330:
		lex.cs = 151
		// line internal/php8/scanner.rl:380
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			lex.cs = 519
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr332:
		// line internal/php8/scanner.rl:327
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_CONCAT_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr333:
		// line internal/php8/scanner.rl:187
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr336:
		// line internal/php8/scanner.rl:330
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DIV_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr337:
		// line internal/php8/scanner.rl:208
		lex.te = (lex.p)
		(lex.p)--
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 151
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr343:
		// line internal/php8/scanner.rl:198
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 8, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 151
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr344:
		// line internal/php8/scanner.rl:188
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 2, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 151
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr345:
		// line internal/php8/scanner.rl:223
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 16, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 151
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr346:
		// line internal/php8/scanner.rl:322
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PAAMAYIM_NEKUDOTAYIM
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr347:
		lex.cs = 151
		// line internal/php8/scanner.rl:185
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 142
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr350:
		lex.cs = 151
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:185
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 142
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr353:
		// line internal/php8/scanner.rl:339
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr354:
		// line internal/php8/scanner.rl:348
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr355:
		// line internal/php8/scanner.rl:343
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SL_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr356:
		lex.cs = 151
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:389
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.heredocLabel = lex.data[lblStart:lblEnd]
			lex.setTokenPosition(tkn)
			tok = token.T_START_HEREDOC

			if lex.isHeredocEnd(lex.p + 1) {
				lex.cs = 549
			} else if lex.data[lblStart-1] == '\'' {
				lex.cs = 526
			} else {
				lex.cs = 530
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr357:
		// line internal/php8/scanner.rl:346
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr358:
		// line internal/php8/scanner.rl:338
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SPACESHIP
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr360:
		// line internal/php8/scanner.rl:337
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOUBLE_ARROW
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr361:
		// line internal/php8/scanner.rl:341
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr362:
		// line internal/php8/scanner.rl:342
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_IDENTICAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr363:
		// line internal/php8/scanner.rl:345
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr365:
		// line internal/php8/scanner.rl:349
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SR
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr366:
		// line internal/php8/scanner.rl:344
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SR_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr370:
		lex.cs = 151
		// line internal/php8/scanner.rl:184
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 142
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr373:
		lex.cs = 151
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:184
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 142
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr374:
		// line internal/php8/scanner.rl:350
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_COALESCE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr375:
		// line internal/php8/scanner.rl:351
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_COALESCE_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr376:
		// line internal/php8/scanner.rl:378
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr382:
		// line internal/php8/scanner.rl:235
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_QUALIFIED
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr458:
		// line internal/php8/scanner.rl:254
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ELSE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr478:
		// line internal/php8/scanner.rl:258
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENDFOR
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr507:
		// line internal/php8/scanner.rl:266
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_FINAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr511:
		// line internal/php8/scanner.rl:268
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_FOR
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr543:
		// line internal/php8/scanner.rl:303
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INCLUDE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr586:
		// line internal/php8/scanner.rl:281
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAMESPACE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr588:
		// line internal/php8/scanner.rl:234
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_RELATIVE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr601:
		// line internal/php8/scanner.rl:282
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PRIVATE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr609:
		// line internal/php8/scanner.rl:287
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PROTECTED
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr615:
		// line internal/php8/scanner.rl:284
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PUBLIC
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr630:
		// line internal/php8/scanner.rl:305
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_REQUIRE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr676:
		// line internal/php8/scanner.rl:302
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr677:
		// line internal/php8/scanner.rl:237
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NS_SEPARATOR
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr678:
		// line internal/php8/scanner.rl:236
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_FULLY_QUALIFIED
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr680:
		// line internal/php8/scanner.rl:333
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_XOR_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr754:
		// line internal/php8/scanner.rl:326
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr755:
		// line internal/php8/scanner.rl:352
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PIPE
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	tr756:
		// line internal/php8/scanner.rl:324
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOLEAN_OR
			{
				(lex.p)++
				lex.cs = 151
				goto _out
			}
		}
		goto st151
	st151:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof151
		}
	st_case_151:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:3946
		switch lex.data[(lex.p)] {
		case 10:
			goto tr241
		case 13:
			goto tr242
		case 32:
			goto st152
		case 33:
			goto st156
		case 34:
			goto tr244
		case 35:
			goto st159
		case 36:
			goto st163
		case 37:
			goto st165
		case 38:
			goto st166
		case 39:
			goto tr249
		case 40:
			goto tr250
		case 42:
			goto st169
		case 43:
			goto st171
		case 45:
			goto st172
		case 46:
			goto tr255
		case 47:
			goto tr256
		case 48:
			goto tr257
		case 58:
			goto st183
		case 59:
			goto tr259
		case 60:
			goto st188
		case 61:
			goto st193
		case 62:
			goto st195
		case 63:
			goto tr263
		case 64:
			goto tr251
		case 65:
			goto tr264
		case 66:
			goto tr265
		case 67:
			goto tr266
		case 68:
			goto tr267
		case 69:
			goto tr268
		case 70:
			goto tr269
		case 71:
			goto tr270
		case 73:
			goto tr272
		case 76:
			goto tr273
		case 77:
			goto tr274
		case 78:
			goto tr275
		case 79:
			goto tr276
		case 80:
			goto tr277
		case 82:
			goto tr278
		case 83:
			goto tr279
		case 84:
			goto tr280
		case 85:
			goto tr281
		case 86:
			goto tr282
		case 87:
			goto tr283
		case 88:
			goto tr284
		case 89:
			goto tr285
		case 92:
			goto st450
		case 94:
			goto st452
		case 95:
			goto tr288
		case 96:
			goto tr289
		case 97:
			goto tr264
		case 98:
			goto tr265
		case 99:
			goto tr266
		case 100:
			goto tr267
		case 101:
			goto tr268
		case 102:
			goto tr269
		case 103:
			goto tr270
		case 105:
			goto tr272
		case 108:
			goto tr273
		case 109:
			goto tr274
		case 110:
			goto tr275
		case 111:
			goto tr276
		case 112:
			goto tr277
		case 114:
			goto tr278
		case 115:
			goto tr279
		case 116:
			goto tr280
		case 117:
			goto tr281
		case 118:
			goto tr282
		case 119:
			goto tr283
		case 120:
			goto tr284
		case 121:
			goto tr285
		case 123:
			goto tr290
		case 124:
			goto st518
		case 125:
			goto tr292
		case 126:
			goto tr251
		case 127:
			goto tr239
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st152
				}
			default:
				goto tr239
			}
		case lex.data[(lex.p)] > 31:
			switch {
			case lex.data[(lex.p)] < 49:
				if 41 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 44 {
					goto tr251
				}
			case lex.data[(lex.p)] > 57:
				if 91 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 93 {
					goto tr251
				}
			default:
				goto tr101
			}
		default:
			goto tr239
		}
		goto tr271
	tr296:
		// line internal/php8/scanner.rl:54

		goto st152
	st152:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		// line internal/php8/scanner.go:4135
		switch lex.data[(lex.p)] {
		case 10:
			goto tr241
		case 13:
			goto tr294
		case 32:
			goto st152
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st152
		}
		goto tr293
	tr241:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st153
	tr297:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st153
	tr299:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st153
	st153:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		// line internal/php8/scanner.go:4193
		switch lex.data[(lex.p)] {
		case 10:
			goto tr297
		case 13:
			goto tr298
		case 32:
			goto tr296
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr296
		}
		goto tr295
	tr294:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st154
	tr298:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st154
	st154:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof154
		}
	st_case_154:
		// line internal/php8/scanner.go:4237
		switch lex.data[(lex.p)] {
		case 10:
			goto tr299
		case 13:
			goto tr298
		case 32:
			goto tr296
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr296
		}
		goto tr295
	tr242:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st155
	st155:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof155
		}
	st_case_155:
		// line internal/php8/scanner.go:4267
		switch lex.data[(lex.p)] {
		case 10:
			goto tr299
		case 13:
			goto tr298
		case 32:
			goto tr296
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr296
		}
		goto tr295
	st156:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof156
		}
	st_case_156:
		if lex.data[(lex.p)] == 61 {
			goto st157
		}
		goto tr300
	st157:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof157
		}
	st_case_157:
		if lex.data[(lex.p)] == 61 {
			goto tr303
		}
		goto tr302
	tr244:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:404
		lex.act = 154
		goto st158
	st158:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof158
		}
	st_case_158:
		// line internal/php8/scanner.go:4310
		switch lex.data[(lex.p)] {
		case 10:
			goto tr12
		case 13:
			goto tr12
		case 34:
			goto tr13
		case 36:
			goto st6
		case 92:
			goto st7
		case 123:
			goto st8
		}
		goto st5
	tr12:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		// line internal/php8/scanner.go:4343
		switch lex.data[(lex.p)] {
		case 10:
			goto tr12
		case 13:
			goto tr12
		case 34:
			goto tr13
		case 36:
			goto st6
		case 92:
			goto st7
		case 123:
			goto st8
		}
		goto st5
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr12
		case 13:
			goto tr12
		case 34:
			goto tr13
		case 36:
			goto st6
		case 92:
			goto st7
		case 96:
			goto st5
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st5
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr10
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr12
		case 13:
			goto tr12
		}
		goto st5
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr12
		case 13:
			goto tr12
		case 34:
			goto tr13
		case 36:
			goto tr10
		case 92:
			goto st7
		}
		goto st5
	st159:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof159
		}
	st_case_159:
		if lex.data[(lex.p)] == 91 {
			goto tr307
		}
		goto st160
	tr312:
		// line internal/php8/scanner.rl:54

		goto st160
	st160:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof160
		}
	st_case_160:
		// line internal/php8/scanner.go:4439
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr309
		case 525:
			goto tr310
		}
		if 512 <= _widec && _widec <= 767 {
			goto st160
		}
		goto tr308
	tr309:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st161
	tr313:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st161
	tr315:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st161
	st161:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		// line internal/php8/scanner.go:4533
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr313
		case 525:
			goto tr314
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr312
		}
		goto tr311
	tr310:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st162
	tr314:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st162
	st162:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof162
		}
	st_case_162:
		// line internal/php8/scanner.go:4613
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr315
		case 525:
			goto tr314
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr312
		}
		goto tr311
	st163:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		if lex.data[(lex.p)] == 96 {
			goto tr300
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr300
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr300
			}
		default:
			goto tr300
		}
		goto st164
	st164:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		if lex.data[(lex.p)] == 96 {
			goto tr317
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr317
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr317
				}
			case lex.data[(lex.p)] >= 91:
				goto tr317
			}
		default:
			goto tr317
		}
		goto st164
	st165:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		if lex.data[(lex.p)] == 61 {
			goto tr318
		}
		goto tr300
	st166:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof166
		}
	st_case_166:
		switch lex.data[(lex.p)] {
		case 38:
			goto tr319
		case 61:
			goto tr320
		}
		goto tr300
	tr249:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st167
	st167:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		// line internal/php8/scanner.go:4740
		switch lex.data[(lex.p)] {
		case 10:
			goto tr19
		case 13:
			goto tr19
		case 39:
			goto tr13
		case 92:
			goto st10
		}
		goto st9
	tr19:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st9
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		// line internal/php8/scanner.go:4769
		switch lex.data[(lex.p)] {
		case 10:
			goto tr19
		case 13:
			goto tr19
		case 39:
			goto tr13
		case 92:
			goto st10
		}
		goto st9
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr19
		case 13:
			goto tr19
		}
		goto st9
	tr250:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st168
	st168:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof168
		}
	st_case_168:
		// line internal/php8/scanner.go:4803
		switch lex.data[(lex.p)] {
		case 9:
			goto st11
		case 32:
			goto st11
		case 65:
			goto st12
		case 66:
			goto st17
		case 68:
			goto st29
		case 70:
			goto st35
		case 73:
			goto st39
		case 79:
			goto st46
		case 83:
			goto st52
		case 85:
			goto st57
		case 86:
			goto st62
		case 97:
			goto st12
		case 98:
			goto st17
		case 100:
			goto st29
		case 102:
			goto st35
		case 105:
			goto st39
		case 111:
			goto st46
		case 115:
			goto st52
		case 117:
			goto st57
		case 118:
			goto st62
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st11
		}
		goto tr300
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch lex.data[(lex.p)] {
		case 9:
			goto st11
		case 32:
			goto st11
		case 65:
			goto st12
		case 66:
			goto st17
		case 68:
			goto st29
		case 70:
			goto st35
		case 73:
			goto st39
		case 79:
			goto st46
		case 83:
			goto st52
		case 85:
			goto st57
		case 86:
			goto st62
		case 97:
			goto st12
		case 98:
			goto st17
		case 100:
			goto st29
		case 102:
			goto st35
		case 105:
			goto st39
		case 111:
			goto st46
		case 115:
			goto st52
		case 117:
			goto st57
		case 118:
			goto st62
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st11
		}
		goto tr21
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch lex.data[(lex.p)] {
		case 82:
			goto st13
		case 114:
			goto st13
		}
		goto tr21
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch lex.data[(lex.p)] {
		case 82:
			goto st14
		case 114:
			goto st14
		}
		goto tr21
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch lex.data[(lex.p)] {
		case 65:
			goto st15
		case 97:
			goto st15
		}
		goto tr21
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 89:
			goto st16
		case 121:
			goto st16
		}
		goto tr21
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch lex.data[(lex.p)] {
		case 9:
			goto st16
		case 32:
			goto st16
		case 41:
			goto tr36
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st16
		}
		goto tr21
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch lex.data[(lex.p)] {
		case 73:
			goto st18
		case 79:
			goto st23
		case 105:
			goto st18
		case 111:
			goto st23
		}
		goto tr21
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 78:
			goto st19
		case 110:
			goto st19
		}
		goto tr21
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 65:
			goto st20
		case 97:
			goto st20
		}
		goto tr21
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 82:
			goto st21
		case 114:
			goto st21
		}
		goto tr21
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 89:
			goto st22
		case 121:
			goto st22
		}
		goto tr21
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 9:
			goto st22
		case 32:
			goto st22
		case 41:
			goto tr43
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st22
		}
		goto tr21
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 79:
			goto st24
		case 111:
			goto st24
		}
		goto tr21
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 76:
			goto st25
		case 108:
			goto st25
		}
		goto tr21
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 9:
			goto st26
		case 32:
			goto st26
		case 41:
			goto tr47
		case 69:
			goto st27
		case 101:
			goto st27
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st26
		}
		goto tr21
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 9:
			goto st26
		case 32:
			goto st26
		case 41:
			goto tr47
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st26
		}
		goto tr21
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 65:
			goto st28
		case 97:
			goto st28
		}
		goto tr21
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 78:
			goto st26
		case 110:
			goto st26
		}
		goto tr21
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 79:
			goto st30
		case 111:
			goto st30
		}
		goto tr21
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 85:
			goto st31
		case 117:
			goto st31
		}
		goto tr21
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 66:
			goto st32
		case 98:
			goto st32
		}
		goto tr21
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 76:
			goto st33
		case 108:
			goto st33
		}
		goto tr21
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 69:
			goto st34
		case 101:
			goto st34
		}
		goto tr21
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 9:
			goto st34
		case 32:
			goto st34
		case 41:
			goto tr55
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st34
		}
		goto tr21
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 76:
			goto st36
		case 108:
			goto st36
		}
		goto tr21
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 79:
			goto st37
		case 111:
			goto st37
		}
		goto tr21
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 65:
			goto st38
		case 97:
			goto st38
		}
		goto tr21
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 84:
			goto st34
		case 116:
			goto st34
		}
		goto tr21
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 78:
			goto st40
		case 110:
			goto st40
		}
		goto tr21
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 84:
			goto st41
		case 116:
			goto st41
		}
		goto tr21
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 9:
			goto st42
		case 32:
			goto st42
		case 41:
			goto tr62
		case 69:
			goto st43
		case 101:
			goto st43
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st42
		}
		goto tr21
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 9:
			goto st42
		case 32:
			goto st42
		case 41:
			goto tr62
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st42
		}
		goto tr21
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 71:
			goto st44
		case 103:
			goto st44
		}
		goto tr21
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 69:
			goto st45
		case 101:
			goto st45
		}
		goto tr21
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 82:
			goto st42
		case 114:
			goto st42
		}
		goto tr21
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 66:
			goto st47
		case 98:
			goto st47
		}
		goto tr21
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 74:
			goto st48
		case 106:
			goto st48
		}
		goto tr21
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 69:
			goto st49
		case 101:
			goto st49
		}
		goto tr21
	st49:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch lex.data[(lex.p)] {
		case 67:
			goto st50
		case 99:
			goto st50
		}
		goto tr21
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 84:
			goto st51
		case 116:
			goto st51
		}
		goto tr21
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 9:
			goto st51
		case 32:
			goto st51
		case 41:
			goto tr71
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st51
		}
		goto tr21
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 84:
			goto st53
		case 116:
			goto st53
		}
		goto tr21
	st53:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch lex.data[(lex.p)] {
		case 82:
			goto st54
		case 114:
			goto st54
		}
		goto tr21
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 73:
			goto st55
		case 105:
			goto st55
		}
		goto tr21
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 78:
			goto st56
		case 110:
			goto st56
		}
		goto tr21
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 71:
			goto st22
		case 103:
			goto st22
		}
		goto tr21
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 78:
			goto st58
		case 110:
			goto st58
		}
		goto tr21
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 83:
			goto st59
		case 115:
			goto st59
		}
		goto tr21
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 69:
			goto st60
		case 101:
			goto st60
		}
		goto tr21
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 84:
			goto st61
		case 116:
			goto st61
		}
		goto tr21
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 9:
			goto st61
		case 32:
			goto st61
		case 41:
			goto tr80
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st61
		}
		goto tr21
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 79:
			goto st63
		case 111:
			goto st63
		}
		goto tr21
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 73:
			goto st64
		case 105:
			goto st64
		}
		goto tr21
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 68:
			goto st65
		case 100:
			goto st65
		}
		goto tr21
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 9:
			goto st65
		case 32:
			goto st65
		case 41:
			goto tr84
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st65
		}
		goto tr21
	st169:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof169
		}
	st_case_169:
		switch lex.data[(lex.p)] {
		case 42:
			goto st170
		case 61:
			goto tr323
		}
		goto tr300
	st170:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof170
		}
	st_case_170:
		if lex.data[(lex.p)] == 61 {
			goto tr325
		}
		goto tr324
	st171:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof171
		}
	st_case_171:
		switch lex.data[(lex.p)] {
		case 43:
			goto tr326
		case 61:
			goto tr327
		}
		goto tr300
	st172:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		switch lex.data[(lex.p)] {
		case 45:
			goto tr328
		case 61:
			goto tr329
		case 62:
			goto tr330
		}
		goto tr300
	tr255:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st173
	st173:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof173
		}
	st_case_173:
		// line internal/php8/scanner.go:5668
		switch lex.data[(lex.p)] {
		case 46:
			goto st66
		case 61:
			goto tr332
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr300
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		if lex.data[(lex.p)] == 46 {
			goto tr85
		}
		goto tr21
	tr89:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:187
		lex.act = 10
		goto st174
	st174:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof174
		}
	st_case_174:
		// line internal/php8/scanner.go:5700
		switch lex.data[(lex.p)] {
		case 69:
			goto st67
		case 95:
			goto st69
		case 101:
			goto st67
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr333
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch lex.data[(lex.p)] {
		case 43:
			goto st68
		case 45:
			goto st68
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr87
		}
		goto tr10
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr87
		}
		goto tr10
	tr87:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:187
		lex.act = 10
		goto st175
	st175:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof175
		}
	st_case_175:
		// line internal/php8/scanner.go:5749
		if lex.data[(lex.p)] == 95 {
			goto st68
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr87
		}
		goto tr333
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr88
	tr256:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st176
	st176:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof176
		}
	st_case_176:
		// line internal/php8/scanner.go:5776
		switch lex.data[(lex.p)] {
		case 42:
			goto st70
		case 47:
			goto st160
		case 61:
			goto tr336
		}
		goto tr300
	tr94:
		// line internal/php8/scanner.rl:54

		goto st70
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		// line internal/php8/scanner.go:5795
		switch lex.data[(lex.p)] {
		case 10:
			goto tr91
		case 13:
			goto tr92
		case 42:
			goto st73
		}
		goto st70
	tr91:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st71
	tr95:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st71
	tr98:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st71
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		// line internal/php8/scanner.go:5850
		switch lex.data[(lex.p)] {
		case 10:
			goto tr95
		case 13:
			goto tr96
		case 42:
			goto tr97
		}
		goto tr94
	tr92:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st72
	tr96:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st72
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		// line internal/php8/scanner.go:5891
		switch lex.data[(lex.p)] {
		case 10:
			goto tr98
		case 13:
			goto tr96
		case 42:
			goto tr97
		}
		goto tr94
	tr97:
		// line internal/php8/scanner.rl:54

		goto st73
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		// line internal/php8/scanner.go:5910
		switch lex.data[(lex.p)] {
		case 10:
			goto tr91
		case 13:
			goto tr92
		case 42:
			goto st73
		case 47:
			goto tr99
		}
		goto st70
	tr257:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:208
		lex.act = 13
		goto st177
	st177:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof177
		}
	st_case_177:
		// line internal/php8/scanner.go:5934
		switch lex.data[(lex.p)] {
		case 46:
			goto tr338
		case 69:
			goto st67
		case 79:
			goto st75
		case 95:
			goto st74
		case 98:
			goto st76
		case 101:
			goto st67
		case 111:
			goto st75
		case 120:
			goto st77
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr101
		}
		goto tr337
	tr338:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:187
		lex.act = 10
		goto st178
	st178:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof178
		}
	st_case_178:
		// line internal/php8/scanner.go:5969
		switch lex.data[(lex.p)] {
		case 69:
			goto st67
		case 101:
			goto st67
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr333
	tr101:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:208
		lex.act = 13
		goto st179
	st179:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof179
		}
	st_case_179:
		// line internal/php8/scanner.go:5992
		switch lex.data[(lex.p)] {
		case 46:
			goto tr338
		case 69:
			goto st67
		case 95:
			goto st74
		case 101:
			goto st67
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr101
		}
		goto tr337
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr101
		}
		goto tr100
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr102
		}
		goto tr10
	tr102:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:198
		lex.act = 12
		goto st180
	st180:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof180
		}
	st_case_180:
		// line internal/php8/scanner.go:6037
		if lex.data[(lex.p)] == 95 {
			goto st75
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr102
		}
		goto tr343
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr103
		}
		goto tr10
	tr103:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:188
		lex.act = 11
		goto st181
	st181:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof181
		}
	st_case_181:
		// line internal/php8/scanner.go:6066
		if lex.data[(lex.p)] == 95 {
			goto st76
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr103
		}
		goto tr344
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr104
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto tr10
	tr104:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:223
		lex.act = 14
		goto st182
	st182:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof182
		}
	st_case_182:
		// line internal/php8/scanner.go:6104
		if lex.data[(lex.p)] == 95 {
			goto st77
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr104
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto tr345
	st183:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof183
		}
	st_case_183:
		if lex.data[(lex.p)] == 58 {
			goto tr346
		}
		goto tr300
	tr259:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st184
	st184:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof184
		}
	st_case_184:
		// line internal/php8/scanner.go:6140
		switch lex.data[(lex.p)] {
		case 10:
			goto tr106
		case 13:
			goto tr107
		case 32:
			goto st78
		case 63:
			goto st81
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st78
		}
		goto tr300
	tr109:
		// line internal/php8/scanner.rl:54

		goto st78
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		// line internal/php8/scanner.go:6164
		switch lex.data[(lex.p)] {
		case 10:
			goto tr106
		case 13:
			goto tr107
		case 32:
			goto st78
		case 63:
			goto st81
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st78
		}
		goto tr21
	tr106:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st79
	tr110:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st79
	tr113:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st79
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		// line internal/php8/scanner.go:6224
		switch lex.data[(lex.p)] {
		case 10:
			goto tr110
		case 13:
			goto tr111
		case 32:
			goto tr109
		case 63:
			goto tr112
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr109
		}
		goto tr21
	tr107:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st80
	tr111:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st80
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		// line internal/php8/scanner.go:6270
		switch lex.data[(lex.p)] {
		case 10:
			goto tr113
		case 13:
			goto tr111
		case 32:
			goto tr109
		case 63:
			goto tr112
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr109
		}
		goto tr21
	tr112:
		// line internal/php8/scanner.rl:54

		goto st81
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		// line internal/php8/scanner.go:6294
		if lex.data[(lex.p)] == 62 {
			goto st185
		}
		goto tr21
	st185:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof185
		}
	st_case_185:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr348
		case 13:
			goto tr349
		}
		goto tr347
	tr348:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st186
	st186:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof186
		}
	st_case_186:
		// line internal/php8/scanner.go:6328
		goto tr350
	tr349:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st187
	st187:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof187
		}
	st_case_187:
		// line internal/php8/scanner.go:6347
		if lex.data[(lex.p)] == 10 {
			goto tr348
		}
		goto tr350
	st188:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof188
		}
	st_case_188:
		switch lex.data[(lex.p)] {
		case 60:
			goto tr351
		case 61:
			goto st192
		case 62:
			goto tr353
		}
		goto tr300
	tr351:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:348
		lex.act = 128
		goto st189
	st189:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof189
		}
	st_case_189:
		// line internal/php8/scanner.go:6378
		switch lex.data[(lex.p)] {
		case 60:
			goto st82
		case 61:
			goto tr355
		}
		goto tr354
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch lex.data[(lex.p)] {
		case 9:
			goto st82
		case 32:
			goto st82
		case 34:
			goto st83
		case 39:
			goto st86
		case 96:
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr10
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr118
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		if lex.data[(lex.p)] == 96 {
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr10
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr119
	tr119:
		// line internal/php8/scanner.rl:35
		lblStart = lex.p
		goto st84
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		// line internal/php8/scanner.go:6446
		switch lex.data[(lex.p)] {
		case 34:
			goto tr120
		case 96:
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr10
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr10
				}
			case lex.data[(lex.p)] >= 91:
				goto tr10
			}
		default:
			goto tr10
		}
		goto st84
	tr120:
		// line internal/php8/scanner.rl:36
		lblEnd = lex.p
		goto st85
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		// line internal/php8/scanner.go:6480
		switch lex.data[(lex.p)] {
		case 10:
			goto tr122
		case 13:
			goto tr123
		}
		goto tr10
	tr122:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st190
	tr126:
		// line internal/php8/scanner.rl:36
		lblEnd = lex.p
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st190
	st190:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof190
		}
	st_case_190:
		// line internal/php8/scanner.go:6519
		goto tr356
	tr123:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st191
	tr127:
		// line internal/php8/scanner.rl:36
		lblEnd = lex.p
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st191
	st191:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof191
		}
	st_case_191:
		// line internal/php8/scanner.go:6552
		if lex.data[(lex.p)] == 10 {
			goto tr122
		}
		goto tr356
	st86:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		if lex.data[(lex.p)] == 96 {
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr10
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr124
	tr124:
		// line internal/php8/scanner.rl:35
		lblStart = lex.p
		goto st87
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		// line internal/php8/scanner.go:6587
		switch lex.data[(lex.p)] {
		case 39:
			goto tr120
		case 96:
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr10
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr10
				}
			case lex.data[(lex.p)] >= 91:
				goto tr10
			}
		default:
			goto tr10
		}
		goto st87
	tr118:
		// line internal/php8/scanner.rl:35
		lblStart = lex.p
		goto st88
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		// line internal/php8/scanner.go:6621
		switch lex.data[(lex.p)] {
		case 10:
			goto tr126
		case 13:
			goto tr127
		case 96:
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr10
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr10
				}
			case lex.data[(lex.p)] >= 91:
				goto tr10
			}
		default:
			goto tr10
		}
		goto st88
	st192:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof192
		}
	st_case_192:
		if lex.data[(lex.p)] == 62 {
			goto tr358
		}
		goto tr357
	st193:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof193
		}
	st_case_193:
		switch lex.data[(lex.p)] {
		case 61:
			goto st194
		case 62:
			goto tr360
		}
		goto tr300
	st194:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof194
		}
	st_case_194:
		if lex.data[(lex.p)] == 61 {
			goto tr362
		}
		goto tr361
	st195:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof195
		}
	st_case_195:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr363
		case 62:
			goto st196
		}
		goto tr300
	st196:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof196
		}
	st_case_196:
		if lex.data[(lex.p)] == 61 {
			goto tr366
		}
		goto tr365
	tr263:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st197
	st197:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof197
		}
	st_case_197:
		// line internal/php8/scanner.go:6709
		switch lex.data[(lex.p)] {
		case 45:
			goto st89
		case 62:
			goto st198
		case 63:
			goto st201
		}
		goto tr300
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		if lex.data[(lex.p)] == 62 {
			goto tr129
		}
		goto tr21
	st198:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof198
		}
	st_case_198:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr371
		case 13:
			goto tr372
		}
		goto tr370
	tr371:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st199
	st199:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof199
		}
	st_case_199:
		// line internal/php8/scanner.go:6757
		goto tr373
	tr372:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st200
	st200:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof200
		}
	st_case_200:
		// line internal/php8/scanner.go:6776
		if lex.data[(lex.p)] == 10 {
			goto tr371
		}
		goto tr373
	st201:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof201
		}
	st_case_201:
		if lex.data[(lex.p)] == 61 {
			goto tr375
		}
		goto tr374
	tr264:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st202
	st202:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof202
		}
	st_case_202:
		// line internal/php8/scanner.go:6802
		switch lex.data[(lex.p)] {
		case 66:
			goto tr377
		case 78:
			goto tr378
		case 82:
			goto tr379
		case 83:
			goto tr380
		case 92:
			goto st90
		case 96:
			goto tr376
		case 98:
			goto tr377
		case 110:
			goto tr378
		case 114:
			goto tr379
		case 115:
			goto tr380
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr271:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st203
	tr380:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:241
		lex.act = 21
		goto st203
	tr388:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:239
		lex.act = 19
		goto st203
	tr389:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:317
		lex.act = 97
		goto st203
	tr392:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:240
		lex.act = 20
		goto st203
	tr397:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:242
		lex.act = 22
		goto st203
	tr409:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:243
		lex.act = 23
		goto st203
	tr410:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:244
		lex.act = 24
		goto st203
	tr412:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:245
		lex.act = 25
		goto st203
	tr419:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:270
		lex.act = 50
		goto st203
	tr423:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:246
		lex.act = 26
		goto st203
	tr425:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:247
		lex.act = 27
		goto st203
	tr429:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:248
		lex.act = 28
		goto st203
	tr433:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:249
		lex.act = 29
		goto st203
	tr436:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:252
		lex.act = 32
		goto st203
	tr442:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:250
		lex.act = 30
		goto st203
	tr446:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:251
		lex.act = 31
		goto st203
	tr447:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:264
		lex.act = 44
		goto st203
	tr455:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:253
		lex.act = 33
		goto st203
	tr460:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:255
		lex.act = 35
		goto st203
	tr463:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:256
		lex.act = 36
		goto st203
	tr475:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:257
		lex.act = 37
		goto st203
	tr482:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:259
		lex.act = 39
		goto st203
	tr483:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:260
		lex.act = 40
		goto st203
	tr488:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:261
		lex.act = 41
		goto st203
	tr492:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:262
		lex.act = 42
		goto st203
	tr494:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:263
		lex.act = 43
		goto st203
	tr500:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:265
		lex.act = 45
		goto st203
	tr502:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:271
		lex.act = 51
		goto st203
	tr509:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:267
		lex.act = 47
		goto st203
	tr515:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:269
		lex.act = 49
		goto st203
	tr521:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:272
		lex.act = 52
		goto st203
	tr523:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:273
		lex.act = 53
		goto st203
	tr524:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:274
		lex.act = 54
		goto st203
	tr535:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:276
		lex.act = 56
		goto st203
	tr548:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:304
		lex.act = 84
		goto st203
	tr556:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:277
		lex.act = 57
		goto st203
	tr560:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:278
		lex.act = 58
		goto st203
	tr566:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:279
		lex.act = 59
		goto st203
	tr569:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:275
		lex.act = 55
		goto st203
	tr572:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:280
		lex.act = 60
		goto st203
	tr576:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:293
		lex.act = 73
		goto st203
	tr589:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:316
		lex.act = 96
		goto st203
	tr590:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:318
		lex.act = 98
		goto st203
	tr597:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:286
		lex.act = 66
		goto st203
	tr625:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:289
		lex.act = 69
		goto st203
	tr635:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:306
		lex.act = 86
		goto st203
	tr638:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:290
		lex.act = 70
		goto st203
	tr644:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:291
		lex.act = 71
		goto st203
	tr648:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:292
		lex.act = 72
		goto st203
	tr653:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:294
		lex.act = 74
		goto st203
	tr655:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:296
		lex.act = 76
		goto st203
	tr657:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:295
		lex.act = 75
		goto st203
	tr662:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:297
		lex.act = 77
		goto st203
	tr663:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:298
		lex.act = 78
		goto st203
	tr665:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:299
		lex.act = 79
		goto st203
	tr669:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:300
		lex.act = 80
		goto st203
	tr671:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:319
		lex.act = 99
		goto st203
	tr695:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:307
		lex.act = 87
		goto st203
	tr699:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:308
		lex.act = 88
		goto st203
	tr705:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:309
		lex.act = 89
		goto st203
	tr713:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:310
		lex.act = 90
		goto st203
	tr725:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:315
		lex.act = 95
		goto st203
	tr730:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:311
		lex.act = 91
		goto st203
	tr737:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:313
		lex.act = 93
		goto st203
	tr747:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:312
		lex.act = 92
		goto st203
	tr753:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:314
		lex.act = 94
		goto st203
	st203:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof203
		}
	st_case_203:
		// line internal/php8/scanner.go:7317
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 96:
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr10
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr10
				}
			case lex.data[(lex.p)] >= 91:
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr271
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		if lex.data[(lex.p)] == 96 {
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr10
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr130
	tr130:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:235
		lex.act = 16
		goto st204
	st204:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof204
		}
	st_case_204:
		// line internal/php8/scanner.go:7375
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 96:
			goto tr382
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr382
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr382
				}
			case lex.data[(lex.p)] >= 91:
				goto tr382
			}
		default:
			goto tr382
		}
		goto tr130
	tr377:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st205
	st205:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof205
		}
	st_case_205:
		// line internal/php8/scanner.go:7412
		switch lex.data[(lex.p)] {
		case 83:
			goto tr383
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr383
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr383:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st206
	st206:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof206
		}
	st_case_206:
		// line internal/php8/scanner.go:7453
		switch lex.data[(lex.p)] {
		case 84:
			goto tr384
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr384
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr384:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st207
	st207:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof207
		}
	st_case_207:
		// line internal/php8/scanner.go:7494
		switch lex.data[(lex.p)] {
		case 82:
			goto tr385
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr385
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr385:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st208
	st208:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof208
		}
	st_case_208:
		// line internal/php8/scanner.go:7535
		switch lex.data[(lex.p)] {
		case 65:
			goto tr386
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr386
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr386:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st209
	st209:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof209
		}
	st_case_209:
		// line internal/php8/scanner.go:7576
		switch lex.data[(lex.p)] {
		case 67:
			goto tr387
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr387
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr387:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st210
	st210:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof210
		}
	st_case_210:
		// line internal/php8/scanner.go:7617
		switch lex.data[(lex.p)] {
		case 84:
			goto tr388
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr388
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr378:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st211
	st211:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof211
		}
	st_case_211:
		// line internal/php8/scanner.go:7658
		switch lex.data[(lex.p)] {
		case 68:
			goto tr389
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr389
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr379:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st212
	st212:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof212
		}
	st_case_212:
		// line internal/php8/scanner.go:7699
		switch lex.data[(lex.p)] {
		case 82:
			goto tr390
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr390
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr390:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st213
	st213:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof213
		}
	st_case_213:
		// line internal/php8/scanner.go:7740
		switch lex.data[(lex.p)] {
		case 65:
			goto tr391
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr391
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr391:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st214
	st214:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof214
		}
	st_case_214:
		// line internal/php8/scanner.go:7781
		switch lex.data[(lex.p)] {
		case 89:
			goto tr392
		case 92:
			goto st90
		case 96:
			goto tr376
		case 121:
			goto tr392
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr265:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st215
	st215:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof215
		}
	st_case_215:
		// line internal/php8/scanner.go:7822
		switch lex.data[(lex.p)] {
		case 34:
			goto st5
		case 60:
			goto st91
		case 82:
			goto tr394
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr394
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	st91:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if lex.data[(lex.p)] == 60 {
			goto st92
		}
		goto tr131
	st92:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		if lex.data[(lex.p)] == 60 {
			goto st82
		}
		goto tr131
	tr394:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st216
	st216:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof216
		}
	st_case_216:
		// line internal/php8/scanner.go:7885
		switch lex.data[(lex.p)] {
		case 69:
			goto tr395
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr395
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr395:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st217
	st217:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof217
		}
	st_case_217:
		// line internal/php8/scanner.go:7926
		switch lex.data[(lex.p)] {
		case 65:
			goto tr396
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr396
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr396:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st218
	st218:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof218
		}
	st_case_218:
		// line internal/php8/scanner.go:7967
		switch lex.data[(lex.p)] {
		case 75:
			goto tr397
		case 92:
			goto st90
		case 96:
			goto tr376
		case 107:
			goto tr397
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr266:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st219
	st219:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof219
		}
	st_case_219:
		// line internal/php8/scanner.go:8008
		switch lex.data[(lex.p)] {
		case 65:
			goto tr398
		case 70:
			goto tr399
		case 76:
			goto tr400
		case 79:
			goto tr401
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr398
		case 102:
			goto tr399
		case 108:
			goto tr400
		case 111:
			goto tr401
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr398:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st220
	st220:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof220
		}
	st_case_220:
		// line internal/php8/scanner.go:8061
		switch lex.data[(lex.p)] {
		case 76:
			goto tr402
		case 83:
			goto tr403
		case 84:
			goto tr404
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr402
		case 115:
			goto tr403
		case 116:
			goto tr404
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr402:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st221
	st221:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof221
		}
	st_case_221:
		// line internal/php8/scanner.go:8110
		switch lex.data[(lex.p)] {
		case 76:
			goto tr405
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr405
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr405:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st222
	st222:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof222
		}
	st_case_222:
		// line internal/php8/scanner.go:8151
		switch lex.data[(lex.p)] {
		case 65:
			goto tr406
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr406
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr406:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st223
	st223:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof223
		}
	st_case_223:
		// line internal/php8/scanner.go:8192
		switch lex.data[(lex.p)] {
		case 66:
			goto tr407
		case 92:
			goto st90
		case 96:
			goto tr376
		case 98:
			goto tr407
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr407:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st224
	st224:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof224
		}
	st_case_224:
		// line internal/php8/scanner.go:8233
		switch lex.data[(lex.p)] {
		case 76:
			goto tr408
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr408
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr408:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st225
	st225:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof225
		}
	st_case_225:
		// line internal/php8/scanner.go:8274
		switch lex.data[(lex.p)] {
		case 69:
			goto tr409
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr409
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr403:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st226
	st226:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof226
		}
	st_case_226:
		// line internal/php8/scanner.go:8315
		switch lex.data[(lex.p)] {
		case 69:
			goto tr410
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr410
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr404:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st227
	st227:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof227
		}
	st_case_227:
		// line internal/php8/scanner.go:8356
		switch lex.data[(lex.p)] {
		case 67:
			goto tr411
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr411
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr411:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st228
	st228:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof228
		}
	st_case_228:
		// line internal/php8/scanner.go:8397
		switch lex.data[(lex.p)] {
		case 72:
			goto tr412
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr412
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr399:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st229
	st229:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof229
		}
	st_case_229:
		// line internal/php8/scanner.go:8438
		switch lex.data[(lex.p)] {
		case 85:
			goto tr413
		case 92:
			goto st90
		case 96:
			goto tr376
		case 117:
			goto tr413
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr413:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st230
	st230:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof230
		}
	st_case_230:
		// line internal/php8/scanner.go:8479
		switch lex.data[(lex.p)] {
		case 78:
			goto tr414
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr414
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr414:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st231
	st231:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof231
		}
	st_case_231:
		// line internal/php8/scanner.go:8520
		switch lex.data[(lex.p)] {
		case 67:
			goto tr415
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr415
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr415:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st232
	st232:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof232
		}
	st_case_232:
		// line internal/php8/scanner.go:8561
		switch lex.data[(lex.p)] {
		case 84:
			goto tr416
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr416
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr416:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st233
	st233:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof233
		}
	st_case_233:
		// line internal/php8/scanner.go:8602
		switch lex.data[(lex.p)] {
		case 73:
			goto tr417
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr417
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr417:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st234
	st234:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof234
		}
	st_case_234:
		// line internal/php8/scanner.go:8643
		switch lex.data[(lex.p)] {
		case 79:
			goto tr418
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr418
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr418:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st235
	st235:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof235
		}
	st_case_235:
		// line internal/php8/scanner.go:8684
		switch lex.data[(lex.p)] {
		case 78:
			goto tr419
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr419
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr400:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st236
	st236:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof236
		}
	st_case_236:
		// line internal/php8/scanner.go:8725
		switch lex.data[(lex.p)] {
		case 65:
			goto tr420
		case 79:
			goto tr421
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr420
		case 111:
			goto tr421
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr420:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st237
	st237:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof237
		}
	st_case_237:
		// line internal/php8/scanner.go:8770
		switch lex.data[(lex.p)] {
		case 83:
			goto tr422
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr422
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr422:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st238
	st238:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof238
		}
	st_case_238:
		// line internal/php8/scanner.go:8811
		switch lex.data[(lex.p)] {
		case 83:
			goto tr423
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr423
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr421:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st239
	st239:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof239
		}
	st_case_239:
		// line internal/php8/scanner.go:8852
		switch lex.data[(lex.p)] {
		case 78:
			goto tr424
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr424
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr424:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st240
	st240:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof240
		}
	st_case_240:
		// line internal/php8/scanner.go:8893
		switch lex.data[(lex.p)] {
		case 69:
			goto tr425
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr425
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr401:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st241
	st241:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof241
		}
	st_case_241:
		// line internal/php8/scanner.go:8934
		switch lex.data[(lex.p)] {
		case 78:
			goto tr426
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr426
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr426:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st242
	st242:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof242
		}
	st_case_242:
		// line internal/php8/scanner.go:8975
		switch lex.data[(lex.p)] {
		case 83:
			goto tr427
		case 84:
			goto tr428
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr427
		case 116:
			goto tr428
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr427:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st243
	st243:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof243
		}
	st_case_243:
		// line internal/php8/scanner.go:9020
		switch lex.data[(lex.p)] {
		case 84:
			goto tr429
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr429
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr428:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st244
	st244:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof244
		}
	st_case_244:
		// line internal/php8/scanner.go:9061
		switch lex.data[(lex.p)] {
		case 73:
			goto tr430
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr430
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr430:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st245
	st245:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof245
		}
	st_case_245:
		// line internal/php8/scanner.go:9102
		switch lex.data[(lex.p)] {
		case 78:
			goto tr431
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr431
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr431:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st246
	st246:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof246
		}
	st_case_246:
		// line internal/php8/scanner.go:9143
		switch lex.data[(lex.p)] {
		case 85:
			goto tr432
		case 92:
			goto st90
		case 96:
			goto tr376
		case 117:
			goto tr432
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr432:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st247
	st247:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof247
		}
	st_case_247:
		// line internal/php8/scanner.go:9184
		switch lex.data[(lex.p)] {
		case 69:
			goto tr433
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr433
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr267:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st248
	st248:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof248
		}
	st_case_248:
		// line internal/php8/scanner.go:9225
		switch lex.data[(lex.p)] {
		case 69:
			goto tr434
		case 73:
			goto tr435
		case 79:
			goto tr436
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr434
		case 105:
			goto tr435
		case 111:
			goto tr436
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr434:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st249
	st249:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof249
		}
	st_case_249:
		// line internal/php8/scanner.go:9274
		switch lex.data[(lex.p)] {
		case 67:
			goto tr437
		case 70:
			goto tr438
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr437
		case 102:
			goto tr438
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr437:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st250
	st250:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof250
		}
	st_case_250:
		// line internal/php8/scanner.go:9319
		switch lex.data[(lex.p)] {
		case 76:
			goto tr439
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr439
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr439:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st251
	st251:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof251
		}
	st_case_251:
		// line internal/php8/scanner.go:9360
		switch lex.data[(lex.p)] {
		case 65:
			goto tr440
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr440
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr440:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st252
	st252:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof252
		}
	st_case_252:
		// line internal/php8/scanner.go:9401
		switch lex.data[(lex.p)] {
		case 82:
			goto tr441
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr441
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr441:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st253
	st253:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof253
		}
	st_case_253:
		// line internal/php8/scanner.go:9442
		switch lex.data[(lex.p)] {
		case 69:
			goto tr442
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr442
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr438:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st254
	st254:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof254
		}
	st_case_254:
		// line internal/php8/scanner.go:9483
		switch lex.data[(lex.p)] {
		case 65:
			goto tr443
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr443
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr443:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st255
	st255:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof255
		}
	st_case_255:
		// line internal/php8/scanner.go:9524
		switch lex.data[(lex.p)] {
		case 85:
			goto tr444
		case 92:
			goto st90
		case 96:
			goto tr376
		case 117:
			goto tr444
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr444:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st256
	st256:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof256
		}
	st_case_256:
		// line internal/php8/scanner.go:9565
		switch lex.data[(lex.p)] {
		case 76:
			goto tr445
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr445
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr445:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st257
	st257:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof257
		}
	st_case_257:
		// line internal/php8/scanner.go:9606
		switch lex.data[(lex.p)] {
		case 84:
			goto tr446
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr446
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr435:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st258
	st258:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof258
		}
	st_case_258:
		// line internal/php8/scanner.go:9647
		switch lex.data[(lex.p)] {
		case 69:
			goto tr447
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr447
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr268:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st259
	st259:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof259
		}
	st_case_259:
		// line internal/php8/scanner.go:9688
		switch lex.data[(lex.p)] {
		case 67:
			goto tr448
		case 76:
			goto tr449
		case 77:
			goto tr450
		case 78:
			goto tr451
		case 86:
			goto tr452
		case 88:
			goto tr453
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr448
		case 108:
			goto tr449
		case 109:
			goto tr450
		case 110:
			goto tr451
		case 118:
			goto tr452
		case 120:
			goto tr453
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr448:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st260
	st260:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof260
		}
	st_case_260:
		// line internal/php8/scanner.go:9749
		switch lex.data[(lex.p)] {
		case 72:
			goto tr454
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr454
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr454:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st261
	st261:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof261
		}
	st_case_261:
		// line internal/php8/scanner.go:9790
		switch lex.data[(lex.p)] {
		case 79:
			goto tr455
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr455
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr449:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st262
	st262:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof262
		}
	st_case_262:
		// line internal/php8/scanner.go:9831
		switch lex.data[(lex.p)] {
		case 83:
			goto tr456
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr456
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr456:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st263
	st263:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof263
		}
	st_case_263:
		// line internal/php8/scanner.go:9872
		switch lex.data[(lex.p)] {
		case 69:
			goto tr457
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr457
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr457:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:254
		lex.act = 34
		goto st264
	st264:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof264
		}
	st_case_264:
		// line internal/php8/scanner.go:9913
		switch lex.data[(lex.p)] {
		case 73:
			goto tr459
		case 92:
			goto st90
		case 96:
			goto tr458
		case 105:
			goto tr459
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr458
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr458
				}
			case lex.data[(lex.p)] >= 91:
				goto tr458
			}
		default:
			goto tr458
		}
		goto tr271
	tr459:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st265
	st265:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof265
		}
	st_case_265:
		// line internal/php8/scanner.go:9954
		switch lex.data[(lex.p)] {
		case 70:
			goto tr460
		case 92:
			goto st90
		case 96:
			goto tr376
		case 102:
			goto tr460
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr450:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st266
	st266:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof266
		}
	st_case_266:
		// line internal/php8/scanner.go:9995
		switch lex.data[(lex.p)] {
		case 80:
			goto tr461
		case 92:
			goto st90
		case 96:
			goto tr376
		case 112:
			goto tr461
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr461:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st267
	st267:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof267
		}
	st_case_267:
		// line internal/php8/scanner.go:10036
		switch lex.data[(lex.p)] {
		case 84:
			goto tr462
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr462
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr462:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st268
	st268:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof268
		}
	st_case_268:
		// line internal/php8/scanner.go:10077
		switch lex.data[(lex.p)] {
		case 89:
			goto tr463
		case 92:
			goto st90
		case 96:
			goto tr376
		case 121:
			goto tr463
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr451:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st269
	st269:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof269
		}
	st_case_269:
		// line internal/php8/scanner.go:10118
		switch lex.data[(lex.p)] {
		case 68:
			goto tr464
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr464
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr464:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st270
	st270:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof270
		}
	st_case_270:
		// line internal/php8/scanner.go:10159
		switch lex.data[(lex.p)] {
		case 68:
			goto tr465
		case 70:
			goto tr466
		case 73:
			goto tr467
		case 83:
			goto tr468
		case 87:
			goto tr469
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr465
		case 102:
			goto tr466
		case 105:
			goto tr467
		case 115:
			goto tr468
		case 119:
			goto tr469
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr465:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st271
	st271:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof271
		}
	st_case_271:
		// line internal/php8/scanner.go:10216
		switch lex.data[(lex.p)] {
		case 69:
			goto tr470
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr470
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr470:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st272
	st272:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof272
		}
	st_case_272:
		// line internal/php8/scanner.go:10257
		switch lex.data[(lex.p)] {
		case 67:
			goto tr471
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr471
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr471:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st273
	st273:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof273
		}
	st_case_273:
		// line internal/php8/scanner.go:10298
		switch lex.data[(lex.p)] {
		case 76:
			goto tr472
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr472
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr472:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st274
	st274:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof274
		}
	st_case_274:
		// line internal/php8/scanner.go:10339
		switch lex.data[(lex.p)] {
		case 65:
			goto tr473
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr473
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr473:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st275
	st275:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof275
		}
	st_case_275:
		// line internal/php8/scanner.go:10380
		switch lex.data[(lex.p)] {
		case 82:
			goto tr474
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr474
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr474:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st276
	st276:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof276
		}
	st_case_276:
		// line internal/php8/scanner.go:10421
		switch lex.data[(lex.p)] {
		case 69:
			goto tr475
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr475
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr466:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st277
	st277:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof277
		}
	st_case_277:
		// line internal/php8/scanner.go:10462
		switch lex.data[(lex.p)] {
		case 79:
			goto tr476
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr476
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr476:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st278
	st278:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof278
		}
	st_case_278:
		// line internal/php8/scanner.go:10503
		switch lex.data[(lex.p)] {
		case 82:
			goto tr477
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr477
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr477:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:258
		lex.act = 38
		goto st279
	st279:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof279
		}
	st_case_279:
		// line internal/php8/scanner.go:10544
		switch lex.data[(lex.p)] {
		case 69:
			goto tr479
		case 92:
			goto st90
		case 96:
			goto tr478
		case 101:
			goto tr479
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr478
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr478
				}
			case lex.data[(lex.p)] >= 91:
				goto tr478
			}
		default:
			goto tr478
		}
		goto tr271
	tr479:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st280
	st280:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof280
		}
	st_case_280:
		// line internal/php8/scanner.go:10585
		switch lex.data[(lex.p)] {
		case 65:
			goto tr480
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr480
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr480:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st281
	st281:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof281
		}
	st_case_281:
		// line internal/php8/scanner.go:10626
		switch lex.data[(lex.p)] {
		case 67:
			goto tr481
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr481
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr481:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st282
	st282:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof282
		}
	st_case_282:
		// line internal/php8/scanner.go:10667
		switch lex.data[(lex.p)] {
		case 72:
			goto tr482
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr482
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr467:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st283
	st283:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof283
		}
	st_case_283:
		// line internal/php8/scanner.go:10708
		switch lex.data[(lex.p)] {
		case 70:
			goto tr483
		case 92:
			goto st90
		case 96:
			goto tr376
		case 102:
			goto tr483
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr468:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st284
	st284:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof284
		}
	st_case_284:
		// line internal/php8/scanner.go:10749
		switch lex.data[(lex.p)] {
		case 87:
			goto tr484
		case 92:
			goto st90
		case 96:
			goto tr376
		case 119:
			goto tr484
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr484:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st285
	st285:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof285
		}
	st_case_285:
		// line internal/php8/scanner.go:10790
		switch lex.data[(lex.p)] {
		case 73:
			goto tr485
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr485
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr485:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st286
	st286:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof286
		}
	st_case_286:
		// line internal/php8/scanner.go:10831
		switch lex.data[(lex.p)] {
		case 84:
			goto tr486
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr486
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr486:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st287
	st287:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof287
		}
	st_case_287:
		// line internal/php8/scanner.go:10872
		switch lex.data[(lex.p)] {
		case 67:
			goto tr487
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr487
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr487:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st288
	st288:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof288
		}
	st_case_288:
		// line internal/php8/scanner.go:10913
		switch lex.data[(lex.p)] {
		case 72:
			goto tr488
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr488
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr469:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st289
	st289:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof289
		}
	st_case_289:
		// line internal/php8/scanner.go:10954
		switch lex.data[(lex.p)] {
		case 72:
			goto tr489
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr489
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr489:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st290
	st290:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof290
		}
	st_case_290:
		// line internal/php8/scanner.go:10995
		switch lex.data[(lex.p)] {
		case 73:
			goto tr490
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr490
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr490:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st291
	st291:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof291
		}
	st_case_291:
		// line internal/php8/scanner.go:11036
		switch lex.data[(lex.p)] {
		case 76:
			goto tr491
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr491
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr491:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st292
	st292:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof292
		}
	st_case_292:
		// line internal/php8/scanner.go:11077
		switch lex.data[(lex.p)] {
		case 69:
			goto tr492
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr492
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr452:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st293
	st293:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof293
		}
	st_case_293:
		// line internal/php8/scanner.go:11118
		switch lex.data[(lex.p)] {
		case 65:
			goto tr493
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr493
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr493:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st294
	st294:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof294
		}
	st_case_294:
		// line internal/php8/scanner.go:11159
		switch lex.data[(lex.p)] {
		case 76:
			goto tr494
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr494
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr453:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st295
	st295:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof295
		}
	st_case_295:
		// line internal/php8/scanner.go:11200
		switch lex.data[(lex.p)] {
		case 73:
			goto tr495
		case 84:
			goto tr496
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr495
		case 116:
			goto tr496
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr495:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st296
	st296:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof296
		}
	st_case_296:
		// line internal/php8/scanner.go:11245
		switch lex.data[(lex.p)] {
		case 84:
			goto tr447
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr447
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr496:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st297
	st297:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof297
		}
	st_case_297:
		// line internal/php8/scanner.go:11286
		switch lex.data[(lex.p)] {
		case 69:
			goto tr497
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr497
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr497:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st298
	st298:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof298
		}
	st_case_298:
		// line internal/php8/scanner.go:11327
		switch lex.data[(lex.p)] {
		case 78:
			goto tr498
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr498
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr498:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st299
	st299:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof299
		}
	st_case_299:
		// line internal/php8/scanner.go:11368
		switch lex.data[(lex.p)] {
		case 68:
			goto tr499
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr499
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr499:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st300
	st300:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof300
		}
	st_case_300:
		// line internal/php8/scanner.go:11409
		switch lex.data[(lex.p)] {
		case 83:
			goto tr500
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr500
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr269:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st301
	st301:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof301
		}
	st_case_301:
		// line internal/php8/scanner.go:11450
		switch lex.data[(lex.p)] {
		case 73:
			goto tr501
		case 78:
			goto tr502
		case 79:
			goto tr503
		case 85:
			goto tr413
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr501
		case 110:
			goto tr502
		case 111:
			goto tr503
		case 117:
			goto tr413
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr501:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st302
	st302:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof302
		}
	st_case_302:
		// line internal/php8/scanner.go:11503
		switch lex.data[(lex.p)] {
		case 78:
			goto tr504
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr504
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr504:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st303
	st303:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof303
		}
	st_case_303:
		// line internal/php8/scanner.go:11544
		switch lex.data[(lex.p)] {
		case 65:
			goto tr505
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr505
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr505:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st304
	st304:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof304
		}
	st_case_304:
		// line internal/php8/scanner.go:11585
		switch lex.data[(lex.p)] {
		case 76:
			goto tr506
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr506
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr506:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:266
		lex.act = 46
		goto st305
	st305:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof305
		}
	st_case_305:
		// line internal/php8/scanner.go:11626
		switch lex.data[(lex.p)] {
		case 76:
			goto tr508
		case 92:
			goto st90
		case 96:
			goto tr507
		case 108:
			goto tr508
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr507
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr507
				}
			case lex.data[(lex.p)] >= 91:
				goto tr507
			}
		default:
			goto tr507
		}
		goto tr271
	tr508:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st306
	st306:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof306
		}
	st_case_306:
		// line internal/php8/scanner.go:11667
		switch lex.data[(lex.p)] {
		case 89:
			goto tr509
		case 92:
			goto st90
		case 96:
			goto tr376
		case 121:
			goto tr509
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr503:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st307
	st307:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof307
		}
	st_case_307:
		// line internal/php8/scanner.go:11708
		switch lex.data[(lex.p)] {
		case 82:
			goto tr510
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr510
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr510:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:268
		lex.act = 48
		goto st308
	st308:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof308
		}
	st_case_308:
		// line internal/php8/scanner.go:11749
		switch lex.data[(lex.p)] {
		case 69:
			goto tr512
		case 92:
			goto st90
		case 96:
			goto tr511
		case 101:
			goto tr512
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr511
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr511
				}
			case lex.data[(lex.p)] >= 91:
				goto tr511
			}
		default:
			goto tr511
		}
		goto tr271
	tr512:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st309
	st309:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof309
		}
	st_case_309:
		// line internal/php8/scanner.go:11790
		switch lex.data[(lex.p)] {
		case 65:
			goto tr513
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr513
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr513:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st310
	st310:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof310
		}
	st_case_310:
		// line internal/php8/scanner.go:11831
		switch lex.data[(lex.p)] {
		case 67:
			goto tr514
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr514
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr514:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st311
	st311:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof311
		}
	st_case_311:
		// line internal/php8/scanner.go:11872
		switch lex.data[(lex.p)] {
		case 72:
			goto tr515
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr515
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr270:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st312
	st312:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof312
		}
	st_case_312:
		// line internal/php8/scanner.go:11913
		switch lex.data[(lex.p)] {
		case 76:
			goto tr516
		case 79:
			goto tr517
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr516
		case 111:
			goto tr517
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr516:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st313
	st313:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof313
		}
	st_case_313:
		// line internal/php8/scanner.go:11958
		switch lex.data[(lex.p)] {
		case 79:
			goto tr518
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr518
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr518:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st314
	st314:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof314
		}
	st_case_314:
		// line internal/php8/scanner.go:11999
		switch lex.data[(lex.p)] {
		case 66:
			goto tr519
		case 92:
			goto st90
		case 96:
			goto tr376
		case 98:
			goto tr519
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr519:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st315
	st315:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof315
		}
	st_case_315:
		// line internal/php8/scanner.go:12040
		switch lex.data[(lex.p)] {
		case 65:
			goto tr520
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr520
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr520:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st316
	st316:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof316
		}
	st_case_316:
		// line internal/php8/scanner.go:12081
		switch lex.data[(lex.p)] {
		case 76:
			goto tr521
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr521
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr517:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st317
	st317:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof317
		}
	st_case_317:
		// line internal/php8/scanner.go:12122
		switch lex.data[(lex.p)] {
		case 84:
			goto tr522
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr522
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr522:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st318
	st318:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof318
		}
	st_case_318:
		// line internal/php8/scanner.go:12163
		switch lex.data[(lex.p)] {
		case 79:
			goto tr523
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr523
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr272:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st319
	st319:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof319
		}
	st_case_319:
		// line internal/php8/scanner.go:12204
		switch lex.data[(lex.p)] {
		case 70:
			goto tr524
		case 77:
			goto tr525
		case 78:
			goto tr526
		case 83:
			goto tr527
		case 92:
			goto st90
		case 96:
			goto tr376
		case 102:
			goto tr524
		case 109:
			goto tr525
		case 110:
			goto tr526
		case 115:
			goto tr527
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr525:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st320
	st320:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof320
		}
	st_case_320:
		// line internal/php8/scanner.go:12257
		switch lex.data[(lex.p)] {
		case 80:
			goto tr528
		case 92:
			goto st90
		case 96:
			goto tr376
		case 112:
			goto tr528
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr528:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st321
	st321:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof321
		}
	st_case_321:
		// line internal/php8/scanner.go:12298
		switch lex.data[(lex.p)] {
		case 76:
			goto tr529
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr529
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr529:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st322
	st322:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof322
		}
	st_case_322:
		// line internal/php8/scanner.go:12339
		switch lex.data[(lex.p)] {
		case 69:
			goto tr530
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr530
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr530:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st323
	st323:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof323
		}
	st_case_323:
		// line internal/php8/scanner.go:12380
		switch lex.data[(lex.p)] {
		case 77:
			goto tr531
		case 92:
			goto st90
		case 96:
			goto tr376
		case 109:
			goto tr531
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr531:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st324
	st324:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof324
		}
	st_case_324:
		// line internal/php8/scanner.go:12421
		switch lex.data[(lex.p)] {
		case 69:
			goto tr532
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr532
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr532:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st325
	st325:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof325
		}
	st_case_325:
		// line internal/php8/scanner.go:12462
		switch lex.data[(lex.p)] {
		case 78:
			goto tr533
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr533
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr533:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st326
	st326:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof326
		}
	st_case_326:
		// line internal/php8/scanner.go:12503
		switch lex.data[(lex.p)] {
		case 84:
			goto tr534
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr534
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr534:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st327
	st327:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof327
		}
	st_case_327:
		// line internal/php8/scanner.go:12544
		switch lex.data[(lex.p)] {
		case 83:
			goto tr535
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr535
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr526:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st328
	st328:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof328
		}
	st_case_328:
		// line internal/php8/scanner.go:12585
		switch lex.data[(lex.p)] {
		case 67:
			goto tr536
		case 83:
			goto tr537
		case 84:
			goto tr538
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr536
		case 115:
			goto tr537
		case 116:
			goto tr538
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr536:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st329
	st329:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof329
		}
	st_case_329:
		// line internal/php8/scanner.go:12634
		switch lex.data[(lex.p)] {
		case 76:
			goto tr539
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr539
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr539:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st330
	st330:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof330
		}
	st_case_330:
		// line internal/php8/scanner.go:12675
		switch lex.data[(lex.p)] {
		case 85:
			goto tr540
		case 92:
			goto st90
		case 96:
			goto tr376
		case 117:
			goto tr540
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr540:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st331
	st331:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof331
		}
	st_case_331:
		// line internal/php8/scanner.go:12716
		switch lex.data[(lex.p)] {
		case 68:
			goto tr541
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr541
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr541:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st332
	st332:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof332
		}
	st_case_332:
		// line internal/php8/scanner.go:12757
		switch lex.data[(lex.p)] {
		case 69:
			goto tr542
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr542
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr542:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:303
		lex.act = 83
		goto st333
	st333:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof333
		}
	st_case_333:
		// line internal/php8/scanner.go:12798
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr544
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr543
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr543
				}
			case lex.data[(lex.p)] >= 91:
				goto tr543
			}
		default:
			goto tr543
		}
		goto tr271
	tr544:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st334
	st334:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof334
		}
	st_case_334:
		// line internal/php8/scanner.go:12835
		switch lex.data[(lex.p)] {
		case 79:
			goto tr545
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr545
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr545:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st335
	st335:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof335
		}
	st_case_335:
		// line internal/php8/scanner.go:12876
		switch lex.data[(lex.p)] {
		case 78:
			goto tr546
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr546
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr546:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st336
	st336:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof336
		}
	st_case_336:
		// line internal/php8/scanner.go:12917
		switch lex.data[(lex.p)] {
		case 67:
			goto tr547
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr547
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr547:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st337
	st337:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof337
		}
	st_case_337:
		// line internal/php8/scanner.go:12958
		switch lex.data[(lex.p)] {
		case 69:
			goto tr548
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr548
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr537:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st338
	st338:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof338
		}
	st_case_338:
		// line internal/php8/scanner.go:12999
		switch lex.data[(lex.p)] {
		case 84:
			goto tr549
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr549
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr549:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st339
	st339:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof339
		}
	st_case_339:
		// line internal/php8/scanner.go:13040
		switch lex.data[(lex.p)] {
		case 65:
			goto tr550
		case 69:
			goto tr551
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr550
		case 101:
			goto tr551
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr550:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st340
	st340:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof340
		}
	st_case_340:
		// line internal/php8/scanner.go:13085
		switch lex.data[(lex.p)] {
		case 78:
			goto tr552
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr552
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr552:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st341
	st341:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof341
		}
	st_case_341:
		// line internal/php8/scanner.go:13126
		switch lex.data[(lex.p)] {
		case 67:
			goto tr553
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr553
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr553:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st342
	st342:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof342
		}
	st_case_342:
		// line internal/php8/scanner.go:13167
		switch lex.data[(lex.p)] {
		case 69:
			goto tr554
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr554
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr554:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st343
	st343:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof343
		}
	st_case_343:
		// line internal/php8/scanner.go:13208
		switch lex.data[(lex.p)] {
		case 79:
			goto tr555
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr555
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr555:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st344
	st344:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof344
		}
	st_case_344:
		// line internal/php8/scanner.go:13249
		switch lex.data[(lex.p)] {
		case 70:
			goto tr556
		case 92:
			goto st90
		case 96:
			goto tr376
		case 102:
			goto tr556
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr551:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st345
	st345:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof345
		}
	st_case_345:
		// line internal/php8/scanner.go:13290
		switch lex.data[(lex.p)] {
		case 65:
			goto tr557
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr557
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr557:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st346
	st346:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof346
		}
	st_case_346:
		// line internal/php8/scanner.go:13331
		switch lex.data[(lex.p)] {
		case 68:
			goto tr558
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr558
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr558:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st347
	st347:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof347
		}
	st_case_347:
		// line internal/php8/scanner.go:13372
		switch lex.data[(lex.p)] {
		case 79:
			goto tr559
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr559
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr559:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st348
	st348:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof348
		}
	st_case_348:
		// line internal/php8/scanner.go:13413
		switch lex.data[(lex.p)] {
		case 70:
			goto tr560
		case 92:
			goto st90
		case 96:
			goto tr376
		case 102:
			goto tr560
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr538:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st349
	st349:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof349
		}
	st_case_349:
		// line internal/php8/scanner.go:13454
		switch lex.data[(lex.p)] {
		case 69:
			goto tr561
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr561
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr561:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st350
	st350:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof350
		}
	st_case_350:
		// line internal/php8/scanner.go:13495
		switch lex.data[(lex.p)] {
		case 82:
			goto tr562
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr562
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr562:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st351
	st351:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof351
		}
	st_case_351:
		// line internal/php8/scanner.go:13536
		switch lex.data[(lex.p)] {
		case 70:
			goto tr563
		case 92:
			goto st90
		case 96:
			goto tr376
		case 102:
			goto tr563
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr563:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st352
	st352:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof352
		}
	st_case_352:
		// line internal/php8/scanner.go:13577
		switch lex.data[(lex.p)] {
		case 65:
			goto tr564
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr564
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr564:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st353
	st353:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof353
		}
	st_case_353:
		// line internal/php8/scanner.go:13618
		switch lex.data[(lex.p)] {
		case 67:
			goto tr565
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr565
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr565:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st354
	st354:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof354
		}
	st_case_354:
		// line internal/php8/scanner.go:13659
		switch lex.data[(lex.p)] {
		case 69:
			goto tr566
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr566
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr527:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st355
	st355:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof355
		}
	st_case_355:
		// line internal/php8/scanner.go:13700
		switch lex.data[(lex.p)] {
		case 83:
			goto tr567
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr567
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr567:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st356
	st356:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof356
		}
	st_case_356:
		// line internal/php8/scanner.go:13741
		switch lex.data[(lex.p)] {
		case 69:
			goto tr568
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr568
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr568:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st357
	st357:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof357
		}
	st_case_357:
		// line internal/php8/scanner.go:13782
		switch lex.data[(lex.p)] {
		case 84:
			goto tr569
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr569
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr273:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st358
	st358:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof358
		}
	st_case_358:
		// line internal/php8/scanner.go:13823
		switch lex.data[(lex.p)] {
		case 73:
			goto tr570
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr570
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr570:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st359
	st359:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof359
		}
	st_case_359:
		// line internal/php8/scanner.go:13864
		switch lex.data[(lex.p)] {
		case 83:
			goto tr571
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr571
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr571:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st360
	st360:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof360
		}
	st_case_360:
		// line internal/php8/scanner.go:13905
		switch lex.data[(lex.p)] {
		case 84:
			goto tr572
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr572
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr274:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st361
	st361:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof361
		}
	st_case_361:
		// line internal/php8/scanner.go:13946
		switch lex.data[(lex.p)] {
		case 65:
			goto tr573
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr573
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr573:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st362
	st362:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof362
		}
	st_case_362:
		// line internal/php8/scanner.go:13987
		switch lex.data[(lex.p)] {
		case 84:
			goto tr574
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr574
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr574:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st363
	st363:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof363
		}
	st_case_363:
		// line internal/php8/scanner.go:14028
		switch lex.data[(lex.p)] {
		case 67:
			goto tr575
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr575
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr575:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st364
	st364:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof364
		}
	st_case_364:
		// line internal/php8/scanner.go:14069
		switch lex.data[(lex.p)] {
		case 72:
			goto tr576
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr576
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr275:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st365
	st365:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof365
		}
	st_case_365:
		// line internal/php8/scanner.go:14110
		switch lex.data[(lex.p)] {
		case 65:
			goto tr577
		case 69:
			goto tr578
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr577
		case 101:
			goto tr578
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr577:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st366
	st366:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof366
		}
	st_case_366:
		// line internal/php8/scanner.go:14155
		switch lex.data[(lex.p)] {
		case 77:
			goto tr579
		case 92:
			goto st90
		case 96:
			goto tr376
		case 109:
			goto tr579
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr579:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st367
	st367:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof367
		}
	st_case_367:
		// line internal/php8/scanner.go:14196
		switch lex.data[(lex.p)] {
		case 69:
			goto tr580
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr580
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr580:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st368
	st368:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof368
		}
	st_case_368:
		// line internal/php8/scanner.go:14237
		switch lex.data[(lex.p)] {
		case 83:
			goto tr581
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr581
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr581:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st369
	st369:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof369
		}
	st_case_369:
		// line internal/php8/scanner.go:14278
		switch lex.data[(lex.p)] {
		case 80:
			goto tr582
		case 92:
			goto st90
		case 96:
			goto tr376
		case 112:
			goto tr582
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr582:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st370
	st370:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof370
		}
	st_case_370:
		// line internal/php8/scanner.go:14319
		switch lex.data[(lex.p)] {
		case 65:
			goto tr583
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr583
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr583:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st371
	st371:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof371
		}
	st_case_371:
		// line internal/php8/scanner.go:14360
		switch lex.data[(lex.p)] {
		case 67:
			goto tr584
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr584
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr584:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st372
	st372:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof372
		}
	st_case_372:
		// line internal/php8/scanner.go:14401
		switch lex.data[(lex.p)] {
		case 69:
			goto tr585
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr585
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr585:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:281
		lex.act = 61
		goto st373
	st373:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof373
		}
	st_case_373:
		// line internal/php8/scanner.go:14442
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 96:
			goto tr586
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr586
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr586
				}
			case lex.data[(lex.p)] >= 91:
				goto tr586
			}
		default:
			goto tr586
		}
		goto tr271
	st93:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		if lex.data[(lex.p)] == 96 {
			goto tr10
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr10
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr133
	tr133:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:234
		lex.act = 15
		goto st374
	st374:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof374
		}
	st_case_374:
		// line internal/php8/scanner.go:14500
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 96:
			goto tr588
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr588
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr588
				}
			case lex.data[(lex.p)] >= 91:
				goto tr588
			}
		default:
			goto tr588
		}
		goto tr133
	tr578:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st375
	st375:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof375
		}
	st_case_375:
		// line internal/php8/scanner.go:14537
		switch lex.data[(lex.p)] {
		case 87:
			goto tr589
		case 92:
			goto st90
		case 96:
			goto tr376
		case 119:
			goto tr589
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr276:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st376
	st376:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof376
		}
	st_case_376:
		// line internal/php8/scanner.go:14578
		switch lex.data[(lex.p)] {
		case 82:
			goto tr590
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr590
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr277:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st377
	st377:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof377
		}
	st_case_377:
		// line internal/php8/scanner.go:14619
		switch lex.data[(lex.p)] {
		case 82:
			goto tr591
		case 85:
			goto tr592
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr591
		case 117:
			goto tr592
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr591:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st378
	st378:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof378
		}
	st_case_378:
		// line internal/php8/scanner.go:14664
		switch lex.data[(lex.p)] {
		case 73:
			goto tr593
		case 79:
			goto tr594
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr593
		case 111:
			goto tr594
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr593:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st379
	st379:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof379
		}
	st_case_379:
		// line internal/php8/scanner.go:14709
		switch lex.data[(lex.p)] {
		case 78:
			goto tr595
		case 86:
			goto tr596
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr595
		case 118:
			goto tr596
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr595:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st380
	st380:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof380
		}
	st_case_380:
		// line internal/php8/scanner.go:14754
		switch lex.data[(lex.p)] {
		case 84:
			goto tr597
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr597
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr596:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st381
	st381:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof381
		}
	st_case_381:
		// line internal/php8/scanner.go:14795
		switch lex.data[(lex.p)] {
		case 65:
			goto tr598
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr598
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr598:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st382
	st382:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof382
		}
	st_case_382:
		// line internal/php8/scanner.go:14836
		switch lex.data[(lex.p)] {
		case 84:
			goto tr599
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr599
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr599:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st383
	st383:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof383
		}
	st_case_383:
		// line internal/php8/scanner.go:14877
		switch lex.data[(lex.p)] {
		case 69:
			goto tr600
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr600
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr600:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:282
		lex.act = 62
		goto st384
	st384:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof384
		}
	st_case_384:
		// line internal/php8/scanner.go:14918
		switch lex.data[(lex.p)] {
		case 40:
			goto st94
		case 92:
			goto st90
		case 96:
			goto tr601
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr601
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr601
				}
			case lex.data[(lex.p)] >= 91:
				goto tr601
			}
		default:
			goto tr601
		}
		goto tr271
	st94:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch lex.data[(lex.p)] {
		case 83:
			goto st95
		case 115:
			goto st95
		}
		goto tr134
	st95:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch lex.data[(lex.p)] {
		case 69:
			goto st96
		case 101:
			goto st96
		}
		goto tr134
	st96:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch lex.data[(lex.p)] {
		case 84:
			goto st97
		case 116:
			goto st97
		}
		goto tr134
	st97:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		if lex.data[(lex.p)] == 41 {
			goto tr138
		}
		goto tr134
	tr594:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st385
	st385:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof385
		}
	st_case_385:
		// line internal/php8/scanner.go:15002
		switch lex.data[(lex.p)] {
		case 84:
			goto tr603
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr603
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr603:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st386
	st386:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof386
		}
	st_case_386:
		// line internal/php8/scanner.go:15043
		switch lex.data[(lex.p)] {
		case 69:
			goto tr604
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr604
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr604:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st387
	st387:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof387
		}
	st_case_387:
		// line internal/php8/scanner.go:15084
		switch lex.data[(lex.p)] {
		case 67:
			goto tr605
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr605
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr605:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st388
	st388:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof388
		}
	st_case_388:
		// line internal/php8/scanner.go:15125
		switch lex.data[(lex.p)] {
		case 84:
			goto tr606
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr606
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr606:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st389
	st389:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof389
		}
	st_case_389:
		// line internal/php8/scanner.go:15166
		switch lex.data[(lex.p)] {
		case 69:
			goto tr607
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr607
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr607:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st390
	st390:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof390
		}
	st_case_390:
		// line internal/php8/scanner.go:15207
		switch lex.data[(lex.p)] {
		case 68:
			goto tr608
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr608
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr608:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:287
		lex.act = 67
		goto st391
	st391:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof391
		}
	st_case_391:
		// line internal/php8/scanner.go:15248
		switch lex.data[(lex.p)] {
		case 40:
			goto st98
		case 92:
			goto st90
		case 96:
			goto tr609
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr609
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr609
				}
			case lex.data[(lex.p)] >= 91:
				goto tr609
			}
		default:
			goto tr609
		}
		goto tr271
	st98:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch lex.data[(lex.p)] {
		case 83:
			goto st99
		case 115:
			goto st99
		}
		goto tr139
	st99:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch lex.data[(lex.p)] {
		case 69:
			goto st100
		case 101:
			goto st100
		}
		goto tr139
	st100:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch lex.data[(lex.p)] {
		case 84:
			goto st101
		case 116:
			goto st101
		}
		goto tr139
	st101:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if lex.data[(lex.p)] == 41 {
			goto tr143
		}
		goto tr139
	tr592:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st392
	st392:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof392
		}
	st_case_392:
		// line internal/php8/scanner.go:15332
		switch lex.data[(lex.p)] {
		case 66:
			goto tr611
		case 92:
			goto st90
		case 96:
			goto tr376
		case 98:
			goto tr611
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr611:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st393
	st393:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof393
		}
	st_case_393:
		// line internal/php8/scanner.go:15373
		switch lex.data[(lex.p)] {
		case 76:
			goto tr612
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr612
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr612:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st394
	st394:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof394
		}
	st_case_394:
		// line internal/php8/scanner.go:15414
		switch lex.data[(lex.p)] {
		case 73:
			goto tr613
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr613
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr613:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st395
	st395:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof395
		}
	st_case_395:
		// line internal/php8/scanner.go:15455
		switch lex.data[(lex.p)] {
		case 67:
			goto tr614
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr614
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr614:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:284
		lex.act = 64
		goto st396
	st396:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof396
		}
	st_case_396:
		// line internal/php8/scanner.go:15496
		switch lex.data[(lex.p)] {
		case 40:
			goto st102
		case 92:
			goto st90
		case 96:
			goto tr615
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr615
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr615
				}
			case lex.data[(lex.p)] >= 91:
				goto tr615
			}
		default:
			goto tr615
		}
		goto tr271
	st102:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch lex.data[(lex.p)] {
		case 83:
			goto st103
		case 115:
			goto st103
		}
		goto tr144
	st103:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch lex.data[(lex.p)] {
		case 69:
			goto st104
		case 101:
			goto st104
		}
		goto tr144
	st104:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch lex.data[(lex.p)] {
		case 84:
			goto st105
		case 116:
			goto st105
		}
		goto tr144
	st105:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		if lex.data[(lex.p)] == 41 {
			goto tr148
		}
		goto tr144
	tr278:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st397
	st397:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof397
		}
	st_case_397:
		// line internal/php8/scanner.go:15580
		switch lex.data[(lex.p)] {
		case 69:
			goto tr617
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr617
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr617:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st398
	st398:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof398
		}
	st_case_398:
		// line internal/php8/scanner.go:15621
		switch lex.data[(lex.p)] {
		case 65:
			goto tr618
		case 81:
			goto tr619
		case 84:
			goto tr620
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr618
		case 113:
			goto tr619
		case 116:
			goto tr620
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr618:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st399
	st399:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof399
		}
	st_case_399:
		// line internal/php8/scanner.go:15670
		switch lex.data[(lex.p)] {
		case 68:
			goto tr621
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr621
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr621:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st400
	st400:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof400
		}
	st_case_400:
		// line internal/php8/scanner.go:15711
		switch lex.data[(lex.p)] {
		case 79:
			goto tr622
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr622
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr622:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st401
	st401:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof401
		}
	st_case_401:
		// line internal/php8/scanner.go:15752
		switch lex.data[(lex.p)] {
		case 78:
			goto tr623
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr623
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr623:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st402
	st402:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof402
		}
	st_case_402:
		// line internal/php8/scanner.go:15793
		switch lex.data[(lex.p)] {
		case 76:
			goto tr624
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr624
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr624:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st403
	st403:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof403
		}
	st_case_403:
		// line internal/php8/scanner.go:15834
		switch lex.data[(lex.p)] {
		case 89:
			goto tr625
		case 92:
			goto st90
		case 96:
			goto tr376
		case 121:
			goto tr625
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr619:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st404
	st404:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof404
		}
	st_case_404:
		// line internal/php8/scanner.go:15875
		switch lex.data[(lex.p)] {
		case 85:
			goto tr626
		case 92:
			goto st90
		case 96:
			goto tr376
		case 117:
			goto tr626
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr626:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st405
	st405:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof405
		}
	st_case_405:
		// line internal/php8/scanner.go:15916
		switch lex.data[(lex.p)] {
		case 73:
			goto tr627
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr627
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr627:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st406
	st406:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof406
		}
	st_case_406:
		// line internal/php8/scanner.go:15957
		switch lex.data[(lex.p)] {
		case 82:
			goto tr628
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr628
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr628:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st407
	st407:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof407
		}
	st_case_407:
		// line internal/php8/scanner.go:15998
		switch lex.data[(lex.p)] {
		case 69:
			goto tr629
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr629
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr629:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:305
		lex.act = 85
		goto st408
	st408:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof408
		}
	st_case_408:
		// line internal/php8/scanner.go:16039
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr631
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr630
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr630
				}
			case lex.data[(lex.p)] >= 91:
				goto tr630
			}
		default:
			goto tr630
		}
		goto tr271
	tr631:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st409
	st409:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof409
		}
	st_case_409:
		// line internal/php8/scanner.go:16076
		switch lex.data[(lex.p)] {
		case 79:
			goto tr632
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr632
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr632:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st410
	st410:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof410
		}
	st_case_410:
		// line internal/php8/scanner.go:16117
		switch lex.data[(lex.p)] {
		case 78:
			goto tr633
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr633
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr633:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st411
	st411:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof411
		}
	st_case_411:
		// line internal/php8/scanner.go:16158
		switch lex.data[(lex.p)] {
		case 67:
			goto tr634
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr634
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr634:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st412
	st412:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof412
		}
	st_case_412:
		// line internal/php8/scanner.go:16199
		switch lex.data[(lex.p)] {
		case 69:
			goto tr635
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr635
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr620:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st413
	st413:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof413
		}
	st_case_413:
		// line internal/php8/scanner.go:16240
		switch lex.data[(lex.p)] {
		case 85:
			goto tr636
		case 92:
			goto st90
		case 96:
			goto tr376
		case 117:
			goto tr636
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr636:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st414
	st414:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof414
		}
	st_case_414:
		// line internal/php8/scanner.go:16281
		switch lex.data[(lex.p)] {
		case 82:
			goto tr637
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr637
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr637:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st415
	st415:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof415
		}
	st_case_415:
		// line internal/php8/scanner.go:16322
		switch lex.data[(lex.p)] {
		case 78:
			goto tr638
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr638
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr279:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st416
	st416:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof416
		}
	st_case_416:
		// line internal/php8/scanner.go:16363
		switch lex.data[(lex.p)] {
		case 84:
			goto tr639
		case 87:
			goto tr640
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr639
		case 119:
			goto tr640
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr639:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st417
	st417:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof417
		}
	st_case_417:
		// line internal/php8/scanner.go:16408
		switch lex.data[(lex.p)] {
		case 65:
			goto tr641
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr641
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr641:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st418
	st418:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof418
		}
	st_case_418:
		// line internal/php8/scanner.go:16449
		switch lex.data[(lex.p)] {
		case 84:
			goto tr642
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr642
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr642:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st419
	st419:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof419
		}
	st_case_419:
		// line internal/php8/scanner.go:16490
		switch lex.data[(lex.p)] {
		case 73:
			goto tr643
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr643
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr643:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st420
	st420:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof420
		}
	st_case_420:
		// line internal/php8/scanner.go:16531
		switch lex.data[(lex.p)] {
		case 67:
			goto tr644
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr644
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr640:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st421
	st421:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof421
		}
	st_case_421:
		// line internal/php8/scanner.go:16572
		switch lex.data[(lex.p)] {
		case 73:
			goto tr645
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr645
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr645:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st422
	st422:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof422
		}
	st_case_422:
		// line internal/php8/scanner.go:16613
		switch lex.data[(lex.p)] {
		case 84:
			goto tr646
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr646
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr646:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st423
	st423:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof423
		}
	st_case_423:
		// line internal/php8/scanner.go:16654
		switch lex.data[(lex.p)] {
		case 67:
			goto tr647
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr647
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr647:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st424
	st424:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof424
		}
	st_case_424:
		// line internal/php8/scanner.go:16695
		switch lex.data[(lex.p)] {
		case 72:
			goto tr648
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr648
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr280:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st425
	st425:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof425
		}
	st_case_425:
		// line internal/php8/scanner.go:16736
		switch lex.data[(lex.p)] {
		case 72:
			goto tr649
		case 82:
			goto tr650
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr649
		case 114:
			goto tr650
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr649:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st426
	st426:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof426
		}
	st_case_426:
		// line internal/php8/scanner.go:16781
		switch lex.data[(lex.p)] {
		case 82:
			goto tr651
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr651
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr651:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st427
	st427:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof427
		}
	st_case_427:
		// line internal/php8/scanner.go:16822
		switch lex.data[(lex.p)] {
		case 79:
			goto tr652
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr652
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr652:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st428
	st428:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof428
		}
	st_case_428:
		// line internal/php8/scanner.go:16863
		switch lex.data[(lex.p)] {
		case 87:
			goto tr653
		case 92:
			goto st90
		case 96:
			goto tr376
		case 119:
			goto tr653
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr650:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st429
	st429:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof429
		}
	st_case_429:
		// line internal/php8/scanner.go:16904
		switch lex.data[(lex.p)] {
		case 65:
			goto tr654
		case 89:
			goto tr655
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr654
		case 121:
			goto tr655
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr654:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st430
	st430:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof430
		}
	st_case_430:
		// line internal/php8/scanner.go:16949
		switch lex.data[(lex.p)] {
		case 73:
			goto tr656
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr656
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr656:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st431
	st431:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof431
		}
	st_case_431:
		// line internal/php8/scanner.go:16990
		switch lex.data[(lex.p)] {
		case 84:
			goto tr657
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr657
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr281:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st432
	st432:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof432
		}
	st_case_432:
		// line internal/php8/scanner.go:17031
		switch lex.data[(lex.p)] {
		case 78:
			goto tr658
		case 83:
			goto tr659
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr658
		case 115:
			goto tr659
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr658:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st433
	st433:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof433
		}
	st_case_433:
		// line internal/php8/scanner.go:17076
		switch lex.data[(lex.p)] {
		case 83:
			goto tr660
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr660
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr660:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st434
	st434:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof434
		}
	st_case_434:
		// line internal/php8/scanner.go:17117
		switch lex.data[(lex.p)] {
		case 69:
			goto tr661
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr661
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr661:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st435
	st435:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof435
		}
	st_case_435:
		// line internal/php8/scanner.go:17158
		switch lex.data[(lex.p)] {
		case 84:
			goto tr662
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr662
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr659:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st436
	st436:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof436
		}
	st_case_436:
		// line internal/php8/scanner.go:17199
		switch lex.data[(lex.p)] {
		case 69:
			goto tr663
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr663
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr282:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st437
	st437:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof437
		}
	st_case_437:
		// line internal/php8/scanner.go:17240
		switch lex.data[(lex.p)] {
		case 65:
			goto tr664
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr664
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr664:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st438
	st438:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof438
		}
	st_case_438:
		// line internal/php8/scanner.go:17281
		switch lex.data[(lex.p)] {
		case 82:
			goto tr665
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr665
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr283:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st439
	st439:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof439
		}
	st_case_439:
		// line internal/php8/scanner.go:17322
		switch lex.data[(lex.p)] {
		case 72:
			goto tr666
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr666
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr666:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st440
	st440:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof440
		}
	st_case_440:
		// line internal/php8/scanner.go:17363
		switch lex.data[(lex.p)] {
		case 73:
			goto tr667
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr667
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr667:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st441
	st441:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof441
		}
	st_case_441:
		// line internal/php8/scanner.go:17404
		switch lex.data[(lex.p)] {
		case 76:
			goto tr668
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr668
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr668:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st442
	st442:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof442
		}
	st_case_442:
		// line internal/php8/scanner.go:17445
		switch lex.data[(lex.p)] {
		case 69:
			goto tr669
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr669
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr284:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st443
	st443:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof443
		}
	st_case_443:
		// line internal/php8/scanner.go:17486
		switch lex.data[(lex.p)] {
		case 79:
			goto tr670
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr670
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr670:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st444
	st444:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof444
		}
	st_case_444:
		// line internal/php8/scanner.go:17527
		switch lex.data[(lex.p)] {
		case 82:
			goto tr671
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr671
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr285:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st445
	st445:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof445
		}
	st_case_445:
		// line internal/php8/scanner.go:17568
		switch lex.data[(lex.p)] {
		case 73:
			goto tr672
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr672
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr672:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st446
	st446:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof446
		}
	st_case_446:
		// line internal/php8/scanner.go:17609
		switch lex.data[(lex.p)] {
		case 69:
			goto tr673
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr673
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr673:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st447
	st447:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof447
		}
	st_case_447:
		// line internal/php8/scanner.go:17650
		switch lex.data[(lex.p)] {
		case 76:
			goto tr674
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr674
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr674:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st448
	st448:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof448
		}
	st_case_448:
		// line internal/php8/scanner.go:17691
		switch lex.data[(lex.p)] {
		case 68:
			goto tr675
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr675
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr675:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:302
		lex.act = 82
		goto st449
	st449:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof449
		}
	st_case_449:
		// line internal/php8/scanner.go:17732
		switch lex.data[(lex.p)] {
		case 10:
			goto tr151
		case 13:
			goto tr152
		case 32:
			goto st106
		case 92:
			goto st90
		case 96:
			goto tr676
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st106
				}
			default:
				goto tr676
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 91:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr676
				}
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr676
				}
			default:
				goto tr676
			}
		default:
			goto tr676
		}
		goto tr271
	tr154:
		// line internal/php8/scanner.rl:54

		goto st106
	st106:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		// line internal/php8/scanner.go:17781
		switch lex.data[(lex.p)] {
		case 10:
			goto tr151
		case 13:
			goto tr152
		case 32:
			goto st106
		case 70:
			goto st109
		case 102:
			goto st109
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st106
		}
		goto tr149
	tr151:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st107
	tr155:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st107
	tr158:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st107
	st107:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		// line internal/php8/scanner.go:17843
		switch lex.data[(lex.p)] {
		case 10:
			goto tr155
		case 13:
			goto tr156
		case 32:
			goto tr154
		case 70:
			goto tr157
		case 102:
			goto tr157
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr154
		}
		goto tr149
	tr152:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st108
	tr156:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st108
	st108:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		// line internal/php8/scanner.go:17891
		switch lex.data[(lex.p)] {
		case 10:
			goto tr158
		case 13:
			goto tr156
		case 32:
			goto tr154
		case 70:
			goto tr157
		case 102:
			goto tr157
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr154
		}
		goto tr149
	tr157:
		// line internal/php8/scanner.rl:54

		goto st109
	st109:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		// line internal/php8/scanner.go:17917
		switch lex.data[(lex.p)] {
		case 82:
			goto st110
		case 114:
			goto st110
		}
		goto tr149
	st110:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		switch lex.data[(lex.p)] {
		case 79:
			goto st111
		case 111:
			goto st111
		}
		goto tr149
	st111:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr161
		case 109:
			goto tr161
		}
		goto tr149
	st450:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof450
		}
	st_case_450:
		if lex.data[(lex.p)] == 96 {
			goto tr677
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr677
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr677
			}
		default:
			goto tr677
		}
		goto tr163
	tr163:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st451
	st451:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof451
		}
	st_case_451:
		// line internal/php8/scanner.go:17980
		switch lex.data[(lex.p)] {
		case 92:
			goto st112
		case 96:
			goto tr678
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr678
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr678
				}
			case lex.data[(lex.p)] >= 91:
				goto tr678
			}
		default:
			goto tr678
		}
		goto tr163
	st112:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		if lex.data[(lex.p)] == 96 {
			goto tr162
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr162
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr162
			}
		default:
			goto tr162
		}
		goto tr163
	st452:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof452
		}
	st_case_452:
		if lex.data[(lex.p)] == 61 {
			goto tr680
		}
		goto tr300
	tr288:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st453
	st453:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof453
		}
	st_case_453:
		// line internal/php8/scanner.go:18047
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr681
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr681:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st454
	st454:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof454
		}
	st_case_454:
		// line internal/php8/scanner.go:18084
		switch lex.data[(lex.p)] {
		case 67:
			goto tr682
		case 68:
			goto tr683
		case 70:
			goto tr684
		case 72:
			goto tr685
		case 76:
			goto tr686
		case 77:
			goto tr687
		case 78:
			goto tr688
		case 84:
			goto tr689
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr682
		case 100:
			goto tr683
		case 102:
			goto tr684
		case 104:
			goto tr685
		case 108:
			goto tr686
		case 109:
			goto tr687
		case 110:
			goto tr688
		case 116:
			goto tr689
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr682:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st455
	st455:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof455
		}
	st_case_455:
		// line internal/php8/scanner.go:18153
		switch lex.data[(lex.p)] {
		case 76:
			goto tr690
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr690
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr690:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st456
	st456:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof456
		}
	st_case_456:
		// line internal/php8/scanner.go:18194
		switch lex.data[(lex.p)] {
		case 65:
			goto tr691
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr691
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr691:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st457
	st457:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof457
		}
	st_case_457:
		// line internal/php8/scanner.go:18235
		switch lex.data[(lex.p)] {
		case 83:
			goto tr692
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr692
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr692:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st458
	st458:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof458
		}
	st_case_458:
		// line internal/php8/scanner.go:18276
		switch lex.data[(lex.p)] {
		case 83:
			goto tr693
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr693
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr693:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st459
	st459:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof459
		}
	st_case_459:
		// line internal/php8/scanner.go:18317
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr694
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr694:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st460
	st460:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof460
		}
	st_case_460:
		// line internal/php8/scanner.go:18354
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr695
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr683:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st461
	st461:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof461
		}
	st_case_461:
		// line internal/php8/scanner.go:18391
		switch lex.data[(lex.p)] {
		case 73:
			goto tr696
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr696
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr696:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st462
	st462:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof462
		}
	st_case_462:
		// line internal/php8/scanner.go:18432
		switch lex.data[(lex.p)] {
		case 82:
			goto tr697
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr697
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr697:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st463
	st463:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof463
		}
	st_case_463:
		// line internal/php8/scanner.go:18473
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr698
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr698:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st464
	st464:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof464
		}
	st_case_464:
		// line internal/php8/scanner.go:18510
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr699
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr684:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st465
	st465:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof465
		}
	st_case_465:
		// line internal/php8/scanner.go:18547
		switch lex.data[(lex.p)] {
		case 73:
			goto tr700
		case 85:
			goto tr701
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr700
		case 117:
			goto tr701
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr700:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st466
	st466:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof466
		}
	st_case_466:
		// line internal/php8/scanner.go:18592
		switch lex.data[(lex.p)] {
		case 76:
			goto tr702
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr702
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr702:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st467
	st467:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof467
		}
	st_case_467:
		// line internal/php8/scanner.go:18633
		switch lex.data[(lex.p)] {
		case 69:
			goto tr703
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr703
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr703:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st468
	st468:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof468
		}
	st_case_468:
		// line internal/php8/scanner.go:18674
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr704
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr704:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st469
	st469:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof469
		}
	st_case_469:
		// line internal/php8/scanner.go:18711
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr705
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr701:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st470
	st470:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof470
		}
	st_case_470:
		// line internal/php8/scanner.go:18748
		switch lex.data[(lex.p)] {
		case 78:
			goto tr706
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr706
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr706:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st471
	st471:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof471
		}
	st_case_471:
		// line internal/php8/scanner.go:18789
		switch lex.data[(lex.p)] {
		case 67:
			goto tr707
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr707
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr707:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st472
	st472:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof472
		}
	st_case_472:
		// line internal/php8/scanner.go:18830
		switch lex.data[(lex.p)] {
		case 84:
			goto tr708
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr708
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr708:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st473
	st473:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof473
		}
	st_case_473:
		// line internal/php8/scanner.go:18871
		switch lex.data[(lex.p)] {
		case 73:
			goto tr709
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr709
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr709:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st474
	st474:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof474
		}
	st_case_474:
		// line internal/php8/scanner.go:18912
		switch lex.data[(lex.p)] {
		case 79:
			goto tr710
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr710
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr710:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st475
	st475:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof475
		}
	st_case_475:
		// line internal/php8/scanner.go:18953
		switch lex.data[(lex.p)] {
		case 78:
			goto tr711
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr711
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr711:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st476
	st476:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof476
		}
	st_case_476:
		// line internal/php8/scanner.go:18994
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr712
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr712:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st477
	st477:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof477
		}
	st_case_477:
		// line internal/php8/scanner.go:19031
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr713
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr685:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st478
	st478:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof478
		}
	st_case_478:
		// line internal/php8/scanner.go:19068
		switch lex.data[(lex.p)] {
		case 65:
			goto tr714
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr714
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr714:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st479
	st479:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof479
		}
	st_case_479:
		// line internal/php8/scanner.go:19109
		switch lex.data[(lex.p)] {
		case 76:
			goto tr715
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr715
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr715:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st480
	st480:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof480
		}
	st_case_480:
		// line internal/php8/scanner.go:19150
		switch lex.data[(lex.p)] {
		case 84:
			goto tr716
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr716
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr716:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st481
	st481:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof481
		}
	st_case_481:
		// line internal/php8/scanner.go:19191
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr717
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr717:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st482
	st482:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof482
		}
	st_case_482:
		// line internal/php8/scanner.go:19228
		switch lex.data[(lex.p)] {
		case 67:
			goto tr718
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr718
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr718:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st483
	st483:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof483
		}
	st_case_483:
		// line internal/php8/scanner.go:19269
		switch lex.data[(lex.p)] {
		case 79:
			goto tr719
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr719
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr719:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st484
	st484:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof484
		}
	st_case_484:
		// line internal/php8/scanner.go:19310
		switch lex.data[(lex.p)] {
		case 77:
			goto tr720
		case 92:
			goto st90
		case 96:
			goto tr376
		case 109:
			goto tr720
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr720:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st485
	st485:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof485
		}
	st_case_485:
		// line internal/php8/scanner.go:19351
		switch lex.data[(lex.p)] {
		case 80:
			goto tr721
		case 92:
			goto st90
		case 96:
			goto tr376
		case 112:
			goto tr721
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr721:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st486
	st486:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof486
		}
	st_case_486:
		// line internal/php8/scanner.go:19392
		switch lex.data[(lex.p)] {
		case 73:
			goto tr722
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr722
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr722:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st487
	st487:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof487
		}
	st_case_487:
		// line internal/php8/scanner.go:19433
		switch lex.data[(lex.p)] {
		case 76:
			goto tr723
		case 92:
			goto st90
		case 96:
			goto tr376
		case 108:
			goto tr723
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr723:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st488
	st488:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof488
		}
	st_case_488:
		// line internal/php8/scanner.go:19474
		switch lex.data[(lex.p)] {
		case 69:
			goto tr724
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr724
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr724:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st489
	st489:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof489
		}
	st_case_489:
		// line internal/php8/scanner.go:19515
		switch lex.data[(lex.p)] {
		case 82:
			goto tr725
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr725
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr686:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st490
	st490:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof490
		}
	st_case_490:
		// line internal/php8/scanner.go:19556
		switch lex.data[(lex.p)] {
		case 73:
			goto tr726
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr726
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr726:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st491
	st491:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof491
		}
	st_case_491:
		// line internal/php8/scanner.go:19597
		switch lex.data[(lex.p)] {
		case 78:
			goto tr727
		case 92:
			goto st90
		case 96:
			goto tr376
		case 110:
			goto tr727
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr727:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st492
	st492:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof492
		}
	st_case_492:
		// line internal/php8/scanner.go:19638
		switch lex.data[(lex.p)] {
		case 69:
			goto tr728
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr728
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr728:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st493
	st493:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof493
		}
	st_case_493:
		// line internal/php8/scanner.go:19679
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr729
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr729:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st494
	st494:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof494
		}
	st_case_494:
		// line internal/php8/scanner.go:19716
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr730
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr687:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st495
	st495:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof495
		}
	st_case_495:
		// line internal/php8/scanner.go:19753
		switch lex.data[(lex.p)] {
		case 69:
			goto tr731
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr731
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr731:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st496
	st496:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof496
		}
	st_case_496:
		// line internal/php8/scanner.go:19794
		switch lex.data[(lex.p)] {
		case 84:
			goto tr732
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr732
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr732:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st497
	st497:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof497
		}
	st_case_497:
		// line internal/php8/scanner.go:19835
		switch lex.data[(lex.p)] {
		case 72:
			goto tr733
		case 92:
			goto st90
		case 96:
			goto tr376
		case 104:
			goto tr733
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr733:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st498
	st498:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof498
		}
	st_case_498:
		// line internal/php8/scanner.go:19876
		switch lex.data[(lex.p)] {
		case 79:
			goto tr734
		case 92:
			goto st90
		case 96:
			goto tr376
		case 111:
			goto tr734
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr734:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st499
	st499:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof499
		}
	st_case_499:
		// line internal/php8/scanner.go:19917
		switch lex.data[(lex.p)] {
		case 68:
			goto tr735
		case 92:
			goto st90
		case 96:
			goto tr376
		case 100:
			goto tr735
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr735:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st500
	st500:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof500
		}
	st_case_500:
		// line internal/php8/scanner.go:19958
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr736
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr736:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st501
	st501:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof501
		}
	st_case_501:
		// line internal/php8/scanner.go:19995
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr737
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr688:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st502
	st502:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof502
		}
	st_case_502:
		// line internal/php8/scanner.go:20032
		switch lex.data[(lex.p)] {
		case 65:
			goto tr738
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr738
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr738:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st503
	st503:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof503
		}
	st_case_503:
		// line internal/php8/scanner.go:20073
		switch lex.data[(lex.p)] {
		case 77:
			goto tr739
		case 92:
			goto st90
		case 96:
			goto tr376
		case 109:
			goto tr739
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr739:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st504
	st504:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof504
		}
	st_case_504:
		// line internal/php8/scanner.go:20114
		switch lex.data[(lex.p)] {
		case 69:
			goto tr740
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr740
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr740:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st505
	st505:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof505
		}
	st_case_505:
		// line internal/php8/scanner.go:20155
		switch lex.data[(lex.p)] {
		case 83:
			goto tr741
		case 92:
			goto st90
		case 96:
			goto tr376
		case 115:
			goto tr741
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr741:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st506
	st506:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof506
		}
	st_case_506:
		// line internal/php8/scanner.go:20196
		switch lex.data[(lex.p)] {
		case 80:
			goto tr742
		case 92:
			goto st90
		case 96:
			goto tr376
		case 112:
			goto tr742
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr742:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st507
	st507:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof507
		}
	st_case_507:
		// line internal/php8/scanner.go:20237
		switch lex.data[(lex.p)] {
		case 65:
			goto tr743
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr743
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr743:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st508
	st508:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof508
		}
	st_case_508:
		// line internal/php8/scanner.go:20278
		switch lex.data[(lex.p)] {
		case 67:
			goto tr744
		case 92:
			goto st90
		case 96:
			goto tr376
		case 99:
			goto tr744
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr744:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st509
	st509:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof509
		}
	st_case_509:
		// line internal/php8/scanner.go:20319
		switch lex.data[(lex.p)] {
		case 69:
			goto tr745
		case 92:
			goto st90
		case 96:
			goto tr376
		case 101:
			goto tr745
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr745:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st510
	st510:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof510
		}
	st_case_510:
		// line internal/php8/scanner.go:20360
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr746
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr746:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st511
	st511:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof511
		}
	st_case_511:
		// line internal/php8/scanner.go:20397
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr747
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr689:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st512
	st512:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof512
		}
	st_case_512:
		// line internal/php8/scanner.go:20434
		switch lex.data[(lex.p)] {
		case 82:
			goto tr748
		case 92:
			goto st90
		case 96:
			goto tr376
		case 114:
			goto tr748
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr748:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st513
	st513:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof513
		}
	st_case_513:
		// line internal/php8/scanner.go:20475
		switch lex.data[(lex.p)] {
		case 65:
			goto tr749
		case 92:
			goto st90
		case 96:
			goto tr376
		case 97:
			goto tr749
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr749:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st514
	st514:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof514
		}
	st_case_514:
		// line internal/php8/scanner.go:20516
		switch lex.data[(lex.p)] {
		case 73:
			goto tr750
		case 92:
			goto st90
		case 96:
			goto tr376
		case 105:
			goto tr750
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr750:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st515
	st515:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof515
		}
	st_case_515:
		// line internal/php8/scanner.go:20557
		switch lex.data[(lex.p)] {
		case 84:
			goto tr751
		case 92:
			goto st90
		case 96:
			goto tr376
		case 116:
			goto tr751
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr751:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st516
	st516:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof516
		}
	st_case_516:
		// line internal/php8/scanner.go:20598
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr752
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	tr752:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:378
		lex.act = 148
		goto st517
	st517:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof517
		}
	st_case_517:
		// line internal/php8/scanner.go:20635
		switch lex.data[(lex.p)] {
		case 92:
			goto st90
		case 95:
			goto tr753
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr376
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr376
				}
			case lex.data[(lex.p)] >= 91:
				goto tr376
			}
		default:
			goto tr376
		}
		goto tr271
	st518:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof518
		}
	st_case_518:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr754
		case 62:
			goto tr755
		case 124:
			goto tr756
		}
		goto tr300
	tr164:
		// line internal/php8/scanner.rl:417
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				goto st151
			}
		}
		goto st519
	tr165:
		// line internal/php8/scanner.rl:415
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 519
				goto _out
			}
		}
		goto st519
	tr757:
		// line internal/php8/scanner.rl:417
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st151
			}
		}
		goto st519
	tr764:
		// line internal/php8/scanner.rl:413
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st519
	tr765:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:413
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st519
	tr770:
		// line internal/php8/scanner.rl:417
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st151
			}
		}
		goto st519
	tr771:
		// line internal/php8/scanner.rl:414
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 519
				goto _out
			}
		}
		goto st519
	tr773:
		lex.cs = 519
		// line internal/php8/scanner.rl:416
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			lex.cs = 151
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st519:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof519
		}
	st_case_519:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:20732
		switch lex.data[(lex.p)] {
		case 10:
			goto tr759
		case 13:
			goto tr760
		case 32:
			goto st520
		case 45:
			goto st523
		case 63:
			goto tr762
		case 96:
			goto tr757
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st520
				}
			default:
				goto tr757
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr757
				}
			case lex.data[(lex.p)] >= 91:
				goto tr757
			}
		default:
			goto tr757
		}
		goto st525
	tr766:
		// line internal/php8/scanner.rl:54

		goto st520
	st520:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof520
		}
	st_case_520:
		// line internal/php8/scanner.go:20779
		switch lex.data[(lex.p)] {
		case 10:
			goto tr759
		case 13:
			goto tr760
		case 32:
			goto st520
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st520
		}
		goto tr764
	tr759:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st521
	tr767:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st521
	tr769:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st521
	st521:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof521
		}
	st_case_521:
		// line internal/php8/scanner.go:20837
		switch lex.data[(lex.p)] {
		case 10:
			goto tr767
		case 13:
			goto tr768
		case 32:
			goto tr766
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr766
		}
		goto tr765
	tr760:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st522
	tr768:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st522
	st522:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof522
		}
	st_case_522:
		// line internal/php8/scanner.go:20881
		switch lex.data[(lex.p)] {
		case 10:
			goto tr769
		case 13:
			goto tr768
		case 32:
			goto tr766
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr766
		}
		goto tr765
	st523:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof523
		}
	st_case_523:
		if lex.data[(lex.p)] == 62 {
			goto tr771
		}
		goto tr770
	tr762:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st524
	st524:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof524
		}
	st_case_524:
		// line internal/php8/scanner.go:20913
		if lex.data[(lex.p)] == 45 {
			goto st113
		}
		goto tr770
	st113:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		if lex.data[(lex.p)] == 62 {
			goto tr165
		}
		goto tr164
	st525:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof525
		}
	st_case_525:
		if lex.data[(lex.p)] == 96 {
			goto tr773
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr773
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr773
				}
			case lex.data[(lex.p)] >= 91:
				goto tr773
			}
		default:
			goto tr773
		}
		goto st525
	tr777:
		lex.cs = 526
		// line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 161:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				lex.cs = 549
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr778:
		lex.cs = 526
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:421
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.cs = 549
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st526:
		// line NONE:1
		lex.ts = 0

		// line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof526
		}
	st_case_526:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:20998
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr775
		case 1037:
			goto tr776
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr774
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	tr774:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:421
		lex.act = 161
		goto st527
	tr779:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:421
		lex.act = 161
		goto st527
	st527:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof527
		}
	st_case_527:
		// line internal/php8/scanner.go:21072
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr775
		case 1037:
			goto tr776
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr774
		}
		goto tr777
	tr775:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st528
	tr780:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st528
	tr782:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st528
	st528:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof528
		}
	st_case_528:
		// line internal/php8/scanner.go:21166
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr780
		case 1037:
			goto tr781
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr779
		}
		goto tr778
	tr776:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st529
	tr781:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st529
	st529:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof529
		}
	st_case_529:
		// line internal/php8/scanner.go:21246
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr782
		case 1037:
			goto tr781
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr779
		}
		goto tr778
	tr166:
		// line internal/php8/scanner.rl:430
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(530, 151)
			goto _out
		}
		goto st530
	tr790:
		// line internal/php8/scanner.rl:432
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 530
					lex.top++
					goto st551
				}
			}
		}
		goto st530
	tr791:
		// line internal/php8/scanner.rl:431
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(530, 567)
			goto _out
		}
		goto st530
	tr792:
		lex.cs = 530
		// line NONE:1
		switch lex.act {
		case 162:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(530, 151)
				goto _out
			}
		case 163:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(530, 567)
				goto _out
			}
		case 165:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE

				if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
					lex.cs = 549
				}
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr793:
		lex.cs = 530
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:433
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 549
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr799:
		lex.cs = 530
		// line internal/php8/scanner.rl:433
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 549
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st530:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof530
		}
	st_case_530:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:21378
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto st531
		case 1403:
			goto st114
		case 1546:
			goto tr786
		case 1549:
			goto tr787
		case 1572:
			goto st535
		case 1659:
			goto st536
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr785
		}
		goto st0
	st531:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof531
		}
	st_case_531:
		if lex.data[(lex.p)] == 123 {
			goto tr791
		}
		goto tr790
	st114:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof114
		}
	st_case_114:
		if lex.data[(lex.p)] == 36 {
			goto tr166
		}
		goto st0
	tr785:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:433
		lex.act = 165
		goto st532
	tr794:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:433
		lex.act = 165
		goto st532
	tr798:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:431
		lex.act = 163
		goto st532
	tr800:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:430
		lex.act = 162
		goto st532
	st532:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof532
		}
	st_case_532:
		// line internal/php8/scanner.go:21488
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr786
		case 1549:
			goto tr787
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr785
		}
		goto tr792
	tr786:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st533
	tr795:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st533
	tr797:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st533
	st533:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof533
		}
	st_case_533:
		// line internal/php8/scanner.go:21582
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr795
		case 1549:
			goto tr796
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr794
		}
		goto tr793
	tr787:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st534
	tr796:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st534
	st534:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof534
		}
	st_case_534:
		// line internal/php8/scanner.go:21662
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr797
		case 1549:
			goto tr796
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr794
		}
		goto tr793
	st535:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof535
		}
	st_case_535:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1403:
			goto tr791
		case 1546:
			goto tr786
		case 1549:
			goto tr787
		case 1659:
			goto tr798
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr785
		}
		goto tr790
	st536:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof536
		}
	st_case_536:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto tr166
		case 1546:
			goto tr786
		case 1549:
			goto tr787
		case 1572:
			goto tr800
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr785
		}
		goto tr799
	tr168:
		// line internal/php8/scanner.rl:447
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(2)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 537
					lex.top++
					goto st551
				}
			}
		}
		goto st537
	tr169:
		// line internal/php8/scanner.rl:446
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(537, 567)
			goto _out
		}
		goto st537
	tr170:
		// line internal/php8/scanner.rl:445
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(537, 151)
			goto _out
		}
		goto st537
	tr802:
		lex.cs = 537
		// line internal/php8/scanner.rl:448
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('`'))
			lex.cs = 151
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr810:
		lex.cs = 537
		// line NONE:1
		switch lex.act {
		case 166:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(537, 151)
				goto _out
			}
		case 167:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(537, 567)
				goto _out
			}
		case 168:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(2)
				{
					lex.growCallStack()
					{
						lex.stack[lex.top] = 537
						lex.top++
						goto st551
					}
				}
			}
		case 169:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('`'))
				lex.cs = 151
				{
					(lex.p)++
					goto _out
				}
			}
		case 170:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr811:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:449
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 537
				goto _out
			}
		}
		goto st537
	tr816:
		// line internal/php8/scanner.rl:449
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 537
				goto _out
			}
		}
		goto st537
	st537:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof537
		}
	st_case_537:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:21905
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto st115
		case 1888:
			goto tr802
		case 1915:
			goto st116
		case 2058:
			goto tr805
		case 2061:
			goto tr806
		case 2084:
			goto st541
		case 2144:
			goto tr808
		case 2171:
			goto st542
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr804
		}
		goto st0
	st115:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch lex.data[(lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr169
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr168
	st116:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof116
		}
	st_case_116:
		if lex.data[(lex.p)] == 36 {
			goto tr170
		}
		goto st0
	tr804:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:449
		lex.act = 170
		goto st538
	tr808:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:448
		lex.act = 169
		goto st538
	tr812:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:449
		lex.act = 170
		goto st538
	tr817:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:447
		lex.act = 168
		goto st538
	tr818:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:446
		lex.act = 167
		goto st538
	tr819:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:445
		lex.act = 166
		goto st538
	st538:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof538
		}
	st_case_538:
		// line internal/php8/scanner.go:22048
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr805
		case 2061:
			goto tr806
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr804
		}
		goto tr810
	tr805:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st539
	tr813:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st539
	tr815:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st539
	st539:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof539
		}
	st_case_539:
		// line internal/php8/scanner.go:22142
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr813
		case 2061:
			goto tr814
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr812
		}
		goto tr811
	tr806:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st540
	tr814:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st540
	st540:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof540
		}
	st_case_540:
		// line internal/php8/scanner.go:22222
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr815
		case 2061:
			goto tr814
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr812
		}
		goto tr811
	st541:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof541
		}
	st_case_541:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1887:
			goto tr168
		case 1915:
			goto tr169
		case 2058:
			goto tr805
		case 2061:
			goto tr806
		case 2143:
			goto tr817
		case 2171:
			goto tr818
		}
		switch {
		case _widec < 2113:
			switch {
			case _widec < 1889:
				if 1857 <= _widec && _widec <= 1882 {
					goto tr168
				}
			case _widec > 1914:
				switch {
				case _widec > 2047:
					if 2048 <= _widec && _widec <= 2112 {
						goto tr804
					}
				case _widec >= 1920:
					goto tr168
				}
			default:
				goto tr168
			}
		case _widec > 2138:
			switch {
			case _widec < 2145:
				if 2139 <= _widec && _widec <= 2144 {
					goto tr804
				}
			case _widec > 2170:
				switch {
				case _widec > 2175:
					if 2176 <= _widec && _widec <= 2303 {
						goto tr817
					}
				case _widec >= 2172:
					goto tr804
				}
			default:
				goto tr817
			}
		default:
			goto tr817
		}
		goto tr816
	st542:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof542
		}
	st_case_542:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto tr170
		case 2058:
			goto tr805
		case 2061:
			goto tr806
		case 2084:
			goto tr819
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr804
		}
		goto tr816
	tr171:
		// line internal/php8/scanner.rl:459
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(2)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 543
					lex.top++
					goto st551
				}
			}
		}
		goto st543
	tr172:
		// line internal/php8/scanner.rl:458
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(543, 567)
			goto _out
		}
		goto st543
	tr173:
		// line internal/php8/scanner.rl:457
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(543, 151)
			goto _out
		}
		goto st543
	tr820:
		lex.cs = 543
		// line internal/php8/scanner.rl:460
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('"'))
			lex.cs = 151
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr829:
		lex.cs = 543
		// line NONE:1
		switch lex.act {
		case 171:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(543, 151)
				goto _out
			}
		case 172:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(543, 567)
				goto _out
			}
		case 173:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(2)
				{
					lex.growCallStack()
					{
						lex.stack[lex.top] = 543
						lex.top++
						goto st551
					}
				}
			}
		case 174:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('"'))
				lex.cs = 151
				{
					(lex.p)++
					goto _out
				}
			}
		case 175:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr830:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:461
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 543
				goto _out
			}
		}
		goto st543
	tr835:
		// line internal/php8/scanner.rl:461
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 543
				goto _out
			}
		}
		goto st543
	st543:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof543
		}
	st_case_543:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:22506
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2338:
			goto tr820
		case 2340:
			goto st117
		case 2427:
			goto st118
		case 2570:
			goto tr824
		case 2573:
			goto tr825
		case 2594:
			goto tr826
		case 2596:
			goto st547
		case 2683:
			goto st548
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr823
		}
		goto st0
	st117:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		switch lex.data[(lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr172
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr171
	st118:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		if lex.data[(lex.p)] == 36 {
			goto tr173
		}
		goto st0
	tr823:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:461
		lex.act = 175
		goto st544
	tr826:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:460
		lex.act = 174
		goto st544
	tr831:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:461
		lex.act = 175
		goto st544
	tr836:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:459
		lex.act = 173
		goto st544
	tr837:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:458
		lex.act = 172
		goto st544
	tr838:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:457
		lex.act = 171
		goto st544
	st544:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof544
		}
	st_case_544:
		// line internal/php8/scanner.go:22649
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr824
		case 2573:
			goto tr825
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr823
		}
		goto tr829
	tr824:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st545
	tr832:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st545
	tr834:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st545
	st545:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof545
		}
	st_case_545:
		// line internal/php8/scanner.go:22743
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr832
		case 2573:
			goto tr833
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr831
		}
		goto tr830
	tr825:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st546
	tr833:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st546
	st546:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof546
		}
	st_case_546:
		// line internal/php8/scanner.go:22823
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr834
		case 2573:
			goto tr833
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr831
		}
		goto tr830
	st547:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof547
		}
	st_case_547:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2399:
			goto tr171
		case 2427:
			goto tr172
		case 2570:
			goto tr824
		case 2573:
			goto tr825
		case 2655:
			goto tr836
		case 2683:
			goto tr837
		}
		switch {
		case _widec < 2625:
			switch {
			case _widec < 2401:
				if 2369 <= _widec && _widec <= 2394 {
					goto tr171
				}
			case _widec > 2426:
				switch {
				case _widec > 2559:
					if 2560 <= _widec && _widec <= 2624 {
						goto tr823
					}
				case _widec >= 2432:
					goto tr171
				}
			default:
				goto tr171
			}
		case _widec > 2650:
			switch {
			case _widec < 2657:
				if 2651 <= _widec && _widec <= 2656 {
					goto tr823
				}
			case _widec > 2682:
				switch {
				case _widec > 2687:
					if 2688 <= _widec && _widec <= 2815 {
						goto tr836
					}
				case _widec >= 2684:
					goto tr823
				}
			default:
				goto tr836
			}
		default:
			goto tr836
		}
		goto tr835
	st548:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof548
		}
	st_case_548:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2340:
			goto tr173
		case 2570:
			goto tr824
		case 2573:
			goto tr825
		case 2596:
			goto tr838
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr823
		}
		goto tr835
	tr840:
		lex.cs = 549
		// line internal/php8/scanner.rl:469
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_END_HEREDOC
			lex.cs = 151
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st549:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof549
		}
	st_case_549:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23050
		if lex.data[(lex.p)] == 96 {
			goto st0
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto st550
	st550:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof550
		}
	st_case_550:
		if lex.data[(lex.p)] == 96 {
			goto tr840
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr840
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr840
				}
			case lex.data[(lex.p)] >= 91:
				goto tr840
			}
		default:
			goto tr840
		}
		goto st550
	tr174:
		// line internal/php8/scanner.rl:489
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st551
	tr175:
		// line internal/php8/scanner.rl:485
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 551
				goto _out
			}
		}
		goto st551
	tr177:
		// line internal/php8/scanner.rl:486
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 551
				goto _out
			}
		}
		goto st551
	tr841:
		// line internal/php8/scanner.rl:489
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st551
	tr846:
		// line internal/php8/scanner.rl:488
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('['))
			lex.call(551, 557)
			goto _out
		}
		goto st551
	tr847:
		// line internal/php8/scanner.rl:489
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st551
	tr849:
		// line internal/php8/scanner.rl:484
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 551
				goto _out
			}
		}
		goto st551
	tr852:
		// line internal/php8/scanner.rl:487
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			lex.ret(1)
			goto _out
		}
		goto st551
	st551:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof551
		}
	st_case_551:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23147
		switch lex.data[(lex.p)] {
		case 36:
			goto st552
		case 45:
			goto tr843
		case 63:
			goto tr844
		case 91:
			goto tr846
		case 96:
			goto tr841
		}
		switch {
		case lex.data[(lex.p)] < 92:
			if lex.data[(lex.p)] <= 64 {
				goto tr841
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr841
			}
		default:
			goto tr841
		}
		goto st556
	st552:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof552
		}
	st_case_552:
		if lex.data[(lex.p)] == 96 {
			goto tr847
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr847
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr847
			}
		default:
			goto tr847
		}
		goto st553
	st553:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof553
		}
	st_case_553:
		if lex.data[(lex.p)] == 96 {
			goto tr849
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr849
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr849
				}
			case lex.data[(lex.p)] >= 91:
				goto tr849
			}
		default:
			goto tr849
		}
		goto st553
	tr843:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st554
	st554:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof554
		}
	st_case_554:
		// line internal/php8/scanner.go:23230
		if lex.data[(lex.p)] == 62 {
			goto st119
		}
		goto tr847
	st119:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof119
		}
	st_case_119:
		if lex.data[(lex.p)] == 96 {
			goto tr174
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr174
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr174
			}
		default:
			goto tr174
		}
		goto tr175
	tr844:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st555
	st555:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof555
		}
	st_case_555:
		// line internal/php8/scanner.go:23266
		if lex.data[(lex.p)] == 45 {
			goto st120
		}
		goto tr847
	st120:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof120
		}
	st_case_120:
		if lex.data[(lex.p)] == 62 {
			goto st121
		}
		goto tr174
	st121:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof121
		}
	st_case_121:
		if lex.data[(lex.p)] == 96 {
			goto tr174
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr174
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr174
			}
		default:
			goto tr174
		}
		goto tr177
	st556:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof556
		}
	st_case_556:
		if lex.data[(lex.p)] == 96 {
			goto tr852
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr852
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr852
				}
			case lex.data[(lex.p)] >= 91:
				goto tr852
			}
		default:
			goto tr852
		}
		goto st556
	tr178:
		// line internal/php8/scanner.rl:493
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 557
				goto _out
			}
		}
		goto st557
	tr853:
		// line internal/php8/scanner.rl:499
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st557
	tr854:
		// line internal/php8/scanner.rl:496
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st557
	tr857:
		// line internal/php8/scanner.rl:497
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 557
				goto _out
			}
		}
		goto st557
	tr861:
		// line internal/php8/scanner.rl:498
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(']'))
			lex.ret(2)
			goto _out
		}
		goto st557
	tr862:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:496
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st557
	tr863:
		// line internal/php8/scanner.rl:497
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 557
				goto _out
			}
		}
		goto st557
	tr865:
		// line internal/php8/scanner.rl:494
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 557
				goto _out
			}
		}
		goto st557
	tr866:
		// line internal/php8/scanner.rl:493
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 557
				goto _out
			}
		}
		goto st557
	tr870:
		// line internal/php8/scanner.rl:495
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 557
				goto _out
			}
		}
		goto st557
	st557:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof557
		}
	st_case_557:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23398
		switch lex.data[(lex.p)] {
		case 10:
			goto tr855
		case 13:
			goto tr856
		case 32:
			goto tr854
		case 33:
			goto tr857
		case 35:
			goto tr854
		case 36:
			goto st560
		case 39:
			goto tr854
		case 48:
			goto tr859
		case 92:
			goto tr854
		case 93:
			goto tr861
		case 96:
			goto tr853
		case 124:
			goto tr857
		case 126:
			goto tr857
		}
		switch {
		case lex.data[(lex.p)] < 37:
			switch {
			case lex.data[(lex.p)] < 9:
				if lex.data[(lex.p)] <= 8 {
					goto tr853
				}
			case lex.data[(lex.p)] > 12:
				if 14 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 34 {
					goto tr853
				}
			default:
				goto tr854
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 58:
				if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
					goto tr179
				}
			case lex.data[(lex.p)] > 64:
				switch {
				case lex.data[(lex.p)] > 94:
					if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
						goto tr853
					}
				case lex.data[(lex.p)] >= 91:
					goto tr857
				}
			default:
				goto tr857
			}
		default:
			goto tr857
		}
		goto st566
	tr855:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st558
	st558:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof558
		}
	st_case_558:
		// line internal/php8/scanner.go:23480
		goto tr862
	tr856:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st559
	st559:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof559
		}
	st_case_559:
		// line internal/php8/scanner.go:23499
		if lex.data[(lex.p)] == 10 {
			goto tr855
		}
		goto tr862
	st560:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof560
		}
	st_case_560:
		if lex.data[(lex.p)] == 96 {
			goto tr863
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr863
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr863
			}
		default:
			goto tr863
		}
		goto st561
	st561:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof561
		}
	st_case_561:
		if lex.data[(lex.p)] == 96 {
			goto tr865
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr865
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr865
				}
			case lex.data[(lex.p)] >= 91:
				goto tr865
			}
		default:
			goto tr865
		}
		goto st561
	tr859:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st562
	st562:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof562
		}
	st_case_562:
		// line internal/php8/scanner.go:23561
		switch lex.data[(lex.p)] {
		case 95:
			goto st122
		case 98:
			goto st123
		case 120:
			goto st124
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr179
		}
		goto tr866
	tr179:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st563
	st563:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof563
		}
	st_case_563:
		// line internal/php8/scanner.go:23584
		if lex.data[(lex.p)] == 95 {
			goto st122
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr179
		}
		goto tr866
	st122:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof122
		}
	st_case_122:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr179
		}
		goto tr178
	st123:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof123
		}
	st_case_123:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr180
		}
		goto tr178
	tr180:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st564
	st564:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof564
		}
	st_case_564:
		// line internal/php8/scanner.go:23620
		if lex.data[(lex.p)] == 95 {
			goto st123
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr180
		}
		goto tr866
	st124:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr181
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr181
			}
		default:
			goto tr181
		}
		goto tr178
	tr181:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st565
	st565:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof565
		}
	st_case_565:
		// line internal/php8/scanner.go:23656
		if lex.data[(lex.p)] == 95 {
			goto st124
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr181
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr181
			}
		default:
			goto tr181
		}
		goto tr866
	st566:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof566
		}
	st_case_566:
		if lex.data[(lex.p)] == 96 {
			goto tr870
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr870
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr870
				}
			case lex.data[(lex.p)] >= 91:
				goto tr870
			}
		default:
			goto tr870
		}
		goto st566
	tr182:
		lex.cs = 567
		// line internal/php8/scanner.rl:507
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr184:
		lex.cs = 567
		// line internal/php8/scanner.rl:506
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_STRING_VARNAME
			lex.cs = 151
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr871:
		lex.cs = 567
		// line internal/php8/scanner.rl:507
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr873:
		lex.cs = 567
		// line internal/php8/scanner.rl:507
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	st567:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof567
		}
	st_case_567:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23735
		if lex.data[(lex.p)] == 96 {
			goto tr871
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr871
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr871
			}
		default:
			goto tr871
		}
		goto tr872
	tr872:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st568
	st568:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof568
		}
	st_case_568:
		// line internal/php8/scanner.go:23762
		switch lex.data[(lex.p)] {
		case 91:
			goto tr184
		case 96:
			goto tr873
		case 125:
			goto tr184
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr873
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr873
				}
			case lex.data[(lex.p)] >= 92:
				goto tr873
			}
		default:
			goto tr873
		}
		goto st125
	st125:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		switch lex.data[(lex.p)] {
		case 91:
			goto tr184
		case 96:
			goto tr182
		case 125:
			goto tr184
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr182
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr182
				}
			case lex.data[(lex.p)] >= 92:
				goto tr182
			}
		default:
			goto tr182
		}
		goto st125
	tr185:
		lex.cs = 569
		// line internal/php8/scanner.rl:515
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr195:
		// line internal/php8/scanner.rl:67
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st569
	tr874:
		lex.cs = 569
		// line internal/php8/scanner.rl:515
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr879:
		lex.cs = 569
		// line internal/php8/scanner.rl:514
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('('))
			lex.cs = 578
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr881:
		// line internal/php8/scanner.rl:511
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st569
	tr882:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:511
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st569
	tr887:
		lex.cs = 569
		// line internal/php8/scanner.rl:515
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr889:
		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st569
	tr892:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st569
	st569:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof569
		}
	st_case_569:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23906
		switch lex.data[(lex.p)] {
		case 10:
			goto tr876
		case 13:
			goto tr877
		case 32:
			goto st570
		case 35:
			goto st573
		case 40:
			goto tr879
		case 47:
			goto tr880
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st570
		}
		goto tr874
	tr883:
		// line internal/php8/scanner.rl:54

		goto st570
	st570:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof570
		}
	st_case_570:
		// line internal/php8/scanner.go:23934
		switch lex.data[(lex.p)] {
		case 10:
			goto tr876
		case 13:
			goto tr877
		case 32:
			goto st570
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st570
		}
		goto tr881
	tr876:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st571
	tr884:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st571
	tr886:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st571
	st571:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof571
		}
	st_case_571:
		// line internal/php8/scanner.go:23992
		switch lex.data[(lex.p)] {
		case 10:
			goto tr884
		case 13:
			goto tr885
		case 32:
			goto tr883
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr883
		}
		goto tr882
	tr877:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st572
	tr885:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st572
	st572:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof572
		}
	st_case_572:
		// line internal/php8/scanner.go:24036
		switch lex.data[(lex.p)] {
		case 10:
			goto tr886
		case 13:
			goto tr885
		case 32:
			goto tr883
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr883
		}
		goto tr882
	st573:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof573
		}
	st_case_573:
		if lex.data[(lex.p)] == 91 {
			goto tr887
		}
		goto st574
	tr893:
		// line internal/php8/scanner.rl:54

		goto st574
	st574:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof574
		}
	st_case_574:
		// line internal/php8/scanner.go:24067
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr890
		case 525:
			goto tr891
		}
		if 512 <= _widec && _widec <= 767 {
			goto st574
		}
		goto tr889
	tr890:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st575
	tr894:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st575
	tr896:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st575
	st575:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof575
		}
	st_case_575:
		// line internal/php8/scanner.go:24161
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr894
		case 525:
			goto tr895
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr893
		}
		goto tr892
	tr891:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st576
	tr895:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st576
	st576:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof576
		}
	st_case_576:
		// line internal/php8/scanner.go:24241
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr896
		case 525:
			goto tr895
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr893
		}
		goto tr892
	tr880:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st577
	st577:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof577
		}
	st_case_577:
		// line internal/php8/scanner.go:24300
		switch lex.data[(lex.p)] {
		case 42:
			goto st126
		case 47:
			goto st574
		}
		goto tr887
	tr190:
		// line internal/php8/scanner.rl:54

		goto st126
	st126:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		// line internal/php8/scanner.go:24317
		switch lex.data[(lex.p)] {
		case 10:
			goto tr187
		case 13:
			goto tr188
		case 42:
			goto st129
		}
		goto st126
	tr187:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st127
	tr191:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st127
	tr194:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st127
	st127:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof127
		}
	st_case_127:
		// line internal/php8/scanner.go:24372
		switch lex.data[(lex.p)] {
		case 10:
			goto tr191
		case 13:
			goto tr192
		case 42:
			goto tr193
		}
		goto tr190
	tr188:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st128
	tr192:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st128
	st128:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof128
		}
	st_case_128:
		// line internal/php8/scanner.go:24413
		switch lex.data[(lex.p)] {
		case 10:
			goto tr194
		case 13:
			goto tr192
		case 42:
			goto tr193
		}
		goto tr190
	tr193:
		// line internal/php8/scanner.rl:54

		goto st129
	st129:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		// line internal/php8/scanner.go:24432
		switch lex.data[(lex.p)] {
		case 10:
			goto tr187
		case 13:
			goto tr188
		case 42:
			goto st129
		case 47:
			goto tr195
		}
		goto st126
	tr196:
		lex.cs = 578
		// line internal/php8/scanner.rl:523
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr206:
		// line internal/php8/scanner.rl:67
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st578
	tr897:
		lex.cs = 578
		// line internal/php8/scanner.rl:523
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr902:
		lex.cs = 578
		// line internal/php8/scanner.rl:522
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(')'))
			lex.cs = 587
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr904:
		// line internal/php8/scanner.rl:519
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st578
	tr905:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:519
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st578
	tr910:
		lex.cs = 578
		// line internal/php8/scanner.rl:523
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr912:
		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st578
	tr915:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st578
	st578:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof578
		}
	st_case_578:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:24530
		switch lex.data[(lex.p)] {
		case 10:
			goto tr899
		case 13:
			goto tr900
		case 32:
			goto st579
		case 35:
			goto st582
		case 41:
			goto tr902
		case 47:
			goto tr903
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st579
		}
		goto tr897
	tr906:
		// line internal/php8/scanner.rl:54

		goto st579
	st579:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof579
		}
	st_case_579:
		// line internal/php8/scanner.go:24558
		switch lex.data[(lex.p)] {
		case 10:
			goto tr899
		case 13:
			goto tr900
		case 32:
			goto st579
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st579
		}
		goto tr904
	tr899:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st580
	tr907:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st580
	tr909:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st580
	st580:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof580
		}
	st_case_580:
		// line internal/php8/scanner.go:24616
		switch lex.data[(lex.p)] {
		case 10:
			goto tr907
		case 13:
			goto tr908
		case 32:
			goto tr906
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr906
		}
		goto tr905
	tr900:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st581
	tr908:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st581
	st581:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof581
		}
	st_case_581:
		// line internal/php8/scanner.go:24660
		switch lex.data[(lex.p)] {
		case 10:
			goto tr909
		case 13:
			goto tr908
		case 32:
			goto tr906
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr906
		}
		goto tr905
	st582:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof582
		}
	st_case_582:
		if lex.data[(lex.p)] == 91 {
			goto tr910
		}
		goto st583
	tr916:
		// line internal/php8/scanner.rl:54

		goto st583
	st583:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof583
		}
	st_case_583:
		// line internal/php8/scanner.go:24691
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr913
		case 525:
			goto tr914
		}
		if 512 <= _widec && _widec <= 767 {
			goto st583
		}
		goto tr912
	tr913:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st584
	tr917:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st584
	tr919:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st584
	st584:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof584
		}
	st_case_584:
		// line internal/php8/scanner.go:24785
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr917
		case 525:
			goto tr918
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr916
		}
		goto tr915
	tr914:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st585
	tr918:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st585
	st585:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof585
		}
	st_case_585:
		// line internal/php8/scanner.go:24865
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr919
		case 525:
			goto tr918
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr916
		}
		goto tr915
	tr903:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st586
	st586:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof586
		}
	st_case_586:
		// line internal/php8/scanner.go:24924
		switch lex.data[(lex.p)] {
		case 42:
			goto st130
		case 47:
			goto st583
		}
		goto tr910
	tr201:
		// line internal/php8/scanner.rl:54

		goto st130
	st130:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		// line internal/php8/scanner.go:24941
		switch lex.data[(lex.p)] {
		case 10:
			goto tr198
		case 13:
			goto tr199
		case 42:
			goto st133
		}
		goto st130
	tr198:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st131
	tr202:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st131
	tr205:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st131
	st131:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		// line internal/php8/scanner.go:24996
		switch lex.data[(lex.p)] {
		case 10:
			goto tr202
		case 13:
			goto tr203
		case 42:
			goto tr204
		}
		goto tr201
	tr199:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st132
	tr203:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st132
	st132:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		// line internal/php8/scanner.go:25037
		switch lex.data[(lex.p)] {
		case 10:
			goto tr205
		case 13:
			goto tr203
		case 42:
			goto tr204
		}
		goto tr201
	tr204:
		// line internal/php8/scanner.rl:54

		goto st133
	st133:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof133
		}
	st_case_133:
		// line internal/php8/scanner.go:25056
		switch lex.data[(lex.p)] {
		case 10:
			goto tr198
		case 13:
			goto tr199
		case 42:
			goto st133
		case 47:
			goto tr206
		}
		goto st130
	tr207:
		lex.cs = 587
		// line internal/php8/scanner.rl:531
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr217:
		// line internal/php8/scanner.rl:67
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st587
	tr920:
		lex.cs = 587
		// line internal/php8/scanner.rl:531
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr926:
		lex.cs = 587
		// line internal/php8/scanner.rl:530
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 596
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr927:
		// line internal/php8/scanner.rl:527
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st587
	tr928:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:527
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st587
	tr933:
		lex.cs = 587
		// line internal/php8/scanner.rl:531
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 151
		}
		goto _again
	tr935:
		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st587
	tr938:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:62
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st587
	st587:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof587
		}
	st_case_587:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:25154
		switch lex.data[(lex.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr923
		case 32:
			goto st588
		case 35:
			goto st591
		case 47:
			goto tr925
		case 59:
			goto tr926
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st588
		}
		goto tr920
	tr929:
		// line internal/php8/scanner.rl:54

		goto st588
	st588:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof588
		}
	st_case_588:
		// line internal/php8/scanner.go:25182
		switch lex.data[(lex.p)] {
		case 10:
			goto tr922
		case 13:
			goto tr923
		case 32:
			goto st588
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st588
		}
		goto tr927
	tr922:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st589
	tr930:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st589
	tr932:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st589
	st589:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof589
		}
	st_case_589:
		// line internal/php8/scanner.go:25240
		switch lex.data[(lex.p)] {
		case 10:
			goto tr930
		case 13:
			goto tr931
		case 32:
			goto tr929
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr929
		}
		goto tr928
	tr923:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st590
	tr931:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st590
	st590:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof590
		}
	st_case_590:
		// line internal/php8/scanner.go:25284
		switch lex.data[(lex.p)] {
		case 10:
			goto tr932
		case 13:
			goto tr931
		case 32:
			goto tr929
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr929
		}
		goto tr928
	st591:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof591
		}
	st_case_591:
		if lex.data[(lex.p)] == 91 {
			goto tr933
		}
		goto st592
	tr939:
		// line internal/php8/scanner.rl:54

		goto st592
	st592:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof592
		}
	st_case_592:
		// line internal/php8/scanner.go:25315
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr936
		case 525:
			goto tr937
		}
		if 512 <= _widec && _widec <= 767 {
			goto st592
		}
		goto tr935
	tr936:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st593
	tr940:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st593
	tr942:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st593
	st593:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof593
		}
	st_case_593:
		// line internal/php8/scanner.go:25409
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr940
		case 525:
			goto tr941
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr939
		}
		goto tr938
	tr937:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st594
	tr941:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st594
	st594:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof594
		}
	st_case_594:
		// line internal/php8/scanner.go:25489
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr942
		case 525:
			goto tr941
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr939
		}
		goto tr938
	tr925:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st595
	st595:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof595
		}
	st_case_595:
		// line internal/php8/scanner.go:25548
		switch lex.data[(lex.p)] {
		case 42:
			goto st134
		case 47:
			goto st592
		}
		goto tr933
	tr212:
		// line internal/php8/scanner.rl:54

		goto st134
	st134:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		// line internal/php8/scanner.go:25565
		switch lex.data[(lex.p)] {
		case 10:
			goto tr209
		case 13:
			goto tr210
		case 42:
			goto st137
		}
		goto st134
	tr209:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st135
	tr213:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st135
	tr216:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st135
	st135:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		// line internal/php8/scanner.go:25620
		switch lex.data[(lex.p)] {
		case 10:
			goto tr213
		case 13:
			goto tr214
		case 42:
			goto tr215
		}
		goto tr212
	tr210:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st136
	tr214:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st136
	st136:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		// line internal/php8/scanner.go:25661
		switch lex.data[(lex.p)] {
		case 10:
			goto tr216
		case 13:
			goto tr214
		case 42:
			goto tr215
		}
		goto tr212
	tr215:
		// line internal/php8/scanner.rl:54

		goto st137
	st137:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof137
		}
	st_case_137:
		// line internal/php8/scanner.go:25680
		switch lex.data[(lex.p)] {
		case 10:
			goto tr209
		case 13:
			goto tr210
		case 42:
			goto st137
		case 47:
			goto tr217
		}
		goto st134
	tr946:
		// line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 208:
			{
				(lex.p) = (lex.te) - 1
				lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
			}
		}

		goto st596
	tr947:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:535
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
		}
		goto st596
	st596:
		// line NONE:1
		lex.ts = 0

		// line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof596
		}
	st_case_596:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:25725
		switch lex.data[(lex.p)] {
		case 10:
			goto tr944
		case 13:
			goto tr945
		}
		goto tr943
	tr943:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:535
		lex.act = 208
		goto st597
	tr948:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:535
		lex.act = 208
		goto st597
	st597:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof597
		}
	st_case_597:
		// line internal/php8/scanner.go:25754
		switch lex.data[(lex.p)] {
		case 10:
			goto tr944
		case 13:
			goto tr945
		}
		goto tr943
	tr944:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st598
	tr949:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st598
	tr951:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		// line internal/php8/scanner.rl:54

		goto st598
	st598:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof598
		}
	st_case_598:
		// line internal/php8/scanner.go:25807
		switch lex.data[(lex.p)] {
		case 10:
			goto tr949
		case 13:
			goto tr950
		}
		goto tr948
	tr945:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st599
	tr950:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && (lex.p+1 == lex.pe || lex.data[lex.p+1] != '\n') {
			lex.newLines.Append(lex.p + 1)
		}

		goto st599
	st599:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof599
		}
	st_case_599:
		// line internal/php8/scanner.go:25846
		switch lex.data[(lex.p)] {
		case 10:
			goto tr951
		case 13:
			goto tr950
		}
		goto tr948
	st_out:
	_test_eof138:
		lex.cs = 138
		goto _test_eof
	_test_eof139:
		lex.cs = 139
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof140:
		lex.cs = 140
		goto _test_eof
	_test_eof141:
		lex.cs = 141
		goto _test_eof
	_test_eof142:
		lex.cs = 142
		goto _test_eof
	_test_eof143:
		lex.cs = 143
		goto _test_eof
	_test_eof144:
		lex.cs = 144
		goto _test_eof
	_test_eof145:
		lex.cs = 145
		goto _test_eof
	_test_eof146:
		lex.cs = 146
		goto _test_eof
	_test_eof147:
		lex.cs = 147
		goto _test_eof
	_test_eof148:
		lex.cs = 148
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof149:
		lex.cs = 149
		goto _test_eof
	_test_eof150:
		lex.cs = 150
		goto _test_eof
	_test_eof151:
		lex.cs = 151
		goto _test_eof
	_test_eof152:
		lex.cs = 152
		goto _test_eof
	_test_eof153:
		lex.cs = 153
		goto _test_eof
	_test_eof154:
		lex.cs = 154
		goto _test_eof
	_test_eof155:
		lex.cs = 155
		goto _test_eof
	_test_eof156:
		lex.cs = 156
		goto _test_eof
	_test_eof157:
		lex.cs = 157
		goto _test_eof
	_test_eof158:
		lex.cs = 158
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof159:
		lex.cs = 159
		goto _test_eof
	_test_eof160:
		lex.cs = 160
		goto _test_eof
	_test_eof161:
		lex.cs = 161
		goto _test_eof
	_test_eof162:
		lex.cs = 162
		goto _test_eof
	_test_eof163:
		lex.cs = 163
		goto _test_eof
	_test_eof164:
		lex.cs = 164
		goto _test_eof
	_test_eof165:
		lex.cs = 165
		goto _test_eof
	_test_eof166:
		lex.cs = 166
		goto _test_eof
	_test_eof167:
		lex.cs = 167
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof168:
		lex.cs = 168
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
	_test_eof31:
		lex.cs = 31
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof
	_test_eof36:
		lex.cs = 36
		goto _test_eof
	_test_eof37:
		lex.cs = 37
		goto _test_eof
	_test_eof38:
		lex.cs = 38
		goto _test_eof
	_test_eof39:
		lex.cs = 39
		goto _test_eof
	_test_eof40:
		lex.cs = 40
		goto _test_eof
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof43:
		lex.cs = 43
		goto _test_eof
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof46:
		lex.cs = 46
		goto _test_eof
	_test_eof47:
		lex.cs = 47
		goto _test_eof
	_test_eof48:
		lex.cs = 48
		goto _test_eof
	_test_eof49:
		lex.cs = 49
		goto _test_eof
	_test_eof50:
		lex.cs = 50
		goto _test_eof
	_test_eof51:
		lex.cs = 51
		goto _test_eof
	_test_eof52:
		lex.cs = 52
		goto _test_eof
	_test_eof53:
		lex.cs = 53
		goto _test_eof
	_test_eof54:
		lex.cs = 54
		goto _test_eof
	_test_eof55:
		lex.cs = 55
		goto _test_eof
	_test_eof56:
		lex.cs = 56
		goto _test_eof
	_test_eof57:
		lex.cs = 57
		goto _test_eof
	_test_eof58:
		lex.cs = 58
		goto _test_eof
	_test_eof59:
		lex.cs = 59
		goto _test_eof
	_test_eof60:
		lex.cs = 60
		goto _test_eof
	_test_eof61:
		lex.cs = 61
		goto _test_eof
	_test_eof62:
		lex.cs = 62
		goto _test_eof
	_test_eof63:
		lex.cs = 63
		goto _test_eof
	_test_eof64:
		lex.cs = 64
		goto _test_eof
	_test_eof65:
		lex.cs = 65
		goto _test_eof
	_test_eof169:
		lex.cs = 169
		goto _test_eof
	_test_eof170:
		lex.cs = 170
		goto _test_eof
	_test_eof171:
		lex.cs = 171
		goto _test_eof
	_test_eof172:
		lex.cs = 172
		goto _test_eof
	_test_eof173:
		lex.cs = 173
		goto _test_eof
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof174:
		lex.cs = 174
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof175:
		lex.cs = 175
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof176:
		lex.cs = 176
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof73:
		lex.cs = 73
		goto _test_eof
	_test_eof177:
		lex.cs = 177
		goto _test_eof
	_test_eof178:
		lex.cs = 178
		goto _test_eof
	_test_eof179:
		lex.cs = 179
		goto _test_eof
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof180:
		lex.cs = 180
		goto _test_eof
	_test_eof76:
		lex.cs = 76
		goto _test_eof
	_test_eof181:
		lex.cs = 181
		goto _test_eof
	_test_eof77:
		lex.cs = 77
		goto _test_eof
	_test_eof182:
		lex.cs = 182
		goto _test_eof
	_test_eof183:
		lex.cs = 183
		goto _test_eof
	_test_eof184:
		lex.cs = 184
		goto _test_eof
	_test_eof78:
		lex.cs = 78
		goto _test_eof
	_test_eof79:
		lex.cs = 79
		goto _test_eof
	_test_eof80:
		lex.cs = 80
		goto _test_eof
	_test_eof81:
		lex.cs = 81
		goto _test_eof
	_test_eof185:
		lex.cs = 185
		goto _test_eof
	_test_eof186:
		lex.cs = 186
		goto _test_eof
	_test_eof187:
		lex.cs = 187
		goto _test_eof
	_test_eof188:
		lex.cs = 188
		goto _test_eof
	_test_eof189:
		lex.cs = 189
		goto _test_eof
	_test_eof82:
		lex.cs = 82
		goto _test_eof
	_test_eof83:
		lex.cs = 83
		goto _test_eof
	_test_eof84:
		lex.cs = 84
		goto _test_eof
	_test_eof85:
		lex.cs = 85
		goto _test_eof
	_test_eof190:
		lex.cs = 190
		goto _test_eof
	_test_eof191:
		lex.cs = 191
		goto _test_eof
	_test_eof86:
		lex.cs = 86
		goto _test_eof
	_test_eof87:
		lex.cs = 87
		goto _test_eof
	_test_eof88:
		lex.cs = 88
		goto _test_eof
	_test_eof192:
		lex.cs = 192
		goto _test_eof
	_test_eof193:
		lex.cs = 193
		goto _test_eof
	_test_eof194:
		lex.cs = 194
		goto _test_eof
	_test_eof195:
		lex.cs = 195
		goto _test_eof
	_test_eof196:
		lex.cs = 196
		goto _test_eof
	_test_eof197:
		lex.cs = 197
		goto _test_eof
	_test_eof89:
		lex.cs = 89
		goto _test_eof
	_test_eof198:
		lex.cs = 198
		goto _test_eof
	_test_eof199:
		lex.cs = 199
		goto _test_eof
	_test_eof200:
		lex.cs = 200
		goto _test_eof
	_test_eof201:
		lex.cs = 201
		goto _test_eof
	_test_eof202:
		lex.cs = 202
		goto _test_eof
	_test_eof203:
		lex.cs = 203
		goto _test_eof
	_test_eof90:
		lex.cs = 90
		goto _test_eof
	_test_eof204:
		lex.cs = 204
		goto _test_eof
	_test_eof205:
		lex.cs = 205
		goto _test_eof
	_test_eof206:
		lex.cs = 206
		goto _test_eof
	_test_eof207:
		lex.cs = 207
		goto _test_eof
	_test_eof208:
		lex.cs = 208
		goto _test_eof
	_test_eof209:
		lex.cs = 209
		goto _test_eof
	_test_eof210:
		lex.cs = 210
		goto _test_eof
	_test_eof211:
		lex.cs = 211
		goto _test_eof
	_test_eof212:
		lex.cs = 212
		goto _test_eof
	_test_eof213:
		lex.cs = 213
		goto _test_eof
	_test_eof214:
		lex.cs = 214
		goto _test_eof
	_test_eof215:
		lex.cs = 215
		goto _test_eof
	_test_eof91:
		lex.cs = 91
		goto _test_eof
	_test_eof92:
		lex.cs = 92
		goto _test_eof
	_test_eof216:
		lex.cs = 216
		goto _test_eof
	_test_eof217:
		lex.cs = 217
		goto _test_eof
	_test_eof218:
		lex.cs = 218
		goto _test_eof
	_test_eof219:
		lex.cs = 219
		goto _test_eof
	_test_eof220:
		lex.cs = 220
		goto _test_eof
	_test_eof221:
		lex.cs = 221
		goto _test_eof
	_test_eof222:
		lex.cs = 222
		goto _test_eof
	_test_eof223:
		lex.cs = 223
		goto _test_eof
	_test_eof224:
		lex.cs = 224
		goto _test_eof
	_test_eof225:
		lex.cs = 225
		goto _test_eof
	_test_eof226:
		lex.cs = 226
		goto _test_eof
	_test_eof227:
		lex.cs = 227
		goto _test_eof
	_test_eof228:
		lex.cs = 228
		goto _test_eof
	_test_eof229:
		lex.cs = 229
		goto _test_eof
	_test_eof230:
		lex.cs = 230
		goto _test_eof
	_test_eof231:
		lex.cs = 231
		goto _test_eof
	_test_eof232:
		lex.cs = 232
		goto _test_eof
	_test_eof233:
		lex.cs = 233
		goto _test_eof
	_test_eof234:
		lex.cs = 234
		goto _test_eof
	_test_eof235:
		lex.cs = 235
		goto _test_eof
	_test_eof236:
		lex.cs = 236
		goto _test_eof
	_test_eof237:
		lex.cs = 237
		goto _test_eof
	_test_eof238:
		lex.cs = 238
		goto _test_eof
	_test_eof239:
		lex.cs = 239
		goto _test_eof
	_test_eof240:
		lex.cs = 240
		goto _test_eof
	_test_eof241:
		lex.cs = 241
		goto _test_eof
	_test_eof242:
		lex.cs = 242
		goto _test_eof
	_test_eof243:
		lex.cs = 243
		goto _test_eof
	_test_eof244:
		lex.cs = 244
		goto _test_eof
	_test_eof245:
		lex.cs = 245
		goto _test_eof
	_test_eof246:
		lex.cs = 246
		goto _test_eof
	_test_eof247:
		lex.cs = 247
		goto _test_eof
	_test_eof248:
		lex.cs = 248
		goto _test_eof
	_test_eof249:
		lex.cs = 249
		goto _test_eof
	_test_eof250:
		lex.cs = 250
		goto _test_eof
	_test_eof251:
		lex.cs = 251
		goto _test_eof
	_test_eof252:
		lex.cs = 252
		goto _test_eof
	_test_eof253:
		lex.cs = 253
		goto _test_eof
	_test_eof254:
		lex.cs = 254
		goto _test_eof
	_test_eof255:
		lex.cs = 255
		goto _test_eof
	_test_eof256:
		lex.cs = 256
		goto _test_eof
	_test_eof257:
		lex.cs = 257
		goto _test_eof
	_test_eof258:
		lex.cs = 258
		goto _test_eof
	_test_eof259:
		lex.cs = 259
		goto _test_eof
	_test_eof260:
		lex.cs = 260
		goto _test_eof
	_test_eof261:
		lex.cs = 261
		goto _test_eof
	_test_eof262:
		lex.cs = 262
		goto _test_eof
	_test_eof263:
		lex.cs = 263
		goto _test_eof
	_test_eof264:
		lex.cs = 264
		goto _test_eof
	_test_eof265:
		lex.cs = 265
		goto _test_eof
	_test_eof266:
		lex.cs = 266
		goto _test_eof
	_test_eof267:
		lex.cs = 267
		goto _test_eof
	_test_eof268:
		lex.cs = 268
		goto _test_eof
	_test_eof269:
		lex.cs = 269
		goto _test_eof
	_test_eof270:
		lex.cs = 270
		goto _test_eof
	_test_eof271:
		lex.cs = 271
		goto _test_eof
	_test_eof272:
		lex.cs = 272
		goto _test_eof
	_test_eof273:
		lex.cs = 273
		goto _test_eof
	_test_eof274:
		lex.cs = 274
		goto _test_eof
	_test_eof275:
		lex.cs = 275
		goto _test_eof
	_test_eof276:
		lex.cs = 276
		goto _test_eof
	_test_eof277:
		lex.cs = 277
		goto _test_eof
	_test_eof278:
		lex.cs = 278
		goto _test_eof
	_test_eof279:
		lex.cs = 279
		goto _test_eof
	_test_eof280:
		lex.cs = 280
		goto _test_eof
	_test_eof281:
		lex.cs = 281
		goto _test_eof
	_test_eof282:
		lex.cs = 282
		goto _test_eof
	_test_eof283:
		lex.cs = 283
		goto _test_eof
	_test_eof284:
		lex.cs = 284
		goto _test_eof
	_test_eof285:
		lex.cs = 285
		goto _test_eof
	_test_eof286:
		lex.cs = 286
		goto _test_eof
	_test_eof287:
		lex.cs = 287
		goto _test_eof
	_test_eof288:
		lex.cs = 288
		goto _test_eof
	_test_eof289:
		lex.cs = 289
		goto _test_eof
	_test_eof290:
		lex.cs = 290
		goto _test_eof
	_test_eof291:
		lex.cs = 291
		goto _test_eof
	_test_eof292:
		lex.cs = 292
		goto _test_eof
	_test_eof293:
		lex.cs = 293
		goto _test_eof
	_test_eof294:
		lex.cs = 294
		goto _test_eof
	_test_eof295:
		lex.cs = 295
		goto _test_eof
	_test_eof296:
		lex.cs = 296
		goto _test_eof
	_test_eof297:
		lex.cs = 297
		goto _test_eof
	_test_eof298:
		lex.cs = 298
		goto _test_eof
	_test_eof299:
		lex.cs = 299
		goto _test_eof
	_test_eof300:
		lex.cs = 300
		goto _test_eof
	_test_eof301:
		lex.cs = 301
		goto _test_eof
	_test_eof302:
		lex.cs = 302
		goto _test_eof
	_test_eof303:
		lex.cs = 303
		goto _test_eof
	_test_eof304:
		lex.cs = 304
		goto _test_eof
	_test_eof305:
		lex.cs = 305
		goto _test_eof
	_test_eof306:
		lex.cs = 306
		goto _test_eof
	_test_eof307:
		lex.cs = 307
		goto _test_eof
	_test_eof308:
		lex.cs = 308
		goto _test_eof
	_test_eof309:
		lex.cs = 309
		goto _test_eof
	_test_eof310:
		lex.cs = 310
		goto _test_eof
	_test_eof311:
		lex.cs = 311
		goto _test_eof
	_test_eof312:
		lex.cs = 312
		goto _test_eof
	_test_eof313:
		lex.cs = 313
		goto _test_eof
	_test_eof314:
		lex.cs = 314
		goto _test_eof
	_test_eof315:
		lex.cs = 315
		goto _test_eof
	_test_eof316:
		lex.cs = 316
		goto _test_eof
	_test_eof317:
		lex.cs = 317
		goto _test_eof
	_test_eof318:
		lex.cs = 318
		goto _test_eof
	_test_eof319:
		lex.cs = 319
		goto _test_eof
	_test_eof320:
		lex.cs = 320
		goto _test_eof
	_test_eof321:
		lex.cs = 321
		goto _test_eof
	_test_eof322:
		lex.cs = 322
		goto _test_eof
	_test_eof323:
		lex.cs = 323
		goto _test_eof
	_test_eof324:
		lex.cs = 324
		goto _test_eof
	_test_eof325:
		lex.cs = 325
		goto _test_eof
	_test_eof326:
		lex.cs = 326
		goto _test_eof
	_test_eof327:
		lex.cs = 327
		goto _test_eof
	_test_eof328:
		lex.cs = 328
		goto _test_eof
	_test_eof329:
		lex.cs = 329
		goto _test_eof
	_test_eof330:
		lex.cs = 330
		goto _test_eof
	_test_eof331:
		lex.cs = 331
		goto _test_eof
	_test_eof332:
		lex.cs = 332
		goto _test_eof
	_test_eof333:
		lex.cs = 333
		goto _test_eof
	_test_eof334:
		lex.cs = 334
		goto _test_eof
	_test_eof335:
		lex.cs = 335
		goto _test_eof
	_test_eof336:
		lex.cs = 336
		goto _test_eof
	_test_eof337:
		lex.cs = 337
		goto _test_eof
	_test_eof338:
		lex.cs = 338
		goto _test_eof
	_test_eof339:
		lex.cs = 339
		goto _test_eof
	_test_eof340:
		lex.cs = 340
		goto _test_eof
	_test_eof341:
		lex.cs = 341
		goto _test_eof
	_test_eof342:
		lex.cs = 342
		goto _test_eof
	_test_eof343:
		lex.cs = 343
		goto _test_eof
	_test_eof344:
		lex.cs = 344
		goto _test_eof
	_test_eof345:
		lex.cs = 345
		goto _test_eof
	_test_eof346:
		lex.cs = 346
		goto _test_eof
	_test_eof347:
		lex.cs = 347
		goto _test_eof
	_test_eof348:
		lex.cs = 348
		goto _test_eof
	_test_eof349:
		lex.cs = 349
		goto _test_eof
	_test_eof350:
		lex.cs = 350
		goto _test_eof
	_test_eof351:
		lex.cs = 351
		goto _test_eof
	_test_eof352:
		lex.cs = 352
		goto _test_eof
	_test_eof353:
		lex.cs = 353
		goto _test_eof
	_test_eof354:
		lex.cs = 354
		goto _test_eof
	_test_eof355:
		lex.cs = 355
		goto _test_eof
	_test_eof356:
		lex.cs = 356
		goto _test_eof
	_test_eof357:
		lex.cs = 357
		goto _test_eof
	_test_eof358:
		lex.cs = 358
		goto _test_eof
	_test_eof359:
		lex.cs = 359
		goto _test_eof
	_test_eof360:
		lex.cs = 360
		goto _test_eof
	_test_eof361:
		lex.cs = 361
		goto _test_eof
	_test_eof362:
		lex.cs = 362
		goto _test_eof
	_test_eof363:
		lex.cs = 363
		goto _test_eof
	_test_eof364:
		lex.cs = 364
		goto _test_eof
	_test_eof365:
		lex.cs = 365
		goto _test_eof
	_test_eof366:
		lex.cs = 366
		goto _test_eof
	_test_eof367:
		lex.cs = 367
		goto _test_eof
	_test_eof368:
		lex.cs = 368
		goto _test_eof
	_test_eof369:
		lex.cs = 369
		goto _test_eof
	_test_eof370:
		lex.cs = 370
		goto _test_eof
	_test_eof371:
		lex.cs = 371
		goto _test_eof
	_test_eof372:
		lex.cs = 372
		goto _test_eof
	_test_eof373:
		lex.cs = 373
		goto _test_eof
	_test_eof93:
		lex.cs = 93
		goto _test_eof
	_test_eof374:
		lex.cs = 374
		goto _test_eof
	_test_eof375:
		lex.cs = 375
		goto _test_eof
	_test_eof376:
		lex.cs = 376
		goto _test_eof
	_test_eof377:
		lex.cs = 377
		goto _test_eof
	_test_eof378:
		lex.cs = 378
		goto _test_eof
	_test_eof379:
		lex.cs = 379
		goto _test_eof
	_test_eof380:
		lex.cs = 380
		goto _test_eof
	_test_eof381:
		lex.cs = 381
		goto _test_eof
	_test_eof382:
		lex.cs = 382
		goto _test_eof
	_test_eof383:
		lex.cs = 383
		goto _test_eof
	_test_eof384:
		lex.cs = 384
		goto _test_eof
	_test_eof94:
		lex.cs = 94
		goto _test_eof
	_test_eof95:
		lex.cs = 95
		goto _test_eof
	_test_eof96:
		lex.cs = 96
		goto _test_eof
	_test_eof97:
		lex.cs = 97
		goto _test_eof
	_test_eof385:
		lex.cs = 385
		goto _test_eof
	_test_eof386:
		lex.cs = 386
		goto _test_eof
	_test_eof387:
		lex.cs = 387
		goto _test_eof
	_test_eof388:
		lex.cs = 388
		goto _test_eof
	_test_eof389:
		lex.cs = 389
		goto _test_eof
	_test_eof390:
		lex.cs = 390
		goto _test_eof
	_test_eof391:
		lex.cs = 391
		goto _test_eof
	_test_eof98:
		lex.cs = 98
		goto _test_eof
	_test_eof99:
		lex.cs = 99
		goto _test_eof
	_test_eof100:
		lex.cs = 100
		goto _test_eof
	_test_eof101:
		lex.cs = 101
		goto _test_eof
	_test_eof392:
		lex.cs = 392
		goto _test_eof
	_test_eof393:
		lex.cs = 393
		goto _test_eof
	_test_eof394:
		lex.cs = 394
		goto _test_eof
	_test_eof395:
		lex.cs = 395
		goto _test_eof
	_test_eof396:
		lex.cs = 396
		goto _test_eof
	_test_eof102:
		lex.cs = 102
		goto _test_eof
	_test_eof103:
		lex.cs = 103
		goto _test_eof
	_test_eof104:
		lex.cs = 104
		goto _test_eof
	_test_eof105:
		lex.cs = 105
		goto _test_eof
	_test_eof397:
		lex.cs = 397
		goto _test_eof
	_test_eof398:
		lex.cs = 398
		goto _test_eof
	_test_eof399:
		lex.cs = 399
		goto _test_eof
	_test_eof400:
		lex.cs = 400
		goto _test_eof
	_test_eof401:
		lex.cs = 401
		goto _test_eof
	_test_eof402:
		lex.cs = 402
		goto _test_eof
	_test_eof403:
		lex.cs = 403
		goto _test_eof
	_test_eof404:
		lex.cs = 404
		goto _test_eof
	_test_eof405:
		lex.cs = 405
		goto _test_eof
	_test_eof406:
		lex.cs = 406
		goto _test_eof
	_test_eof407:
		lex.cs = 407
		goto _test_eof
	_test_eof408:
		lex.cs = 408
		goto _test_eof
	_test_eof409:
		lex.cs = 409
		goto _test_eof
	_test_eof410:
		lex.cs = 410
		goto _test_eof
	_test_eof411:
		lex.cs = 411
		goto _test_eof
	_test_eof412:
		lex.cs = 412
		goto _test_eof
	_test_eof413:
		lex.cs = 413
		goto _test_eof
	_test_eof414:
		lex.cs = 414
		goto _test_eof
	_test_eof415:
		lex.cs = 415
		goto _test_eof
	_test_eof416:
		lex.cs = 416
		goto _test_eof
	_test_eof417:
		lex.cs = 417
		goto _test_eof
	_test_eof418:
		lex.cs = 418
		goto _test_eof
	_test_eof419:
		lex.cs = 419
		goto _test_eof
	_test_eof420:
		lex.cs = 420
		goto _test_eof
	_test_eof421:
		lex.cs = 421
		goto _test_eof
	_test_eof422:
		lex.cs = 422
		goto _test_eof
	_test_eof423:
		lex.cs = 423
		goto _test_eof
	_test_eof424:
		lex.cs = 424
		goto _test_eof
	_test_eof425:
		lex.cs = 425
		goto _test_eof
	_test_eof426:
		lex.cs = 426
		goto _test_eof
	_test_eof427:
		lex.cs = 427
		goto _test_eof
	_test_eof428:
		lex.cs = 428
		goto _test_eof
	_test_eof429:
		lex.cs = 429
		goto _test_eof
	_test_eof430:
		lex.cs = 430
		goto _test_eof
	_test_eof431:
		lex.cs = 431
		goto _test_eof
	_test_eof432:
		lex.cs = 432
		goto _test_eof
	_test_eof433:
		lex.cs = 433
		goto _test_eof
	_test_eof434:
		lex.cs = 434
		goto _test_eof
	_test_eof435:
		lex.cs = 435
		goto _test_eof
	_test_eof436:
		lex.cs = 436
		goto _test_eof
	_test_eof437:
		lex.cs = 437
		goto _test_eof
	_test_eof438:
		lex.cs = 438
		goto _test_eof
	_test_eof439:
		lex.cs = 439
		goto _test_eof
	_test_eof440:
		lex.cs = 440
		goto _test_eof
	_test_eof441:
		lex.cs = 441
		goto _test_eof
	_test_eof442:
		lex.cs = 442
		goto _test_eof
	_test_eof443:
		lex.cs = 443
		goto _test_eof
	_test_eof444:
		lex.cs = 444
		goto _test_eof
	_test_eof445:
		lex.cs = 445
		goto _test_eof
	_test_eof446:
		lex.cs = 446
		goto _test_eof
	_test_eof447:
		lex.cs = 447
		goto _test_eof
	_test_eof448:
		lex.cs = 448
		goto _test_eof
	_test_eof449:
		lex.cs = 449
		goto _test_eof
	_test_eof106:
		lex.cs = 106
		goto _test_eof
	_test_eof107:
		lex.cs = 107
		goto _test_eof
	_test_eof108:
		lex.cs = 108
		goto _test_eof
	_test_eof109:
		lex.cs = 109
		goto _test_eof
	_test_eof110:
		lex.cs = 110
		goto _test_eof
	_test_eof111:
		lex.cs = 111
		goto _test_eof
	_test_eof450:
		lex.cs = 450
		goto _test_eof
	_test_eof451:
		lex.cs = 451
		goto _test_eof
	_test_eof112:
		lex.cs = 112
		goto _test_eof
	_test_eof452:
		lex.cs = 452
		goto _test_eof
	_test_eof453:
		lex.cs = 453
		goto _test_eof
	_test_eof454:
		lex.cs = 454
		goto _test_eof
	_test_eof455:
		lex.cs = 455
		goto _test_eof
	_test_eof456:
		lex.cs = 456
		goto _test_eof
	_test_eof457:
		lex.cs = 457
		goto _test_eof
	_test_eof458:
		lex.cs = 458
		goto _test_eof
	_test_eof459:
		lex.cs = 459
		goto _test_eof
	_test_eof460:
		lex.cs = 460
		goto _test_eof
	_test_eof461:
		lex.cs = 461
		goto _test_eof
	_test_eof462:
		lex.cs = 462
		goto _test_eof
	_test_eof463:
		lex.cs = 463
		goto _test_eof
	_test_eof464:
		lex.cs = 464
		goto _test_eof
	_test_eof465:
		lex.cs = 465
		goto _test_eof
	_test_eof466:
		lex.cs = 466
		goto _test_eof
	_test_eof467:
		lex.cs = 467
		goto _test_eof
	_test_eof468:
		lex.cs = 468
		goto _test_eof
	_test_eof469:
		lex.cs = 469
		goto _test_eof
	_test_eof470:
		lex.cs = 470
		goto _test_eof
	_test_eof471:
		lex.cs = 471
		goto _test_eof
	_test_eof472:
		lex.cs = 472
		goto _test_eof
	_test_eof473:
		lex.cs = 473
		goto _test_eof
	_test_eof474:
		lex.cs = 474
		goto _test_eof
	_test_eof475:
		lex.cs = 475
		goto _test_eof
	_test_eof476:
		lex.cs = 476
		goto _test_eof
	_test_eof477:
		lex.cs = 477
		goto _test_eof
	_test_eof478:
		lex.cs = 478
		goto _test_eof
	_test_eof479:
		lex.cs = 479
		goto _test_eof
	_test_eof480:
		lex.cs = 480
		goto _test_eof
	_test_eof481:
		lex.cs = 481
		goto _test_eof
	_test_eof482:
		lex.cs = 482
		goto _test_eof
	_test_eof483:
		lex.cs = 483
		goto _test_eof
	_test_eof484:
		lex.cs = 484
		goto _test_eof
	_test_eof485:
		lex.cs = 485
		goto _test_eof
	_test_eof486:
		lex.cs = 486
		goto _test_eof
	_test_eof487:
		lex.cs = 487
		goto _test_eof
	_test_eof488:
		lex.cs = 488
		goto _test_eof
	_test_eof489:
		lex.cs = 489
		goto _test_eof
	_test_eof490:
		lex.cs = 490
		goto _test_eof
	_test_eof491:
		lex.cs = 491
		goto _test_eof
	_test_eof492:
		lex.cs = 492
		goto _test_eof
	_test_eof493:
		lex.cs = 493
		goto _test_eof
	_test_eof494:
		lex.cs = 494
		goto _test_eof
	_test_eof495:
		lex.cs = 495
		goto _test_eof
	_test_eof496:
		lex.cs = 496
		goto _test_eof
	_test_eof497:
		lex.cs = 497
		goto _test_eof
	_test_eof498:
		lex.cs = 498
		goto _test_eof
	_test_eof499:
		lex.cs = 499
		goto _test_eof
	_test_eof500:
		lex.cs = 500
		goto _test_eof
	_test_eof501:
		lex.cs = 501
		goto _test_eof
	_test_eof502:
		lex.cs = 502
		goto _test_eof
	_test_eof503:
		lex.cs = 503
		goto _test_eof
	_test_eof504:
		lex.cs = 504
		goto _test_eof
	_test_eof505:
		lex.cs = 505
		goto _test_eof
	_test_eof506:
		lex.cs = 506
		goto _test_eof
	_test_eof507:
		lex.cs = 507
		goto _test_eof
	_test_eof508:
		lex.cs = 508
		goto _test_eof
	_test_eof509:
		lex.cs = 509
		goto _test_eof
	_test_eof510:
		lex.cs = 510
		goto _test_eof
	_test_eof511:
		lex.cs = 511
		goto _test_eof
	_test_eof512:
		lex.cs = 512
		goto _test_eof
	_test_eof513:
		lex.cs = 513
		goto _test_eof
	_test_eof514:
		lex.cs = 514
		goto _test_eof
	_test_eof515:
		lex.cs = 515
		goto _test_eof
	_test_eof516:
		lex.cs = 516
		goto _test_eof
	_test_eof517:
		lex.cs = 517
		goto _test_eof
	_test_eof518:
		lex.cs = 518
		goto _test_eof
	_test_eof519:
		lex.cs = 519
		goto _test_eof
	_test_eof520:
		lex.cs = 520
		goto _test_eof
	_test_eof521:
		lex.cs = 521
		goto _test_eof
	_test_eof522:
		lex.cs = 522
		goto _test_eof
	_test_eof523:
		lex.cs = 523
		goto _test_eof
	_test_eof524:
		lex.cs = 524
		goto _test_eof
	_test_eof113:
		lex.cs = 113
		goto _test_eof
	_test_eof525:
		lex.cs = 525
		goto _test_eof
	_test_eof526:
		lex.cs = 526
		goto _test_eof
	_test_eof527:
		lex.cs = 527
		goto _test_eof
	_test_eof528:
		lex.cs = 528
		goto _test_eof
	_test_eof529:
		lex.cs = 529
		goto _test_eof
	_test_eof530:
		lex.cs = 530
		goto _test_eof
	_test_eof531:
		lex.cs = 531
		goto _test_eof
	_test_eof114:
		lex.cs = 114
		goto _test_eof
	_test_eof532:
		lex.cs = 532
		goto _test_eof
	_test_eof533:
		lex.cs = 533
		goto _test_eof
	_test_eof534:
		lex.cs = 534
		goto _test_eof
	_test_eof535:
		lex.cs = 535
		goto _test_eof
	_test_eof536:
		lex.cs = 536
		goto _test_eof
	_test_eof537:
		lex.cs = 537
		goto _test_eof
	_test_eof115:
		lex.cs = 115
		goto _test_eof
	_test_eof116:
		lex.cs = 116
		goto _test_eof
	_test_eof538:
		lex.cs = 538
		goto _test_eof
	_test_eof539:
		lex.cs = 539
		goto _test_eof
	_test_eof540:
		lex.cs = 540
		goto _test_eof
	_test_eof541:
		lex.cs = 541
		goto _test_eof
	_test_eof542:
		lex.cs = 542
		goto _test_eof
	_test_eof543:
		lex.cs = 543
		goto _test_eof
	_test_eof117:
		lex.cs = 117
		goto _test_eof
	_test_eof118:
		lex.cs = 118
		goto _test_eof
	_test_eof544:
		lex.cs = 544
		goto _test_eof
	_test_eof545:
		lex.cs = 545
		goto _test_eof
	_test_eof546:
		lex.cs = 546
		goto _test_eof
	_test_eof547:
		lex.cs = 547
		goto _test_eof
	_test_eof548:
		lex.cs = 548
		goto _test_eof
	_test_eof549:
		lex.cs = 549
		goto _test_eof
	_test_eof550:
		lex.cs = 550
		goto _test_eof
	_test_eof551:
		lex.cs = 551
		goto _test_eof
	_test_eof552:
		lex.cs = 552
		goto _test_eof
	_test_eof553:
		lex.cs = 553
		goto _test_eof
	_test_eof554:
		lex.cs = 554
		goto _test_eof
	_test_eof119:
		lex.cs = 119
		goto _test_eof
	_test_eof555:
		lex.cs = 555
		goto _test_eof
	_test_eof120:
		lex.cs = 120
		goto _test_eof
	_test_eof121:
		lex.cs = 121
		goto _test_eof
	_test_eof556:
		lex.cs = 556
		goto _test_eof
	_test_eof557:
		lex.cs = 557
		goto _test_eof
	_test_eof558:
		lex.cs = 558
		goto _test_eof
	_test_eof559:
		lex.cs = 559
		goto _test_eof
	_test_eof560:
		lex.cs = 560
		goto _test_eof
	_test_eof561:
		lex.cs = 561
		goto _test_eof
	_test_eof562:
		lex.cs = 562
		goto _test_eof
	_test_eof563:
		lex.cs = 563
		goto _test_eof
	_test_eof122:
		lex.cs = 122
		goto _test_eof
	_test_eof123:
		lex.cs = 123
		goto _test_eof
	_test_eof564:
		lex.cs = 564
		goto _test_eof
	_test_eof124:
		lex.cs = 124
		goto _test_eof
	_test_eof565:
		lex.cs = 565
		goto _test_eof
	_test_eof566:
		lex.cs = 566
		goto _test_eof
	_test_eof567:
		lex.cs = 567
		goto _test_eof
	_test_eof568:
		lex.cs = 568
		goto _test_eof
	_test_eof125:
		lex.cs = 125
		goto _test_eof
	_test_eof569:
		lex.cs = 569
		goto _test_eof
	_test_eof570:
		lex.cs = 570
		goto _test_eof
	_test_eof571:
		lex.cs = 571
		goto _test_eof
	_test_eof572:
		lex.cs = 572
		goto _test_eof
	_test_eof573:
		lex.cs = 573
		goto _test_eof
	_test_eof574:
		lex.cs = 574
		goto _test_eof
	_test_eof575:
		lex.cs = 575
		goto _test_eof
	_test_eof576:
		lex.cs = 576
		goto _test_eof
	_test_eof577:
		lex.cs = 577
		goto _test_eof
	_test_eof126:
		lex.cs = 126
		goto _test_eof
	_test_eof127:
		lex.cs = 127
		goto _test_eof
	_test_eof128:
		lex.cs = 128
		goto _test_eof
	_test_eof129:
		lex.cs = 129
		goto _test_eof
	_test_eof578:
		lex.cs = 578
		goto _test_eof
	_test_eof579:
		lex.cs = 579
		goto _test_eof
	_test_eof580:
		lex.cs = 580
		goto _test_eof
	_test_eof581:
		lex.cs = 581
		goto _test_eof
	_test_eof582:
		lex.cs = 582
		goto _test_eof
	_test_eof583:
		lex.cs = 583
		goto _test_eof
	_test_eof584:
		lex.cs = 584
		goto _test_eof
	_test_eof585:
		lex.cs = 585
		goto _test_eof
	_test_eof586:
		lex.cs = 586
		goto _test_eof
	_test_eof130:
		lex.cs = 130
		goto _test_eof
	_test_eof131:
		lex.cs = 131
		goto _test_eof
	_test_eof132:
		lex.cs = 132
		goto _test_eof
	_test_eof133:
		lex.cs = 133
		goto _test_eof
	_test_eof587:
		lex.cs = 587
		goto _test_eof
	_test_eof588:
		lex.cs = 588
		goto _test_eof
	_test_eof589:
		lex.cs = 589
		goto _test_eof
	_test_eof590:
		lex.cs = 590
		goto _test_eof
	_test_eof591:
		lex.cs = 591
		goto _test_eof
	_test_eof592:
		lex.cs = 592
		goto _test_eof
	_test_eof593:
		lex.cs = 593
		goto _test_eof
	_test_eof594:
		lex.cs = 594
		goto _test_eof
	_test_eof595:
		lex.cs = 595
		goto _test_eof
	_test_eof134:
		lex.cs = 134
		goto _test_eof
	_test_eof135:
		lex.cs = 135
		goto _test_eof
	_test_eof136:
		lex.cs = 136
		goto _test_eof
	_test_eof137:
		lex.cs = 137
		goto _test_eof
	_test_eof596:
		lex.cs = 596
		goto _test_eof
	_test_eof597:
		lex.cs = 597
		goto _test_eof
	_test_eof598:
		lex.cs = 598
		goto _test_eof
	_test_eof599:
		lex.cs = 599
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 139:
				goto tr220
			case 1:
				goto tr0
			case 140:
				goto tr221
			case 141:
				goto tr221
			case 143:
				goto tr226
			case 144:
				goto tr228
			case 145:
				goto tr228
			case 146:
				goto tr226
			case 147:
				goto tr226
			case 148:
				goto tr235
			case 2:
				goto tr4
			case 3:
				goto tr4
			case 4:
				goto tr4
			case 149:
				goto tr238
			case 150:
				goto tr238
			case 152:
				goto tr293
			case 153:
				goto tr295
			case 154:
				goto tr295
			case 155:
				goto tr295
			case 156:
				goto tr300
			case 157:
				goto tr302
			case 158:
				goto tr304
			case 5:
				goto tr10
			case 6:
				goto tr10
			case 7:
				goto tr10
			case 8:
				goto tr10
			case 159:
				goto tr305
			case 160:
				goto tr308
			case 161:
				goto tr311
			case 162:
				goto tr311
			case 163:
				goto tr300
			case 164:
				goto tr317
			case 165:
				goto tr300
			case 166:
				goto tr300
			case 167:
				goto tr321
			case 9:
				goto tr17
			case 10:
				goto tr17
			case 168:
				goto tr300
			case 11:
				goto tr21
			case 12:
				goto tr21
			case 13:
				goto tr21
			case 14:
				goto tr21
			case 15:
				goto tr21
			case 16:
				goto tr21
			case 17:
				goto tr21
			case 18:
				goto tr21
			case 19:
				goto tr21
			case 20:
				goto tr21
			case 21:
				goto tr21
			case 22:
				goto tr21
			case 23:
				goto tr21
			case 24:
				goto tr21
			case 25:
				goto tr21
			case 26:
				goto tr21
			case 27:
				goto tr21
			case 28:
				goto tr21
			case 29:
				goto tr21
			case 30:
				goto tr21
			case 31:
				goto tr21
			case 32:
				goto tr21
			case 33:
				goto tr21
			case 34:
				goto tr21
			case 35:
				goto tr21
			case 36:
				goto tr21
			case 37:
				goto tr21
			case 38:
				goto tr21
			case 39:
				goto tr21
			case 40:
				goto tr21
			case 41:
				goto tr21
			case 42:
				goto tr21
			case 43:
				goto tr21
			case 44:
				goto tr21
			case 45:
				goto tr21
			case 46:
				goto tr21
			case 47:
				goto tr21
			case 48:
				goto tr21
			case 49:
				goto tr21
			case 50:
				goto tr21
			case 51:
				goto tr21
			case 52:
				goto tr21
			case 53:
				goto tr21
			case 54:
				goto tr21
			case 55:
				goto tr21
			case 56:
				goto tr21
			case 57:
				goto tr21
			case 58:
				goto tr21
			case 59:
				goto tr21
			case 60:
				goto tr21
			case 61:
				goto tr21
			case 62:
				goto tr21
			case 63:
				goto tr21
			case 64:
				goto tr21
			case 65:
				goto tr21
			case 169:
				goto tr300
			case 170:
				goto tr324
			case 171:
				goto tr300
			case 172:
				goto tr300
			case 173:
				goto tr300
			case 66:
				goto tr21
			case 174:
				goto tr333
			case 67:
				goto tr10
			case 68:
				goto tr10
			case 175:
				goto tr333
			case 69:
				goto tr88
			case 176:
				goto tr300
			case 70:
				goto tr21
			case 71:
				goto tr21
			case 72:
				goto tr21
			case 73:
				goto tr21
			case 177:
				goto tr337
			case 178:
				goto tr333
			case 179:
				goto tr337
			case 74:
				goto tr100
			case 75:
				goto tr10
			case 180:
				goto tr343
			case 76:
				goto tr10
			case 181:
				goto tr344
			case 77:
				goto tr10
			case 182:
				goto tr345
			case 183:
				goto tr300
			case 184:
				goto tr300
			case 78:
				goto tr21
			case 79:
				goto tr21
			case 80:
				goto tr21
			case 81:
				goto tr21
			case 185:
				goto tr347
			case 186:
				goto tr350
			case 187:
				goto tr350
			case 188:
				goto tr300
			case 189:
				goto tr354
			case 82:
				goto tr10
			case 83:
				goto tr10
			case 84:
				goto tr10
			case 85:
				goto tr10
			case 190:
				goto tr356
			case 191:
				goto tr356
			case 86:
				goto tr10
			case 87:
				goto tr10
			case 88:
				goto tr10
			case 192:
				goto tr357
			case 193:
				goto tr300
			case 194:
				goto tr361
			case 195:
				goto tr300
			case 196:
				goto tr365
			case 197:
				goto tr300
			case 89:
				goto tr21
			case 198:
				goto tr370
			case 199:
				goto tr373
			case 200:
				goto tr373
			case 201:
				goto tr374
			case 202:
				goto tr376
			case 203:
				goto tr10
			case 90:
				goto tr10
			case 204:
				goto tr382
			case 205:
				goto tr376
			case 206:
				goto tr376
			case 207:
				goto tr376
			case 208:
				goto tr376
			case 209:
				goto tr376
			case 210:
				goto tr376
			case 211:
				goto tr376
			case 212:
				goto tr376
			case 213:
				goto tr376
			case 214:
				goto tr376
			case 215:
				goto tr376
			case 91:
				goto tr131
			case 92:
				goto tr131
			case 216:
				goto tr376
			case 217:
				goto tr376
			case 218:
				goto tr376
			case 219:
				goto tr376
			case 220:
				goto tr376
			case 221:
				goto tr376
			case 222:
				goto tr376
			case 223:
				goto tr376
			case 224:
				goto tr376
			case 225:
				goto tr376
			case 226:
				goto tr376
			case 227:
				goto tr376
			case 228:
				goto tr376
			case 229:
				goto tr376
			case 230:
				goto tr376
			case 231:
				goto tr376
			case 232:
				goto tr376
			case 233:
				goto tr376
			case 234:
				goto tr376
			case 235:
				goto tr376
			case 236:
				goto tr376
			case 237:
				goto tr376
			case 238:
				goto tr376
			case 239:
				goto tr376
			case 240:
				goto tr376
			case 241:
				goto tr376
			case 242:
				goto tr376
			case 243:
				goto tr376
			case 244:
				goto tr376
			case 245:
				goto tr376
			case 246:
				goto tr376
			case 247:
				goto tr376
			case 248:
				goto tr376
			case 249:
				goto tr376
			case 250:
				goto tr376
			case 251:
				goto tr376
			case 252:
				goto tr376
			case 253:
				goto tr376
			case 254:
				goto tr376
			case 255:
				goto tr376
			case 256:
				goto tr376
			case 257:
				goto tr376
			case 258:
				goto tr376
			case 259:
				goto tr376
			case 260:
				goto tr376
			case 261:
				goto tr376
			case 262:
				goto tr376
			case 263:
				goto tr376
			case 264:
				goto tr458
			case 265:
				goto tr376
			case 266:
				goto tr376
			case 267:
				goto tr376
			case 268:
				goto tr376
			case 269:
				goto tr376
			case 270:
				goto tr376
			case 271:
				goto tr376
			case 272:
				goto tr376
			case 273:
				goto tr376
			case 274:
				goto tr376
			case 275:
				goto tr376
			case 276:
				goto tr376
			case 277:
				goto tr376
			case 278:
				goto tr376
			case 279:
				goto tr478
			case 280:
				goto tr376
			case 281:
				goto tr376
			case 282:
				goto tr376
			case 283:
				goto tr376
			case 284:
				goto tr376
			case 285:
				goto tr376
			case 286:
				goto tr376
			case 287:
				goto tr376
			case 288:
				goto tr376
			case 289:
				goto tr376
			case 290:
				goto tr376
			case 291:
				goto tr376
			case 292:
				goto tr376
			case 293:
				goto tr376
			case 294:
				goto tr376
			case 295:
				goto tr376
			case 296:
				goto tr376
			case 297:
				goto tr376
			case 298:
				goto tr376
			case 299:
				goto tr376
			case 300:
				goto tr376
			case 301:
				goto tr376
			case 302:
				goto tr376
			case 303:
				goto tr376
			case 304:
				goto tr376
			case 305:
				goto tr507
			case 306:
				goto tr376
			case 307:
				goto tr376
			case 308:
				goto tr511
			case 309:
				goto tr376
			case 310:
				goto tr376
			case 311:
				goto tr376
			case 312:
				goto tr376
			case 313:
				goto tr376
			case 314:
				goto tr376
			case 315:
				goto tr376
			case 316:
				goto tr376
			case 317:
				goto tr376
			case 318:
				goto tr376
			case 319:
				goto tr376
			case 320:
				goto tr376
			case 321:
				goto tr376
			case 322:
				goto tr376
			case 323:
				goto tr376
			case 324:
				goto tr376
			case 325:
				goto tr376
			case 326:
				goto tr376
			case 327:
				goto tr376
			case 328:
				goto tr376
			case 329:
				goto tr376
			case 330:
				goto tr376
			case 331:
				goto tr376
			case 332:
				goto tr376
			case 333:
				goto tr543
			case 334:
				goto tr376
			case 335:
				goto tr376
			case 336:
				goto tr376
			case 337:
				goto tr376
			case 338:
				goto tr376
			case 339:
				goto tr376
			case 340:
				goto tr376
			case 341:
				goto tr376
			case 342:
				goto tr376
			case 343:
				goto tr376
			case 344:
				goto tr376
			case 345:
				goto tr376
			case 346:
				goto tr376
			case 347:
				goto tr376
			case 348:
				goto tr376
			case 349:
				goto tr376
			case 350:
				goto tr376
			case 351:
				goto tr376
			case 352:
				goto tr376
			case 353:
				goto tr376
			case 354:
				goto tr376
			case 355:
				goto tr376
			case 356:
				goto tr376
			case 357:
				goto tr376
			case 358:
				goto tr376
			case 359:
				goto tr376
			case 360:
				goto tr376
			case 361:
				goto tr376
			case 362:
				goto tr376
			case 363:
				goto tr376
			case 364:
				goto tr376
			case 365:
				goto tr376
			case 366:
				goto tr376
			case 367:
				goto tr376
			case 368:
				goto tr376
			case 369:
				goto tr376
			case 370:
				goto tr376
			case 371:
				goto tr376
			case 372:
				goto tr376
			case 373:
				goto tr586
			case 93:
				goto tr10
			case 374:
				goto tr588
			case 375:
				goto tr376
			case 376:
				goto tr376
			case 377:
				goto tr376
			case 378:
				goto tr376
			case 379:
				goto tr376
			case 380:
				goto tr376
			case 381:
				goto tr376
			case 382:
				goto tr376
			case 383:
				goto tr376
			case 384:
				goto tr601
			case 94:
				goto tr134
			case 95:
				goto tr134
			case 96:
				goto tr134
			case 97:
				goto tr134
			case 385:
				goto tr376
			case 386:
				goto tr376
			case 387:
				goto tr376
			case 388:
				goto tr376
			case 389:
				goto tr376
			case 390:
				goto tr376
			case 391:
				goto tr609
			case 98:
				goto tr139
			case 99:
				goto tr139
			case 100:
				goto tr139
			case 101:
				goto tr139
			case 392:
				goto tr376
			case 393:
				goto tr376
			case 394:
				goto tr376
			case 395:
				goto tr376
			case 396:
				goto tr615
			case 102:
				goto tr144
			case 103:
				goto tr144
			case 104:
				goto tr144
			case 105:
				goto tr144
			case 397:
				goto tr376
			case 398:
				goto tr376
			case 399:
				goto tr376
			case 400:
				goto tr376
			case 401:
				goto tr376
			case 402:
				goto tr376
			case 403:
				goto tr376
			case 404:
				goto tr376
			case 405:
				goto tr376
			case 406:
				goto tr376
			case 407:
				goto tr376
			case 408:
				goto tr630
			case 409:
				goto tr376
			case 410:
				goto tr376
			case 411:
				goto tr376
			case 412:
				goto tr376
			case 413:
				goto tr376
			case 414:
				goto tr376
			case 415:
				goto tr376
			case 416:
				goto tr376
			case 417:
				goto tr376
			case 418:
				goto tr376
			case 419:
				goto tr376
			case 420:
				goto tr376
			case 421:
				goto tr376
			case 422:
				goto tr376
			case 423:
				goto tr376
			case 424:
				goto tr376
			case 425:
				goto tr376
			case 426:
				goto tr376
			case 427:
				goto tr376
			case 428:
				goto tr376
			case 429:
				goto tr376
			case 430:
				goto tr376
			case 431:
				goto tr376
			case 432:
				goto tr376
			case 433:
				goto tr376
			case 434:
				goto tr376
			case 435:
				goto tr376
			case 436:
				goto tr376
			case 437:
				goto tr376
			case 438:
				goto tr376
			case 439:
				goto tr376
			case 440:
				goto tr376
			case 441:
				goto tr376
			case 442:
				goto tr376
			case 443:
				goto tr376
			case 444:
				goto tr376
			case 445:
				goto tr376
			case 446:
				goto tr376
			case 447:
				goto tr376
			case 448:
				goto tr376
			case 449:
				goto tr676
			case 106:
				goto tr149
			case 107:
				goto tr149
			case 108:
				goto tr149
			case 109:
				goto tr149
			case 110:
				goto tr149
			case 111:
				goto tr149
			case 450:
				goto tr677
			case 451:
				goto tr678
			case 112:
				goto tr162
			case 452:
				goto tr300
			case 453:
				goto tr376
			case 454:
				goto tr376
			case 455:
				goto tr376
			case 456:
				goto tr376
			case 457:
				goto tr376
			case 458:
				goto tr376
			case 459:
				goto tr376
			case 460:
				goto tr376
			case 461:
				goto tr376
			case 462:
				goto tr376
			case 463:
				goto tr376
			case 464:
				goto tr376
			case 465:
				goto tr376
			case 466:
				goto tr376
			case 467:
				goto tr376
			case 468:
				goto tr376
			case 469:
				goto tr376
			case 470:
				goto tr376
			case 471:
				goto tr376
			case 472:
				goto tr376
			case 473:
				goto tr376
			case 474:
				goto tr376
			case 475:
				goto tr376
			case 476:
				goto tr376
			case 477:
				goto tr376
			case 478:
				goto tr376
			case 479:
				goto tr376
			case 480:
				goto tr376
			case 481:
				goto tr376
			case 482:
				goto tr376
			case 483:
				goto tr376
			case 484:
				goto tr376
			case 485:
				goto tr376
			case 486:
				goto tr376
			case 487:
				goto tr376
			case 488:
				goto tr376
			case 489:
				goto tr376
			case 490:
				goto tr376
			case 491:
				goto tr376
			case 492:
				goto tr376
			case 493:
				goto tr376
			case 494:
				goto tr376
			case 495:
				goto tr376
			case 496:
				goto tr376
			case 497:
				goto tr376
			case 498:
				goto tr376
			case 499:
				goto tr376
			case 500:
				goto tr376
			case 501:
				goto tr376
			case 502:
				goto tr376
			case 503:
				goto tr376
			case 504:
				goto tr376
			case 505:
				goto tr376
			case 506:
				goto tr376
			case 507:
				goto tr376
			case 508:
				goto tr376
			case 509:
				goto tr376
			case 510:
				goto tr376
			case 511:
				goto tr376
			case 512:
				goto tr376
			case 513:
				goto tr376
			case 514:
				goto tr376
			case 515:
				goto tr376
			case 516:
				goto tr376
			case 517:
				goto tr376
			case 518:
				goto tr300
			case 520:
				goto tr764
			case 521:
				goto tr765
			case 522:
				goto tr765
			case 523:
				goto tr770
			case 524:
				goto tr770
			case 113:
				goto tr164
			case 525:
				goto tr773
			case 527:
				goto tr777
			case 528:
				goto tr778
			case 529:
				goto tr778
			case 531:
				goto tr790
			case 532:
				goto tr792
			case 533:
				goto tr793
			case 534:
				goto tr793
			case 535:
				goto tr790
			case 536:
				goto tr799
			case 538:
				goto tr810
			case 539:
				goto tr811
			case 540:
				goto tr811
			case 541:
				goto tr816
			case 542:
				goto tr816
			case 544:
				goto tr829
			case 545:
				goto tr830
			case 546:
				goto tr830
			case 547:
				goto tr835
			case 548:
				goto tr835
			case 550:
				goto tr840
			case 552:
				goto tr847
			case 553:
				goto tr849
			case 554:
				goto tr847
			case 119:
				goto tr174
			case 555:
				goto tr847
			case 120:
				goto tr174
			case 121:
				goto tr174
			case 556:
				goto tr852
			case 558:
				goto tr862
			case 559:
				goto tr862
			case 560:
				goto tr863
			case 561:
				goto tr865
			case 562:
				goto tr866
			case 563:
				goto tr866
			case 122:
				goto tr178
			case 123:
				goto tr178
			case 564:
				goto tr866
			case 124:
				goto tr178
			case 565:
				goto tr866
			case 566:
				goto tr870
			case 568:
				goto tr873
			case 125:
				goto tr182
			case 570:
				goto tr881
			case 571:
				goto tr882
			case 572:
				goto tr882
			case 573:
				goto tr887
			case 574:
				goto tr889
			case 575:
				goto tr892
			case 576:
				goto tr892
			case 577:
				goto tr887
			case 126:
				goto tr185
			case 127:
				goto tr185
			case 128:
				goto tr185
			case 129:
				goto tr185
			case 579:
				goto tr904
			case 580:
				goto tr905
			case 581:
				goto tr905
			case 582:
				goto tr910
			case 583:
				goto tr912
			case 584:
				goto tr915
			case 585:
				goto tr915
			case 586:
				goto tr910
			case 130:
				goto tr196
			case 131:
				goto tr196
			case 132:
				goto tr196
			case 133:
				goto tr196
			case 588:
				goto tr927
			case 589:
				goto tr928
			case 590:
				goto tr928
			case 591:
				goto tr933
			case 592:
				goto tr935
			case 593:
				goto tr938
			case 594:
				goto tr938
			case 595:
				goto tr933
			case 134:
				goto tr207
			case 135:
				goto tr207
			case 136:
				goto tr207
			case 137:
				goto tr207
			case 597:
				goto tr946
			case 598:
				goto tr947
			case 599:
				goto tr947
			}
		}

	_out:
		{
		}
	}

	// line internal/php8/scanner.rl:539

	tkn.Value = lex.data[lex.ts:lex.te]
	tkn.ID = token.ID(tok)

	return tkn
}
