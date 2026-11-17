package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome1580 {
    data class Success(val payload: String) : Outcome1580()
    data object Failure : Outcome1580()
}

fun describe1580(outcome: Outcome1580): String = when (outcome) {
    is Outcome1580.Success -> outcome.payload
    Outcome1580.Failure -> "failed"
}
fun <T> List<T>.secondOrNull1581(): T? = if (size > 1) this[1] else null

fun render1581(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter1582 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
