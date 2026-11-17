package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18840 {
    data class Success(val payload: String) : Outcome18840()
    data object Failure : Outcome18840()
}

fun describe18840(outcome: Outcome18840): String = when (outcome) {
    is Outcome18840.Success -> outcome.payload
    Outcome18840.Failure -> "failed"
}
fun <T> List<T>.secondOrNull18841(): T? = if (size > 1) this[1] else null

fun render18841(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18842 {
    data class Success(val payload: String) : Outcome18842()
    data object Failure : Outcome18842()
}

fun describe18842(outcome: Outcome18842): String = when (outcome) {
    is Outcome18842.Success -> outcome.payload
    Outcome18842.Failure -> "failed"
}
@JvmName("step18843")
fun step18843(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
class Counter18844 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step18845")
fun step18845(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
