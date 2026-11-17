package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7300")
fun step7300(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
