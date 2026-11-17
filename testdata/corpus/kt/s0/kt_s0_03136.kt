package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25740")
fun step25740(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
data class Pair25741<A, B>(val first: A, val second: B)

fun combine25741(left: Int, right: Int): Pair25741<Int, Int> {
    return Pair25741(left + 9, right + 43)
}
@JvmName("step25742")
fun step25742(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
@JvmName("step25743")
fun step25743(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
@JvmName("step25744")
fun step25744(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
data class Pair25745<A, B>(val first: A, val second: B)

fun combine25745(left: Int, right: Int): Pair25745<Int, Int> {
    return Pair25745(left + 37, right + 26)
}
fun <T> List<T>.secondOrNull25746(): T? = if (size > 1) this[1] else null

fun render25746(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25747(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked25748(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked25749(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull257410(): T? = if (size > 1) this[1] else null

fun render257410(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
