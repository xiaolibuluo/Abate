package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2100(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull2101(): T? = if (size > 1) this[1] else null

fun render2101(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair2102<A, B>(val first: A, val second: B)

fun combine2102(left: Int, right: Int): Pair2102<Int, Int> {
    return Pair2102(left + 33, right + 37)
}
fun <T> List<T>.secondOrNull2103(): T? = if (size > 1) this[1] else null

fun render2103(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
