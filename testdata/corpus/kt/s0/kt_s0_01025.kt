package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8420 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
sealed class Outcome8421 {
    data class Success(val payload: String) : Outcome8421()
    data object Failure : Outcome8421()
}

fun describe8421(outcome: Outcome8421): String = when (outcome) {
    is Outcome8421.Success -> outcome.payload
    Outcome8421.Failure -> "failed"
}
fun <T> List<T>.secondOrNull8422(): T? = if (size > 1) this[1] else null

fun render8422(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
