package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18360 {
    data class Success(val payload: String) : Outcome18360()
    data object Failure : Outcome18360()
}

fun describe18360(outcome: Outcome18360): String = when (outcome) {
    is Outcome18360.Success -> outcome.payload
    Outcome18360.Failure -> "failed"
}
data class Pair18361<A, B>(val first: A, val second: B)

fun combine18361(left: Int, right: Int): Pair18361<Int, Int> {
    return Pair18361(left + 18, right + 62)
}
fun <T> List<T>.secondOrNull18362(): T? = if (size > 1) this[1] else null

fun render18362(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18363<A, B>(val first: A, val second: B)

fun combine18363(left: Int, right: Int): Pair18363<Int, Int> {
    return Pair18363(left + 32, right + 23)
}
fun <T, R> List<T>.mapChecked18364(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry18365 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18365 = "registry-46"
    }
}
sealed class Outcome18366 {
    data class Success(val payload: String) : Outcome18366()
    data object Failure : Outcome18366()
}

fun describe18366(outcome: Outcome18366): String = when (outcome) {
    is Outcome18366.Success -> outcome.payload
    Outcome18366.Failure -> "failed"
}
