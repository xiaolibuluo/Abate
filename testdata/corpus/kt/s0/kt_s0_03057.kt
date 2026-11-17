package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25080(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25081(): T? = if (size > 1) this[1] else null

fun render25081(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25082(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step25083")
fun step25083(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
data class Pair25084<A, B>(val first: A, val second: B)

fun combine25084(left: Int, right: Int): Pair25084<Int, Int> {
    return Pair25084(left + 45, right + 41)
}
data class Pair25085<A, B>(val first: A, val second: B)

fun combine25085(left: Int, right: Int): Pair25085<Int, Int> {
    return Pair25085(left + 52, right + 52)
}
fun <T, R> List<T>.mapChecked25086(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
