package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8440 {
    var current: Int by Delegates.observable(16) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull8441(): T? = if (size > 1) this[1] else null

fun render8441(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome8442 {
    data class Success(val payload: String) : Outcome8442()
    data object Failure : Outcome8442()
}

fun describe8442(outcome: Outcome8442): String = when (outcome) {
    is Outcome8442.Success -> outcome.payload
    Outcome8442.Failure -> "failed"
}
