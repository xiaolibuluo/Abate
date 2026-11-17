package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25690(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25691(): T? = if (size > 1) this[1] else null

fun render25691(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25692(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry25693 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25693 = "registry-55"
    }
}
sealed class Outcome25694 {
    data class Success(val payload: String) : Outcome25694()
    data object Failure : Outcome25694()
}

fun describe25694(outcome: Outcome25694): String = when (outcome) {
    is Outcome25694.Success -> outcome.payload
    Outcome25694.Failure -> "failed"
}
data class Pair25695<A, B>(val first: A, val second: B)

fun combine25695(left: Int, right: Int): Pair25695<Int, Int> {
    return Pair25695(left + 69, right + 52)
}
data class Pair25696<A, B>(val first: A, val second: B)

fun combine25696(left: Int, right: Int): Pair25696<Int, Int> {
    return Pair25696(left + 76, right + 63)
}
fun <T> List<T>.secondOrNull25697(): T? = if (size > 1) this[1] else null

fun render25697(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
