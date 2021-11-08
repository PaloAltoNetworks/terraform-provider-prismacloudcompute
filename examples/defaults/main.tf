terraform {
  required_providers {
    prismacloudcompute = {
      source = "PaloAltoNetworks/prismacloudcompute"
      version = "0.0.1"
    }
  }
}

provider "prismacloudcompute" {
  config_file = "creds.json"
}

# These policy resources represent the default values for Prisma Cloud Compute.

resource "prismacloudcompute_ci_image_compliance_policy" "ruleset" {
  rule {
    name        = "Default - alert on critical and high"
    effect      = "alert"
    collections = ["All"]
    conditions {
      compliance_check {
        block = false
        id    = 41
      }
      compliance_check {
        block = false
        id    = 422
      }
      compliance_check {
        block = false
        id    = 424
      }
      compliance_check {
        block = false
        id    = 425
      }
      compliance_check {
        block = false
        id    = 426
      }
      compliance_check {
        block = false
        id    = 448
      }
      compliance_check {
        block = false
        id    = 5041
      }
    }
  }
}

resource "prismacloudcompute_container_compliance_policy" "ruleset" {
  rule {
    name        = "Default - ignore Twistlock components"
    effect      = "alert"
    collections = ["All"]
    conditions {
      compliance_check {
        block = false
        id    = 56
      }
      compliance_check {
        block = false
        id    = 57
      }
      compliance_check {
        block = false
        id    = 422
      }
      compliance_check {
        block = false
        id    = 424
      }
      compliance_check {
        block = false
        id    = 425
      }
      compliance_check {
        block = false
        id    = 426
      }
      compliance_check {
        block = false
        id    = 427
      }
      compliance_check {
        block = false
        id    = 434
      }
      compliance_check {
        block = false
        id    = 435
      }
      compliance_check {
        block = false
        id    = 436
      }
      compliance_check {
        block = false
        id    = 437
      }
      compliance_check {
        block = false
        id    = 439
      }
      compliance_check {
        block = false
        id    = 441
      }
      compliance_check {
        block = false
        id    = 442
      }
      compliance_check {
        block = false
        id    = 443
      }
      compliance_check {
        block = false
        id    = 444
      }
      compliance_check {
        block = false
        id    = 445
      }
      compliance_check {
        block = false
        id    = 446
      }
      compliance_check {
        block = false
        id    = 447
      }
      compliance_check {
        block = false
        id    = 448
      }
      compliance_check {
        block = false
        id    = 451
      }
      compliance_check {
        block = false
        id    = 452
      }
      compliance_check {
        block = false
        id    = 511
      }
      compliance_check {
        block = false
        id    = 516
      }
      compliance_check {
        block = false
        id    = 517
      }
      compliance_check {
        block = false
        id    = 519
      }
      compliance_check {
        block = false
        id    = 524
      }
      compliance_check {
        block = false
        id    = 597
      }
      compliance_check {
        block = false
        id    = 598
      }
      compliance_check {
        block = false
        id    = 5056
      }
      compliance_check {
        block = false
        id    = 5511
      }
      compliance_check {
        block = false
        id    = 5516
      }
      compliance_check {
        block = false
        id    = 5519
      }
      compliance_check {
        block = false
        id    = 5524
      }
      compliance_check {
        block = false
        id    = 100001
      }
      compliance_check {
        block = false
        id    = 100002
      }
      compliance_check {
        block = false
        id    = 100003
      }
      compliance_check {
        block = false
        id    = 100013
      }
    }
  }
  rule {
    name        = "Default - alert on critical and high"
    effect      = "alert"
    collections = ["All"]
    conditions {
      compliance_check {
        block = false

        id = 41
      }
      compliance_check {
        block = false
        id    = 51
      }
      compliance_check {
        block = false
        id    = 52
      }
      compliance_check {
        block = false
        id    = 54
      }
      compliance_check {
        block = false
        id    = 55
      }
      compliance_check {
        block = false
        id    = 56
      }
      compliance_check {
        block = false
        id    = 57
      }
      compliance_check {
        block = false
        id    = 58
      }
      compliance_check {
        block = false
        id    = 59
      }
      compliance_check {
        block = false
        id    = 422
      }
      compliance_check {
        block = false
        id    = 424
      }
      compliance_check {
        block = false
        id    = 425
      }
      compliance_check {
        block = false
        id    = 426
      }
      compliance_check {
        block = false
        id    = 427
      }
      compliance_check {
        block = false
        id    = 434
      }
      compliance_check {
        block = false
        id    = 435
      }
      compliance_check {
        block = false
        id    = 436
      }
      compliance_check {
        block = false
        id    = 437
      }
      compliance_check {
        block = false
        id    = 439
      }
      compliance_check {
        block = false
        id    = 441
      }
      compliance_check {
        block = false
        id    = 442
      }
      compliance_check {
        block = false
        id    = 443
      }
      compliance_check {
        block = false
        id    = 444
      }
      compliance_check {
        block = false
        id    = 445
      }
      compliance_check {
        block = false
        id    = 446
      }
      compliance_check {
        block = false
        id    = 447
      }
      compliance_check {
        block = false
        id    = 448
      }
      compliance_check {
        block = false
        id    = 451
      }
      compliance_check {
        block = false
        id    = 452
      }
      compliance_check {
        block = false
        id    = 510
      }
      compliance_check {
        block = false
        id    = 511
      }
      compliance_check {
        block = false
        id    = 512
      }
      compliance_check {
        block = false
        id    = 514
      }
      compliance_check {
        block = false
        id    = 515
      }
      compliance_check {
        block = false
        id    = 516
      }
      compliance_check {
        block = false
        id    = 517
      }
      compliance_check {
        block = false
        id    = 519
      }
      compliance_check {
        block = false
        id    = 520
      }
      compliance_check {
        block = false
        id    = 521
      }
      compliance_check {
        block = false
        id    = 524
      }
      compliance_check {
        block = false
        id    = 525
      }
      compliance_check {
        block = false
        id    = 528
      }
      compliance_check {
        block = false
        id    = 530
      }
      compliance_check {
        block = false
        id    = 531
      }
      compliance_check {
        block = false
        id    = 597
      }
      compliance_check {
        block = false
        id    = 598
      }
      compliance_check {
        block = false
        id    = 599
      }
      compliance_check {
        block = false
        id    = 5041
      }
      compliance_check {
        block = false
        id    = 5051
      }
      compliance_check {
        block = false
        id    = 5052
      }
      compliance_check {
        block = false
        id    = 5054
      }
      compliance_check {
        block = false
        id    = 5055
      }
      compliance_check {
        block = false
        id    = 5056
      }
      compliance_check {
        block = false
        id    = 5059
      }
      compliance_check {
        block = false
        id    = 5510
      }
      compliance_check {
        block = false
        id    = 5511
      }
      compliance_check {
        block = false
        id    = 5512
      }
      compliance_check {
        block = false
        id    = 5515
      }
      compliance_check {
        block = false
        id    = 5516
      }
      compliance_check {
        block = false
        id    = 5519
      }
      compliance_check {
        block = false
        id    = 5520
      }
      compliance_check {
        block = false
        id    = 5521
      }
      compliance_check {
        block = false
        id    = 5524
      }
      compliance_check {
        block = false
        id    = 5525
      }
      compliance_check {
        block = false
        id    = 5528
      }
      compliance_check {
        block = false
        id    = 5531
      }
      compliance_check {
        block = false
        id    = 100001
      }
      compliance_check {
        block = false
        id    = 100002
      }
      compliance_check {
        block = false
        id    = 100003
      }
      compliance_check {
        block = false
        id    = 100013
      }
    }
  }
}

