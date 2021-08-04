coverage
========

This project aims to parse data from one of these two programs:

- [Mosdepth](https://github.com/brentp/mosdepth)

        mosdepth --by ${intervals_bed} \
                 --thresholds 5,10,20,30  \
                 ${outut_prefix} ${bam}

- [GATK DiagnoseTagets](https://github.com/broadgsa/gatk-protected/blob/master/protected/gatk-tools-protected/src/main/java/org/broadinstitute/gatk/tools/walkers/diagnostics/diagnosetargets/DiagnoseTargets.java#L101-L108)

            java -jar /usr/gitc/GATK35.jar \
                -T DiagnoseTargets \
                -R ${ref_fasta} \
                -o ${output_name} \
                -I ${input_bam} \
                -L ${intervals} \
                -min ${minCov}

And organize it in a way that one could easily inspect how is the sequencing coverage in genes of interest.


## Quickstart

Download and extract the [proper release](https://github.com/mendelics/coverage/releases) for your system, and then:

```bash
# For GATK DiagnoseTargets output
./coverage --vcf <coverage.vcf.gz> 1>stdout.tsv

# For Mosdepth '.thresholds.bed.gz' output
./coverage --bed <thresholds.bed.gz> 1>stdout.tsv
```

After that you'll have two files:

- stdout.tsv

```tsv
ENSG00000187634 SAMD11  Count: 15   Total Bases: 3246   Coverage 5x: 2831   10x: 2514   20x: 2252   30x: 2137
ENSG00000188976 NOC2L   Count: 19   Total Bases: 2250   Coverage 5x: 2250   10x: 2250   20x: 2250   30x: 2250
ENSG00000187961 KLHL17  Count: 12   Total Bases: 2134   Coverage 5x: 2071   10x: 2027   20x: 2012   30x: 1992
ENSG00000187583 PLEKHN1 Count: 15   Total Bases: 2327   Coverage 5x: 2327   10x: 2327   20x: 2316   30x: 2298
ENSG00000187642 PERM1   Count: 5    Total Bases: 2463   Coverage 5x: 2463   10x: 2463   20x: 2463   30x: 2447
...
```

- *_coverage.json.gz

```json
  "Identifier": "no_sample_id",
  "GitVersion": "",
  "GlobalTotalBases": 36726635,
  "GlobalCoveredBases5x": 36462766,
  "GlobalCoveredBases10x": 36290670,
  "GlobalCoveredBases20x": 35708523,
  "GlobalCoveredBases30x": 34512655,
  "PerGeneCoverage": {
    "ENSG00000000003": {
      "ENSG": "ENSG00000000003",
      "Symbol": "TSPAN6",
      "TotalBases": 753,
      "CoveredBases": 740,
      "BasesCovered5x": 753,
      "BasesCovered10x": 740,
      "BasesCovered20x": 722,
      "BasesCovered30x": 676
    },
    "ENSG00000000005": {
      "ENSG": "ENSG00000000005",
      "Symbol": "TNMD",
      "TotalBases": 954,
      "CoveredBases": 954,
      "BasesCovered5x": 954,
      "BasesCovered10x": 954,
      "BasesCovered20x": 954,
      "BasesCovered30x": 954
    },
    ...
```