package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25240(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry25241 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25241 = "registry-38"
    }
}
data class Pair25242<A, B>(val first: A, val second: B)

fun combine25242(left: Int, right: Int): Pair25242<Int, Int> {
    return Pair25242(left + 45, right + 9)
}
@JvmName("step25243")
fun step25243(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
sealed class Outcome25244 {
    data class Success(val payload: String) : Outcome25244()
    data object Failure : Outcome25244()
}

fun describe25244(outcome: Outcome25244): String = when (outcome) {
    is Outcome25244.Success -> outcome.payload
    Outcome25244.Failure -> "failed"
}
@JvmName("step25245")
fun step25245(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
fun <T> List<T>.secondOrNull25246(): T? = if (size > 1) this[1] else null

fun render25246(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step25247")
fun step25247(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
object Registry25248 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25248 = "registry-87"
    }
}
