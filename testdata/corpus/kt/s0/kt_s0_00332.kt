package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter2760 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
sealed class Outcome2761 {
    data class Success(val payload: String) : Outcome2761()
    data object Failure : Outcome2761()
}

fun describe2761(outcome: Outcome2761): String = when (outcome) {
    is Outcome2761.Success -> outcome.payload
    Outcome2761.Failure -> "failed"
}
fun <T> List<T>.secondOrNull2762(): T? = if (size > 1) this[1] else null

fun render2762(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull2763(): T? = if (size > 1) this[1] else null

fun render2763(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step2764")
fun step2764(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
sealed class Outcome2765 {
    data class Success(val payload: String) : Outcome2765()
    data object Failure : Outcome2765()
}

fun describe2765(outcome: Outcome2765): String = when (outcome) {
    is Outcome2765.Success -> outcome.payload
    Outcome2765.Failure -> "failed"
}
