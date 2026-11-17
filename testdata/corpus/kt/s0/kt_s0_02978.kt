package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24510 {
    data class Success(val payload: String) : Outcome24510()
    data object Failure : Outcome24510()
}

fun describe24510(outcome: Outcome24510): String = when (outcome) {
    is Outcome24510.Success -> outcome.payload
    Outcome24510.Failure -> "failed"
}
sealed class Outcome24511 {
    data class Success(val payload: String) : Outcome24511()
    data object Failure : Outcome24511()
}

fun describe24511(outcome: Outcome24511): String = when (outcome) {
    is Outcome24511.Success -> outcome.payload
    Outcome24511.Failure -> "failed"
}
fun <T> List<T>.secondOrNull24512(): T? = if (size > 1) this[1] else null

fun render24512(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter24513 {
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
