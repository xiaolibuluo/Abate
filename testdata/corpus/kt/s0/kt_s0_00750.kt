package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6190")
fun step6190(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
fun <T, R> List<T>.mapChecked6191(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6192(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6193(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step6194")
fun step6194(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
data class Pair6195<A, B>(val first: A, val second: B)

fun combine6195(left: Int, right: Int): Pair6195<Int, Int> {
    return Pair6195(left + 36, right + 66)
}
fun <T, R> List<T>.mapChecked6196(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull6197(): T? = if (size > 1) this[1] else null

fun render6197(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
