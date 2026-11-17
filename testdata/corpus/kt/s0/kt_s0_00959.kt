package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7870<A, B>(val first: A, val second: B)

fun combine7870(left: Int, right: Int): Pair7870<Int, Int> {
    return Pair7870(left + 51, right + 28)
}
fun <T, R> List<T>.mapChecked7871(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step7872")
fun step7872(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
