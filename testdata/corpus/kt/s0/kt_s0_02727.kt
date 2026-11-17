package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22500 {
    data class Success(val payload: String) : Outcome22500()
    data object Failure : Outcome22500()
}

fun describe22500(outcome: Outcome22500): String = when (outcome) {
    is Outcome22500.Success -> outcome.payload
    Outcome22500.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22501(): T? = if (size > 1) this[1] else null

fun render22501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter22502 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
