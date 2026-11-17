package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25070 {
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
fun <T> List<T>.secondOrNull25071(): T? = if (size > 1) this[1] else null

fun render25071(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome25072 {
    data class Success(val payload: String) : Outcome25072()
    data object Failure : Outcome25072()
}

fun describe25072(outcome: Outcome25072): String = when (outcome) {
    is Outcome25072.Success -> outcome.payload
    Outcome25072.Failure -> "failed"
}
sealed class Outcome25073 {
    data class Success(val payload: String) : Outcome25073()
    data object Failure : Outcome25073()
}

fun describe25073(outcome: Outcome25073): String = when (outcome) {
    is Outcome25073.Success -> outcome.payload
    Outcome25073.Failure -> "failed"
}
