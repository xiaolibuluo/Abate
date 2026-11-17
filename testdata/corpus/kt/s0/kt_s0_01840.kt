package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15120")
fun step15120(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
@JvmName("step15121")
fun step15121(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
