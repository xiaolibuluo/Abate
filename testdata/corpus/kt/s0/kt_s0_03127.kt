package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25660 {
    data class Success(val payload: String) : Outcome25660()
    data object Failure : Outcome25660()
}

fun describe25660(outcome: Outcome25660): String = when (outcome) {
    is Outcome25660.Success -> outcome.payload
    Outcome25660.Failure -> "failed"
}
data class Pair25661<A, B>(val first: A, val second: B)

fun combine25661(left: Int, right: Int): Pair25661<Int, Int> {
    return Pair25661(left + 99, right + 48)
}
fun <T, R> List<T>.mapChecked25662(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25663<A, B>(val first: A, val second: B)

fun combine25663(left: Int, right: Int): Pair25663<Int, Int> {
    return Pair25663(left + 16, right + 9)
}
object Registry25664 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25664 = "registry-23"
    }
}
data class Pair25665<A, B>(val first: A, val second: B)

fun combine25665(left: Int, right: Int): Pair25665<Int, Int> {
    return Pair25665(left + 30, right + 31)
}
@JvmName("step25666")
fun step25666(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
fun <T> List<T>.secondOrNull25667(): T? = if (size > 1) this[1] else null

fun render25667(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
