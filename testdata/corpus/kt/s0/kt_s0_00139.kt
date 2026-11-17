package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked1120(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair1121<A, B>(val first: A, val second: B)

fun combine1121(left: Int, right: Int): Pair1121<Int, Int> {
    return Pair1121(left + 13, right + 11)
}
@JvmName("step1122")
fun step1122(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
data class Pair1123<A, B>(val first: A, val second: B)

fun combine1123(left: Int, right: Int): Pair1123<Int, Int> {
    return Pair1123(left + 27, right + 33)
}
