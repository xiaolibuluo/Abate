package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter9490 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull9491(): T? = if (size > 1) this[1] else null

fun render9491(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9492 {
    data class Success(val payload: String) : Outcome9492()
    data object Failure : Outcome9492()
}

fun describe9492(outcome: Outcome9492): String = when (outcome) {
    is Outcome9492.Success -> outcome.payload
    Outcome9492.Failure -> "failed"
}
sealed class Outcome9493 {
    data class Success(val payload: String) : Outcome9493()
    data object Failure : Outcome9493()
}

fun describe9493(outcome: Outcome9493): String = when (outcome) {
    is Outcome9493.Success -> outcome.payload
    Outcome9493.Failure -> "failed"
}
