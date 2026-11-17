package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8830(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8831<A, B>(val first: A, val second: B)

fun combine8831(left: Int, right: Int): Pair8831<Int, Int> {
    return Pair8831(left + 45, right + 40)
}
fun <T> List<T>.secondOrNull8832(): T? = if (size > 1) this[1] else null

fun render8832(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
