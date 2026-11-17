package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15740 {
    data class Success(val payload: String) : Outcome15740()
    data object Failure : Outcome15740()
}

fun describe15740(outcome: Outcome15740): String = when (outcome) {
    is Outcome15740.Success -> outcome.payload
    Outcome15740.Failure -> "failed"
}
object Registry15741 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15741 = "registry-7"
    }
}
fun <T> List<T>.secondOrNull15742(): T? = if (size > 1) this[1] else null

fun render15742(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15743<A, B>(val first: A, val second: B)

fun combine15743(left: Int, right: Int): Pair15743<Int, Int> {
    return Pair15743(left + 21, right + 19)
}
