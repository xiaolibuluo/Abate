package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6140(): T? = if (size > 1) this[1] else null

fun render6140(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6141(): T? = if (size > 1) this[1] else null

fun render6141(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome6142 {
    data class Success(val payload: String) : Outcome6142()
    data object Failure : Outcome6142()
}

fun describe6142(outcome: Outcome6142): String = when (outcome) {
    is Outcome6142.Success -> outcome.payload
    Outcome6142.Failure -> "failed"
}
class Counter6143 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
sealed class Outcome6144 {
    data class Success(val payload: String) : Outcome6144()
    data object Failure : Outcome6144()
}

fun describe6144(outcome: Outcome6144): String = when (outcome) {
    is Outcome6144.Success -> outcome.payload
    Outcome6144.Failure -> "failed"
}
