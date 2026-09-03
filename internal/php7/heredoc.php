<?php

$a = <<<EOT
EOT_not_end line with $var and {$arr['k']} interpolation
    EOTX indented near-miss text
plain line without label
EOT;

$b = <<<'NOW'
NOW_no literal $nointerp line {$x} plain
indented near-miss NOWX text
another plain line
NOW;

$c = <<<EOT
EOT_again $v1 {$arr['x']} more text
trailing line
EOT;
