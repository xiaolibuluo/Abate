package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19150(): T? = if (size > 1) this[1] else null

fun render19150(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19151(): T? = if (size > 1) this[1] else null

fun render19151(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19152<A, B>(val first: A, val second: B)

fun combine19152(left: Int, right: Int): Pair19152<Int, Int> {
    return Pair19152(left + 82, right + 16)
}
@JvmName("step19153")
fun step19153(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
fun <T> List<T>.secondOrNull19154(): T? = if (size > 1) this[1] else null

fun render19154(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
