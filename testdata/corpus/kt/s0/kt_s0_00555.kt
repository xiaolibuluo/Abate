package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4510")
fun step4510(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
fun <T, R> List<T>.mapChecked4511(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step4512")
fun step4512(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
@JvmName("step4513")
fun step4513(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
