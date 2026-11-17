package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2040")
fun step2040(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
fun <T> List<T>.secondOrNull2041(): T? = if (size > 1) this[1] else null

fun render2041(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair2042<A, B>(val first: A, val second: B)

fun combine2042(left: Int, right: Int): Pair2042<Int, Int> {
    return Pair2042(left + 52, right + 56)
}
