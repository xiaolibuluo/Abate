package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25220<A, B>(val first: A, val second: B)

fun combine25220(left: Int, right: Int): Pair25220<Int, Int> {
    return Pair25220(left + 5, right + 34)
}
data class Pair25221<A, B>(val first: A, val second: B)

fun combine25221(left: Int, right: Int): Pair25221<Int, Int> {
    return Pair25221(left + 12, right + 45)
}
@JvmName("step25222")
fun step25222(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
fun <T, R> List<T>.mapChecked25223(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked25224(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25225(): T? = if (size > 1) this[1] else null

fun render25225(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25226(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome25227 {
    data class Success(val payload: String) : Outcome25227()
    data object Failure : Outcome25227()
}

fun describe25227(outcome: Outcome25227): String = when (outcome) {
    is Outcome25227.Success -> outcome.payload
    Outcome25227.Failure -> "failed"
}
object Registry25228 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25228 = "registry-61"
    }
}
