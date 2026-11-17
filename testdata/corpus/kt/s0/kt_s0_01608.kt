package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13270")
fun step13270(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
fun <T, R> List<T>.mapChecked13271(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step13272")
fun step13272(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
@JvmName("step13273")
fun step13273(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
