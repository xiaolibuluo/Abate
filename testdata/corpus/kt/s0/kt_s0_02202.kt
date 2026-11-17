package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18130(): T? = if (size > 1) this[1] else null

fun render18130(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18131")
fun step18131(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
data class Pair18132<A, B>(val first: A, val second: B)

fun combine18132(left: Int, right: Int): Pair18132<Int, Int> {
    return Pair18132(left + 17, right + 34)
}
fun <T, R> List<T>.mapChecked18133(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
