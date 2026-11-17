package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15220<A, B>(val first: A, val second: B)

fun combine15220(left: Int, right: Int): Pair15220<Int, Int> {
    return Pair15220(left + 100, right + 49)
}
fun <T, R> List<T>.mapChecked15221(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15222<A, B>(val first: A, val second: B)

fun combine15222(left: Int, right: Int): Pair15222<Int, Int> {
    return Pair15222(left + 17, right + 10)
}
@JvmName("step15223")
fun step15223(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
