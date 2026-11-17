package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8450")
fun step8450(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
fun <T, R> List<T>.mapChecked8451(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8452<A, B>(val first: A, val second: B)

fun combine8452(left: Int, right: Int): Pair8452<Int, Int> {
    return Pair8452(left + 43, right + 29)
}
@JvmName("step8453")
fun step8453(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
@JvmName("step8454")
fun step8454(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