resource "prismacloudcompute_host_compliance_policy" "ruleset" {
  rule {
    name        = "Default - alert on critical and high"
    effect      = "alert, block"
    collections = ["All"]
    conditions {
      compliance_check {
        block = false
        id    = 16
      }
      compliance_check {
        block = false
        id    = 21
      }
      compliance_check {
        block = false
        id    = 24
      }
      compliance_check {
        block = false
        id    = 26
      }
      compliance_check {
        block = false
        id    = 28
      }
      compliance_check {
        block = false
        id    = 31
      }
      compliance_check {
        block = false
        id    = 32
      }
      compliance_check {
        block = false
        id    = 33
      }
      compliance_check {
        block = false
        id    = 34
      }
      compliance_check {
        block = false
        id    = 35
      }
      compliance_check {
        block = false
        id    = 36
      }
      compliance_check {
        block = false
        id    = 37
      }
      compliance_check {
        block = false
        id    = 38
      }
      compliance_check {
        block = false
        id    = 39
      }
      compliance_check {
        block = false
        id    = 211
      }
      compliance_check {
        block = false
        id    = 213
      }
      compliance_check {
        block = false
        id    = 215
      }
      compliance_check {
        block = false
        id    = 217
      }
      compliance_check {
        block = false
        id    = 219
      }
      compliance_check {
        block = false
        id    = 221
      }
      compliance_check {
        block = false
        id    = 224
      }
      compliance_check {
        block = false
        id    = 311
      }
      compliance_check {
        block = false
        id    = 313
      }
      compliance_check {
        block = false
        id    = 315
      }
      compliance_check {
        block = false
        id    = 316
      }
      compliance_check {
        block = false
        id    = 317
      }
      compliance_check {
        block = false
        id    = 318
      }
      compliance_check {
        block = false
        id    = 319
      }
      compliance_check {
        block = false
        id    = 320
      }
      compliance_check {
        block = false
        id    = 321
      }
      compliance_check {
        block = false
        id    = 322
      }
      compliance_check {
        block = false
        id    = 449
      }
      compliance_check {
        block = false
        id    = 5024
      }
      compliance_check {
        block = false
        id    = 5026
      }
      compliance_check {
        block = false
        id    = 5031
      }
      compliance_check {
        block = false
        id    = 5032
      }
      compliance_check {
        block = false
        id    = 5035
      }
      compliance_check {
        block = false
        id    = 5036
      }
      compliance_check {
        block = false
        id    = 5037
      }
      compliance_check {
        block = false
        id    = 5038
      }
      compliance_check {
        block = false
        id    = 5315
      }
      compliance_check {
        block = false
        id    = 5316
      }
      compliance_check {
        block = false
        id    = 5317
      }
      compliance_check {
        block = false
        id    = 5318
      }
      compliance_check {
        block = false
        id    = 5319
      }
      compliance_check {
        block = false
        id    = 5320
      }
      compliance_check {
        block = false
        id    = 6112
      }
      compliance_check {
        block = false
        id    = 6141
      }
      compliance_check {
        block = false
        id    = 6143
      }
      compliance_check {
        block = false
        id    = 6151
      }
      compliance_check {
        block = false
        id    = 6152
      }
      compliance_check {
        block = false
        id    = 6153
      }
      compliance_check {
        block = false
        id    = 6216
      }
      compliance_check {
        block = false
        id    = 6218
      }
      compliance_check {
        block = false
        id    = 6219
      }
      compliance_check {
        block = false
        id    = 6223
      }
      compliance_check {
        block = false
        id    = 6224
      }
      compliance_check {
        block = false
        id    = 6225
      }
      compliance_check {
        block = false
        id    = 6226
      }
      compliance_check {
        block = false
        id    = 6227
      }
      compliance_check {
        block = false
        id    = 6228
      }
      compliance_check {
        block = false
        id    = 6229
      }
      compliance_check {
        block = false
        id    = 6321
      }
      compliance_check {
        block = false
        id    = 6344
      }
      compliance_check {
        block = false
        id    = 6345
      }
      compliance_check {
        block = false
        id    = 6361
      }
      compliance_check {
        block = false
        id    = 6412
      }
      compliance_check {
        block = false
        id    = 6518
      }
      compliance_check {
        block = false
        id    = 6521
      }
      compliance_check {
        block = false
        id    = 6522
      }
      compliance_check {
        block = false
        id    = 6525
      }
      compliance_check {
        block = false
        id    = 6528
      }
      compliance_check {
        block = false
        id    = 6529
      }
      compliance_check {
        block = false
        id    = 6612
      }
      compliance_check {
        block = false
        id    = 6613
      }
      compliance_check {
        block = false
        id    = 6614
      }
      compliance_check {
        block = false
        id    = 6615
      }
      compliance_check {
        block = false
        id    = 6616
      }
      compliance_check {
        block = false
        id    = 6617
      }
      compliance_check {
        block = false
        id    = 6618
      }
      compliance_check {
        block = false
        id    = 6619
      }
      compliance_check {
        block = false
        id    = 6621
      }
      compliance_check {
        block = false
        id    = 6625
      }
      compliance_check {
        block = false
        id    = 6627
      }
      compliance_check {
        block = false
        id    = 6628
      }
      compliance_check {
        block = false
        id    = 6629
      }
      compliance_check {
        block = false
        id    = 7162
      }
      compliance_check {
        block = false
        id    = 8111
      }
      compliance_check {
        block = false
        id    = 8112
      }
      compliance_check {
        block = false
        id    = 8114
      }
      compliance_check {
        block = false
        id    = 8115
      }
      compliance_check {
        block = false
        id    = 8116
      }
      compliance_check {
        block = false
        id    = 8117
      }
      compliance_check {
        block = false
        id    = 8118
      }
      compliance_check {
        block = false
        id    = 8133
      }
      compliance_check {
        block = false
        id    = 8134
      }
      compliance_check {
        block = false
        id    = 8135
      }
      compliance_check {
        block = false
        id    = 8136
      }
      compliance_check {
        block = false
        id    = 8141
      }
      compliance_check {
        block = false
        id    = 8142
      }
      compliance_check {
        block = false
        id    = 8143
      }
      compliance_check {
        block = false
        id    = 8144
      }
      compliance_check {
        block = false
        id    = 8145
      }
      compliance_check {
        block = false
        id    = 8146
      }
      compliance_check {
        block = false
        id    = 8147
      }
      compliance_check {
        block = false
        id    = 8148
      }
      compliance_check {
        block = false
        id    = 8149
      }
      compliance_check {
        block = false
        id    = 8151
      }
      compliance_check {
        block = false
        id    = 8152
      }
      compliance_check {
        block = false
        id    = 8153
      }
      compliance_check {
        block = false
        id    = 8154
      }
      compliance_check {
        block = false
        id    = 8155
      }
      compliance_check {
        block = false
        id    = 8156
      }
      compliance_check {
        block = false
        id    = 8211
      }
      compliance_check {
        block = false
        id    = 8212
      }
      compliance_check {
        block = false
        id    = 8213
      }
      compliance_check {
        block = false
        id    = 8214
      }
      compliance_check {
        block = false
        id    = 8215
      }
      compliance_check {
        block = false
        id    = 8223
      }
      compliance_check {
        block = false
        id    = 8224
      }
      compliance_check {
        block = false
        id    = 8225
      }
      compliance_check {
        block = false
        id    = 8226
      }
      compliance_check {
        block = false
        id    = 8227
      }
      compliance_check {
        block = false
        id    = 8228
      }
      compliance_check {
        block = false
        id    = 8229
      }
      compliance_check {
        block = false
        id    = 8230
      }
      compliance_check {
        block = false
        id    = 8231
      }
      compliance_check {
        block = false
        id    = 8232
      }
      compliance_check {
        block = false
        id    = 8233
      }
      compliance_check {
        block = false
        id    = 8234
      }
      compliance_check {
        block = false
        id    = 8311
      }
      compliance_check {
        block = false
        id    = 8313
      }
      compliance_check {
        block = false
        id    = 8314
      }
      compliance_check {
        block = false
        id    = 8315
      }
      compliance_check {
        block = false
        id    = 8316
      }
      compliance_check {
        block = false
        id    = 8318
      }
      compliance_check {
        block = false
        id    = 60522
      }
      compliance_check {
        block = false
        id    = 60523
      }
      compliance_check {
        block = false
        id    = 61611
      }
      compliance_check {
        block = false
        id    = 61612
      }
      compliance_check {
        block = false
        id    = 61613
      }
      compliance_check {
        block = false
        id    = 62110
      }
      compliance_check {
        block = false
        id    = 62210
      }
      compliance_check {
        block = false
        id    = 62211
      }
      compliance_check {
        block = false
        id    = 62212
      }
      compliance_check {
        block = false
        id    = 62213
      }
      compliance_check {
        block = false
        id    = 62214
      }
      compliance_check {
        block = false
        id    = 62216
      }
      compliance_check {
        block = false
        id    = 62217
      }
      compliance_check {
        block = false
        id    = 64115
      }
      compliance_check {
        block = false
        id    = 64116
      }
      compliance_check {
        block = false
        id    = 64117
      }
      compliance_check {
        block = false
        id    = 64118
      }
      compliance_check {
        block = false
        id    = 65414
      }
      compliance_check {
        block = false
        id    = 66210
      }
      compliance_check {
        block = false
        id    = 66213
      }
      compliance_check {
        block = false
        id    = 66215
      }
      compliance_check {
        block = false
        id    = 66216
      }
      compliance_check {
        block = false
        id    = 66217
      }
      compliance_check {
        block = false
        id    = 66218
      }
      compliance_check {
        block = false
        id    = 81111
      }
      compliance_check {
        block = false
        id    = 81112
      }
      compliance_check {
        block = false
        id    = 81120
      }
      compliance_check {
        block = false
        id    = 81122
      }
      compliance_check {
        block = false
        id    = 81123
      }
      compliance_check {
        block = false
        id    = 81126
      }
      compliance_check {
        block = false
        id    = 81127
      }
      compliance_check {
        block = false
        id    = 81128
      }
      compliance_check {
        block = false
        id    = 81129
      }
      compliance_check {
        block = false
        id    = 81130
      }
      compliance_check {
        block = false
        id    = 81131
      }
      compliance_check {
        block = false
        id    = 81132
      }
      compliance_check {
        block = false
        id    = 81133
      }
      compliance_check {
        block = false
        id    = 81137
      }
      compliance_check {
        block = false
        id    = 81138
      }
      compliance_check {
        block = false
        id    = 81410
      }
      compliance_check {
        block = false
        id    = 81411
      }
      compliance_check {
        block = false
        id    = 81412
      }
      compliance_check {
        block = false
        id    = 81413
      }
      compliance_check {
        block = false
        id    = 81414
      }
      compliance_check {
        block = false
        id    = 81417
      }
      compliance_check {
        block = false
        id    = 81418
      }
      compliance_check {
        block = false
        id    = 81419
      }
      compliance_check {
        block = false
        id    = 81420
      }
      compliance_check {
        block = false
        id    = 81421
      }
      compliance_check {
        block = false
        id    = 81422
      }
      compliance_check {
        block = false
        id    = 81423
      }
      compliance_check {
        block = false
        id    = 81424
      }
      compliance_check {
        block = false
        id    = 81425
      }
      compliance_check {
        block = false
        id    = 81426
      }
      compliance_check {
        block = false
        id    = 81427
      }
      compliance_check {
        block = false
        id    = 81428
      }
      compliance_check {
        block = false
        id    = 81429
      }
      compliance_check {
        block = false
        id    = 82112
      }
      compliance_check {
        block = false
        id    = 82113
      }
      compliance_check {
        block = false
        id    = 83114
      }
      compliance_check {
        block = false
        id    = 83117
      }
      compliance_check {
        block = false
        id    = 83118
      }
      compliance_check {
        block = false
        id    = 83119
      }
      compliance_check {
        block = false
        id    = 200001
      }
      compliance_check {
        block = false
        id    = 200002
      }
      compliance_check {
        block = false
        id    = 200004
      }
      compliance_check {
        block = false
        id    = 200005
      }
      compliance_check {
        block = false
        id    = 200006
      }
      compliance_check {
        block = false
        id    = 200007
      }
      compliance_check {
        block = false
        id    = 200201
      }
      compliance_check {
        block = false
        id    = 200202
      }
      compliance_check {
        block = false
        id    = 200203
      }
      compliance_check {
        block = false
        id    = 200400
      }
      compliance_check {
        block = false
        id    = 200401
      }
      compliance_check {
        block = false
        id    = 641113
      }
    }
  }
}

