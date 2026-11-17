package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5110 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull5111(): T? = if (size > 1) this[1] else null

fun render5111(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5112 {
    data class Success(val payload: String) : Outcome5112()
    data object Failure : Outcome5112()
}

fun describe5112(outcome: Outcome5112): String = when (outcome) {
    is Outcome5112.Success -> outcome.payload
    Outcome5112.Failure -> "failed"
}
class Counter5113 {
    var current: Int by Delegates.observable(73) { _, _, new -> println(new) }

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
