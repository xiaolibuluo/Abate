package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19740 {
    data class Success(val payload: String) : Outcome19740()
    data object Failure : Outcome19740()
}

fun describe19740(outcome: Outcome19740): String = when (outcome) {
    is Outcome19740.Success -> outcome.payload
    Outcome19740.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19741(): T? = if (size > 1) this[1] else null

fun render19741(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19742<A, B>(val first: A, val second: B)

fun combine19742(left: Int, right: Int): Pair19742<Int, Int> {
    return Pair19742(left + 73, right + 63)
}
fun <T> List<T>.secondOrNull19743(): T? = if (size > 1) this[1] else null

fun render19743(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry19744 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19744 = "registry-87"
    }
}
object Registry19745 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19745 = "registry-94"
    }
}
