package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8850(): T? = if (size > 1) this[1] else null

fun render8850(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step8851")
fun step8851(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
data class Pair8852<A, B>(val first: A, val second: B)

fun combine8852(left: Int, right: Int): Pair8852<Int, Int> {
    return Pair8852(left + 78, right + 65)
}