resource "prismacloudcompute_ci_image_vulnerability_policy" "ruleset" {
  rule {
    name        = "Default - alert all components"
    effect      = "alert"
    collections = ["All"]
    alert_threshold = {
      disabled = false
      value    = 0
    }
    block_threshold = {
      enabled = false
      value   = 0
    }
  }
}

resource "prismacloudcompute_host_vulnerability_policy" "ruleset" {
  rule {
    name        = "Default - alert all components"
    effect      = "alert"
    collections = ["All"]
    alert_threshold = {
      disabled = false
      value    = 0
    }
  }
}

resource "prismacloudcompute_image_vulnerability_policy" "ruleset" {
  rule {
    name        = "Default - ignore Twistlock components"
    effect      = "alert"
    collections = ["Prisma Cloud resources"]
    alert_threshold = {
      disabled = false
      value    = 4
    }
    block_threshold = {
      enabled = false
      value   = 0
    }
    cve_rule {
      description = "Not Affected"
      effect      = "ignore"
      expiration  = {
        date = "0001-01-01T00:00:00Z"
        enabled = false
      }
      id          = "CVE-2021-29923"
    }
    cve_rule {
      description = "Not Affected"
      effect      = "ignore"
      expiration  = {
        date = "0001-01-01T00:00:00Z"
        enabled = false
      }
      id          = "CVE-2021-36221"
    }
    only_fixed = true
  }
  rule {
    name        = "Default - alert all components"
    effect      = "alert"
    collections = ["All"]
    alert_threshold = {
      disabled = false
      value    = 0
    }
    block_threshold = {
      enabled = false
      value   = 0
    }
  }
}
