package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4130(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair4131<A, B>(val first: A, val second: B)

fun combine4131(left: Int, right: Int): Pair4131<Int, Int> {
    return Pair4131(left + 46, right + 44)
}
fun <T> List<T>.secondOrNull4132(): T? = if (size > 1) this[1] else null

fun render4132(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4133<A, B>(val first: A, val second: B)

fun combine4133(left: Int, right: Int): Pair4133<Int, Int> {
    return Pair4133(left + 60, right + 66)
}
fun <T> List<T>.secondOrNull4134(): T? = if (size > 1) this[1] else null

fun render4134(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4135(): T? = if (size > 1) this[1] else null

fun render4135(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
