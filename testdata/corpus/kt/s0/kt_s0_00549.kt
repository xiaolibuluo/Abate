package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4460(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step4461")
fun step4461(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
data class Pair4462<A, B>(val first: A, val second: B)

fun combine4462(left: Int, right: Int): Pair4462<Int, Int> {
    return Pair4462(left + 94, right + 42)
}
