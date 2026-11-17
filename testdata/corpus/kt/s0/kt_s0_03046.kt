package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25000 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
sealed class Outcome25001 {
    data class Success(val payload: String) : Outcome25001()
    data object Failure : Outcome25001()
}

fun describe25001(outcome: Outcome25001): String = when (outcome) {
    is Outcome25001.Success -> outcome.payload
    Outcome25001.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25002(): T? = if (size > 1) this[1] else null

fun render25002(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
