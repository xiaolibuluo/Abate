package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23000(): T? = if (size > 1) this[1] else null

fun render23000(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23001")
fun step23001(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
data class Pair23002<A, B>(val first: A, val second: B)

fun combine23002(left: Int, right: Int): Pair23002<Int, Int> {
    return Pair23002(left + 43, right + 27)
}
fun <T, R> List<T>.mapChecked23003(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
